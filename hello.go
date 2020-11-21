package main

import (
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	defer webcam.Close() // Close the cam once execution completed
	window := gocv.NewWindow("Hello")
	defer window.Close() // Close the GUI once execution completed, though it is done automatically
	img := gocv.NewMat()
	defer img.Close() // Close the img once execution completed

	for {
		webcam.Read(&img)
		window.IMShow(img)
		if window.WaitKey(1) == 27 { // 27 => Esc
			break
		}
	}
}
