package mysql

import (
	"fmt"
	"simple-proj/stores"
)

type Conn struct {
	db string
}

func New(db string) Conn {
	return Conn{db: db}
}

func (m Conn) Create(usr stores.User) error {
	fmt.Println("adding to mysql", usr)
	return nil
}
func (m Conn) Update(usr stores.User) error {
	fmt.Println("updating in mysql", usr)
	return nil
}

func (m Conn) Delete() error {
	fmt.Println("deleting in mysql")
	return nil
}
