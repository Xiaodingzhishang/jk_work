package web

import (
	"errors"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/domain"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/service"
	regexp "github.com/dlclark/regexp2"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

type UserHandle struct {
	svc         *service.UserService
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
}

func NewUserHandle(svc *service.UserService) *UserHandle {
	const (
		emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
	return &UserHandle{
		svc:         svc,
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

//func (u *UserHandle) RegisterRoutesV1(ug *gin.RouterGroup) {
//	ug.POST("/signup", u.SignUp)
//	ug.POST("/login", u.Login)
//	ug.POST("/edit", u.Edit)
//	ug.GET("/profile", u.Profile)
//}

func (u *UserHandle) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.LoginJWT)
	ug.POST("/edit", u.Edit)
	ug.GET("/profile", u.Profile)
}
func (u *UserHandle) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}
	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	ok, err := u.emailExp.MatchString(req.Email)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "你的邮箱格式不对")
		return
	}

	if req.ConfirmPassword != req.Password {
		ctx.String(http.StatusOK, "两次输入的密码不相同")
		return
	}

	ok, err = u.passwordExp.MatchString(req.Password)
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	if !ok {
		ctx.String(http.StatusOK, "密码必须包含数字、特殊字符，并且长度不能小于 8 位")
		return
	}
	err = u.svc.SignUp(ctx, domain.User{
		Email:    req.Email,
		Password: req.Password,
	})
	if errors.Is(err, service.ErrUserDuplicateEmail) {
		ctx.String(http.StatusOK, "邮箱冲突")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统异常")
		return
	}
	ctx.String(http.StatusOK, "响应成功")
}
func (u *UserHandle) LoginJWT(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名/邮箱或密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//在这里登入成功了
	token := jwt.New(jwt.SigningMethodHS512)
	tokenstr, err := token.SignedString([]byte("Mwz4s3qHZjMpXJdUTn7QT5TZ7ztWWWaY"))
	if err != nil {
		ctx.String(http.StatusInternalServerError, "系统错误")
		return
	}
	ctx.Header("x-jwt-token", tokenstr)

	ctx.String(http.StatusOK, "登入成功")
}
func (u *UserHandle) Login(ctx *gin.Context) {
	type LoginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req LoginReq
	if err := ctx.Bind(&req); err != nil {
		return
	}
	err := u.svc.Login(ctx, req.Email, req.Password)
	if err == service.ErrInvalidUserOrPassword {
		ctx.String(http.StatusOK, "用户名/邮箱或密码错误")
		return
	}
	if err != nil {
		ctx.String(http.StatusOK, "系统错误")
		return
	}
	//在这里登入成功了
	sess := sessions.Default(ctx)
	sess.Set("userId", "hello word")
	sess.Options(sessions.Options{
		//生产环境
		//Secure: true,
		//HttpOnly: true,
		MaxAge: 60,
	})
	sess.Save()
	ctx.String(http.StatusOK, "登入成功")
}
func (u *UserHandle) Logout(ctx *gin.Context) {
	sess := sessions.Default(ctx)
	sess.Options(sessions.Options{
		//生产环境
		//Secure: true,
		//HttpOnly: true,
		MaxAge: -1,
	})
	sess.Save()
	ctx.String(http.StatusOK, "退出登入成功")

}

func (u *UserHandle) Edit(ctx *gin.Context) {

}
func (u *UserHandle) Profile(ctx *gin.Context) {
	ctx.String(http.StatusOK, "这是你的profile")
	return
}
