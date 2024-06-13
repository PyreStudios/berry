# Berry

Your own personal little Li-berry. A media server for digital reading -- comics, books, pdfs.

Disclaimer: Everything under this line is the pipe-dream -- this does not work as is today, but this is what we're building towards.

## Setup

TODO: Add installation instructions for Berry, ideally including Docker, native Go, and a binary version.

A configuration file is used by Berry to help kickstart your media server and manage it's configuration. The configuration file is written in Yaml and you'll want to give your server a name and specify a list of "media_roots". These are root folders that you can store your media. Each root should follow the following structure (let's assume my "M" drive is my media root):

TODO: Does this structure make sense?

M:
| > comics
| | > Marvel (this is a publisher)
| | | > Iron Man (this is a comic series)
| | | | > 44 (this is the series number)
| | | | | > iron_man_44.pdf (whatever the name of the file is)
| | | | | > iron_man_44_cover.jpg (the cover image for this comic)
| > books
| | > Henry James (this is an author)
| | | > The Portrait of a Lady (this is a book)
| | | | > portrait.epub (whatever the name of the file is)
| | | | > portrait_cover.jpg (the cover image for this book)