package main

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time" // Import time มาเพื่อจับเวลาในการ Run
)

// step คือเวลาที่ใช้ต่อหนึ่งงาน (เทสใช้ค่าที่สั้นกว่านี้)
const step = 1 * time.Second

func main() {
	took := shop(os.Stdout, step)
	fmt.Println("ใช้เวลาในการ Run ทั้งสิ้น :", took)
}

// shop ทำงาน 5 อย่าง แต่ละอย่างใช้เวลา step: 3 อย่างทำพร้อมกันด้วย goroutine
// อีก 2 อย่างทำต่อกันใน goroutine หลัก จึงใช้เวลารวมประมาณ 2 step ไม่ใช่ 5
func shop(out io.Writer, step time.Duration) time.Duration {
	start := time.Now()     // เริ่มจับเวลาตอนโปรแกรมทำงาน (3)
	ch := make(chan string) // สร้างท่อ Channel เอาไว้ส่งข้อมูล (4)
	var wg sync.WaitGroup   // รอให้ goroutine ทุกตัวทำงานเสร็จก่อนจบ
	wg.Add(2)
	go func() { defer wg.Done(); buyGlassesAtSevenEleven(out, step) }()
	go func() { defer wg.Done(); buyWatchAtCentral(out, step) }()
	go sendToMisterA(ch, step) // ส่งท่อ ch เข้าไปใน Function (5)
	buyFruitAtSiamParagon(out, step)
	buyCarAtToyota(out, step)
	messageFromMisterB := <-ch                        // ค่าจากท่อ Channel จะออกตรงนี้ (6)
	if messageFromMisterB == "กำลังส่งของให้ นาย A" { // (7)
		fmt.Fprintln(out, "นาย A ได้รับของแล้ว")
	}
	wg.Wait()
	return time.Since(start)
}

func sendToMisterA(message chan<- string, step time.Duration) { //สร้าง Function (1)
	time.Sleep(step)
	message <- "กำลังส่งของให้ นาย A" // นำค่าเข้าท่อ Channel (2)
}

func buyGlassesAtSevenEleven(out io.Writer, step time.Duration) {
	time.Sleep(step)
	fmt.Fprintln(out, "ซื้อแว่น : ที่เซเว่น : เสร็จแล้ว")
}

func buyWatchAtCentral(out io.Writer, step time.Duration) {
	time.Sleep(step)
	fmt.Fprintln(out, "ซื้อนาฬิกา : ที่เซ็นทรัล : เสร็จแล้ว")
}

func buyFruitAtSiamParagon(out io.Writer, step time.Duration) {
	time.Sleep(step)
	fmt.Fprintln(out, "ซื้อผลไม้ : ที่สยามพารากอน : เสร็จแล้ว")
}

func buyCarAtToyota(out io.Writer, step time.Duration) {
	time.Sleep(step)
	fmt.Fprintln(out, "ซื้อรถ : ที่ศูนย์โตโยต้า : เสร็จแล้ว")
}
