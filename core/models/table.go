package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Id       int    `gorm:"primary_key" json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RepositoryPool struct {
	gorm.Model
	Id       int    `gorm:"primary_key" json:"id"`
	Identity string `json:"identity"`
	Hash     string `json:"hash"`
	Name     string `json:"name"`
	Ext      string `json:"ext"`
	Size     int64  `json:"size"`
	Path     string `json:"path"`
}

type UserRepository struct {
	gorm.Model
	Id                 int    `gorm:"primary_key" json:"id"`
	Identity           string `json:"identity"`
	ParentId           int64  `json:"parent_id"`
	UserIdentity       string `json:"user_identity"`
	RepositoryIdentity string `json:"repository_identity"`
	Ext                string `json:"ext"`
	Name               string `json:"name"`
}
