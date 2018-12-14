package main

import (
	"fmt"
	"strings"
)

func q(ch chan string) {
	var str string
	for {
		fmt.Scan(&str)
		ch <- str
	}

}

func a(ch chan string) {
	for {
		s := <-ch
		s = strings.Replace(s, "吗", "", -1)
		s = strings.Replace(s, "?", "!", -1)
		s = strings.Replace(s, "? ", "! ", -1)

		fmt.Println("AI:", s)
	}

}

func main() {
	ch := make(chan string)
	go q(ch)
	go a(ch)

	var done chan bool
	<- done
}
