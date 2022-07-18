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

//go:embed *.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	dirEntries, err := path.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println("Nama File: ", entry.Name())

			file, err := path.ReadFile("" + entry.Name())
			if err != nil {
				panic(err)
			}

			fmt.Println(string(file))
		}
	}
}
