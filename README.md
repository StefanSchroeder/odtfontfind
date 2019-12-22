# odtfontfind
Determine fonts that are used in an ODT LibreOffice document.

http://docs.oasis-open.org/office/v1.2/os/OpenDocument-v1.2-os-part1.html#property-fo_font-family

Usage:

    go run odtfontfind.go filename.odt

Will print a list of fonts referenced in that document. (Not necessarily
used.)

