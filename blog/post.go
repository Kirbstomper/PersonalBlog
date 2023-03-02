package blog

import (
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

type Post struct {
	Title   string `yaml:"title"`
	Date    string `yaml:"date"`
	Content string `yaml:"content"`
}

func ReadPostFromFile(filepath string) Post {
	p := &Post{}
	f, err := os.ReadFile(filepath)
	if err != nil {
		panic("error reading files " + filepath)
	}
	err = yaml.Unmarshal(f, p)
	if err != nil {
		panic("error unmarshalling yaml " + filepath)

	}
	return *p
}

func ReadAllPostsInDirectory(dirpath string) []Post {
	posts = make([]Post, 0)
	files, err := os.ReadDir(dirpath)
	if err != nil {
		panic("Error reading directory")
	}
	for _, f := range files {
		posts = append(posts, ReadPostFromFile(dirpath+"/"+f.Name()))
	}
	return posts
}

func (p Post) GetTrucatedTitle() string {
	return strings.ReplaceAll(p.Title, " ", "")
}
