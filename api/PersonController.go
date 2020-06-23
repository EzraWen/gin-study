package api

import (
	. "com.ezra/gin-test/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "it works")
}

func AddPersonApi(c *gin.Context) {
	var person Person

	if err := c.BindJSON(&person); err != nil {
		c.String(http.StatusInternalServerError, "参数绑定异常")
	}
	id, err := person.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,})
}
