// Read the fonts from the content.xml of an ODT File
// to figure out what fonts were used.
// Stefan Schröder 2019
package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
//	"launchpad.net/xmlpath-2"
	"gopkg.in/xmlpath.v2"
	"log"
)

const xxp = "/document-content/font-face-decls/font-face/@name"
const file_to_analyze = "content.xml"

func main() {
	flag.Parse()
	for _, filename := range flag.Args() {
		ExampleReader(filename)
	}
}

func ExampleReader(fname string) {
	r, err := zip.OpenReader(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == file_to_analyze {
			rc, _ := f.Open()
			parseXML(rc)
			rc.Close()
		}
	}
}

func parseXML(fd io.Reader) {
	root, err := xmlpath.Parse(fd)
	if err != nil {
		log.Fatal(err)
	}

	path := xmlpath.MustCompile(xxp)

	iter := path.Iter(root)
	for iter.Next() {
		strVal := iter.Node().String()

		fmt.Print(strVal)
		fmt.Print("\n")
	}
}
