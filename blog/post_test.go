package blog_test

import (
	"log"
	"testing"

	"github.com/Kirbstomper/PersonalBlog/blog"
)

func Test_ReadFromFile(t *testing.T) {
	p := blog.ReadPostFromFile("../test/testpost.yml")
	log.Println(p.Title)
	log.Println(p.Date)
	log.Println(p.Content)
}
