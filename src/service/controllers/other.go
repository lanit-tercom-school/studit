package controllers

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Success string `json:"success"`
}