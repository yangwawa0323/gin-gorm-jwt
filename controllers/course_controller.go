package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewCourse(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"result": "Not implemented yet",
	})
}

func CourseList(ctx *gin.Context) {
	MockJsonResponse(ctx, "mock/list-course.json")
}

type course struct {
	CourseID int32 `uri:"course_id" binding:"required"`
}

// 2023-01-06
// uri /course/detail/:course_id
func CourseDetail(ctx *gin.Context) {
	var course course
	if err := ctx.ShouldBindUri(&course); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	MockJsonResponse(ctx, "mock/course-detail.json")
}

type userFavroite struct {
	CourseID int32 `uri:"course_id" binding:"required"`
	UserID   int32 `uri:"user_id" binding:"required"`
}

// uri /course/favorite/:user_id/:course_id
func ToggleUserFavoriteCourse(ctx *gin.Context) {
	var userFavroite userFavroite
	if err := ctx.ShouldBindUri(&userFavroite); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": "toggle user favorite course",
	})
}
