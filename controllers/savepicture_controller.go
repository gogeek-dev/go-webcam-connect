package controller

import (
	"fmt"
	"io/ioutil"
	mysqldb "webcam/connection"

	"github.com/gin-gonic/gin"
)

func Viewpicture(c *gin.Context) {

	Username, _ := c.Get("username")

	tempFile, err := ioutil.TempFile("assets/image", "upload-*.jpg")

	if err != nil {
		fmt.Println(err)
	}

	filepath := tempFile.Name()

	fmt.Println(filepath)

	defer tempFile.Close()

	tempFile.Write(frame)

	c.HTML(200, "save-picture.html", gin.H{"filepath": filepath, "Username": Username})
}
func Savepicture(c *gin.Context) {

	db := mysqldb.SetupDB()

	Userid, _ := c.Get("userid")

	Path := c.Query("path")

	updt, err := db.Prepare("UPDATE `user` SET image_path=? WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	updt.Exec(Path, Userid)

	c.Redirect(301, "/index")
}
