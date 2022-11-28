package models

import (
	"errors"
	"time"
)

type MessageStatus int64

const (
	UnReaded MessageStatus = 1 << iota
	MoveToTrash
	Deleted
	Readed
	Stared
	Banned // Filter by sensitive word or by administrator operate
)

type Message struct {
	ID        int32         `json:"id" binding:"required"`
	Content   string        `json:"content"`
	PostTime  *time.Time    `json:"post_time" binding:"required"`
	Sender    *User         `json:"sender"`
	Recipient *User         `json:"recipient"`
	Status    MessageStatus `json:"status" gorm:"type:tinyint"`
}

func (msg *Message) Send() error {
	return errors.New("not implemented yet")
}

func (msg *Message) Open(user *User) error {
	return errors.New("not implemented yet")
}

func (msg *Message) MarkReaded() error {
	msg.Status = msg.Status | Readed
	return errors.New("not implemented yet")
}

func (msg *Message) MarkUnReaded() error {
	msg.Status = msg.Status &^ Readed // AND NOT: This is a bit clear operator
	return errors.New("not implemented yet")
}

func (msg *Message) MoveToTrash() error {
	msg.Status = msg.Status | MoveToTrash
	return errors.New("not implemented yet")
}

func (msg *Message) MoveFromTrash() error {
	msg.Status = msg.Status &^ MoveToTrash
	return errors.New("not implemented yet")
}
