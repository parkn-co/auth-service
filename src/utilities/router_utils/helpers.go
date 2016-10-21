package routerutils

import "net/http"

// Response formats a response to keep responses uniform
func Response(data interface{}) interface{} {
	return map[string]interface{}{"success": true, "data": data}
}

// ErrorResponse is a helper func to return a single error
// Ex: 401 {
//   "error": "Unauthorized"
// }
func ErrorResponse(err interface{}) interface{} {
	return map[string]interface{}{"error": err}
}

// FormErrorResponse is a helper func to return errors that will be returned when a form
// is submitted and has errors with specific fields
// Ex: SignUp {
//   "errors": {
//     "email": "Email is associated with another account",
//     "firstName": "First name is required",
//   }
// }
func FormErrorResponse(errs map[string]string) map[string]map[string]string {
	return map[string]map[string]string{"errors": errs}
}

// Unauthorized is a helper func to return a 401
func Unauthorized() (int, interface{}) {
	return http.StatusUnauthorized, ErrorResponse("Unauthorized")
}

// InternalError is a helper func to return an internal server error
func InternalError() (int, interface{}) {
	return http.StatusInternalServerError, ErrorResponse("Internal Server Error")
}

// NotFound is a helper func to return a not found error
func NotFound() (int, interface{}) {
	return http.StatusNotFound, ErrorResponse("Not found")
}
