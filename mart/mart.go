package mart //ใน Folder เดียวกัน หรือใน Level  เดียว ห้ามชื่อ package ซ้ำกันและมีได้ package เดียว.

import (
	"fmt"
	"github.com/mart/go-example/mart/internal/chatchapon"


)

func SayHelloMart() {
	fmt.Println("Hello Mart")
	mart.SayHelloChatchapon()
	
}
