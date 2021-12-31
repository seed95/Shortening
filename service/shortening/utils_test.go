package shortening

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {

	data1 := "test"

	tests := []struct {
		data   string
		result [md5.Size]byte
	}{
		{
			data:   data1,
			result: md5.Sum([]byte(data1)),
		},
		{
			data:   "",
			result: md5.Sum([]byte("")),
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("hash %v", tt.data)
		t.Run(name, func(t *testing.T) {
			if result := hash(tt.data); result != tt.result {
				t.Fatalf("got resutl: %v, want result: %v", result, tt.result)
			}
		})
	}

}

//TODO REMOVE THIS FUNCTION
func TestEncode(t *testing.T) {

	tests := []struct {
		data   []byte
		result string
	}{
		{
			data:   []byte(""),
			result: "",
		},
		{
			data:   []byte("276J21#!kadfhoi23[hrlknckSCNLKNKLDN lzdvhsadkklsfj dj l;fkajsf; kQWH O'1l;"),
			result: "",
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("encode %v", tt.data)
		t.Run(name, func(t *testing.T) {
			result := encode(tt.data)
			fmt.Println(result)
			//if result != tt.result {
			//	t.Fatalf("got resutl: %v, want result: %v", result, tt.result)
			//}
		})
	}

}

//TODO REMOVE THIS FUNCTION
func TestDecode(t *testing.T) {

	tests := []struct {
		data string
	}{
		{
			data: "kljas-_dfoifjlk",
		},
		{
			data: "kljas-_dfoifjlk##",
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("encode %v", tt.data)
		t.Run(name, func(t *testing.T) {
			result, err := decode(tt.data)
			fmt.Println("result:", result, ", err:", err)
			//if result != tt.result {
			//	t.Fatalf("got resutl: %v, want result: %v", result, tt.result)
			//}
		})
	}

}
