package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Item struct {
	title string
	body  string
}
type API int

var database []Item

func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}
func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item
	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}
	*reply = getItem
	return nil
}

func (a *API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for idx, valu := range database {
		if valu.title == edit.title {
			database[idx] = edit
			changed = edit

		}
	}
	*reply = changed
	return nil

}

func (a *API) DeleteItem(item Item, reply *Item) error {
	var del Item
	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil

}

func main() {
	var api = new(API)
	err := rpc.Register(api)
	if err != nil {
		log.Fatal("error registration Api", err)

	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatal("Listener error", err)
	}
	log.Printf("serving rpc on port %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("listner error", err)
	}
	// fmt.Println("intial database:", database)
	// a := Item{"first", "i am the first"}
	// b := Item{"second", "i am the second"}
	// c := Item{"third", "i am the third"}

	// AddItem(a)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("database:", database)
	// DeleteItem(c)
	// fmt.Println("database:", database)

}
