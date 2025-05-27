package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//inicia o router utilizando as configurações default do gin
	router := gin.Default()

	// Definindo a Rota
	router.GET("/ping", func(c *gin.Context) {
		// Retorna um JSON com a mensagem "pong"
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.Run(":8080")
}
