# odtfontfind

Determine font-families that are used in a LibreOffice document.
Or in an SVG Inkscape image.

# Usage

There is a tiny client to show how it works. The main point was to
create a tool to install the fonts for the documents you have.

    go run odtfontfind/client.go filename.odt

Will print a list of fonts referenced in that document. (Not necessarily
used.)

# Caveats

You could easily construct input files that make the library barf.
The library will bail out with a log.Fatalf. Let me know if you 
come across this case. 

# References

http://docs.oasis-open.org/office/v1.2/os/OpenDocument-v1.2-os-part1.html#property-fo_font-family

