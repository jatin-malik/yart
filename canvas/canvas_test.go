package canvas

import (
	"github.com/jatin-malik/yart/color"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCreation(t *testing.T) {
	canvas := New(10, 20)

	assert.Equal(t, canvas.width, 10)
	assert.Equal(t, canvas.height, 20)

	// Assert every pixel is init to black(0,0,0)
	for r := 0; r < canvas.height; r++ {
		for c := 0; c < canvas.width; c++ {
			assert.Equal(t, canvas.screen[r][c], color.New(0, 0, 0))
		}
	}
}

func TestPixelWriting(t *testing.T) {
	canvas := New(10, 20)
	red := color.New(1, 0, 0)

	canvas.WritePixel(3, 2, red)
	assert.Equal(t, red, canvas.PixelAt(3, 2))
}

func TestToPPMHeader(t *testing.T) {
	canvas := New(5, 3)
	ppm := canvas.ToPPM()
	ppmLines := strings.Split(ppm, "\n")
	assert.Equal(t, "P3", ppmLines[0])
	assert.Equal(t, "5 3", ppmLines[1])
	assert.Equal(t, "255", ppmLines[2])
}

func TestToPPM(t *testing.T) {
	canvas := New(5, 3)
	c1 := color.New(1.5, 0, 0)
	c2 := color.New(0, 0.5, 0)
	c3 := color.New(-0.5, 0, 1)

	canvas.WritePixel(0, 0, c1)
	canvas.WritePixel(2, 1, c2)
	canvas.WritePixel(4, 2, c3)

	ppm := canvas.ToPPM()

	expectedLines := `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`

	assert.Equal(t, expectedLines, ppm)
}

func TestToPPMWith70CharsLimit(t *testing.T) {
	canvas := New(10, 2)
	c := color.New(1, 0.8, 0.6)
	canvas.Fill(c)

	ppm := canvas.ToPPM()
	ppmLines := strings.Split(ppm, "\n")
	for _, line := range ppmLines {
		assert.True(t, len(line) <= 70)
	}
}

func TestPPMEndsWithNewlineCharacter(t *testing.T) {
	canvas := New(10, 2)
	c := color.New(1, 0.8, 0.6)
	canvas.Fill(c)

	ppm := canvas.ToPPM()
	assert.NotEmpty(t, ppm, "PPM output should not be empty")

	lastChar := ppm[len(ppm)-1]
	assert.EqualValuesf(t, '\n', lastChar, "PPM output should end with a newline, got %q", lastChar)
}
