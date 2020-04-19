package graph

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRContractionInput(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")

	cnt := 0
	for i := 0; i < len(vertices); i++ {
		if vertices[i].edges != nil {
			cnt++
		}
	}

	if cnt != 200 {
		t.Error("RContractionInput error !")
	}

	fmt.Println(len(edges))

	if len(edges) != 2517 {
		t.Error("RContractionInput error !")
	}
}

func TestContracted(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")

	cnt := 0
	for i := 0; i < len(edges); i++ {
		if edges[i].contracted == true {
			cnt++
		}
	}

	fmt.Println(cnt)

	for i := 0; i < len(vertices[0].edges); i++ {
		vertices[0].edges[i].contracted = true
	}

	cnt = 0
	for i := 0; i < len(edges); i++ {
		if edges[i].contracted == true {
			cnt++
		}
	}

	fmt.Println(cnt)
}

func TestPointer2(t *testing.T) {
	vertices := make([]Vertex, 8)

	vertices[0] = MakeVertex(0)
	vertices[0].edges = make([]*Edge, 3)

	vertices[1] = MakeVertex(1)
	vertices[1].edges = make([]*Edge, 3)
	vertices[2] = MakeVertex(2)
	vertices[2].edges = make([]*Edge, 3)
	vertices[3] = MakeVertex(3)
	vertices[3].edges = make([]*Edge, 3)
	// vertices[4] = MakeVertex(4)
	// vertices[5] = MakeVertex(5)
	// vertices[6] = MakeVertex(6)
	// vertices[7] = MakeVertex(7)

	edges := make([]Edge, 6)

	pos := 0

	edges[pos] = MakeEdge(vertices, 0, 1)

	vertices[0].edges[vertices[0].pos] = &edges[pos]
	vertices[0].pos++
	vertices[1].edges[vertices[1].pos] = &edges[pos]
	vertices[1].pos++
	pos++

	edges[pos] = MakeEdge(vertices, 0, 2)

	vertices[0].edges[vertices[0].pos] = &edges[pos]
	vertices[0].pos++
	vertices[2].edges[vertices[2].pos] = &edges[pos]
	vertices[2].pos++
	pos++

	edges[pos] = MakeEdge(vertices, 0, 3)

	vertices[0].edges[vertices[0].pos] = &edges[pos]
	vertices[0].pos++
	vertices[3].edges[vertices[3].pos] = &edges[pos]
	vertices[3].pos++
	pos++

	edges[pos] = MakeEdge(vertices, 1, 2)

	vertices[1].edges[vertices[1].pos] = &edges[pos]
	vertices[1].pos++
	vertices[2].edges[vertices[2].pos] = &edges[pos]
	vertices[2].pos++
	pos++

	edges[pos] = MakeEdge(vertices, 2, 3)

	vertices[2].edges[vertices[2].pos] = &edges[pos]
	vertices[2].pos++
	vertices[3].edges[vertices[3].pos] = &edges[pos]
	vertices[3].pos++
	pos++
	edges[pos] = MakeEdge(vertices, 1, 3)

	vertices[1].edges[vertices[1].pos] = &edges[pos]
	vertices[1].pos++
	vertices[3].edges[vertices[3].pos] = &edges[pos]
	vertices[3].pos++
	pos++
	r = rand.New(rand.NewSource(time.Now().Unix()))

	contract := edges[0]

	contract.head.edges[0].contracted = true
	contract.head.edges[1].contracted = true

	contract = edges[len(edges)-1]
	contract.head.edges[len(contract.head.edges)-1].contracted = true

	fmt.Println(contract.head.edges[0])

	for i := 0; i < len(vertices); i++ {
		if &vertices[i] == contract.head.edges[0].head {
			fmt.Println("true")
		}

		if &vertices[i] == contract.head.edges[0].tail {
			fmt.Println("true")
		}
	}

	for i := 0; i < len(edges); i++ {
		if &edges[i] == contract.head.edges[0] {
			fmt.Println("true")
		}

		if &edges[i] == contract.head.edges[1] {
			fmt.Println("true")
		}

		if &edges[i] == contract.head.edges[2] {
			fmt.Println("true")
		}
	}

	cnt := 0
	for i := 0; i < len(edges); i++ {
		if edges[i].contracted == true {
			cnt++
		}
	}

	fmt.Println(edges)
	fmt.Println(contract.head.edges[0].head)
	fmt.Println(cnt)
}

func TestPointer(t *testing.T) {
	slice := make([]Edge, 0)

	slice = append(slice, Edge{})
	slice = append(slice, Edge{})
	slice = append(slice, Edge{})
	slice = append(slice, Edge{})
	slice = append(slice, Edge{})
	slice = append(slice, Edge{})
	slice = append(slice, Edge{})
	slice = append(slice, Edge{})

	ptr3 := &slice[3]
	ptr3.contracted = true

	ptr1 := &slice[1]
	ptr1.contracted = true

	ptr5 := &slice[5]
	ptr5.contracted = true

	fmt.Println(slice)
}

func TestPointer3(t *testing.T) {
	slice := make([]int, 0)

	slice = append(slice, 0)
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)

	slice2 := make([]*int, 0)

	slice2 = append(slice2, &slice[2])
	slice2 = append(slice2, &slice[3])

	fmt.Println(slice)
	fmt.Println(*slice2[0], *slice2[1])

	*slice2[0] = 182

	fmt.Println(slice)
	fmt.Println(*slice2[0], *slice2[1])

	// slice2[0] = 182

	// fmt.Println(slice)
	// fmt.Println(slice2)

	// slice = append(slice, 4)
	// slice = append(slice, 5)

	// fmt.Println(slice)
	// fmt.Println(slice2)

}

func TestRContraction(t *testing.T) {
	vertices, edges := ConstructGraph("../data/kargerMinCut.txt")

	mincut := RContraction(vertices, edges)

	fmt.Println(mincut)
}
