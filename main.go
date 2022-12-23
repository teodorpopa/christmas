package main

import (
	"bytes"
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/buger/goterm"
	"github.com/logrusorgru/aurora"
)

var au *aurora.Aurora

var bulbs = flag.Int("bulbs", 7, "how many bulbs")
var title = flag.String("title", "Happy Holidays !!!", "title to display")

var colors = []uint8{
	124,
	184,
	105,
	45,
	166,
	225,
	70,
	231,
	21,
}

var baseColor uint8 = 179

func init() {
	flag.Parse()
}

func getBulb() []string {
	return []string{
		"_.--.--._",
		"   _Y_   ",
		"  [___]  ",
		"  /:' \\  ",
		" |::   | ",
		" \\::.  / ",
		"  \\::./  ",
		"   '='   ",
	}
}

func getColors(l int) []uint8 {
	rand.Seed(time.Now().Unix())
	c := []uint8{}

	for i := 0; i < l; i++ {
		c = append(c, colors[rand.Intn(len(colors))])
	}

	return c
}

func main() {
	for {
		goterm.Clear()
		goterm.MoveCursor(1, 1)
		displayTitle()
		displayLights()
		time.Sleep(time.Second)
		goterm.Flush()
	}
}

func displayTitle() {
	fmt.Println(centerText(*bulbs*9, *title))
	fmt.Println("")
	fmt.Println("")
}

func displayLights() {
	lights := make([][]string, *bulbs)

	for i := 0; i < *bulbs; i++ {
		lights[i] = getBulb()
	}

	lr := len(lights[0])
	l := len(lights)
	lc := getColors(l)

	for i := 0; i < lr; i++ {
		for j := 0; j < l; j++ {

			color := baseColor
			if i > 2 {
				color = lc[j]
			}

			fmt.Print(fmt.Sprint(aurora.Index(color, lights[j][i])))
		}
		fmt.Println("")
	}
}

func centerText(width int, s string) *bytes.Buffer {
	const half, space = 2, "\u0020"
	var b bytes.Buffer
	n := (width - utf8.RuneCountInString(s)) / half
	fmt.Fprintf(&b, "%s%s", strings.Repeat(space, n), s)
	return &b
}
