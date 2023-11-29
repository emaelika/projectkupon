package repository

import (
	kupon "projectkupon/features/kupons"
	"projectkupon/features/users"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Nama     string
	Hp       string
	Password string
	Kupons   []kupon.Kupon `gorm:"foreignKey:id"`
}

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.Repository {
	return &userQuery{
		db: db,
	}
}

func (uq *userQuery) Insert(newUser users.User) (users.User, error) {
	var inputDB = new(UserModel)
	inputDB.Hp = newUser.HP
	inputDB.Nama = newUser.Nama
	inputDB.Password = newUser.Password

	if err := uq.db.Create(&inputDB).Error; err != nil {
		return users.User{}, err
	}

	newUser.ID = inputDB.ID

	return newUser, nil
}

func (uq *userQuery) Login(hp string) (users.User, error) {
	var userData = new(UserModel)

	if err := uq.db.Where("hp = ?", hp).First(userData).Error; err != nil {
		return users.User{}, err
	}

	var result = new(users.User)
	result.ID = userData.ID
	result.HP = userData.Hp
	result.Nama = userData.Nama
	result.Password = userData.Password

	return *result, nil
}
