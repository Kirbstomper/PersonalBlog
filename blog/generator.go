package blog

import (
	"html/template"
	"os"
	"sort"
)

var posts []Post

// Generates the entire site
func GenerateSite(templatesPath string, outpath string, postsdir string) {
	generateIndex(templatesPath, outpath)
	generatePosts(templatesPath, outpath, postsdir)
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
func generatePosts(templatesPath, outpath, postsdir string) {
	os.Mkdir(outpath+"/posts", os.ModePerm)

	// For each post, create its page, add the title of page to list
	var posts = ReadAllPostsInDirectory(postsdir)
	//sort posts by date
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})

	//Tempate for posts
	tmpl, err := template.New("post.html").ParseFiles(templatesPath+"/post.html", templatesPath+"/header.html")
	if err != nil {
		panic("Error reading template files for generating posts")
	}
	for i, p := range posts {
		if i != 0 {
			p.Next = posts[i-1].GetTrucatedTitle()
		}
		if i != len(posts)-1 {
			p.Previous = posts[i+1].GetTrucatedTitle()
		}
		generatePageForPost(outpath, p, tmpl)
	}

	tmpl, err = template.New("posts.html").ParseFiles(templatesPath+"/posts.html", templatesPath+"/header.html")
	if err != nil {
		panic("Error reading template files")
	}
	file, err := os.Create(outpath + "/posts.html")
	if err != nil {
		panic("Error creating posts.html")
	}
	err = tmpl.Execute(file, posts)
	if err != nil {
		panic("Error executing template")
	}

}

//generate each post

func generatePageForPost(outpath string, post Post, tmpl *template.Template) {
	file, err := os.Create(outpath + "/posts/" + post.GetTrucatedTitle() + ".html")
	if err != nil {
		panic("Error creating " + post.Title)
	}
	tmpl.Execute(file, post)
}
