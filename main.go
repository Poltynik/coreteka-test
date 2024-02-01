package main

import (
	"fmt"
	"github.com/poltynik/coreteka-test.git/db"
)

// Uncomment necessary examples if needed
func main() {
	inmDB := db.NewInMemoryDB()

	// Commit transaction
	inmDB.Set("key1", "value1")
	inmDB.StartTransaction()
	inmDB.Set("key1", "value2")
	inmDB.CommitTransaction()
	fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value2”

	//Rollback transaction
	//inmDB.Set("key1", "value1")
	//inmDB.StartTransaction()
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value1”
	//inmDB.Set("key1", "value2")
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value2”
	//inmDB.RollBackTransaction()
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value1”

	// Commit nested transaction
	//inmDB.Set("key1", "value1")
	//inmDB.StartTransaction()
	//inmDB.Set("key1", "value2")
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value2”
	//inmDB.StartTransaction()
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value2”
	//inmDB.Delete("key1")
	//inmDB.CommitTransaction()
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get None
	//inmDB.CommitTransaction()
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get None

	//// Rollback with nested transaction
	//inmDB.Set("key1", "value1")
	//inmDB.StartTransaction()
	//inmDB.Set("key1", "value2")
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value2”
	//inmDB.StartTransaction()
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value2”
	//inmDB.Delete("key1")
	//inmDB.RollBackTransaction()
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value2”
	//inmDB.CommitTransaction()
	//fmt.Println(inmDB.Get("key1")) //-> Expect to get ”value2”
}
