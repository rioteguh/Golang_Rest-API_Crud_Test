package entities

import (
	"fmt"
)

type Nama struct{
	Id 		int64 	`json:"id"`
	Name 	string 	`json:"name"`
}

func (nama Nama) ToString() string {
	return fmt.Sprintf("id: %d\nname: %s", nama.Id, nama.Name)
}