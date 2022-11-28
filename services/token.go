package services

// TODO : Done token need to save into database?

// import (
// 	"errors"
// 	"fmt"

// 	"github.com/yangwawa0323/gin-gorm-jwt/models"
// )

// type tokenService struct {
// 	*dbService
// 	User *models.User
// 	// Token *auth.JWTClaim
// }

// func NewTokenService() *tokenService {
// 	return &tokenService{
// 		dbService: NewDBService(),
// 	}
// }

// func (tksvc *tokenService) Save() error {
// 	if tksvc.IsValid() {
// 		// TODO: save the token string to user table
// 		return errors.New("not implemented yet")
// 	} else {
// 		return fmt.Errorf("token service user: %#v", tksvc.User)
// 	}
// }

// func (tksvc *tokenService) IsValid() bool {
// 	return tksvc.User != nil
// }

// func (tksvc *tokenService) FindByEmail() error {
// 	if tksvc.IsValid() {
// 		// TODO: get the token string from database query
// 		// gorm.DB.First is a scanner function
// 		result := tksvc.DB.Where("email = ?", tksvc.User.Email).First(tksvc.User)
// 		return result.Error
// 	}
// 	return errors.New("not implemented yet")
// }
