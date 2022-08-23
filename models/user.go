package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	Base          `xorm:"-" json:"-"`
	Id            int64  `xorm:"not null pk autoincr comment('主键ID') INT" json:"id,omitempty"`
	Username      string `xorm:"not null default '' comment('用户名') VARCHAR(50)" form:"username" binding:"required,min=6,max=20" json:"username,omitempty"`
	Password      string `xorm:"not null default '' comment('账号密码') VARCHAR(255)" form:"password" binding:"required,min=6,max=50" json:"-" json:"password,omitempty"`
	Email         string `xorm:"not null default '' comment('邮箱地址') VARCHAR(255)" form:"email"  binding:"required,email" json:"email,omitempty"`
	Phone         string `xorm:"not null default '' comment('手机号') CHAR(11)" form:"phone"  binding:"required,len=11" json:"phone,omitempty"`
	Status        int    `xorm:"not null default 0 comment('账号状态 0 正常 1 禁用') SMALLINT" json:"status,omitempty"`
	CreateTime    int64  `xorm:"not null default 0 comment('账号创建时间') INT" json:"create_time,omitempty"`
	UpdateTime    int64  `xorm:"not null default 0 comment('账号信息修改时间') INT" json:"update_time,omitempty"`
	LastLoginTime int64  `xorm:"not null default 0 comment('最后一次登陆时间') INT" json:"last_login_time,omitempty"`
	LastLoginIp   int64  `xorm:"default 0 comment('登陆ip') INT" json:"last_login_ip,omitempty"`
}

// LoginUser 登陆校验Model
type LoginUser struct {
	Id            int64  `xorm:"not null pk autoincr comment('主键ID') INT"`
	Username      string `xorm:"not null default '' comment('用户名') VARCHAR(50)" form:"username" binding:"required_without_all=Email Phone,min=6,max=20"`
	Password      string `xorm:"not null default '' comment('账号密码') VARCHAR(255)" form:"password" binding:"required,min=6,max=50" json:"-"`
	Email         string `xorm:"not null default '' comment('邮箱地址') VARCHAR(255)" form:"email"  binding:"required_without_all=Username Phone"`
	Phone         string `xorm:"not null default '' comment('手机号') CHAR(11)" form:"phone"  binding:"required_without_all=Username Email"`
}

// StatusNormal 正常
const StatusNormal = 0
// StatusDisAble 禁用
const StatusDisAble = 1

//TableName 返回表名
func (u User) TableName() string  {
	return "user"
}

//NewUser 实例化userModel
func NewUser() *User {
    return &User{}
}

// Add 注册
func (u *User) Add() (int64, error) {
	// 校验用户是否已存在
	exists, err := u.GetDb().Table(u.TableName()).Where("username = ? or email = ? or phone = ?", u.Username, u.Email, u.Phone).Exist()
	if err != nil {
		return 0, err
	}
	if exists {
		return 0, errors.New("用户已存在")
	}
	// 加密密码
	pass, err := u.GeneratePassword(u.Password)
	if err != nil {
        return 0, err
    }
	u.Password      = pass
	// 赋值默认字段
	u.Status        = StatusNormal
	u.CreateTime    = time.Now().Unix()
	u.UpdateTime    = time.Now().Unix()
	u.LastLoginTime = time.Now().Unix()
    return u.GetDb().Insert(u)
}

//Login 登陆
func (u *User) Login(user *User) (*User, error){
	// 用户是否存在
	get, err := u.GetDb().Get(u)
	if err != nil || get == false {
		return nil, errors.New("用户名或密码错误")
	}
	// 密码是否正确
	password := u.ComparePassword(user.Password, u.Password)
	if password{
		// 更新用户登陆信息
		err := u.UpdateUserLoginInfo()
		if err != nil {
			return nil, err
		}
		return u, nil
	}
	return nil, errors.New("用户名或密码错误")
}

//UpdateUserLoginInfo 修改用户登录信息
func (u *User) UpdateUserLoginInfo() error{
	u.LastLoginTime = time.Now().Unix()
	update, err := u.GetDb().Update(u)
	if err != nil || update == 0 {
		return err
	}
	return nil
}

//GeneratePassword 密码加密
func (u *User) GeneratePassword(pass string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "nil", err
	}
	return string(password), err
}

// ComparePassword 密码解密
func (u *User) ComparePassword(pass, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
    return err == nil
}

//FindUserOfId 根据用户ID查询用户信息
func (u *User) FindUserOfId(id int64) *User{
	get, err := u.GetDb().Table(u.TableName()).Where("id = ?", id).Get(u)
	if err != nil {
		return nil
	}
	if get{
		return u
	}
	return nil
}
// 修改
// 删除


