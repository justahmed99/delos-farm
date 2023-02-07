package handler

import (
	"net/http"
	"strconv"

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
	id, _ := strconv.ParseInt(context.Param("id"), 10, 64)
	pond, err := h.pondUseCase.GetPondByID(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, pond)
}

func (h *PondHandler) GetPonds(context *gin.Context) {
	pond, err := h.pondUseCase.GetPonds()
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, pond)
}

func (h *PondHandler) GetPondsByFarmID(context *gin.Context) {
	farm_id, _ := strconv.ParseInt(context.Param("farm_id"), 10, 64)
	pond, err := h.pondUseCase.GetPondsByFarmID(farm_id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, pond)
}

func (h *PondHandler) CreatePond(context *gin.Context) {

	var pond domain.Pond

	if err := context.ShouldBindJSON(&pond); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	created_pond, err := h.pondUseCase.CreatePond(pond)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, created_pond)
}

func (h *PondHandler) UpdatePond(context *gin.Context) {

	var pond domain.Pond

	if err := context.ShouldBindJSON(&pond); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	created_pond, err := h.pondUseCase.UpdatePond(pond)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, created_pond)
}

func (h *PondHandler) DeletePond(context *gin.Context) {

	id, err_input := strconv.ParseInt(context.Param("id"), 10, 64)
	if err_input != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "input must be integer"})
		return
	}

	delete_err := h.pondUseCase.DeletePond(id)
	if delete_err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": delete_err.Error()})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Pond data successufully deleted"})
}
