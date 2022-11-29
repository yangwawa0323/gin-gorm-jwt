package models

import (
	"strings"

	"gorm.io/gorm"
)

type courseCategory int64

const (
	_ courseCategory = iota
	Development
	OperationAndMaintenance
	FullStack
)

var LiteralCourseCategory = map[int]string{
	int(Development):             "development",
	int(OperationAndMaintenance): "operation and maintenance",
	int(FullStack):               "full stack engineer",
}

type Tag struct {
	gorm.Model
	Keyword string `json:"keyword"`
}

type CourseCategory struct {
	gorm.Model
	Name string `json:"name"`
}

/**
*	**Course** belong to **CourseCategory**
*	`CategoryID` is the association foreign key
*
*	**Course** has many **Tag**, a **Tag** can mark many **Course**
**  **Course** and **Tag** has Many To Many association.
*   `gorm:"many2many:course_tags"` is the principle created
*   intermediated table name.
*
 */

type Course struct {
	gorm.Model
	Name        string         `json:"class_name"`
	Category    CourseCategory `json:"category" gorm:"foreignKey:CategoryID"`
	CategoryID  uint           `json:"category_id"` // `gorm` has one asscocitaion;
	Description string         `json:"description"`
	Tags        []Tag          `json:"tags" gorm:"many2many:course_tags"`
	TagString   string         `json:"tag_string"`
	// The Tag should append to one string by using `gorm` **BeforeSave** hook
}

/**
*	Course
*	  ^
*     |  gorm belong to
*   Class
 */

type Class struct {
	gorm.Model
	CourseID     int    `json:"course_id"` // gorm belongs to mapping principle
	Title        string `json:"title"`
	DisplayOrder int8   `json:"display_order"`
	Content      string `json:"content"`
}

func (c *Course) BeforeSave(tx *gorm.DB) (err error) {
	var allTags []string = make([]string, 0)
	for _, tag := range c.Tags {
		allTags = append(allTags, tag.Keyword)
	}
	c.TagString = strings.Join(allTags, ";")
	return
}
