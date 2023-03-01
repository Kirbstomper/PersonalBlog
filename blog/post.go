package blog

type Post struct {
	title   string `yaml:"title"`
	date    string `yaml:"date"`
	content string `yaml:"content"`
}
