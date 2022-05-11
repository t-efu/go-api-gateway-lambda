package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/t-efu/go-api-gateway-lambda/usecase"
)

type MemoHandler interface {
	Find(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error)
}

type memoHandler struct {
	memoUsecase usecase.MemoUsecase
}

func NewMemoHandler(memoUsecase usecase.MemoUsecase) MemoHandler {
	return &memoHandler{
		memoUsecase: memoUsecase,
	}
}

func (h *memoHandler) Find(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	memos, err := h.memoUsecase.Find(ctx)
	if err != nil {
		res := &events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}
		return res, nil
	}
	bodyBytes, _ := json.Marshal(memos)
	res := &events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(bodyBytes),
	}
	return res, nil
}
