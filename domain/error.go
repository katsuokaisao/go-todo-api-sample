package domain

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrTodoTitleRequired     = errors.New("todo title is required")
	ErrTodoTitleEmpty        = errors.New("todo title is empty")
	ErrTodoStatusInvalid     = errors.New("todo status is invalid")
	ErrTodoPriorityInvalid   = errors.New("todo priority is invalid")
	ErrTodoBeginAtAfterEndAt = errors.New("todo begin_at is after end_at")
	ErrTodoExpireAtBeforeNow = errors.New("todo expire_at is before now")
	ErrRecordNotFound        = errors.New("record not found")
	ErrConflict              = errors.New("conflict")
	ErrInvalidToken          = errors.New("invalid token")
)

func ToGinResponse(c *gin.Context, err error) {
	if errors.Is(err, ErrConflict) {
		c.JSON(http.StatusConflict, gin.H{"message": http.StatusText(http.StatusConflict)})
		return
	} else if errors.Is(err, ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"message": http.StatusText(http.StatusNotFound)})
		return
	} else if errors.Is(err, ErrInvalidToken) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": http.StatusText(http.StatusUnauthorized)})
		return
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)})
		return
	}
}
