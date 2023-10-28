package schemas

// swagger:response responseMessage
type ResponseMessage struct {
	Message string `json:"message"`
}

// swagger:response successResponse
type SuccessResponse struct {
	Message string `json:"message"`
}

// swagger:response errorResponse
type ErrorResponse struct {
	Message string `json:"message"`
}
