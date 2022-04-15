
package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

var count int
var lock = sync.Mutex{}

//	m := append(S2B("用户"+username+"说："),message...)
func main() {
	flag.Parse()
	r:=gin.Default()
	hub := newHub()

	r.GET("/ws", func(context *gin.Context) {
		lock.Lock()
		count++
		username := strconv.Itoa(count)
		lock.Unlock()
		go hub.run(username)
		 serveWs(hub, context)
	})
	r.Run(":8080")
}
