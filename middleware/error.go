package middleware

import (
	"errors"
	"github.com/ShaoZeMing/going/common"
	"github.com/gin-gonic/gin"
)

func Recovered(c *gin.Context, recovered interface{}) {
	var err error
	switch x := recovered.(type) {
	case string:
		err = errors.New(x)
	case error:
		err = x
	default:
		err = errors.New("unknown panic")
	}

	res := common.Rep{
		Code:    400,
		Message: err.Error(),
	}
	common.Error(c, &res)

}
