package main

import (
	"io/fs"
	"os"

	"github.com/Kirbstomper/PersonalBlog/blog"
)

func main() {
	//Create a directory for the static site
	os.Mkdir("static", fs.ModePerm)

	//generate the site
	blog.GenerateSite("templates", "static", "posts")
}
