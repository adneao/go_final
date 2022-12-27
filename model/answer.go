package model

type Answer struct {
	TaskName string  `json:"taskName"`
	Percent  float32 `json:"percent"`
	Fails    []Fails `json:"fails"`
}

type Fails struct {
	OriginalResult string
	ExternalResult string
	DataSet        int
}
