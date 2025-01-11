package lib

import (
	"net/http"
	"time"
)

type ResponseSchema struct {
	Status     int         `json:"code"`
	StatusText string      `json:"status_text"`
	D          interface{} `json:"d"`
	Time       time.Time `json:"time"`
}

func NewResponse(status int, d interface{}) ResponseSchema {
	return ResponseSchema{
		Status: status,
		StatusText: http.StatusText(status),
		D: d,
		Time: time.Now(),
	}
}

func NewErrorResponse(status int, err error) ResponseSchema {
	d := struct{
		Err string `json:"err"`
	}{
		Err: err.Error(),
	}
	return NewResponse(status, d)
}