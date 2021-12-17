package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"

    "github.com/jung-kurt/gofpdf"
)

func main() {
    args := os.Args
    if len(args) < 2 {
        log.Fatal("Please pass file name as an argument")
    }
    file := args[1]
    content, err := ioutil.ReadFile(file)
    if err != nil {
        log.Fatalf("%s file not found", file)
    }

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    pdf.SetFont("Arial", "B", 14.0)

    pdf.MultiCell(190, 5, string(content), "0", "0", false)

    _ = pdf.OutputFileAndClose("test.pdf")

    fmt.Println("PDF created")
}

