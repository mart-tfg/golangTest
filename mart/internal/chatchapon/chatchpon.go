package mart //ใน Folder เดียวกัน หรือใน Level  เดียว ห้ามชื่อ package ซ้ำกันและมีได้ package เดียว.

import (
	"fmt"
)

func SayHelloChatchapon() { //func นี้จะใช้ได้แค่ใน Folder เดียวกันเท่านั้น แต่่ต้อง import  ด้วย
	fmt.Println("Hello Chatchapon")
}
