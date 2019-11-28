package wad

// String8 represents a 8-byte string
type String8 [8]byte

// String16 represents a 16-byte string
type String16 [16]byte

func (s String8) String() string {
	for i, v := range s {
		if v == 0 {
			return string(s[0:i])
		}
	}
	return string(s[:])
}

func (s String16) String() string {
	for i, v := range s {
		if v == 0 {
			return string(s[0:i])
		}
	}
	return string(s[:])
}

func string8(input string) String8 {
	if len(input) > 8 {
		panic("String is too long")
	}
	var result [8]byte
	for i, c := range []byte(input) {
		result[i] = c
	}
	return result
}
