package main

import (
	"fmt"
	"time"

	"github.com/quantum-craft/go/graph"
)

func report(mincut *int, quitCh chan bool) {
	for {
		time.Sleep(2000 * time.Millisecond)

		fmt.Println(*mincut)

		select {
		case <-quitCh:
			return
		default:
			continue
		}
	}
}

func main() {
	filePath := "../data/kargerMinCut.txt"
	vertices, edges := graph.ConstructGraph(filePath)

	minOfMincut := 100000 // larger than 2517 :P
	quitCh := make(chan bool)

	go report(&minOfMincut, quitCh)

	for i := 0; i < 50000; i++ { // n := len(vertices) // lgn := int(math.Log2(float64(n))) // for i := 0; i < n*n*lgn; i++ {
		vertices, edges = graph.ConstructGraph(filePath)

		mincut := graph.RContraction(vertices, edges)

		if mincut < minOfMincut {
			minOfMincut = mincut
		}
	}

	fmt.Print("min-cut is ")
	fmt.Println(minOfMincut)

	quitCh <- true
}
