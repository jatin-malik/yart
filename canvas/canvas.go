package canvas

import (
	"fmt"
	"github.com/jatin-malik/yart/color"
	"os"
	"strings"
)

type Canvas struct {
	width  int
	height int
	screen [][]color.Color
}

func New(width int, height int) *Canvas {
	screen := make([][]color.Color, height)
	for i := range screen {
		screen[i] = make([]color.Color, width)
	}
	return &Canvas{width, height, screen}
}

// WritePixel writes pixel at (x,y). Assumes y to start from top left.
func (c *Canvas) WritePixel(x int, y int, color color.Color) {
	if y < 0 || y >= c.height {
		return
	}
	if x < 0 || x >= c.width {
		return
	}
	c.screen[y][x] = color
}

// PixelAt gives pixel at (x,y). Assumes y to start from top left.
func (c *Canvas) PixelAt(x int, y int) color.Color {
	if y < 0 || y >= c.height {
		return color.Color{}
	}
	if x < 0 || x >= c.width {
		return color.Color{}
	}
	return c.screen[y][x]
}

func (c *Canvas) ToPPM() string {
	buf := strings.Builder{}
	buf.WriteString(fmt.Sprintf("P3\n%d %d\n255\n", c.width, c.height))

	for y := 0; y < c.height; y++ {
		var rowParts []string
		for x := 0; x < c.width; x++ {
			pixel := c.PixelAt(x, y)
			p := pixel.ToByte()
			// Append each R G B as a chunk
			rowParts = append(rowParts,
				fmt.Sprintf("%d %d %d", int(p.GetRed()), int(p.GetGreen()), int(p.GetBlue())))
		}

		line := ""
		for _, part := range rowParts {
			if len(line)+len(part)+1 > 70 {
				buf.WriteString(strings.TrimSpace(line) + "\n")
				line = ""
			}
			line += part + " "
		}
		if line != "" {
			buf.WriteString(strings.TrimSpace(line) + "\n")
		}
	}

	return buf.String()
}

func (c *Canvas) Fill(color color.Color) {
	for y := 0; y < c.height; y++ {
		for x := 0; x < c.width; x++ {
			c.screen[y][x] = color
		}
	}
}

func (c *Canvas) GetWidth() int {
	return c.width
}

func (c *Canvas) GetHeight() int {
	return c.height
}

// ToDisk saves the canvas to disk in the plain PPM format.
// (Filename should include extension .ppm)
func (c *Canvas) ToDisk(filename string) error {
	ppm := c.ToPPM()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(ppm)
	return err
}
