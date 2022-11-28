package services

import (
	"errors"

	"github.com/yangwawa0323/gin-gorm-jwt/models"
)

type messageService struct {
	*dbService
	Message *models.Message
}

func NewMessageService(msg *models.Message) *messageService {
	return &messageService{
		dbService: NewDBService(),
		Message:   msg,
	}
}

func (ms *messageService) Save() error {
	if ms.Message != nil {
		result := ms.DB.Save(ms.Message)
		return result.Error
	} else {
		return errors.New("message service has not set Message field up")
	}
}

func (ms *messageService) FindMessageByID(msgID int64) (*models.Message, error) {
	return nil, errors.New("not implemented yet")
}
