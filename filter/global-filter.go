package filter

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"gorm-study/common"
	"gorm-study/service"
	"gorm-study/utils/redis"
	"gorm-study/utils/token"
	"net/http"
	"strings"
)

var userService = service.NewUserService()

func GlobalFilter() gin.HandlerFunc {

	return func(context *gin.Context) {

		uris := filterUris()

		index := findWhiteUris(uris, context.Request.URL.Path)

		var currentUri filterChain

		if index == -1 {
			currentUri = defaultFilterChain()
		} else {
			currentUri = uris[index]
		}

		if currentUri.Next {
			context.Next()
			return
		}

		var tokenHeader = context.GetHeader(common.TokenHeader)

		if tokenHeader == "" || !strings.HasPrefix(tokenHeader, common.PrefixToken) {
			context.JSON(http.StatusForbidden, common.BuildFailure("没有权限访问"))
			context.Abort()
			return
		}

		var resultToken = tokenHeader[len(common.PrefixToken):]

		username := token.ParseTokenFormToUsername(resultToken)

		if username == "" {
			context.JSON(http.StatusForbidden, common.BuildFailure("token可能已失效,请重新登录!"))
			context.Abort()
			return
		}

		accessToken := redis.GetString(common.AccessTokenPrefix + base64.StdEncoding.EncodeToString([]byte(username)))

		if accessToken != "" && accessToken == resultToken {
			resultUser := userService.FindByUsername(username)

			var existsRole, existsPermission = false, false

			existsRole = hasRole(currentUri.Role, resultUser.Role.Name)

			var permissions = make([]string, len(currentUri.Permission))

			for i, permission := range resultUser.Permissions {
				permissions[i] = permission.Name
			}

			existsPermission = hasPermission(currentUri.Permission, permissions)

			if existsRole && existsPermission {
				context.Next()
				return
			}

		}

		context.JSON(http.StatusForbidden, common.BuildFailure("你没有权限访问"))

		context.Abort()
	}

}

func hasPermission(permissions []string, userPermission []string) bool {

	if len(permissions) == 0 || len(userPermission) == 0 {
		return false
	}

	var map1, map2 = make(map[string]int), make(map[string]int)

	for i, permission := range permissions {
		map1[permission] = i
	}

	for i, permission := range userPermission {
		map2[permission] = i
	}

	for key, _ := range map1 {
		if _, ok := map2[key]; ok {
			return true
		}
	}

	return false
}

func hasRole(roles []string, roleName string) bool {
	for _, r := range roles {
		if r == roleName {
			return true
		}
	}
	return false
}

func findWhiteUris(uris []filterChain, str string) int {

	for index, result := range uris {

		var r = result.Path

		if strings.HasSuffix(r, "/**") && strings.HasPrefix(str, r[0:strings.LastIndex(r, "/**")+1]) {
			return index
		} else if r == str {
			return index
		}

	}

	return -1
}

type filterChain struct {
	Path       string
	Next       bool
	Role       []string
	Permission []string
}

func defaultFilterChain() filterChain {
	return filterChain{Path: "/api/v1/user/findAll", Next: true, Role: []string{common.RoleUser}, Permission: []string{common.PermissionSelect}}
}

func filterUris() []filterChain {

	var chains []filterChain

	chains = append(chains,
		filterChain{"/api/v1/user/login", true, []string{common.RoleAdmin, common.RoleUser}, []string{common.PermissionInsert, common.PermissionDelete}},
		filterChain{"/api/v1/user/findAll", false, []string{common.RoleAdmin, common.RoleUser}, []string{common.PermissionInsert, common.PermissionDelete}},
	)

	return chains

}
