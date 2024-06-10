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
	go contador(10)
	go contador(10)
	contador(10)
}