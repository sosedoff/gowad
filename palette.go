package wad

import (
	"image/color"
)

// RGB contains pixel color values
type RGB struct {
	R uint8
	G uint8
	B uint8
}

// Palette contains a set of colors
type Palette [256]RGB

// Playpal contains a set of palettes
type Playpal [14]Palette

// Color returns a new RGBA color variable
func (c RGB) Color() color.RGBA {
	return color.RGBA{c.R, c.G, c.B, 0xFF}
}

// GetPlaypal returns a set of palettes from the blob
func GetPlaypal(blob *Blob) (*Playpal, error) {
	result := &Playpal{}
	err := blob.Read(result)
	return result, err
}
