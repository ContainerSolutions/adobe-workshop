package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Entrypoint of our (micro)sevice
func main() {
	fmt.Println("Deals service started...")

	// Connec to DB
	session := connect()
	defer session.Close()

	initData()

	http.HandleFunc("/deals", dealsHandler)

	http.ListenAndServe(":8080", nil)
}

// Handle deal requests
func dealsHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.FormValue("id")

	var deal Deal
	if idStr == "" {
		fmt.Println("No Id passed in")
		w.WriteHeader(400)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Oops: ", err.Error())
		w.WriteHeader(500)
		return
	}
	deal, err = fetchDeal(id)
	if err != nil {
		fmt.Println("Oops: ", err.Error())
		w.WriteHeader(404)
		return
	}

	err = json.NewEncoder(w).Encode(deal)
	if err != nil {
		fmt.Println("Oops: ", err.Error())
		w.WriteHeader(500)
	}
}

// Retrieve Deal by ID
func fetchDeal(id int) (Deal, error) {
	session := connect()
	result := Deal{}
	c := session.DB("test").C("deals")
	err := c.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		fmt.Printf("Could not find deal %d in database. Err: %s\n", id, err.Error())
		return Deal{}, errors.New("Invalid Id")
	}
	fmt.Printf("Found deal %d in datastore.\n", id)
	// Set Expiry date to +1 week
	result.Expires = time.Now().AddDate(0, 0, 7).Format("2006-01-02")
	fmt.Println("Expiry date added.")

	return result, nil
}

// Initialize dummy data
func initData() {
	session := connect()
	c := session.DB("test").C("deals")
	err := c.Insert(&Deal{Id: 1, Name: "Buy 400 pairs, get one unmatched sock free!"},
		&Deal{Id: 2, Name: "Free shipping anywhere in the Andromeda Galaxy"})
	if err != nil {
		fmt.Printf("Error inserting records in database: %s\n", err.Error())
		panic(err)
	}
}

func connect() (session *mgo.Session) {
	connectURL := "deals-db"
	session, err := mgo.Dial(connectURL)
	if err != nil {
		fmt.Printf("Can't connect to database. Error %s\n", err.Error())
		os.Exit(1)
	}
	return session
}

type Deal struct {
	Id      int
	Name    string
	Expires string `bson:"omitempty"`
}
