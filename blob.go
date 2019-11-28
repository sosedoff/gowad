package wad

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"io/ioutil"
)

// Blob contains raw bytes
type Blob struct {
	Size   int32
	data   []byte
	reader io.ReadSeeker
}

// NewBlob returns a new blob initialized from bytes
func NewBlob(data []byte) *Blob {
	return &Blob{
		Size:   int32(len(data)),
		data:   data,
		reader: bytes.NewReader(data),
	}
}

// NewBlobFromFile returns a new blob with contents from a file
func NewBlobFromFile(path string) (*Blob, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewBlob(data), nil
}

// Bytes returns a blob byte slice
func (b Blob) Bytes() []byte {
	return b.data
}

// Seek seeks to an offset within the blob
func (b Blob) Seek(offset int32) error {
	off := int64(offset)
	pos, err := b.reader.Seek(off, 0)
	if err != nil {
		return err
	}
	if pos != off {
		return errors.New("seek position mismatch")
	}
	return nil
}

// MustSeek seeks to an offset within the blob or panics
func (b Blob) MustSeek(offset int32) {
	if err := b.Seek(offset); err != nil {
		panic(err)
	}
}

// Read reads binary data into the provided output
func (b Blob) Read(out interface{}) error {
	return binary.Read(b.reader, binary.LittleEndian, out)
}

// MustRead reads a binary data and panics on errors
func (b Blob) MustRead(out interface{}) {
	if err := b.Read(out); err != nil {
		panic(err)
	}
}

// MustReadByte reads a single byte or panics
func (b Blob) MustReadByte() byte {
	var result byte
	b.MustRead(&result)
	return result
}

// MustReadString8 reads a 8-byte string or panics
func (b Blob) MustReadString8() String8 {
	var result String8
	b.MustRead(&result)
	return result
}

// MustReadInt32 reads a single Int32 value or panics
func (b Blob) MustReadInt32() int32 {
	var result int32
	b.MustRead(&result)
	return result
}

// MustReadInt16 reads a single Int16 value or panics
func (b Blob) MustReadInt16() int16 {
	var result int16
	b.MustRead(&result)
	return result
}

// MustReadInt8 reads a single Int8 value or panics
func (b Blob) MustReadInt8() int8 {
	var result int8
	b.MustRead(&result)
	return result
}
