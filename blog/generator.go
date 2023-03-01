package blog

import (
	"html/template"
	"os"
)

// Generates the entire site
func GenerateSite(path string) {
	generateIndex(path)
}

// Generates the index landing page
func generateIndex(path string) {
	tmpl, err := template.New("index.html").ParseFiles("../templates/index.html", "../templates/header.html")
	if err != nil {
		panic("Error reading template files")
	}
	file, err := os.Create(path + "/index.html")
	if err != nil {
		panic("Error creating index.html")
	}
	err = tmpl.Execute(file, nil)
	if err != nil {
		panic("Error executing template")
	}
}
