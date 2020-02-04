package controllers

import "WeatherAPI/pkg/services"

type ApiError struct {
	Status  int
	Message string
}

func (e *ApiError) Error() string {
	return e.Message
}

func ParseError(e error) ApiError {
	switch e {
	case services.ErrorNoLocations:
		return ApiError{200, e.Error()}
	case services.ErrorLocationNotFound:
		return ApiError{404, e.Error()}
	case services.ErrorLocationAlreadyExists, services.ErrorCreateLocation:
		return ApiError{400, e.Error()}
	default:
		return ApiError{500, e.Error()}
	}
}
