package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"gocv.io/x/gocv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tfaceblur [camera ID] [classifier XML file]")
		return
	}

	//parse args
	deviceID, _ := strconv.Atoi(os.Args[1])
	xmlFile := os.Args[2]

	//open webcam
	webcam, err := gocv.VideoCaptureDevice(deviceID)
	if err != nil {
		fmt.Println("error opening video capture device %V\n", deviceID)
		return
	}

	defer webcam.Close()

	//open display window
	window := gocv.NewWindow("Face Blur")
	defer window.Close()

	//prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	//load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	classifier.Load(xmlFile)

	fmt.Printf("start reading camera device: %v\n", deviceID)
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("can't read device %d\n", deviceID)
			return
		}

		if img.Empty() {
			continue
		}

		//detect faces
		rects := classifier.DetectMultiScale(img)
		fmt.Printf("found %d faces\n", len(rects))

		//blur each face on the original image
		for _, r := range rects {
			imgFace := img.Region(r)
			//blur face
			gocv.GaussianBlur(imgFace, &imgFace, image.Pt(75,75), 0, 0, gocv.BorderDefault)
			imgFace.Close()
		}

		//show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
