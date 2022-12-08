package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	height  int
	visible bool
}

type Wood struct {
	size  int
	trees [][]*Tree
}

func newWood(size int) *Wood {
	w := &Wood{
		size:  size,
		trees: make([][]*Tree, size),
	}
	for i := 0; i < size; i++ {
		w.trees[i] = make([]*Tree, size)
	}
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			w.trees[x][y] = &Tree{
				height:  0,
				visible: false,
			}
		}
	}
	return w
}

func (w *Wood) rotate() {
	for i, j := 0, w.size-1; i < j; i, j = i+1, j-1 {
		w.trees[i], w.trees[j] = w.trees[j], w.trees[i]
	}

	for i := 0; i < w.size; i++ {
		for j := 0; j < i; j++ {
			w.trees[i][j], w.trees[j][i] = w.trees[j][i], w.trees[i][j]
		}
	}
}

func (w *Wood) scan() {
	for y := 0; y < w.size; y++ {
		minHeight := -1
		idx := 0
		for x := 0; x < w.size; x++ {
			if w.trees[x][y].height > minHeight {
				minHeight = w.trees[x][y].height
				idx = x
				// w.trees[x][y].visible = true
			}
		}
		for x := idx; x < w.size; x++ {
			w.trees[x][y].visible = true
		}
	}
}

func (w *Wood) calculate() {
	for i := 0; i < 4; i++ {
		w.print()
		w.scan()
		w.print()
		w.rotate()
	}
	w.print()
}

func (w *Wood) print() {
	for y := 0; y < w.size; y++ {
		b := strings.Builder{}
		for x := 0; x < w.size; x++ {
			if !w.trees[x][y].visible {
				b.WriteString(fmt.Sprintf("%d", w.trees[x][y].height))
			} else {
				b.WriteString(fmt.Sprintf("x")) // w.trees[x][y].height
			}
		}
		log.Println(b.String())
	}
	log.Println()
}

func (w *Wood) getVisibleCount() (result int) {
	for i := 0; i < w.size; i++ {
		for j := 0; j < i; j++ {
			if w.trees[i][j].visible {
				result++
			}
		}
	}
	return
}

func parseInt(s []byte) int {
	i, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var wood *Wood

	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if wood == nil {
			wood = newWood(len(line))
		}
		for x := 0; x < len(line); x++ {
			wood.trees[x][y].height = parseInt([]byte{line[x]})
		}
		y++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	wood.calculate()
	log.Println(wood.getVisibleCount())

}
