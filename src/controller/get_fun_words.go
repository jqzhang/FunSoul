package controller

import (
	"GoFun/src/db"
	"GoFun/src/model"
	"github.com/gin-gonic/gin"
)

func GetFunWords(c *gin.Context) {
	var funWords model.FunWords
	err := db.Instance.QueryRow("select * from fun_words order by rand( ) limit 1").Scan(&funWords.ID, &funWords.Content, &funWords.Hits)
	if err != nil {
		panic(err)
	}

	c.String(200, funWords.Content)
	//c.HTML(200, "http", "Hello")
}
