package chapter12

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed pass.txt
var pass string

func TestEmbedString(t *testing.T) {
	fmt.Println(pass)
}

//go:embed 1.png
var image []byte

func TestEmbedByte(t *testing.T) {
	fmt.Println(string(image))
}

//go:embed pass.txt
//go:embed 1.png
var multi embed.FS

func TestMulti(t *testing.T) {
	fmt.Println(multi)

	a, _ := multi.ReadFile("pass.txt")
	fmt.Println(string(a))
}
