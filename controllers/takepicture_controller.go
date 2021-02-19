package controller

import (
	"image"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"gocv.io/x/gocv"
)

var (
	Err      error
	Webcam   *gocv.VideoCapture
	window   *gocv.Window
	frame_id int
)

var buffer = make(map[int][]byte)
var frame []byte
var mutex = &sync.Mutex{}

func Takepicture(c *gin.Context) {

	Username, _ := c.Get("username")

	c.HTML(200, "take-picture.html", gin.H{"Username": Username})
}

func Video(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary=frame")

	data := ""
	for {
		/*			fmt.Println("Frame ID: ", frame_id)
		 */mutex.Lock()

		data = "--frame\r\n  Content-Type: image/jpeg\r\n\r\n" + string(frame) + "\r\n\r\n"

		mutex.Unlock()

		time.Sleep(33 * time.Millisecond)

		c.Writer.Write([]byte(data))
	}
}
func Getframes() {

	img := gocv.NewMat()

	defer img.Close()

	for {
		Webcam.Read(&img)
		if img.Empty() {
			continue
		}

		frame_id++

		gocv.Resize(img, &img, image.Point{}, float64(0.5), float64(0.5), 0)

		frame, _ = gocv.IMEncode(".jpg", img)
	}

}
