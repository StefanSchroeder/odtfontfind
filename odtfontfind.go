// This package implements a library to retrieve the font-family
// from a small variety of file formats, in particular LibreOffice
// and Inkscape SVG files.
//
// This library comes with a client program that shows how to use it
// and is a valuable tool all by itself.
// Stefan Schröder 2019
package odtfontfind

import (
	"archive/zip"
	"gopkg.in/xmlpath.v2"
	"io"
	"log"
	"os"
	"strings"
)

const xxp = "/document-content/font-face-decls/font-face/@name"
const xp_tspan = "//tspan/@style"
const xp_text = "//text/@style"
const file_to_analyze = "content.xml"

func SvgFontReader(fname string) (ret []string) {
	rc, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	// Provide multiple Xpaths, because there is more than one:
	// Reference SVG 1.1 at w3.org
	for _, xpath := range []string{xp_text, xp_tspan} {
		rc.Seek(0, 0)
		for _, item := range parseXML(rc, xpath) {
			// This is the full 'stlye' element.
			for _, subitem := range strings.Split(item, ";") {
				// find the font-family item:
				a := strings.Split(subitem, ":")
				if a[0] == "font-family" {
					e := a[1]
					// Remove quotes
					e = e[1 : len(e)-1]
					ret = append(ret, e)
				}
			}
		}
	}
	rc.Close()
	return ret
}

func LibreofficeFontReader(fname string) (ret []string) {
	r, err := zip.OpenReader(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for _, f := range r.File {
		if f.Name == file_to_analyze {
			rc, err := f.Open()
			if err != nil {
				log.Fatal(err)
			}
			for _, item := range parseXML(rc, xxp) {
				ret = append(ret, item)
			}
			rc.Close()
		}
	}
	return ret
}

func parseXML(fd io.Reader, pattern string) (result []string) {
	root, err := xmlpath.Parse(fd)
	if err != nil {
		log.Fatal(err)
	}

	path := xmlpath.MustCompile(pattern)

	iter := path.Iter(root)
	for iter.Next() {
		strVal := iter.Node().String()
		result = append(result, strVal)
	}
	return result
}
