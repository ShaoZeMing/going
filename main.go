package main

import (
	"github.com/ShaoZeMing/going/route"
)

func main() {

	r := route.Init()
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
