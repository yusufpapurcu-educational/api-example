package router

import (
	"net/http"

	"gorm.io/gorm"
)

func New(db *gorm.DB) http.Handler {
	return &http.ServeMux{}
}
