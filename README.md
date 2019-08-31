# Face Collector
Face collector is a data augmentation tool.
This tool detect face in a picture, rotate them, and then resize them.

## Requirements

* OpenCV4.0

```
$ go get -u -d gocv.io/x/gocv
$ cd $GOPATH/src/gocv.io/x/gocv
$ make install
```

## How to Use

At first, download [a cascade classifier file](https://github.com/bookun/face-collector/blob/master/data/haarcascade_frontalface_alt.xml)

```
$ go get -u github.com/bookun/face-collector
$ face-collector -h
Usage of face-collector:
  -c string
        cascade classifier filepath
  -height int
        height (default 100)
  -input string
        input directory
  -output string
        output directory
  -width int
        width (default 100)

$ face-collector -input ./sample_input -output ./sample_output -height 100 -width 100 -c ./haarcascade_frontalface_alt.xml
```

| image | description|
| :----: | :---: |
| ![original]("https://github.com/bookun/face-collector/blob/master/testdata/input/person1/person1.jpg")| original |
