package controller

import (
	mysqldb "webcam/connection"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	db := mysqldb.SetupDB()

	Username, _ := c.Get("username")

	Userid, _ := c.Get("userid")

	var imgpath string

	err := db.QueryRow("select image_path from `user` where id=?", Userid).Scan(&imgpath)

	if err != nil {
		imgpath = "assets/img/avatar.svg"
	}

	c.HTML(200, "index.html", gin.H{"imagepath": imgpath, "Username": Username})
}
