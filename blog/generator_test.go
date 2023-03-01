package blog_test

import (
	"os"
	"testing"

	"github.com/Kirbstomper/PersonalBlog/blog"
)

//Tests generation of website. If these tests fail, that means somethings is wrong
//with templates or data

func Test_GenerateSite(t *testing.T) {
	os.Mkdir("test_temp", os.ModePerm)
	blog.GenerateSite("../templates", "test_temp")
	t.Run("Generate Index Test", test_GenerateIndex)

	os.RemoveAll("test_temp")
}
func test_GenerateIndex(t *testing.T) {
	//make directory for temp
	_, err := os.Stat("test_temp/index.html")
	if err != nil {
		t.Fail()
	}

}
