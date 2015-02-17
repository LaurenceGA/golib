/*
Package drawUtil package provides methods of drawing shapes using gl
*/
package drawUtil

import (
	"fmt"
	"math"
	"os"

	"github.com/LaurenceGA/lib/colours"

	"github.com/andrebq/gas"
	"github.com/go-gl/gl"
	"github.com/go-gl/gltext"
)

//Fnt is an Alias to avoid package importation
//type Fnt *gltext.Font

//InitGltext will load in font files
func InitGltext(scale int) *gltext.Font {
	// This file holds the actual glyph shapes.
	imgFile, err := gas.Abs("cpackages/Pong/src/fonts/bitmap_font.png")
	if err != nil {
		fmt.Printf("Find font image file: %v", err)
	}

	// This file is a JSON encoded dataset which describes the font
	// and contains the pixel offsets and sizes for each glyph in
	// bitmap_font.png. Both files are needed to load a bitmap font.
	configFile, err := gas.Abs("cpackages/Pong/src/fonts/bitmap-font.js")
	if err != nil {
		fmt.Printf("Find font config file: %v", err)
	}

	font, err := loadFont(imgFile, configFile, scale)
	if err != nil {
		fmt.Println(err)
	}

	return font
}

//Loads a scaled bitmap font
func loadFont(image, config string, scale int) (*gltext.Font, error) {
	a, err := os.Open(image)
	if err != nil {
		return nil, err
	}

	defer a.Close()

	b, err := os.Open(config)
	if err != nil {
		return nil, err
	}

	defer b.Close()

	return gltext.LoadBitmap(a, b, scale)
}

//GetBounds returns dimensions of string characters
func GetBounds(fnt *gltext.Font) (int, int) {
	var w, h = fnt.GlyphBounds()
	return w, h
}

//DrawString draws a given string at a given location
func DrawString(x, y float32, str string, col colours.Colour, font *gltext.Font) error {
	gl.Color4f(col.R, col.G, col.B, col.A)
	err := font.Printf(x, y, str)
	if err != nil {
		return err
	}

	return nil
}

//DrawSquare will draw a coloured square with a given side length
func DrawSquare(sideLen float64, col colours.Colour) {
	//Lay out square's vertices
	gl.Color4f(col.R, col.G, col.B, col.A)
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(sideLen/2, sideLen/2)
	gl.Vertex2d(-sideLen/2, sideLen/2)
	gl.Vertex2d(-sideLen/2, -sideLen/2)
	gl.Vertex2d(sideLen/2, -sideLen/2)
	gl.End()

	gl.Vertex3f(0, 0, 0)
}

//DrawRect will draw a coloured rectangle with a given side length
func DrawRect(x, y float64, col colours.Colour) {
	gl.Color4f(col.R, col.G, col.B, col.A)
	gl.Begin(gl.POLYGON)
	gl.Vertex2d(x/2, y/2)
	gl.Vertex2d(-x/2, y/2)
	gl.Vertex2d(-x/2, -y/2)
	gl.Vertex2d(x/2, -y/2)
	gl.End()

	gl.Vertex3f(0, 0, 0)
}

//DrawLine with create a line betweeen two points with a given width and colour
func DrawLine(x1, y1, x2, y2 float64, col colours.Colour, w float32) {
	gl.LineWidth(w)
	gl.Color4f(col.R, col.G, col.B, col.A)
	gl.Begin(gl.LINES)
	gl.Vertex2d(x1, y1)
	gl.Vertex2d(x2, y2)
	gl.End()

	gl.Vertex3f(0, 0, 0)
}

//DrawDotLine draws a dashed line
//Too much math
func DrawDotLine(x1, y1, x2, y2 float64, col colours.Colour, w float32, segments int) {
	gl.LineWidth(w)
	gl.Color4f(col.R, col.G, col.B, col.A)
	gl.Begin(gl.LINES)
	lLen := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))
	for i := 0; i < segments/2; i++ {
		gl.Vertex2d(x1+float64(i)*lLen/float64(segments), y1+float64(i)*lLen/float64(segments))
		gl.Vertex2d(x2+float64(i+1)*lLen/float64(segments), y2+float64(i+1)*lLen/float64(segments))
	}
	gl.End()

	gl.Vertex3f(0, 0, 0)
}

//MakeCircle creates circle points
func MakeCircle(radius float64, res int) [][2]float64 {
	tp := (2 * math.Pi)
	var vertexes [][2]float64
	for i := 0; i < res; i++ {
		theta := (tp / float64(res)) * float64(i)
		//vertexes[i] = [2]float64{(math.Cos(theta) * radius), (math.Sin(theta) * radius)}
		vertexes = append(vertexes, [2]float64{(math.Cos(theta) * radius), (math.Sin(theta) * radius)})
	}
	return vertexes
}

//DrawVertexes from an array of points
func DrawVertexes(verts [][2]float64, col colours.Colour) {
	gl.Color4f(col.R, col.G, col.B, col.A)
	gl.Begin(gl.POLYGON)
	// gl.Vertex2d(x1, y1)
	for _, arr := range verts {
		gl.Vertex2d(arr[0], arr[1])
	}
	gl.End()

	gl.Vertex3f(0, 0, 0)
}

//DrawCircle with create a cricle
func DrawCircle(radius float64, col colours.Colour, res int) {
	//var theta float64
	tp := (2 * math.Pi)
	gl.Color4f(col.R, col.G, col.B, col.A)
	gl.Begin(gl.POLYGON)
	for i := 0; i < res; i++ {
		theta := (tp / float64(res)) * float64(i)
		gl.Vertex2d(math.Cos(theta)*radius, math.Sin(theta)*radius)
	}
	gl.End()

	gl.Vertex3f(0, 0, 0)
}
