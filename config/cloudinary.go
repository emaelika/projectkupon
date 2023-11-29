package config

import (
	"github.com/cloudinary/cloudinary-go/v2"
)

func SetupCloudinary() (*cloudinary.Cloudinary, error) {
	cldSecret := "k2XtB4rAn5Zm9KhT6je0sF9BUJo"
	cldName := "daxpcsncf"
	cldKey := "815495445246189"

	cld, err := cloudinary.NewFromParams(cldName, cldKey, cldSecret)
	if err != nil {
		return nil, err
	}

	return cld, nil
}
