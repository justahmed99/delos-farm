package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/justahmed99/delos-farm/core/domain"
	"github.com/justahmed99/delos-farm/core/usecase"
)

type PondHandler struct {
	pondUseCase usecase.PondUseCases
}

func NewPondHandler(pondUseCase usecase.PondUseCases) *PondHandler {
	return &PondHandler{
		pondUseCase: pondUseCase,
	}
}

func (h *PondHandler) GetPondById(context *gin.Context) {
	id := context.Param("id")
	pond, err := h.pondUseCase.GetPondByID(context, id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pond not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, pond)
}

func (h *PondHandler) CreatePond(context *gin.Context) {

	var pond domain.Pond

	if err := context.ShouldBindJSON(&pond); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	created_pond, err := h.pondUseCase.CreatePond(context, pond)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	context.IndentedJSON(http.StatusOK, created_pond)
}

func (h *PondHandler) UpdatePond(context *gin.Context) {

	var pond domain.Pond

	if err := context.ShouldBindJSON(&pond); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	created_pond, err := h.pondUseCase.UpdatePond(context, pond)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	context.IndentedJSON(http.StatusOK, created_pond)
}

func (h *PondHandler) DeletePond(context *gin.Context) {

	id := context.Param("id")
	_, err := h.pondUseCase.GetPondByID(context, id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pond not found"})
		return
	}

	status, err := h.pondUseCase.DeletePond(context, id)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": status})
}
