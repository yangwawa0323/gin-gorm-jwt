package services

import (
	"time"

	"github.com/yangwawa0323/gin-gorm-jwt/models"
)

type auditService struct {
	*dbService
	Audit *models.AuditLog
}

func NewAuditService() *auditService {
	return &auditService{
		dbService: NewDBService(),
	}
}

func (adtsvc *auditService) Save(ctn string) error {
	adtsvc.Audit = &models.AuditLog{
		Content:   ctn,
		Timestamp: time.Now(),
	}
	return adtsvc.DB.Create(adtsvc.Audit).Error
}

func AuditSave(ctn string) error {
	adtsvc := NewAuditService()
	return adtsvc.Save(ctn)
}
