package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pdf/fpdf"
	"net/http"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()

	r.GET("/pdf", func(c *gin.Context) {
		pdf := fpdf.New("P", "mm", "A4", "")
		pdf.AddPage()
		pdf.SetFont("Arial", "B", 16)
		pdf.Cell(40, 10, "Hello, world")

		var buf bytes.Buffer
		err := pdf.Output(&buf)
		fmt.Println("error => ", err)

		c.Data(http.StatusOK, "application/pdf", buf.Bytes())
	})

	r.Run()
}
