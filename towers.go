package main

import (
	"fmt"
	"strings"
	"time"
)

type location int

const (
	left location = iota
	middle
	right
)

func (l location) String() string {
	switch l {
	case left:
		return "left"
	case middle:
		return "middle"
	case right:
		return "right"
	default:
		panic("unreachable")
	}
}

type towers []location

func printDisk(d, width int) string {
	d++
	return strings.Repeat(" ", width-d) + strings.Repeat("=", d) + "|" + strings.Repeat("=", d) + strings.Repeat(" ", width-d)
}

func printPillar(base string, t towers, l location) []string {
	tower := make([]string, len(t))
	for i := range tower {
		tower[i] = base
	}
	h := 0
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] == l {
			tower[h] = printDisk(i, len(t))
			h++
		}
	}
	return tower
}

func printMerged(l, m, r []string) {
	for i := len(l) - 1; i >= 0; i-- {
		fmt.Println(l[i] + " " + m[i] + " " + r[i])
	}
}

func (t towers) _print() {
	base := strings.Repeat(" ", len(t))
	base = base + "|" + base
	l := printPillar(base, t, left)
	m := printPillar(base, t, middle)
	r := printPillar(base, t, right)
	printMerged(l, m, r)
}

func move(src, dest, mid location, n int, t towers) {
	if n > 1 {
		move(src, mid, dest, n-1, t)
	}
	t[n-1] = dest
	time.Sleep(1400 * time.Millisecond)
	fmt.Println("\n")
	fmt.Printf("%s -> %s\n", src, dest)
	t._print()
	//fmt.Printf("%s -> %s\n", src, dest)
	if n > 1 {
		move(mid, dest, src, n-1, t)
	}
}

func main() {
	t := make(towers, 9)
	fmt.Println("starting position")
	t._print()
	move(left, right, middle, 9, t)
	//move("L", "R", "M", 9)
}
