// Package generated provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package generated

const (
	ApiKeyAuthScopes = "ApiKeyAuth.Scopes"
)

// CreateNewUserRequest defines model for CreateNewUserRequest.
type CreateNewUserRequest struct {
	Name string `json:"name"`
}

// CreateNewUserResponse defines model for CreateNewUserResponse.
type CreateNewUserResponse struct {
	Code int32 `json:"code"`

	// User model.
	Data    *User  `json:"data,omitempty"`
	Message string `json:"message"`
}

// ErrorBadRequest defines model for ErrorBadRequest.
type ErrorBadRequest struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
	Trace   *[]struct {
		Line *string `json:"line,omitempty"`
	} `json:"trace,omitempty"`
}

// ErrorInternalServer defines model for ErrorInternalServer.
type ErrorInternalServer struct {
	Code    int32                   `json:"code"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
	Trace   *[]struct {
		Line *string `json:"line,omitempty"`
	} `json:"trace,omitempty"`
}

// ErrorUnauthorized defines model for ErrorUnauthorized.
type ErrorUnauthorized struct {
	Code    int32                   `json:"code"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
	Trace   *[]struct {
		Line *string `json:"line,omitempty"`
	} `json:"trace,omitempty"`
}

// ErrorUnexpected defines model for ErrorUnexpected.
type ErrorUnexpected struct {
	Code    int32                   `json:"code"`
	Data    *map[string]interface{} `json:"data,omitempty"`
	Message string                  `json:"message"`
	Trace   *[]struct {
		Line *string `json:"line,omitempty"`
	} `json:"trace,omitempty"`
}

// GetUserByIdResponse defines model for GetUserByIdResponse.
type GetUserByIdResponse struct {
	Code int32 `json:"code"`

	// User model.
	Data    *User  `json:"data,omitempty"`
	Message string `json:"message"`
}

// GetUsersResponse defines model for GetUsersResponse.
type GetUsersResponse struct {
	Code    int32   `json:"code"`
	Data    *[]User `json:"data,omitempty"`
	Message string  `json:"message"`
}

// User model.
type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// BadRequestError defines model for BadRequestError.
type BadRequestError = ErrorBadRequest

// InternalServerError defines model for InternalServerError.
type InternalServerError = ErrorInternalServer

// UnauthorizedError defines model for UnauthorizedError.
type UnauthorizedError = ErrorUnauthorized

// UnexpectedError defines model for UnexpectedError.
type UnexpectedError = ErrorUnexpected

// GetUsersParams defines parameters for GetUsers.
type GetUsersParams struct {
	// Page
	Page int32 `form:"page" json:"page"`
}

// GetUsersJSONBody defines parameters for GetUsers.
type GetUsersJSONBody = CreateNewUserRequest

// GetUsersJSONRequestBody defines body for GetUsers for application/json ContentType.
type GetUsersJSONRequestBody = GetUsersJSONBody
