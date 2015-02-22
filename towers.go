package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bradfitz/iter"
)

const disks = 5

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

func printDisk(d int, bar ...string) string {
	if len(bar) == 0 {
		bar = []string{"|"}
	}
	return strings.Repeat(" ", disks-d) + strings.Repeat("=", d) + bar[0] + strings.Repeat("=", d) + strings.Repeat(" ", disks-d)
}

func printPillar(base string, t towers, l location) []string {
	tower := make([]string, len(t))
	for i := range tower {
		tower[i] = base
	}
	h := 0
	for i := len(t) - 1; i >= 0; i-- {
		if t[i] == l {
			tower[h] = printDisk(i + 1)
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

type pillar [disks]int

func printPillars(ps []pillar) {
	for i := disks - 1; i >= 0; i-- {
		fmt.Printf("%s %s %s\n", printDisk(ps[0][i]), printDisk(ps[1][i]), printDisk(ps[2][i]))
	}
}

func shift(hover string, dest, src location) string {
	if dest < src {
		return hover[1:] + " "
	}
	return " " + hover[:len(hover)-1]
}

func delta(dest, src location) int {
	if src == left && dest == right || src == right && dest == left {
		return 4*disks + 4
	}
	return 2*disks + 2
}

func (t towers) move(dest, src location) {
	pillars := make([]pillar, 3)
	h := make([]int, 3)
	for i := disks - 1; i >= 0; i-- {
		l := t[i]
		pillars[l][h[l]] = i + 1
		h[l]++
	}
	//fmt.Println()
	//printPillars(pillars)

	for hh := h[src] - 1; hh+1 < disks; hh++ {
		pillars[src][hh], pillars[src][hh+1] = 0, pillars[src][hh]
		time.Sleep(100 * time.Millisecond)
		fmt.Println()
		printPillars(pillars)
	}
	n := pillars[src][disks-1]
	pillars[src][disks-1] = 0

	var hover string
	{
		hs := []string{printDisk(0, " "), printDisk(0, " "), printDisk(0, " ")}
		hs[src] = printDisk(n, " ")
		hover = strings.Join(hs, " ")
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(hover)
	printPillars(pillars)
	for range iter.N(delta(dest, src)) {
		time.Sleep(50 * time.Millisecond)
		hover = shift(hover, dest, src)
		fmt.Println(hover)
		printPillars(pillars)
	}

	pillars[dest][disks-1] = n
	for hhh := disks - 1; hhh-1 > h[dest]; hhh-- {
		time.Sleep(100 * time.Millisecond)
		fmt.Println()
		printPillars(pillars)
		pillars[dest][hhh], pillars[dest][hhh-1] = 0, pillars[dest][hhh]
	}
}

func move(src, dest, mid location, n int, t towers) {
	if n > 1 {
		move(src, mid, dest, n-1, t)
	}
	t.move(dest, src)
	t[n-1] = dest
	if n > 1 {
		move(mid, dest, src, n-1, t)
	}
}

func main() {
	t := make(towers, disks)
	fmt.Println("starting position")
	t._print()
	move(left, right, middle, disks, t)
	t._print()
}
