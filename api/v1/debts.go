package v1

import (
	"net/http"
	"strconv"

	"github.com/SaidovZohid/market-project/api/models"
	"github.com/SaidovZohid/market-project/storage/repo"
	"github.com/gin-gonic/gin"
)

// @Router  [post]
// @Summary Create a debt
// @Description Create a debt
// @Tags debt
// @Accept json
// @Produce json
// @Param debt body models.CreateDebt true "Debt"
// @Success 200 {object} models.GetAfterCreate
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) CreateDebt(ctx *gin.Context) {
	var (
		req models.CreateDebt
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	res, err := h.Storage.Debt().Create(&repo.Debt{
		FirstName:             req.FirstName,
		LastName:              req.LastName,
		PhoneNumber:           req.PhoneNumber,
		AdditionalPhoneNumber: req.AdditionalPhoneNumber,
		AddressWork:           req.AddressWork,
		SellerFullName:        req.SellerFullName,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GetAfterCreate{
		ID:        res.ID,
		CreatedAt: res.CreatedAt,
	})
}

// @Router /{id} [get]
// @Summary get a debt
// @Description get a debt
// @Tags debt
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 200 {object} models.GetDebt
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) GetDebt(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	res, err := h.Storage.Debt().Get(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.GetDebt{
		ID:                    res.ID,
		FirstName:             res.FirstName,
		LastName:              res.LastName,
		PhoneNumber:           res.PhoneNumber,
		AdditionalPhoneNumber: res.AdditionalPhoneNumber,
		AddressWork:           res.AddressWork,
		SellerFullName:        res.SellerFullName,
		CreatedAt:             res.CreatedAt,
		UpdatedAt:             *res.UpdatedAt,
		DeletedAt:             *res.DeletedAt,
	})
}

// @Router /{id} [put]
// @Summary update a debt
// @Description update a debt
// @Tags debt
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Param debt body models.CreateDebt true "Debt"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) UpdateDebt(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	var (
		req models.CreateDebt
	)
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	err = h.Storage.Debt().Update(&repo.Debt{
		ID:                    id,
		FirstName:             req.FirstName,
		LastName:              req.LastName,
		PhoneNumber:           req.PhoneNumber,
		AdditionalPhoneNumber: req.AdditionalPhoneNumber,
		AddressWork:           req.AddressWork,
		SellerFullName:        req.SellerFullName,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseOK{
		Message: "Succesfully updated",
	})
}

// @Router /{id} [delete]
// @Summary get a debt
// @Description delete a debt
// @Tags debt
// @Accept json
// @Produce json
// @Param id path int true "ID"
// @Success 201 {object} models.ResponseOK
// @Failure 500 {object} models.ResponseError
func (h *handlerV1) DeleteDebt(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	err = h.Storage.Debt().Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ResponseError{
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.ResponseOK{
		Message: "Succesfully deleted",
	})
}

func (h *handlerV1) GetAllDebts(ctx *gin.Context) {

}
