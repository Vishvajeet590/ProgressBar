package main

import (
	"Progress/cmd"
	"fmt"
	"time"
)

func main() {
	bar := cmd.NewProgressBar(10, 200)
	fmt.Printf("Downlading file....")
	for i := 1.0; i <= 10; i++ {
		time.Sleep(time.Second * 1)
		bar.Update(i)
	}
	time.Sleep(time.Second)
	bar.CleanUp()
}
