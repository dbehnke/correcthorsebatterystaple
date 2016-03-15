package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, err := os.Open("wordlist.txt")
	if err != nil {
		fmt.Println("ERR Reading wordlist.txt", err)
		os.Exit(1)
	}
	defer in.Close()

	out, err := os.Create("wordlist.go")
	if err != nil {
		fmt.Println("ERR Creating wordlist.go", err)
	}
	defer out.Close()

	out.Write([]byte("package main \n\n var( \n"))
	out.Write([]byte("words=[]string{\"\""))

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		s := scanner.Text()
		t := fmt.Sprintf(",\n\"%s\"", s)
		out.Write([]byte(t))
	}
	out.Write([]byte("}\n)\n"))
}
