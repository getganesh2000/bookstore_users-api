package users

import (
	"github.com/getganesh2000/bookstore_users-api/domain/users"
	"github.com/getganesh2000/bookstore_users-api/services"
	"github.com/getganesh2000/bookstore_users-api/utils/erros"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	var user users.User


	if err:= c.ShouldBindJSON(&user);err!=nil{
		restErr := erros.NewBadRequest("Bad Request")
		c.JSON(restErr.Status,restErr)
		return
	}

	result, saveErr:= services.CreateUser(user)

	if saveErr!=nil {
		c.JSON(saveErr.Status,saveErr)
		return
	}

	c.JSON(http.StatusCreated,result)
}

func GetUser(c *gin.Context) {
	userId,userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr!=nil {
		err:= erros.NewBadRequest("Invalid id")
		c.JSON(err.Status,err)
		return
	}

	user, getErr:= services.FindUser(userId)

	if getErr!=nil {
		c.JSON(getErr.Status,getErr)
		return
	}

	c.JSON(http.StatusOK,user)
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Implement me")
}
