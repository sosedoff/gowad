package wad

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

// transparentIndex defines a color value for transparent pixels
const transparentIndex = 255

// defaultBackground is a color used for background
var defaultBackground = color.RGBA{0, 255, 255, 255}

// ImageHeader contains information about image
type ImageHeader struct {
	Width   int16
	Height  int16
	XOffset int16
	YOffset int16
}

// Image contains header and pixel values
type Image struct {
	ImageHeader
	Pixels []byte
}

// Render returns an image rendered using a palette
func (img Image) Render(pal Palette) (*image.RGBA, error) {
	width := int(img.Width)
	height := int(img.Height)
	rect := image.Rect(0, 0, width, height)
	target := image.NewRGBA(rect)

	// Draw image background
	draw.Draw(
		target,
		target.Bounds(),
		&image.Uniform{defaultBackground},
		image.ZP,
		draw.Src,
	)

	// Set actual pixel colors using the palette
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			val := img.Pixels[y*width+x]
			if val != transparentIndex {
				target.SetRGBA(x, y, pal[val].Color())
			}
		}
	}

	return target, nil
}

// RenderToFile renders image to a file using a palette
func (img Image) RenderToFile(path string, pal Palette) error {
	result, err := img.Render(pal)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, result)
}

// GetImage reads an image from a blob
func GetImage(blob *Blob) (*Image, error) {
	// Read image header
	header := ImageHeader{}
	blob.MustRead(&header)

	// Read image data offsets
	offsets := make([]int32, header.Width)
	blob.MustRead(&offsets)

	// Init pixels array
	size := int(header.Width * header.Height)
	width := int(header.Width)
	pixels := make([]byte, size)

	// Set default background color
	for i := range pixels {
		pixels[i] = transparentIndex
	}

	for col, offset := range offsets {
		blob.Seek(offset)
		rowStart := blob.MustReadByte()

		// Row end flag
		if rowStart == 255 {
			continue
		}

		// Read number of colored pixels in this column
		numPixels := blob.MustReadByte()
		if numPixels == 0 {
			continue
		}

		// Read pixel color values
		colors := make([]byte, numPixels)
		blob.MustReadByte()
		blob.Read(&colors)
		blob.MustReadByte()

		// Assign color values to pixels
		for i := rowStart; i < rowStart+numPixels; i++ {
			pixels[int(i)*width+col] = colors[i-rowStart]
		}
	}

	return &Image{ImageHeader: header, Pixels: pixels}, nil
}
