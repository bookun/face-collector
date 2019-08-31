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

At first, download [a cascade classifier file](https://github.com/bookun/face-collector/blob/master/data/haarcascade_frontalface_default.xml)

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

$ face-collector -input ./testdata/input -output ./testdata/output -height 100 -width 100 -c ./haarcascade_frontalface_default.xml

$ tree -L 2 testdata/output
testdata/output
├ person1
│   ├ 0_Angle0Blur3_person1-1.jpg
│   ├ 0_Angle0Blur3_person1-2.jpg
│   ├ 0_Angle30Blur0_person1-1.jpg
│   ├ 0_Angle30Blur0_person1-2.jpg
│   ├ 0_Angle315Blur0_person1-1.jpg
│   ├ 0_Angle315Blur0_person1-2.jpg
│   ├ 0_Angle330Blur0_person1-1.jpg
│   ├ 0_Angle330Blur0_person1-2.jpg
│   ├ 0_Angle45Blur0_person1-1.jpg
│   ├ 0_Angle45Blur0_person1-2.jpg
│   ├ 0_Original_person1-1.jpg
│   ├ 0_Original_person1-2.jpg
│   ├ 0_person1-1.jpg
│   └ 0_person1-2.jpg
└ person2
    ├ 0_Angle0Blur3_person2-1.jpg
    ├ 0_Angle30Blur0_person2-1.jpg
    ├ 0_Angle315Blur0_person2-1.jpg
    ├ 0_Angle330Blur0_person2-1.jpg
    ├ 0_Angle45Blur0_person2-1.jpg
    ├ 0_Original_person2-1.jpg
    └ 0_person2-1.jpg

2 directories, 21 files
```

| image | description|
| :----: | :---: |
| ![original](https://raw.githubusercontent.com/bookun/face-collector/master/testdata/input/person1/person1-1.jpg)| original |
| ![faceDetect](https://raw.githubusercontent.com/bookun/face-collector/master/testdata/output/person1/0_Original_person1-1.jpg) | original face|
| ![rotate30](https://raw.githubusercontent.com/bookun/face-collector/master/testdata/output/person1/0_Angle30Blur0_person1-1.jpg) | rotate 30|
| ![rotate45](https://raw.githubusercontent.com/bookun/face-collector/master/testdata/output/person1/0_Angle45Blur0_person1-1.jpg) | rotate 45|
| ![rotate315](https://raw.githubusercontent.com/bookun/face-collector/master/testdata/output/person1/0_Angle315Blur0_person1-1.jpg) | rotate 315|
| ![rotate330](https://raw.githubusercontent.com/bookun/face-collector/master/testdata/output/person1/0_Angle330Blur0_person1-1.jpg) | rotate 330|
| ![blur](https://raw.githubusercontent.com/bookun/face-collector/master/testdata/output/person1/0_Angle0Blur3_person1-1.jpg) | blur 3|
