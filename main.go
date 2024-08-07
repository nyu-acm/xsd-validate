package main

import (
	"fmt"
	"io"
	"os"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

func main() {
	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()

	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl(os.Args[1], xsdvalidate.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xsdhandler.Free()

	xmlFile, err := os.Open(os.Args[2])
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	inXml, err := io.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	// Option 1:
	xmlhandler, err := xsdvalidate.NewXmlHandlerMem(inXml, xsdvalidate.ParsErrDefault)
	if err != nil {
		panic(err)
	}
	defer xmlhandler.Free()

	err = xsdhandler.Validate(xmlhandler, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch err.(type) {
		case xsdvalidate.ValidationError:
			fmt.Println(err)
			fmt.Printf("Error in line: %d\n", err.(xsdvalidate.ValidationError).Errors[0].Line)
			fmt.Println(err.(xsdvalidate.ValidationError).Errors[0].Message)
		default:
			fmt.Println(err)
		}
	}
}
