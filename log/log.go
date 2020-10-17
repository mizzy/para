package log

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"io"

	"github.com/gookit/color"
)

type logger struct {
	prefix string
}

func New(arg string) *logger {
	c := determineColor(arg)
	return &logger{
		prefix: c(fmt.Sprintf("[%s]", arg)),
	}
}

func (l *logger) Info(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("%s %s %s\n", l.prefix, color.FgGreen.Render("[stdout]"), scanner.Text())
	}
}

func (l *logger) Error(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Printf("%s %s %s\n", l.prefix, color.FgRed.Render("[stderr]"), scanner.Text())
	}
}

var colorList = []color.Color{
	color.BgRed,
	color.BgGreen,
	color.BgYellow,
	color.BgBlue,
	color.BgMagenta,
	color.BgCyan,
	color.BgDarkGray,
	color.BgLightRed,
	color.BgLightGreen,
	//color.BgLightYellow,
	color.BgLightBlue,
	color.BgLightMagenta,
	color.BgLightCyan,
	color.BgLightWhite,
	color.BgGray,
}

func determineColor(arg string) func(...interface{}) string {
	hash := fnv.New32()
	hash.Write([]byte(arg))
	idx := hash.Sum32() % uint32(len(colorList))
	return color.New(colorList[idx], color.OpBold).Render
}
