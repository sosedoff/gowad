package wad

// GetPNames reads a set of pname strings from the blob
func GetPNames(blob *Blob) ([]String8, error) {
	count := blob.MustReadInt32()
	names := make([]String8, count)
	err := blob.Read(&names)
	return names, err
}
