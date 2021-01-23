package ascii_art_err

import "net/http"

type Error struct {
	Status int
	StatusText string
}

var OK = Error{
	Status: http.StatusOK,
	StatusText: http.StatusText(http.StatusOK),
}

var Created = Error{
	Status: http.StatusCreated,
	StatusText: http.StatusText(http.StatusCreated),
}

var BadRequest = Error{
	Status: http.StatusBadRequest,
	StatusText: http.StatusText(http.StatusBadRequest),
}

var NotFound = Error{
	Status: http.StatusNotFound,
	StatusText: http.StatusText(http.StatusNotFound),
}

var InternalServerError = Error{
	Status: http.StatusInternalServerError,
	StatusText: http.StatusText(http.StatusInternalServerError),
}
