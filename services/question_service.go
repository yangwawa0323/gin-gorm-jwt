package services

import (
	"github.com/yangwawa0323/gin-gorm-jwt/models"
)

type questionService struct {
	*dbService
	Question *models.Question
}

func NewQuestionService(qst *models.Question) *questionService {
	return &questionService{
		dbService: NewDBService(),
		Question:  qst,
	}
}

func (qstsvc *questionService) Save() error {
	return qstsvc.DB.Save(qstsvc.Question).Error
}
