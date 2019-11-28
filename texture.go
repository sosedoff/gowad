package wad

// Texture contains header and set of patches
type Texture struct {
	TextureHeader
	Patches []Patch
}

// TextureHeader contains texture meta data
type TextureHeader struct {
	Name       String8
	Blank1     int16
	Blank2     int16
	Width      int16
	Height     int16
	Blank3     int16
	Blank4     int16
	NumPatches int16
}

// Patch contains texture patch information
type Patch struct {
	XOffset  int16
	YOffset  int16
	Number   int16
	StepDir  int16
	Colormap int16
}

// GetTextures reads textures from the blob
func GetTextures(blob *Blob) ([]Texture, error) {
	// Get number of textures in the blob
	count := blob.MustReadInt32()

	// Read data offsets
	offsets := make([]int32, count)
	blob.MustRead(&offsets)

	textures := make([]Texture, count)
	// Load texture data
	for i, offset := range offsets {
		blob.MustSeek(offset)

		header := TextureHeader{}
		blob.MustRead(&header)

		patches := make([]Patch, header.NumPatches)
		blob.MustRead(&patches)

		textures[i] = Texture{
			TextureHeader: header,
			Patches:       patches,
		}
	}

	return textures, nil
}
