package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Printf("No arguments were provided.\n")
		return
	}

	sentences := os.Args[1:]

	for _, v := range sentences {
		words := strings.Split(v, " ")
		reverseWords(words)
	}
}

func reverseWords(wl []string) {
	n := len(wl)
	nwl := make([]string, n)
	for _, v := range wl {
		n--
		nwl[n] = reverseLetters(v)
	}

	fmt.Printf("%s\n", strings.Join(nwl, " "))
}

func reverseLetters(w string) string {
	n := len(w)
	nw := make([]string, n)
	for _, v := range w {
		n--
		nw[n] = string(v)
	}

	return strings.Join(nw, "")
}
