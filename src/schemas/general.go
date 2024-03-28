package schemas

// swagger:response healthcheckResponse
type HealthcheckResponse struct {
	Message string `json:"message" example:"Logvista API Server is running!"`
}

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
