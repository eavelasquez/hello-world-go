package crypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestNewSHA256(t *testing.T) {
	for i, tt := range []struct {
		in  []byte
		out string
	}{
		{[]byte("Hello, world."), "f8c3bf62a9aa3e6fc1619c250e48abe7519373d3edf41be62eb5dc45199af2ef"},
		{[]byte("abc"), "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"},
		{[]byte(""), "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
	} {
		t.Run(fmt.Sprintf("%v", i), func(t *testing.T) {
			result := NewSHA256(tt.in)
			if hex.EncodeToString(result) != tt.out {
				t.Errorf("NewSHA256(%v) = %v, want %v; got %v", tt.in, result, tt.out, hex.EncodeToString(result))
			}
		})
	}
}
