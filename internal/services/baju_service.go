package services

import (
	"log"
	"sagara-msib-test/internal/entities"
	"sagara-msib-test/internal/repositories"
)

type BajuServices interface {
	// POST
	CreateBaju(baju entities.Baju) error

	// GET
	GetBajuByID(id int) (entities.Baju, error)
	GetAllBaju() ([]entities.Baju, error)
	GetBajuOrderByEmptyStok() ([]entities.Baju, error)
	GetBajuOrderByStok(stok int, kondisi string) (bajuList []entities.Baju, err error)

	// PUT
	UpdateBaju(baju entities.Baju) error

	// DELETE
	DeleteBaju(id int) error
}

type bajuServices struct {
	bajuRepo repositories.BajuRepository
}

func NewInventoryBajuService(r repositories.BajuRepository) (ibs BajuServices) {
	ibs = &bajuServices{
		bajuRepo: r,
	}

	return ibs
}

func (ibs *bajuServices) CreateBaju(baju entities.Baju) (err error) {
	log.Printf("[LOG][Service] Nama Baju Request : %v\n", baju.Nama)
	err = ibs.bajuRepo.Create(baju)

	if err != nil {
		return err
	}

	return err
}

func (s *bajuServices) GetBajuByID(id int) (baju entities.Baju, err error) {
	return s.bajuRepo.GetByID(id)
}

func (s *bajuServices) GetAllBaju() (bajuList []entities.Baju, err error) {
	return s.bajuRepo.GetAll()
}

func (s *bajuServices) GetBajuOrderByEmptyStok() (bajuList []entities.Baju, err error) {
	return s.bajuRepo.GetBajuOrderByEmptyStok()
}

func (s *bajuServices) GetBajuOrderByStok(stok int, kondisi string) (bajuList []entities.Baju, err error) {
	return s.bajuRepo.GetBajuOrderByStok(stok, kondisi)
}

func (s *bajuServices) UpdateBaju(baju entities.Baju) (err error) {
	return s.bajuRepo.Update(baju)
}

func (s *bajuServices) DeleteBaju(id int) (err error) {
	return s.bajuRepo.Delete(id)
}
