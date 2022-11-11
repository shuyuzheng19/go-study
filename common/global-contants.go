package common

import (
	"time"
)

//Token
const (
	PrefixToken            = "Bearer "
	TokenHeader            = "Authorization"
	AccessTokenExpireTime  = time.Hour * 2
	RefreshTokenExpireTime = time.Hour * 3
	TokenSigningKey        = "2528959216@zsy:love you"
	AccessTokenPrefix      = "access_token:"
	RefreshTokenPrefix     = "refresh_token:"
)

//Roles
const (
	RoleAdmin  = "ADMIN"
	RoleUser   = "USER"
	RoleMaster = "MASTER"
)

//Permission
const (
	PermissionInsert = "INSERT"
	PermissionUpdate = "UPDATE"
	PermissionDelete = "DELETE"
	PermissionSelect = "SELECT"
)
