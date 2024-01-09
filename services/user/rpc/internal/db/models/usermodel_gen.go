package models

import (
	"context"
	"gorm.io/gorm"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) error
		FindOne(ctx context.Context, id int64) (*User, error)
		Update(ctx context.Context, data *User) error
		FindByUserIDAndPassword(ctx context.Context, username string, password string) (*User, error)
	}

	defaultUserModel struct {
		db    *gorm.DB
		table string
	}

	User struct {
		BaseModel
		Name           string `gorm:"column:name;type:varchar(30);not null;comment:显示名"`
		UserName       string `gorm:"column:username;type:varchar(30);not null;unique;comment:用户名(唯一)"`
		Password       string `gorm:"column:password;type:varchar(100);not null;comment:密码"`
		Email          string `gorm:"column:email;type:varchar(60);comment:邮箱"`
		EmployeeNumber string `gorm:"column:employee_number;type:varchar(30);comment:员工号"`
		Phone          string `gorm:"column:phone;type:varchar(30);comment:手机号"`
		Avatar         string `gorm:"column:avatar;type:varchar(255);comment:头像"`
		IsAdmin        bool   `gorm:"column:is_admin;type:bool;not null;comment:是否是管理员"`
	}
)

func (emp User) TableName() string {
	return "users"
}

func newUserModel(conn *gorm.DB) *defaultUserModel {
	return &defaultUserModel{
		db:    conn,
		table: "`users`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Create(data).Error
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).Where("id = ?", id).First(&result).Error
	return &result, err
}

func (m *defaultUserModel) Update(ctx context.Context, data *User) error {
	return m.db.WithContext(ctx).Save(data).Error
}

func (m *defaultUserModel) UpdateFields(ctx context.Context, id int64, values map[string]interface{}) error {
	return m.db.WithContext(ctx).Model(&User{}).Where("id = ?", id).Updates(values).Error
}

func (m *defaultUserModel) FindByUserIDAndPassword(ctx context.Context, username string, password string) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).
		Where("username = ? AND password = ?", username, password).
		First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &result, err
}

func (m *defaultUserModel) FindByUserId(ctx context.Context, userId int64, limit int) ([]*User, error) {
	var result []*User
	err := m.db.WithContext(ctx).
		Where("user_id = ? AND follow_status = ?", userId, 1).
		Order("id desc").
		Limit(limit).
		Find(&result).Error

	return result, err
}

func (m *defaultUserModel) FindByFollowedUserIds(ctx context.Context, userId int64, followedUserIds []int64) ([]*User, error) {
	var result []*User
	err := m.db.WithContext(ctx).
		Where("user_id = ?", userId).
		Where("followed_user_id in (?)", followedUserIds).
		Find(&result).Error

	return result, err
}

func (m *defaultUserModel) FindByFollowedUserId(ctx context.Context, userId int64, limit int) ([]*User, error) {
	var result []*User
	err := m.db.WithContext(ctx).
		Where("followed_user_id = ? AND follow_status = ?", userId, 1).
		Order("id desc").
		Limit(limit).
		Find(&result).Error
	return result, err
}
