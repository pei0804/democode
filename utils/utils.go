package Utils

import "fmt"

type Data struct {
	Title string
}

func NewHoge(data string) *Data {
	return &Data{data}
}

func (d *Data) Hoge() {
	fmt.Println(d.Title)
}
