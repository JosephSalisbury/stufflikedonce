package main

import (
	"math/rand"
	"time"

	"github.com/JosephSalisbury/stufflikedonce/cmd"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	cmd.Execute()
}
