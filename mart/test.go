package mart

import "fmt"

func SayTest() {
	fmt.Println("Say Test")
	privteFunc()
}

func privteFunc() {
	fmt.Printf("ถ้าตัวแรกของ Function เป็นตัวเล็กจะเป็น Private Func ใช้ได้แค่ในไฟล์นี้ และถ้าใน Level เดียวกันสามารถเรียกใช้ในไฟล์อื่นได้เลย แต่ถ้าคนละ Layer ต้อง import")
}