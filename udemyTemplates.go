package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

func udemyTemplatesMain() {
	fmt.Println("Huh")

	// Let's make a file
	nf, err := os.Create("somefile.txt")
	if err != nil {
		log.Fatal("couldn't create file", err)
	}
	// Closes the file after the function finishes... I think
	defer nf.Close()

	// Have to pass in a reader object
	//   is it really an object? idk
	io.Copy(nf, strings.NewReader("yolo"))

	/*
		Instead of parsing files explicitly, we can have it parse
		some arbitrary directory (or glob)

		aTemplate, err := template.ParseFiles("html/base.gohtml")
		if err != nil {
			log.Fatal("Couldn't parse template", err)
		}
	*/

	// Here we use "Must", which is a handy dandy wrapper provided
	//  by the teplate package to wrap the error handling for us
	aTemplate := template.Must(template.ParseGlob("html/*.gohtml"))

	err = aTemplate.ExecuteTemplate(os.Stdout, "base.gohtml", 42)
	if err != nil {
		log.Fatal("Couldn't write to stdout", err)
	}

}
