package service

import (
	"context"
	"errors"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/domain"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)
var ErrUserDuplicateEmail=repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword=errors.New("账号/邮箱或密码错误")
type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	hash,err:=bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password=string(hash)
	return svc.repo.Create(ctx, u)
}
func (svc *UserService) Login(ctx context.Context, email,password string) error {
	u,err:=svc.repo.FindByEmail(ctx,email)
	if err==repository.ErrUserNotFound{
		return ErrInvalidUserOrPassword
	}
	if err != nil {
		return err
	}
	//比较密码
	err=bcrypt.CompareHashAndPassword([]byte(u.Password),[]byte(password))
	if err != nil {
		return ErrInvalidUserOrPassword
	}
	return nil
}