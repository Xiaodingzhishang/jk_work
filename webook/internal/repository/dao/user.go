package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	ErrUserDuplicateEmail=errors.New("邮箱冲突")
	ErrUserNotFound=gorm.ErrRecordNotFound
)

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Utime = now
	u.Ctime = now
	err:=dao.db.WithContext(ctx).Create(&u).Error
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		const uniqueConflictsErrNo uint16 = 1062
		if mysqlErr.Number == uniqueConflictsErrNo {
			return ErrUserDuplicateEmail
		}
	}
	return err
}
func (dao *UserDAO) FindByEmail(ctx context.Context, email string) (User,error) {
	var u User
	err:=dao.db.WithContext(ctx).Where("email=?",email).First(&u).Error
	//err:=dao.db.WithContext(ctx).First(&u,"email=?",email).Error
	return u,err
}

// User 直接对应数据库结构
type User struct {
	Id       int64 `gorm:"primaryKey,autoIncrement"`
	Email    string `gorm:"unique"`
	Password string

	//往这里面加

	Ctime    int64
	Utime    int64
}
