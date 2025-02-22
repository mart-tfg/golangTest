package main //

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mart/go-example/mart"
)

func main() {
	var numberRandom int = 2 // var จะสามารถเปลี่ยนค่าได้ภายหลัง
	const intRandom int = 3  // const จะเป็นค่า default จะไมม่สารถเปลี่ยนแลงได้
	id := uuid.New()
	fmt.Println("Hello world:", intRandom)
	fmt.Printf("UUID: %s test: %d", id, numberRandom)
	mart.SayHelloMart()
	mart.SayTest()
}
