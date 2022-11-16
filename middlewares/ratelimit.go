package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

type RATELIMIT_TYPE int32

const (
	POST_NEW_PAGE = iota
	CREATE_USER
)

type ClientActivity struct {
	User           string         `json:"user"`
	RemoteAddr     string         `json:"remote"`
	LatestPostTime time.Time      `json:"latestpost"`
	Type           RATELIMIT_TYPE `json:-`
}

var Cache = cache.New(5*time.Minute, 5*time.Minute)

func RateLimitErrorDetail(limitType RATELIMIT_TYPE) string {
	switch limitType {
	case POST_NEW_PAGE:
		return "Post new page too fast"
	case CREATE_USER:
		return "Create use too fast"
	default:
		return ""
	}
}

// ReteLimit is depend on the remote client address
func RateLimit(limitType RATELIMIT_TYPE, duration time.Duration) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		remote := ctx.ClientIP()

		// First time is not in the cache
		activity, found := Cache.Get(remote)
		if !found {
			// add to cache
			Cache.Set(remote, &ClientActivity{
				RemoteAddr:     remote,
				LatestPostTime: time.Now(),
				Type:           limitType,
			}, cache.DefaultExpiration)
		} else {
			recent := activity.(*ClientActivity)
			if recent.LatestPostTime.Add(duration).After(time.Now()) {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": RateLimitErrorDetail(limitType),
				})
				ctx.Abort()
				return
			}
			// After save the page reset the LatestPostTime
			recent.LatestPostTime = time.Now()
		}
		ctx.Next()
	}
}
