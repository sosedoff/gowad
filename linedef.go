package wad

// Linedef represents a line from one of the VERTEXES to another
type Linedef struct {
	VertexFrom int8
	VertexTo   int8
	Flags      int8
	Types      int8
	Tag        int8
	LeftNum    int8
	RightNum   int8
}

// ReadLinedefs reads a linedef objects from the blob
func ReadLinedefs(blob *Blob) ([]Linedef, error) {
	linedefs := make([]Linedef, blob.Count(Linedef{}))
	err := blob.Read(&linedefs)
	return linedefs, err
}
