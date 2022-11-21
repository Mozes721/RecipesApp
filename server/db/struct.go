package db

import (
	"context"
	"github.com/gin-gonic/gin"
)

type Data struct {
	Recepies []Recepie
	DB       gin.Engine
}

type DbCallers struct {
	ctx context.Context
}

type Recepie struct {
	Made   bool
	Rating int
	Title  string
	Url    string
}

type HTTPError struct {
	Err      error
	Code     int
	FlashMsg string
	URL      string
}
