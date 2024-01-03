package model

import (
	"context"
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) error
		FindOne(ctx context.Context, id int64) (*User, error)
		Update(ctx context.Context, data *User) error
	}

	defaultUserModel struct {
		db    *gorm.DB
		table string
	}

	User struct {
		Id       int64          `db:"id"`
		Name     sql.NullString `db:"name"`     // The username
		Password string         `db:"password"` // The user password
		Mobile   string         `db:"mobile"`   // The mobile phone number
		Gender   string         `db:"gender"`   // gender,male|female|unknown
		Nickname string         `db:"nickname"` // The nickname
		Type     int64          `db:"type"`     // The user type, 0:normal,1:vip, for test golang keyword
		CreateAt sql.NullTime   `db:"create_at"`
		UpdateAt time.Time      `db:"update_at"`
	}
)

func newUserModel(conn *gorm.DB) *defaultUserModel {
	return &defaultUserModel{
		db:    conn,
		table: "`user`",
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

func (m *defaultUserModel) FindByUserIDAndFollowedUserID(ctx context.Context, userId, followedUserId int64) (*User, error) {
	var result User
	err := m.db.WithContext(ctx).
		Where("user_id = ? AND followed_user_id = ?", userId, followedUserId).
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
