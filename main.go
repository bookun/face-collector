package main

import (
	"context"
	"flag"
	"fmt"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/bookun/face-collector/entity"
	"github.com/bookun/face-collector/face_image"
	"github.com/bookun/face-collector/util"
	"github.com/cheggaaa/pb/v3"
)

func main() {
	operation := entity.Operation{
		InputDir:          flag.String("input", "", "input directory"),
		OutputDir:         flag.String("output", "", "output directory"),
		Width:             flag.Int("width", 100, "width"),
		Height:            flag.Int("height", 100, "height"),
		CascadeClassifier: flag.String("classifier", "", "cascade classifier filepath"),
		Gray:              flag.Bool("gray", false, "grayscale"),
		Concurrency:       flag.Int("concurrency", 4, "number of thread"),
		DataArguation:     flag.Bool("da", false, "data arguation"),
	}
	flag.Parse()
	if *operation.CascadeClassifier == "" {
		log.Fatal("specify cascade classifier file path by using 'c' option")
	}

	imagePaths, err := util.Dirwalk(*operation.InputDir)
	if err != nil {
		log.Fatal(err)
	}
	numOfRemainFile := len(imagePaths)
	bar := pb.StartNew(numOfRemainFile)
	ctx := context.Background()
	ctxParent, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	var mux sync.Mutex
	limit := make(chan struct{}, *operation.Concurrency)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	for _, imagePath := range imagePaths {
		wg.Add(1)
		go func(ctx context.Context, path string, op entity.Operation) {
			defer wg.Done()
			limit <- struct{}{}
			err := face_image.SaveFaceImages(ctx, path, operation)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
			mux.Lock()
			bar.Increment()
			mux.Unlock()
			<-limit
		}(ctxParent, imagePath, operation)
	}
	go func() {
		select {
		case <-quit:
			cancel()
			bar.Finish()
			fmt.Println("cancel request from user")
			close(quit)
			return
		}
	}()
	wg.Wait()
}
