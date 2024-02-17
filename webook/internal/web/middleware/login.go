package middleware

import (
	"encoding/gob"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginMiddlewareBuilder struct {
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}
func (l *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	l.paths = append(l.paths, path)
	return l
}

func (l *LoginMiddlewareBuilder) Builder() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		for _, path := range l.paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}

		sess := sessions.Default(ctx)
		id := sess.Get("userId")

		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		updateTime := sess.Get("update_time")
		sess.Set("userId", id)
		sess.Options(sessions.Options{
			//生产环境
			//Secure: true,
			//HttpOnly: true,
			MaxAge: 60,
		})
		now := time.Now().UnixMilli()
		//还没有刷新过 刚登入，还没有刷新
		if updateTime == nil {
			sess.Set("update_time", now)
			sess.Save()
			return
		}
		updateTimeVal, ok := updateTime.(int64)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if now-updateTimeVal > 60*1000 {
			sess.Set("update_time", now)
			sess.Save()
			return
		}
	}
}

func (l *LoginMiddlewareBuilder) BuilderV1() gin.HandlerFunc {
	gob.Register(time.Now())
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
			return
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")

		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		updateTime := sess.Get("update_time")
		sess.Set("userId", id)
		sess.Options(sessions.Options{
			//生产环境
			//Secure: true,
			//HttpOnly: true,
			MaxAge: 60,
		})
		now := time.Now()
		//还没有刷新过 刚登入，还没有刷新
		if updateTime == nil {
			sess.Set("update_time", now)
			sess.Save()
			return
		}
		updateTimeVal, ok := updateTime.(time.Time)
		if !ok {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if now.Sub(updateTimeVal) > time.Second*10 {
			sess.Set("update_time", now)
			sess.Save()
			return
		}
	}
}
func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.Request.URL.Path == "/users/login" || ctx.Request.URL.Path == "/users/signup" {
			return
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

var IgnorePaths []string

func CheckLoginV1() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range IgnorePaths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")

		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func CheckLoginV2(paths []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, path := range paths {
			if ctx.Request.URL.Path == path {
				return
			}
		}
		sess := sessions.Default(ctx)
		id := sess.Get("userId")
		if id == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
