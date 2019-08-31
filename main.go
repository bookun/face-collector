package main

import (
	"context"
	"flag"
	_ "image/jpeg"
	_ "image/png"
	"log"

	"github.com/bookun/face-collector/entity"
	"github.com/bookun/face-collector/face_image"
	"github.com/bookun/face-collector/util"
	"golang.org/x/sync/errgroup"
)

func main() {
	operation := entity.Operation{
		InputDir:          flag.String("input", "", "input directory"),
		OutputDir:         flag.String("output", "", "output directory"),
		Width:             flag.Int("width", 100, "width"),
		Height:            flag.Int("height", 100, "height"),
		CascadeClassifier: flag.String("c", "", "cascade classifier filepath"),
	}
	flag.Parse()
	if *operation.CascadeClassifier == "" {
		log.Fatal("specify cascade classifier file path by using 'c' option")
	}

	imagePaths, err := util.Dirwalk(*operation.InputDir)
	if err != nil {
		log.Fatal(err)
	}

	eg, ctx := errgroup.WithContext(context.TODO())
	for _, imagePath := range imagePaths {
		path := imagePath
		eg.Go(func() error {
			return face_image.SaveFaceImages(ctx, path, operation)
		})
	}
	if err := eg.Wait(); err != nil {
		log.Panicln(err)
	}

}
