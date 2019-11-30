package wad

// Level contains all level objects
type Level struct {
	Linedefs []Linedef
	Sidedefs []Sidedef
	Vertices []Vertex
	Segs     []Seg
}

// Linedef represents a line from one of the VERTEXES to another
type Linedef struct {
	From     int16 // From Vertex
	To       int16 // To Vertex
	Flags    int16
	Types    int16
	Tag      int16
	LeftNum  int16
	RightNum int16
}

// Sidedef is a definition of what wall texture(s) to draw along a
// LINEDEF, and a group of sidedefs outline the space of a SECTOR.
//   There will be one sidedef for a line that borders only one sector
// (and it must be the RIGHT side, as noted in [4-3]). It is not
// necessary to define what the doom player would see from the other
// side of that line because the doom player can't go there. The doom
// player can only go where there is a sector.
type Sidedef struct {
	XOffset       int16
	YOffset       int16
	UpperTexture  String8
	MiddleTexture String8
	LowerTexture  String8
	Sector        int16
}

// Vertex contains beginning and end points for LINEDEFS and SEGS
type Vertex struct {
	X int16
	Y int16
}

// Seg is a definition of a line segment
type Seg struct {
	VStart     int16 // Start vertex number
	VEnd       int16 // End vertex number
	Angle      int16
	LinedefNum int16
	Direction  int16
	Offset     int16
}

// SSector is a sub-sector
type SSector struct {
	NumSegs  int16
	StartSeg int16
}

// BBox is a bounding box
type BBox struct {
	Top    int16
	Bottom int16
	Left   int16
	Right  int16
}

type Node struct {
	X     int16
	Y     int16
	Dx    int16
	Dy    int16
	BBox  [2]BBox
	Child [2]int16
}

type Sector struct {
	FloorHeight   int16
	CeilingHeight int16
	Floorpic      String8
	Ceilingpic    String8
	Lightlevel    int16
	SpecialSector int16
	Tag           int16
}

// ReadLinedefs reads a linedef objects from the blob
func ReadLinedefs(blob *Blob) ([]Linedef, error) {
	linedefs := make([]Linedef, blob.Count(Linedef{}))
	err := blob.Read(&linedefs)
	return linedefs, err
}

// ReadSidedefs reads a set of SIDEDEF's from the blob
func ReadSidedefs(blob *Blob) ([]Sidedef, error) {
	sidedefs := make([]Sidedef, blob.Count(Sidedef{}))
	err := blob.Read(&sidedefs)
	return sidedefs, err
}

// ReadVertices reads a set of VERTEX's from the blob
func ReadVertices(blob *Blob) ([]Vertex, error) {
	vertices := make([]Vertex, blob.Count(Vertex{}))
	err := blob.Read(&vertices)
	return vertices, err
}

// ReadSegs reads a set of SEG's from the blob
func ReadSegs(blob *Blob) ([]Seg, error) {
	segs := make([]Seg, blob.Count(Seg{}))
	err := blob.Read(&segs)
	return segs, err
}
