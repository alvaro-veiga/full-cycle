package main

import (
	"fmt"
	"time"
)

func contador(qtd int) {
	for i := range qtd {
		fmt.println(i)
		time.Sleep(time.Second)
	}
}

func main() {
	go contador(10) // thread 2
	go contador(10) // thread 3
	contador(10)
}