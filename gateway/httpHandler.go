package main

import "net/http"

type httpHandler struct {
	//all gateway routes
}

func NewHandler() *httpHandler {
	return &httpHandler{}
}

func (handler *httpHandler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	//logic to handle order
}

func (handler *httpHandler) registerRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /api/customer/{customerId}/orders", handler.handleCreateOrder)
}
