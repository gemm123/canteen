package controller

import "github.com/gemm123/canteen/service"

type controller struct {
	service service.Service
}

func NewController(service service.Service) *controller {
	return &controller{service: service}
}
