package wad

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

const (
	headerSize = 12
	entrySize  = 16
)

type File struct {
	reader  *os.File
	header  Header
	entries []Entry
}

// Header contains WAD meta data
type Header struct {
	Signature       [4]byte
	Entries         int32
	DirectoryOffset int32
}

// Entry contains WAD blob meta data
type Entry struct {
	Offset int32
	Size   int32
	Name   String8
}

// Header returns WAD header
func (f *File) Header() Header {
	return f.header
}

// Index returns all entries in the WAD file
func (f *File) Index() []Entry {
	return f.entries
}

// Open returns a new WAD file
func Open(path string) (*File, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	file := File{reader: f}

	if err := file.readHeader(); err != nil {
		f.Close()
		return nil, err
	}

	if err := file.readEntries(); err != nil {
		f.Close()
		return nil, err
	}

	return &file, nil
}

// GetEntryBlob returns a blob for the entry
func (f *File) GetEntryBlob(entry Entry) (*Blob, error) {
	if entry.Size == 0 {
		return nil, errors.New("Entry is not a blob")
	}
	data := make([]byte, entry.Size)
	if _, err := f.reader.ReadAt(data, int64(entry.Offset)); err != nil {
		return nil, err
	}

	blob := Blob{
		data:   data,
		reader: bytes.NewReader(data),
		Size:   entry.Size,
	}

	return &blob, nil
}

// GetBlob returns a new blob for a name
func (f *File) GetBlob(name string) (*Blob, error) {
	var entry Entry
	var found bool

	for _, e := range f.entries {
		if e.Name.String() == name {
			if found {
				return nil, errors.New("Found multiple entries with name " + name)
			}
			entry = e
			found = true
			continue
		}
	}

	return f.GetEntryBlob(entry)
}

// Close closes WAD file
func (f *File) Close() error {
	return f.reader.Close()
}

func readHeader(r io.Reader) (Header, error) {
	header := Header{}
	buf := make([]byte, headerSize)
	bufReader := bytes.NewBuffer(buf)

	n, err := r.Read(buf)
	if err != nil {
		return header, err
	}
	if n != headerSize {
		return header, errors.New("size mismatch")
	}

	if err := binary.Read(bufReader, binary.LittleEndian, &header); err != nil {
		return header, err
	}

	return header, nil
}

func (f *File) readHeader() error {
	return binary.Read(f.reader, binary.LittleEndian, &f.header)
}

func (f *File) readEntries() error {
	if _, err := f.reader.Seek(int64(f.header.DirectoryOffset), 0); err != nil {
		return err
	}

	entries := make([]Entry, f.header.Entries)
	for i := int32(0); i < f.header.Entries; i++ {
		if err := binary.Read(f.reader, binary.LittleEndian, &entries[i]); err != nil {
			return err
		}
	}

	f.entries = entries
	return nil
}

// MarkersIndex returns a set of all markers in the file
func (f *File) MarkersIndex() []Entry {
	result := []Entry{}
	for _, e := range f.entries {
		if e.Size == 0 {
			result = append(result, e)
		}
	}
	return result
}

// EntriesForMarkers returns a set of entries between markers
func (f *File) EntriesForMarkers(start, end string) ([]Entry, error) {
	startMarker := string8(start)
	endMarker := string8(end)
	startIndex := -1
	endIndex := -1

	for i, e := range f.entries {
		// Skip actual entries until we find the end marker
		if startIndex >= 0 && e.Size != 0 {
			continue
		}

		if e.Name == startMarker {
			if startIndex >= 0 {
				return nil, errors.New("Found duplicate start marker")
			}
			startIndex = i + 1
			continue
		}

		if startIndex < 0 {
			continue
		}

		if end == "" {
			if e.Size == 0 {
				endIndex = i
				break
			}
		} else {
			if e.Name == endMarker {
				if endIndex > 0 {
					return nil, errors.New("Found duplicate end marker")
				}
				endIndex = i
				break
			}
		}
	}

	if startIndex < 0 {
		return nil, errors.New("Start marker was not found")
	}
	if endIndex < 0 {
		return nil, errors.New("End marker was not found")
	}

	return f.entries[startIndex:endIndex], nil
}

// GetPNames returns patch names from a blob
func (f *File) GetPNames(name string) ([]String8, error) {
	blob, err := f.GetBlob(name)
	if err != nil {
		return nil, err
	}
	return GetPNames(blob)
}

// GetTextures returns a set of textures from a blob
func (f *File) GetTextures(name string) ([]Texture, error) {
	blob, err := f.GetBlob(name)
	if err != nil {
		return nil, err
	}
	return GetTextures(blob)
}

// GetPlaypal returns a play palette
func (f *File) GetPlaypal(name string) (*Playpal, error) {
	blob, err := f.GetBlob(name)
	if err != nil {
		return nil, err
	}
	return GetPlaypal(blob)
}

// GetImage returns an image for matching blob name
func (f *File) GetImage(name string) (*Image, error) {
	blob, err := f.GetBlob(name)
	if err != nil {
		return nil, err
	}
	return GetImage(blob)
}
