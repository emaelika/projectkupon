package kupon

import "mime/multipart"

type KuponRequest struct {
	NamaProgram string         `form:"namaprogram"`
	KodeKupon   string         `form:"kodekupon"`
	LinkKupon   string         `form:"linkkupon"`
	GambarKupon multipart.File `form:"image"`
	UserID      uint
}
type KuponResponse struct {
	NamaProgram string `json:"namaprogram"`
	KodeKupon   string `json:"kodekupon"`
	LinkKupon   string `json:"linkkupon"`
	GambarKupon string `json:"image"`
}

// type Kupon struct {
// 	ID          uint
// 	NamaProgram string
// 	KodeKupon   string
// 	LinkKupon   string
// 	GambarKupon any
// }
