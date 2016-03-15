package main

import (
	"fmt"
	"math/rand"
	"time"
)

//go:generate go run scripts/makewordlist.go

func main() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 10; i++ {
		numbers := []int{rand.Intn(len(words)),
			rand.Intn(len(words)),
			rand.Intn(len(words)),
			rand.Intn(len(words))}
		fmt.Println(fmt.Sprintf("%s %s %s %s", words[numbers[0]], words[numbers[1]],
			words[numbers[2]], words[numbers[3]]))
	}
}
