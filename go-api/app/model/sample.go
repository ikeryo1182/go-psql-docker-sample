package model

type Sample struct {
	Id int `json:"id"`
	Name string `json:"name"`
  }
  
func NewSample(id int, name string) *Sample {
	return &Sample{
		Id: id,
		Name: name,
	}
}