package main

import (
	"fmt"
	"time" // Import time มาเพื่อจับเวลาในการ Run
)

func main() {
	start := time.Now()     // เริ่มจับเวลาตอนโปรแกรมทำงาน (3)
	ch := make(chan string) // สร้างท่อ Channel เอาไว้ส่งข้อมูล (4)
	go buyGlassesAtSevenEleven()
	go buyWatchAtCentral()
	go sendToMisterA(ch) // ส่งท่อ ch เข้าไปใน Function (5)
	buyFruitAtSiamParagon()
	buyCarAtToyota()
	messageFromMisterB := <-ch                        // ค่าจากท่อ Channel จะออกตรงนี้ (6)
	if messageFromMisterB == "กำลังส่งของให้ นาย A" { // (7)
		fmt.Println("นาย A ได้รับของแล้ว")
		fmt.Println("ใช้เวลาในการ Run ทั้งสิ้น : ", time.Since(start), "วินาที")
	}
}

func sendToMisterA(message chan<- string) { //สร้าง Function (1)
	time.Sleep(1 * time.Second)
	message <- "กำลังส่งของให้ นาย A" // นำค่าเข้าท่อ Channel (2)
}

func buyGlassesAtSevenEleven() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อแว่น : ที่เซเว่น : เสร็จแล้ว")
}

func buyWatchAtCentral() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อนาฬิกา : ที่เซ็นทรัล : เสร็จแล้ว")
}

func buyFruitAtSiamParagon() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อผลไม้ : ที่สยามพารากอน : เสร็จแล้ว")
}

func buyCarAtToyota() {
	time.Sleep(1 * time.Second)
	fmt.Println("ซื้อรถ : ที่ศูนย์โตโยต้า : เสร็จแล้ว")
}
