package entity

type Operation struct {
	InputDir          *string
	OutputDir         *string
	CascadeClassifier *string
	Width             *int
	Height            *int
	Gray              *bool
	Concurrency       *int
	DataArguation     *bool
}
