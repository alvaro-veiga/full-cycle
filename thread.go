package main

import "fmt"

func main() {
	canal := make (chan string)

	go func() {
		canal <- "Olá Mundo!" // thread 2 // preenchido
	}()

	fmt.Println(<-canal)// thread 1 // vazio
}