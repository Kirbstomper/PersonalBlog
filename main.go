package main

import (
	"io/fs"
	"os"

	"github.com/Kirbstomper/PersonalBlog/blog"
)

func main() {
	//Create a directory for the static site
	os.Mkdir("static", fs.ModeDir)

	//generate the site
	blog.GenerateSite("static")
}
