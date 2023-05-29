package main

import "fmt"

type Item struct {
	title string
	body  string
}

var database []Item

func AddItem(item Item) Item {
	database = append(database, item)
	return item
}
func GetByName(title string) Item {
	var getItem Item
	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}
	return getItem
}

func EditItem(title string, edit Item) Item {
	var changed Item
	for idx, valu := range database {
		if valu.title == edit.title {
			database[idx] = edit
			changed = edit

		}
	}
	return changed

}

func DeleteItem(item Item) Item {
	var del Item
	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	return del

}

func main() {
	fmt.Println("intial database:", database)
	a := Item{"first", "i am the first"}
	b := Item{"second", "i am the second"}
	c := Item{"third", "i am the third"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("database:", database)
	DeleteItem(c)
	fmt.Println("database:", database)

}
