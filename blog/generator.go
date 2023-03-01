package blog

import (
	"html/template"
	"os"
)

// Generates the entire site
func GenerateSite(templatesPath string, outpath string) {
	generateIndex(templatesPath, outpath)
	generatePosts(templatesPath, outpath)
}

// Generates the index landing page
func generateIndex(templatesPath, outpath string) {
	tmpl, err := template.New("index.html").ParseFiles(templatesPath+"/index.html", templatesPath+"/header.html")
	if err != nil {
		panic("Error reading template files")
	}
	file, err := os.Create(outpath + "/index.html")
	if err != nil {
		panic("Error creating index.html")
	}
	err = tmpl.Execute(file, nil)
	if err != nil {
		panic("Error executing template")
	}
}

// Generates the posts page
func generatePosts(templatesPath, outpath string) {
	tmpl, err := template.New("posts.html").ParseFiles(templatesPath+"/posts.html", templatesPath+"/header.html")
	if err != nil {
		panic("Error reading template files")
	}
	file, err := os.Create(outpath + "/posts.html")
	if err != nil {
		panic("Error creating posts.html")
	}
	err = tmpl.Execute(file, nil)
	if err != nil {
		panic("Error executing template")
	}
}
