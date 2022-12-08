package services

import (
	"time"

	"github.com/yangwawa0323/gin-gorm-jwt/models/audit"
)

type auditService struct {
	*dbService
	Audit *audit.AuditLog
}

func NewAuditService() *auditService {
	return &auditService{
		dbService: NewDBService(),
	}
}

func (adtsvc *auditService) New(ctn string) error {
	adtsvc.Audit = &audit.AuditLog{
		Content:   ctn,
		Timestamp: time.Now(),
	}
	return adtsvc.DB.Create(adtsvc.Audit).Error
}

// func (adtsvc *auditService) Save() error {
// 	return adtsvc.DB.Save(adtsvc.Audit).Error
// }

func AuditSave(ctn string) error {
	adtsvc := NewAuditService()
	return adtsvc.New(ctn)
}
