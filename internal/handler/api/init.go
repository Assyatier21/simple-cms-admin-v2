package api

import (
	"github.com/assyatier21/simple-cms-admin-v2/internal/usecase"
)

type DeliveryHandler interface {
}

type handler struct {
	usecase usecase.UsecaseHandler
}

func NewHandler(usecase usecase.UsecaseHandler) DeliveryHandler {
	return &handler{
		usecase: usecase,
	}
}
