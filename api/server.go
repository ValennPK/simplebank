package api

import (
	db "github.com/ValennPK/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// server serves HTTP requests for our banking service
type Server struct {
	store  *db.Store
	router *gin.Engine
}
