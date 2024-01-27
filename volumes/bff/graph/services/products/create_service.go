package products

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kuroweb/price-monitoring/volumes/bff/config"
	"github.com/kuroweb/price-monitoring/volumes/bff/graph/model"
)

type ICreateService interface {
	Create(ctx context.Context, input model.CreateProductInput) (model.CreateProductResult, error)
}

type CreateService struct{}

func (c *CreateService) Create(ctx context.Context, input model.CreateProductInput) (model.CreateProductResult, error) {
	cfg := config.NewConfig()
	url := fmt.Sprintf("%s/api/v1/products", cfg.BackendUrl)

	// NewProduct 構造体をJSONに変換
	requestBody, err := json.Marshal(input)
	if err != nil {
		return c.handleServerError(), nil
	}

	// POSTリクエストの作成
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return c.handleServerError(), nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.handleServerError(), nil
	}

	var response struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&response); err != nil {
		return c.handleServerError(), nil
	}

	result := model.CreateProductResultSuccess{
		Ok: true,
		Product: &model.Product{
			ID:   strconv.Itoa(response.ID),
			Name: response.Name,
		},
	}

	return result, nil
}

func (c *CreateService) handleServerError() model.CreateProductResultError {
	return model.CreateProductResultError{
		Ok: false,
		Error: model.CreateProductResultValidationFailed{
			Code:    "503",
			Message: "Service is currently unavailable.",
			Details: []*model.ErrorDetail{},
		},
	}
}
