package repository

import (
	"context"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/domain"
	"github.com/Xiaodingzhishang/jk_work/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail=dao.ErrUserDuplicateEmail
	ErrUserNotFound=dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{dao: dao}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	return r.dao.Insert(ctx, dao.User{
		Email:    u.Email,
		Password: u.Password,
	})
	//先从cache里面找
	//再从dao里面找
	//找到了回写cache
}
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User,error) {
	u,err:=r.dao.FindByEmail(ctx,email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		Email:    u.Email,
		Password: u.Password,
	},nil
}

func (r *UserRepository) FindById(int64) {
	//先从cache里面找
	//再从dao里面找
	//找到了回写cache
}
