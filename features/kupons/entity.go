package kupon

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Kupon struct {
	ID          uint
	NamaProgram string
	KodeKupon   string
	LinkKupon   string
	GambarKupon string
}

type Handler interface {
	Add() echo.HandlerFunc
	GetOne() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type Service interface {
	TambahKupon(token *jwt.Token, newKupon Kupon) (Kupon, error)
	LihatKuponUser(token *jwt.Token) ([]Kupon, error)
	LihatSemuaKupon() ([]Kupon, error)
}

type Repository interface {
	InsertKupon(userID uint, newKupon Kupon) (Kupon, error)
	GetKuponUser(userID uint) ([]Kupon, error)
	GetAllKupons() ([]Kupon, error)
}
