package wad

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
	things := make([]Thing, blob.Count(Thing{}))
	err := blob.Read(&things)
	return things, err
}
