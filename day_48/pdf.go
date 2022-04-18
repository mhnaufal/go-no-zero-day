package main

import (
	"log"

	"github.com/jung-kurt/gofpdf"
)

func main() {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Text(40, 10, "Good morning!")

	err := pdf.OutputFileAndClose("./day_48/file.pdf")
	if err != nil {
		log.Println("ERROR", err.Error())
	}
}
