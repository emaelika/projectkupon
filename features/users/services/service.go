package services

import (
	"errors"
	"projectkupon/features/users"
	"projectkupon/helper/enkrip"
	"strings"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	repo users.Repository
	h    enkrip.HashInterface
}

func New(r users.Repository, h enkrip.HashInterface) users.Service {
	return &userService{
		repo: r,
		h:    h,
	}
}

func (us *userService) Register(newUser users.User) (users.User, error) {
	// validasi
	var validate = validator.New()
	if err := validate.Var(newUser.Nama, "required,alpha"); err != nil {
		if strings.Contains(err.Error(), "required") {
			return users.User{}, errors.New("harap masukkan nama")
		} else if strings.Contains(err.Error(), "alpha") {
			return users.User{}, errors.New("input tidak valid")
		}
	}

	if err := validate.Var(newUser.HP, "required,numeric"); err != nil {
		if strings.Contains(err.Error(), "required") {
			return users.User{}, errors.New("nomor hp tidak diisi")
		} else if strings.Contains(err.Error(), "numeric") {
			return users.User{}, errors.New("input tidak valid")
		}
	}

	// enkripsi password
	ePassword, err := us.h.HashPassword(newUser.Password)

	if err != nil {
		return users.User{}, errors.New("terdapat masalah saat memproses data")
	}

	newUser.Password = ePassword

	result, err := us.repo.Insert(newUser)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return users.User{}, errors.New("data telah terdaftar pada sistem")
		}
		return users.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	return result, nil
}

func (us *userService) Login(hp string, password string) (users.User, error) {
	result, err := us.repo.Login(hp)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return users.User{}, errors.New("data tidak ditemukan")
		}
		return users.User{}, errors.New("terjadi kesalahan pada sistem")
	}

	err = us.h.Compare(result.Password, password)

	if err != nil {
		return users.User{}, errors.New("password salah")
	}

	return result, nil
}
