package main

import (
	"fmt"
	"log"
	"os"
	mysqldbcon "webcam/connection"
	controller "webcam/controllers"
	routes "webcam/routes"

	"gocv.io/x/gocv"
)

func main() {

	if len(os.Args) < 2 {

		fmt.Println(">> device /dev/video0 (default)")

		controller.Webcam, controller.Err = gocv.VideoCaptureDevice(0)
		// window = gocv.NewWindow("Webcam")
	} else {

		fmt.Println(">> file/url :: " + os.Args[1])

		controller.Webcam, controller.Err = gocv.VideoCaptureFile(os.Args[1])
	}

	if controller.Err != nil {

		fmt.Printf("Error opening capture device: \n")

	}

	defer controller.Webcam.Close()

	db := mysqldbcon.SetupDB()

	r := routes.SetupRoutes(db)

	log.Println("Server started on: http://localhost:2020")

	r.Run(":2020")

}
