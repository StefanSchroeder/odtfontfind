# odtfontfind

The purpose of odtfontfind is to determine font-families that are used
in a LibreOffice document or in an SVG Inkscape image.

# Installation

To install the client, run

	GOBIN=~/bin go install github.com/StefanSchroeder/odtfontfind/odtfontfind-client@latest
	~/bin/odtfontfind-client sample-file.odt

The result might look like this



But usually you might want to use the library. In this case, no
install is required, just reference the library in your import
statement and let go get do its thing. Visit
https://github.com/StefanSchroeder/apt-font for an example of an
application that uses this library.

# Caveats

You could easily construct input files that make the library barf.
The library will bail out. Let me know if you come across this case. 

The main reason why retrieval of fonts might fail is that the 
input file is corrupt.

# License 

The license is MIT.

# References

http://docs.oasis-open.org/office/v1.2/os/OpenDocument-v1.2-os-part1.html#property-fo_font-family

