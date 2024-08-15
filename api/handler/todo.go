package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/katsuokaisao/gin-play/api/request"
	"github.com/katsuokaisao/gin-play/usecase"
)

type TodoHandler struct {
	todoUseCase *usecase.TodoUseCase
}

func NewTodoHandler(todoUseCase *usecase.TodoUseCase) *TodoHandler {
	return &TodoHandler{
		todoUseCase: todoUseCase,
	}
}

func (h *TodoHandler) Create(c *gin.Context) {
	var req request.TodoCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	if err := h.todoUseCase.Create(req.ToDomain()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": http.StatusCreated})
}

func (h *TodoHandler) List(c *gin.Context) {
	var req request.TodoListRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}
	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	todoList, err := h.todoUseCase.List(req.ToDomain())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "data": todoList})
}

func (h *TodoHandler) Get(c *gin.Context) {
	var req request.TodoGetRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	todo, err := h.todoUseCase.Get(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": http.StatusOK, "data": todo})
}

func (h *TodoHandler) Update(c *gin.Context) {
	var req request.TodoUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	if err := req.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	if err := h.todoUseCase.Update(idInt, req.ToDomain()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": http.StatusNoContent})
}

func (h *TodoHandler) Delete(c *gin.Context) {
	var req request.TodoGetRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	if err := h.todoUseCase.Delete(req.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": http.StatusNoContent})
}
