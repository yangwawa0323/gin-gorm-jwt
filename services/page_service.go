package services

import (
	"errors"

	"github.com/yangwawa0323/gin-gorm-jwt/models"
)

type pageService struct {
	*dbService
	Page *models.Page
}

func NewPageService(page *models.Page) *pageService {
	return &pageService{
		dbService: NewDBService(), // implicit filed initial
		Page:      page,
	}
}

func (pgsvc *pageService) New(page *models.Page) error {
	return errorDebug(pgsvc.DB.Create(page).Error)

}

func (pgsvc *pageService) Save() error {
	if pgsvc.Page != nil {
		return pgsvc.DB.Save(pgsvc.Page).Error
	} else {
		return errors.New("page service has not set Page field up")
	}

}
