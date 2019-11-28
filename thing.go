package wad

import (
	"encoding/binary"
)

// Thing contains information about level object
type Thing struct {
	XPos    int16
	YPos    int16
	Angle   int16
	Type    int16
	Options int16
}

// GetThings reads things from the blob
func GetThings(blob *Blob) ([]Thing, error) {
	count := int(blob.Size) / int(binary.Size(Thing{}))
	things := make([]Thing, count)
	err := blob.Read(&things)
	return things, err
}
