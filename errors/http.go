package errors

// BadRequest is handled by setting the status code in the reply to StatusBadRequest.
type BadRequest struct{ error }

// NotFound is handled by setting the status code in the reply to StatusNotFound.
type NotFound struct{ error }

// NotAuthorized is handled by setting the status code in the reply to StatusUnauthorized
type NotAuthorized struct{ error }

// UnprocessableEntity is handled by setting the status code in the reply to StatusUnprocessableEntity
type UnprocessableEntity struct{ error }

// DuplicateItem is handled by setting the status code in the reply to
type DuplicateItem struct{ error }

// HTTPError model struct
type HTTPError struct {
	Error   error
	Message string
}
