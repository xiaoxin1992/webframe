package main

//func Registry(r gin.IRouter) {
//	r.GET("/test", func(context *gin.Context) {
//		context.JSON(200, gin.H{"uuu": "test"})
//	})
//}
import (
	_ "webframe/apps/all"
	"webframe/cmd"
)

func main() {
	cmd.Execute()
}
