package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-git/go-git"
)

func main() {
	r := gin.Default()
	r.POST("/api/v1/validate_repository", validate_repository)
	r.Run()
}

func ValidateRepositoryEndpoint(c *gin.Context):
        repository := c.PostForm("repository")
        username := c.PostForm("username")
        password := c.PostForm("password")
	validate_status, have_err := ValidateRepository(repository, username, password)
	if have_err true {
                c.JSON(200, gin.H{
                        "status":  false,
                        "messege": "There was a problem checking the repository",
                })
	} else {
		if validate_staus {
                        c.JSON(200, gin.H{
                                "status": true,
                                "Repository approved"
                        })
		} else {
                        c.JSON(200, gin.H{
                                "status":  false,
                                "messege": "Repository not approved",
                        })
                }
	}


func ValidateRepository(repository, username, password string) (bool, bool){
	if username != "" && password == "" {
		_, err = git.PlainClone("./tmp", false, &git.CloneOptions{
			URL:      repository,
			Progress: os.Stdout,
		})
	} else {
		_, err := git.PlainClone("./tmp", false, &git.CloneOptions{
			URL:      repository,
			Progress: os.Stdout,
			Auth: &http.BasicAuth{
				Username: username,
				Password: password,
			},
		})
		if err != nil {
			return false, true
		}
	}
	if fileExists("tmp/" + os.Getenv("FILE_VALIDATE")) == true {
		return true, false
	} else {
		return false, false
	}
	os.RemoveAll("tmp")

}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
