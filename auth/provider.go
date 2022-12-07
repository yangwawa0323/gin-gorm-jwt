package auth

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/casbin/casbin"
	"github.com/yangwawa0323/gin-gorm-jwt/models"
)

const (
	POST     Activity = "POST"
	GET      Activity = "GET"
	READ     Activity = "READ"
	WRITE    Activity = "READ"
	DELETE   Activity = "DELETE"
	UPLOAD   Activity = "UPLOAD"
	DOWNLOAD Activity = "DOWNLOAD"
	FAVORITE Activity = "FAVORITE"
)

type AuthCategory int64

type Role string
type Resource string
type Privilege string
type Unused string
type Activity string

const (
	OAUTH2 AuthCategory = iota
	PAYMENT
)

type Oauth2Provider interface {
	Login(*models.User) error
	Logout(*models.User) error
	ValidateClient(clientID string, clientSecret string) bool
}

type FakerProvider struct {
	Name         string
	Catagory     AuthCategory
	Type         string
	ClientID     string
	ClientSecret string
	ProviderUrl  string
}

func (fp *FakerProvider) Login(*models.User) error {
	// TODO:
	return errors.New("not implemented yet")
}

func (fp *FakerProvider) Logout(*models.User) error {
	// TODO:
	return errors.New("not implemented yet")
}

func (fp *FakerProvider) ValidateClient(clientID string, clientString string) bool {
	// TODO: not implemented yet
	return false
}

type Application struct {
	Description string
	RedirectUrl string
	// SignUpUrl		string
	// SignInUrl		string
	ClientID        string
	ClientSecret    string
	Provider        Oauth2Provider
	Token           JWTClaim
	TokenExpire     time.Time
	RefreshInterval time.Duration
}

/*
Each user each activity will check Casbin rule will cause a lot DB query slow down the whole system.
There are three prefer caching system library.

1. https://github.com/bluele/gcache
An in-memory cache library for golang. It supports multiple eviction policies: LRU, LFU, ARC

2. https://github.com/go-redis/cache
Cache library with Redis backend for Golang

3.https://github.com/dgraph-io/ristretto
A high performance memory-bound Go cache
*/

// support timeout context
type CacheProviderInt interface {
	Set(ctx *context.Context, key string, item *Permission) error
	Get(ctx *context.Context, key string) (item *Permission, err error)
}

type PrivilegeProvider interface {
	GrantToUser(*Permission) error
	RevokeFromUser(*Permission) error

	GrantToRole(*Permission) error
	RevokeFromRole(*Permission) error

	HasPermission(user *models.User, activity string, resource string) bool

	// If not in the cache will query the database, and update the cache with whole rules keep in the
	// DB rules table.
	IsCached() bool
	CacheTo(cacheProvider *CacheProviderInt) error
}

type CasbinPrivilegeProvider struct {
	Enforce *casbin.Enforcer
}

func (casbin *CasbinPrivilegeProvider) GrantToUser(*Permission) error {
	// TODO: AddPolicy parameters , `ptype` to "p"
	if ok := casbin.Enforce.AddPolicy(); !ok {
		return errors.New("can not add policy to casbin")
	}
	return nil
}

func (casbin *CasbinPrivilegeProvider) RevokeFromUser(*Permission) error {
	// TODO: RemovePolicy parameters
	if ok := casbin.Enforce.RemovePolicy(); !ok {
		return errors.New("can not revoke policy from casbin or not such policy")
	}
	return nil
}

func (casbin *CasbinPrivilegeProvider) GrantToRole(*Permission) error {
	// TODO: AddPolicy parameters, `ptype` to "g"
	if ok := casbin.Enforce.AddNamedPolicy("g"); !ok {
		return errors.New("can not add policy to casbin")
	}
	return nil
}

func (casbin *CasbinPrivilegeProvider) RevokeFromRole(*Permission) error {
	// TODO
	if ok := casbin.Enforce.RemoveNamedPolicy("g"); !ok {
		return errors.New("can not revoke policy from casbin")
	}
	return nil
}

func (casbin *CasbinPrivilegeProvider) HasPermission(user *models.User, activity string, resource string) bool {
	// TODO: Find permission according the user.Username
	return casbin.Enforce.HasPermissionForUser(user.Username)
}

func (casbin *CasbinPrivilegeProvider) IsCached() bool {
	// TODO: using the cache system, open the conf/app.conf
	return false
}

func (casbin *CasbinPrivilegeProvider) CacheTo(cacheProvider *CacheProviderInt) error {
	// TODO: cache whole casbin policy to Cache
	return fmt.Errorf("cannot cache the specified cache provider %#v", casbin)
}

// Permission struct  as prototype of casbin_rule table as well, but more readable with column attributes.
// Read only struct. only use for gorm scan data to struct
type Permission struct {
	Ptype string `gorm:"column:ptype:type:text;<-:false"`
	// Username:  joe
	V0 string `gorm:"column:v0;type:text;<-:false"`
	// Role : guest
	V1 Role `gorm:"column:v1;type:text;<-:false"`
	// Activity: POST
	V2 string `gorm:"column:v2;type:text;<-:false"`
	// Resource :  api/page/new
	V3 Resource `gorm:"column:v3;type:text;<-:false"`
	// Privilege : ALLOW
	V4 Privilege `gorm:"column:v4;type:text;<-:false"`
	// Unused:  -
	V5 Unused `gorm:"column:v5;type:text;<-:false"`
}

func LoadFromCasbin() ([]*Permission, error) {
	return nil, errors.New("not implemented yet")
}

func (perm *Permission) SaveToCasbin() error {
	return errors.New("not implemented yet")
}
