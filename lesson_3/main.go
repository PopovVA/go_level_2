package lesson3

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/valyala/fasthttp"
)

func main() {
	fmt.Printf("websocket.BinaryMessage: %v\n", websocket.BinaryMessage)
	fmt.Print(fasthttp.AcquireArgs())
}
