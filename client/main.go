// This is a client program for the odtfindfont library
// Its purpose is to take some files on the command line
// and print the font-families that are used within the
// file.

// Author: Stefan Schröder
// License: See LICENSE file.
package main

import "flag"
import "fmt"
import "strings"

import o "github.com/StefanSchroeder/odtfontfind"

func main() {
	flag.Parse()
	var ret = []string{}
	for _, filename := range flag.Args() {
		if strings.HasSuffix(filename, ".odt") ||
			strings.HasSuffix(filename, ".ods") ||
			strings.HasSuffix(filename, ".odp") {
			ret = append(ret, o.LibreofficeFontReader(filename)...)
		} else if strings.HasSuffix(filename, ".svg") {
			ret = append(ret, o.SvgFontReader(filename)...)
		} else {
			fmt.Println("Unknown suffix for \"" + filename + "\"")
		}
	}
	for _, item := range ret {
		fmt.Println(item)
	}
}
