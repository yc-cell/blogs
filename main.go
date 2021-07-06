package main

import (
	"github.com/gin-gonic/gin"
	"my/blogs/common"
)

func main() {
	db := common.InitDB()
	defer db.Clauses() //todo db.close
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run()) // listen and serve on 0.0.0.0:8080
}
