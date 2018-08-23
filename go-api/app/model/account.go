package model

type Account struct {
	Id int `json:"id"`
	Name string `json:"name"`
  }
  
func NewAccount(id int, name string) *Account {
	return &Account{
		Id: id,
		Name: name,
	}
}