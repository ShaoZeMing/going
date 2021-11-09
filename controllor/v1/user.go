package v1

import (
	"github.com/ShaoZeMing/going/common"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	common.Success(c, &common.Rep{Code: 200, Message: "cad9siad", Data: nil})

}

func UserLogin() {

}

func UserRegister() {

}

func UserResetPwd() {

}

func UserGetSmsCode() {

}
