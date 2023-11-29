package services

import (
	"errors"
	kupon "projectkupon/features/kupons"
	"projectkupon/helper/jwt"

	"strings"

	"github.com/go-playground/validator/v10"
	golangjwt "github.com/golang-jwt/jwt/v5"
)

type KuponServices struct {
	m kupon.Repository
}

func New(kr kupon.Repository) kupon.Service {
	return &KuponServices{
		m: kr,
	}
}

func (ks *KuponServices) TambahKupon(token *golangjwt.Token, newKupon kupon.Kupon) (kupon.Kupon, error) {
	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return kupon.Kupon{}, errors.New("silakan login terlebih dahulu")
	}
	var validate = validator.New()
	if err := validate.Var(newKupon.LinkKupon, "omitempty,url"); err != nil {
		if strings.Contains(err.Error(), "url") {
			return kupon.Kupon{}, errors.New("format link salah")
		}
	}

	var kuponProcess kupon.Kupon

	kuponProcess.KodeKupon = newKupon.KodeKupon
	kuponProcess.LinkKupon = newKupon.LinkKupon
	kuponProcess.NamaProgram = newKupon.NamaProgram
	kuponProcess.GambarKupon = newKupon.GambarKupon

	result, err := ks.m.InsertKupon(userID, kuponProcess)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return kupon.Kupon{}, errors.New("barang sudah pernah diinputkan")
		}
		return kupon.Kupon{}, errors.New("terjadi kesalahan pada server")
	}

	return result, nil
}

func (ks *KuponServices) LihatKuponUser(token *golangjwt.Token) ([]kupon.Kupon, error) {

	userID, err := jwt.ExtractToken(token)
	if err != nil {
		return nil, errors.New("silakan login terlebih dahulu")
	}

	result, err := ks.m.GetKuponUser(userID)
	if err != nil {
		return nil, errors.New("anda tidak memiliki kupon")
	}
	return result, nil
}

func (ks *KuponServices) LihatSemuaKupon() ([]kupon.Kupon, error) {
	result, err := ks.m.GetAllKupons()
	if err != nil {
		return nil, errors.New("tidak ada kupon")
	}
	return result, nil
}
