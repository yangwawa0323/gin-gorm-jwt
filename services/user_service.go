package services

import (
	"errors"
	"fmt"

	"github.com/yangwawa0323/gin-gorm-jwt/models"
	"github.com/yangwawa0323/gin-gorm-jwt/utils"
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

	// return errors.New("not implemented yet")
}

func (us *userService) ReadMessage(msgID int64) error {
	var msg *models.Message
	var err error
	msgsvc := NewMessageService(msg)
	if msg, err = msgsvc.FindMessageByID(msgID); err != nil {
		return err
	}

	msg.Recipient = us.User
	msg.MarkReaded()
	msgsvc.Message = msg
	return errorDebug(msgsvc.Save())

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
func (us *userService) Grant(priv models.Privilege) error {

	us.User.Privilege = us.User.Privilege | priv
	if utils.RequireAudit() {
		// TODO
	}
	return errorDebug(us.Save())
}

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
	err = us.Save()
	return errorDebug(err)
}

func (us *userService) SendNotificationMail() error {
	return errors.New("not implemented yet")
}

func (us *userService) HasPrivilege(priv models.Privilege) bool {
	return us.User.Privilege&priv == priv
}

// Admin change the specified uid User's UserClass attribute
func (us *userService) ChangeUserClass(userID int64, userCls models.UserClass) error {
	if !us.HasPrivilege(models.Admin) {
		var warning = fmt.Sprintf("User [%d:%s] change user id [%d] to %s has not admin privilege",
			us.User.ID, us.User.Username, userID)
		AuditSave(warning)
		return errors.New(warning)
	}
	return nil
}

/**
* Base Database Operation
 */
func (us *userService) FindUserByID(userID int64) (*models.User, error) {
	var user *models.User
	result := us.DB.First(user, userID)
	return user, errorDebug(result.Error)
}

func (us *userService) Save() error {
	if us.User != nil {
		return errorDebug(us.DB.Save(us.User).Error)
	} else {
		return errors.New("user service has not set User field up")
	}
}
