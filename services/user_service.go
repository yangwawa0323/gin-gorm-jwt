package services

import (
	"errors"
	"fmt"

	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
	"gorm.io/gorm"
)

var errorDebug = utils.ErrorDebug

type userService struct {
	*dbService
	User *models.User
}

func NewUserService(user *models.User) *userService {
	return &userService{
		dbService: NewDBService(),
		User:      user,
	}
}

func (us *userService) SendMessage(rcpID int64, msg *models.Message) error {
	rcp, err := us.FindUserByID(rcpID)
	if err != nil {
		return fmt.Errorf("the recipient with ID: %d is not exists", rcpID)
	}

	msg.Sender = us.User
	msg.Recipient = rcp

	msgsvc := NewMessageService(msg)
	return errorDebug(msgsvc.Save())
}

// Not valid the gorm SQL.
func (us *userService) ReadMessage(msgID int64) (*models.Message, error) {
	var msg *models.Message
	msgsvc := NewMessageService(msg)
	result := msgsvc.DB.Model(msg).
		Where("ID = ?", msgID).
		Update("status", gorm.Expr("status | ?", models.Readed)).
		First(msg)
	return msg, errorDebug(result.Error)
}

/**
* user answer / post question
 */
func (us *userService) PostQuestion(qst *models.Question) error {
	qst.Poster = us.User
	qstsvc := NewQuestionService(qst)
	return errorDebug(qstsvc.Save())
}

func (us *userService) AnswerQuestion(qstID int64, content string) error {
	var qst *models.Question
	qstsvc := NewQuestionService(qst)
	err := errorDebug(qstsvc.DB.First(qst, qstID).Error)
	if err != nil {
		return err
	}

	// TODO: answer data
	return errors.New("not implemented yet")
}

/**
* user privilege functions
 */
func (us *userService) Grant(userID int64, priv models.Privilege) error {

	if utils.RequireAudit() && !us.HasPrivilege(models.Grant) {
		var warning = fmt.Sprintf("User [%d:%s] has not [%s] prvilege to grant [%s] to user id [%d]",
			us.User.ID, us.User.Username, models.LiteralPrivilege[int64(models.Grant)],
			models.LiteralPrivilege[int64(priv)], userID)
		AuditSave(warning)
		return errors.New(warning)
	}

	grantTo, err := us.FindUserByID(userID)
	us.User = grantTo
	if errorDebug(err) != nil {
		return err
	}
	// use bit OR grant the new privilege
	us.User.Privilege = us.User.Privilege | priv
	return errorDebug(us.Save()) // change other user info
}

// Not finished.
func (us *userService) ChangePassword(pwd string) error {

	err := us.User.HashPassword(pwd)
	if err := errorDebug(err); err != nil {
		return err
	}

	err = us.SendNotificationMail()
	if err := errorDebug(err); err != nil {
		return err
	}

	if utils.RequireAudit() {
		AuditSave(
			fmt.Sprintf("User %s changed password [**%s**]",
				us.User.Username, pwd[2:len(pwd)-2]), // Save partial password string
		)
	}
	err = us.Save() // save password themself.
	return errorDebug(err)
}

func (us *userService) SendNotificationMail() error {
	return errors.New("not implemented yet")
}

// Done.
func (us *userService) HasPrivilege(priv models.Privilege) bool {
	return us.User.Privilege&priv == priv
}

// Admin change the specified uid User's UserClass attribute
func (us *userService) ChangeUserClass(userID int64, userCls models.UserClass) error {
	if !us.HasPrivilege(models.Admin) {
		var warning = fmt.Sprintf("User [%d:%s] change user id [%d] to %s has not admin privilege",
			us.User.ID, us.User.Username, userID, models.LiteralUserClass[int64(userCls)])
		AuditSave(warning)
		return errors.New(warning)
	}
	return nil
}

/**
* Base Database Operation
 */
// Done.
func (us *userService) FindUserByID(userID int64) (*models.User, error) {
	var user *models.User
	result := us.DB.First(user, userID)
	return user, errorDebug(result.Error)
}

// Done

func (us *userService) New(user *models.User) error {
	return errorDebug(us.DB.Create(user).Error)
}

func (us *userService) Save() error {
	return errorDebug(us.DB.Save(us.User).Error)
}
