package api

import (
	"database/sql"
	"pitsinkr/internal/storage"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Start(db *sql.DB) {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/api/logs", storage.FetchLogs(db))

	router.Run("localhost:4004")
}
