package repository

import (
	kupon "projectkupon/features/kupons"

	"gorm.io/gorm"
)

type KuponModel struct {
	gorm.Model
	NamaProgram string
	KodeKupon   string
	LinkKupon   string
	GambarKupon string
	UserID      uint
}

type kuponQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) kupon.Repository {
	return &kuponQuery{
		db: db,
	}
}

func (kq *kuponQuery) InsertKupon(userID uint, newKupon kupon.Kupon) (kupon.Kupon, error) {
	var inputData = new(KuponModel)
	inputData.UserID = userID
	inputData.NamaProgram = newKupon.NamaProgram
	inputData.KodeKupon = newKupon.KodeKupon
	inputData.LinkKupon = newKupon.LinkKupon
	inputData.GambarKupon = newKupon.GambarKupon

	if err := kq.db.Create(&inputData).Error; err != nil {
		return kupon.Kupon{}, err
	}

	newKupon.ID = inputData.ID

	return newKupon, nil
}

func (kq *kuponQuery) GetKuponUser(userID uint) ([]kupon.Kupon, error) {
	var KuponModels []KuponModel
	if err := kq.db.Where("user_id = ?", userID).Find(&KuponModels); err.Error != nil {
		return nil, err.Error
	}
	var results []kupon.Kupon

	for _, kupondb := range KuponModels {
		// Create a new kupon.Kupon struct
		newKupon := kupon.Kupon{
			NamaProgram: kupondb.NamaProgram,
			GambarKupon: kupondb.GambarKupon,
			LinkKupon:   kupondb.LinkKupon,
			ID:          kupondb.ID,
			KodeKupon:   kupondb.KodeKupon,
		}

		// Append the new kupon.Kupon struct to the results slice
		results = append(results, newKupon)
	}

	return results, nil
}

func (kq *kuponQuery) GetAllKupons() ([]kupon.Kupon, error) {
	var KuponModels []KuponModel
	if err := kq.db.Find(&KuponModels); err.Error != nil {

		return nil, err.Error
	}
	var results []kupon.Kupon

	for _, kupondb := range KuponModels {
		// Create a new kupon.Kupon struct
		newKupon := kupon.Kupon{
			NamaProgram: kupondb.NamaProgram,
			GambarKupon: kupondb.GambarKupon,
			LinkKupon:   kupondb.LinkKupon,
			ID:          kupondb.ID,
			KodeKupon:   kupondb.KodeKupon,
		}

		// Append the new kupon.Kupon struct to the results slice
		results = append(results, newKupon)
	}

	return results, nil
}
