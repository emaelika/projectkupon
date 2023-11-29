package kupon

import (
	"context"
	"net/http"

	kupon "projectkupon/features/kupons"
	"projectkupon/utils/cloudnr"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2"
	gojwt "github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type kuponHandler struct {
	s  kupon.Service
	cl *cloudinary.Cloudinary
	ct context.Context
	up string
}

func New(s kupon.Service, cld *cloudinary.Cloudinary, ctx context.Context, uploadparam string) kupon.Handler {
	return &kuponHandler{
		s:  s,
		cl: cld,
		ct: ctx,
		up: uploadparam,
	}
}

func (kh *kuponHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input = new(KuponRequest)
		if err := c.Bind(input); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "input yang diberikan tidak sesuai",
			})
		}
		formHeader, err := c.FormFile("image")
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError, map[string]any{
					"message": "formheader error",
				})

		}
		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(
				http.StatusInternalServerError, map[string]any{
					"message": "formfile error",
				})
		}

		link, err := cloudnr.UploadImage(kh.cl, kh.ct, formFile, kh.up)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"message": "harap pilih gambar",
					"data":    nil,
				})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"message": "kesalahan pada server",
					"data":    nil,
				})
			}
		}

		var inputProcess = new(kupon.Kupon)
		inputProcess.NamaProgram = input.NamaProgram
		inputProcess.LinkKupon = input.LinkKupon
		inputProcess.KodeKupon = input.KodeKupon
		inputProcess.GambarKupon = link

		result, err := kh.s.TambahKupon(c.Get("user").(*gojwt.Token), *inputProcess)

		if err != nil {
			c.Logger().Error("ERROR Register, explain:", err.Error())
			var statusCode = http.StatusInternalServerError
			var message = "terjadi permasalahan ketika memproses data"

			if strings.Contains(err.Error(), "terdaftar") {
				statusCode = http.StatusBadRequest
				message = "data yang diinputkan sudah terdaftar ada sistem"
			}

			return c.JSON(statusCode, map[string]any{
				"message": message,
			})
		}

		var response = new(KuponResponse)
		response.NamaProgram = result.NamaProgram
		response.KodeKupon = result.KodeKupon
		response.LinkKupon = result.LinkKupon
		response.GambarKupon = result.GambarKupon

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    response,
		})
	}
}

func (kh *kuponHandler) GetOne() echo.HandlerFunc {
	return func(c echo.Context) error {
		paramid := c.Param("id")
		if paramid == "" {
			return c.JSON(http.StatusNotFound, map[string]any{
				"message": "param 0",
				"data":    nil,
			})
		}
		result, err := kh.s.LihatKuponUser(c.Get("user").(*gojwt.Token))
		if err != nil {
			if strings.Contains(err.Error(), "login") {
				return c.JSON(http.StatusUnauthorized, map[string]any{
					"message": "harap login terelebih dahulu",
					"data":    nil,
				})
			} else {
				return c.JSON(http.StatusNotFound, map[string]any{
					"message": "user tidak memiliki kupon",
					"data":    nil,
				})
			}
		}

		var results []KuponResponse

		for _, kupondb := range result {
			// Create a new kupon.Kupon struct
			newKupon := KuponResponse{
				NamaProgram: kupondb.NamaProgram,
				GambarKupon: kupondb.GambarKupon,
				LinkKupon:   kupondb.LinkKupon,

				KodeKupon: kupondb.KodeKupon,
			}

			// Append the new kupon.Kupon struct to the results slice
			results = append(results, newKupon)
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    results,
		})
	}
}
func (kh *kuponHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := kh.s.LihatSemuaKupon()
		if err != nil {
			if strings.Contains(err.Error(), "tidak") {
				return c.JSON(http.StatusNotFound, map[string]any{
					"message": "tidak ada kupon tersedia",
					"data":    nil,
				})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"message": "ada masalah di server",
					"data":    nil,
				})
			}

		}
		var results []KuponResponse

		for _, kupondb := range result {
			// Create a new kupon.Kupon struct
			newKupon := KuponResponse{
				NamaProgram: kupondb.NamaProgram,
				GambarKupon: kupondb.GambarKupon,
				LinkKupon:   kupondb.LinkKupon,

				KodeKupon: kupondb.KodeKupon,
			}

			// Append the new kupon.Kupon struct to the results slice
			results = append(results, newKupon)
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success create data",
			"data":    results,
		})
	}
}
