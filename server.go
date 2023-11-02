package main

import "github.com/gin-gonic/gin"

func StartServer() error {

	router := gin.New()
	router.Use()

	router.GET("/ping", Ping)
	router.POST("/evaluate", EvaluateBoard)

	return router.Run(":8001")

}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{})
}

func EvaluateBoard(c *gin.Context) {
	var board Board
	c.BindJSON(&board)
	eval := board.EvaluateBoard()
	c.JSON(200, eval)
}
