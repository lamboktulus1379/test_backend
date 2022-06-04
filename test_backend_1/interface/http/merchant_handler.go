package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"test_backend_1/domain/model"
	"test_backend_1/usecase"

	"github.com/gin-gonic/gin"
)

type IMerchantHandler interface {
	ReportDaily(c *gin.Context)
	ReportDailyOutlet(c *gin.Context)
}

type MerchantHandler struct {
	merchantUsecase usecase.IMerchantUsecase
}

func NewMerchantHandler(merchantUsecase usecase.IMerchantUsecase) IMerchantHandler {
	return &MerchantHandler{merchantUsecase: merchantUsecase}
}

func (merchantHandler *MerchantHandler) ReportDaily(c *gin.Context) {
	var reqQueryParam model.ReqQueryParamMerchant

	if err := c.ShouldBindQuery(&reqQueryParam); err != nil {
		log.Printf("An error occurred: %v", err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("An error occurred: %v", err.Error()))
	}

	reqQueryParam.ID, _ = strconv.Atoi(c.GetString("user_id"))
	fmt.Println("User ID", c.GetString("user_id"))
	res := merchantHandler.merchantUsecase.ReportDaily(c.Request.Context(), reqQueryParam)

	c.JSON(http.StatusOK, res)
}

func (merchantHandler *MerchantHandler) ReportDailyOutlet(c *gin.Context) {
	var reqQueryParam model.ReqQueryParamMerchant

	if err := c.ShouldBindQuery(&reqQueryParam); err != nil {
		log.Printf("An error occurred: %v", err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("An error occurred: %v", err.Error()))
	}

	reqQueryParam.ID, _ = strconv.Atoi(c.GetString("user_id"))
	fmt.Println("User ID", c.GetString("user_id"))
	res := merchantHandler.merchantUsecase.ReportDailyOutlet(c.Request.Context(), reqQueryParam)

	c.JSON(http.StatusOK, res)
}
