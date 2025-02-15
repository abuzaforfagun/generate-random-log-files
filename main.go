package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var NUMBER_OF_USERS = 5
var NUMBER_OF_EVENTS = []int{1000000, 10000, 5000, 3000}
var EVENT_TYPES = []Event{
	{name: "click", description: "User clicked on a button"},
	{name: "view", description: "User viewed a page"},
	{name: "purchase", description: "User purchased an item"},
	{name: "search", description: "User searched for an item"},
}
var LOG_FORMAT = "%s | EVENT_TYPE=%s | DESCRIPTION=\"%s\" | USER_ID=%d\n"

func main() {

	var wg sync.WaitGroup

	for i := 0; i < NUMBER_OF_USERS; i++ {
		wg.Add(1)
		go generateFile(i, &wg)
	}

	wg.Wait()
}

func generateFile(userId int, wg *sync.WaitGroup) {
	defer wg.Done()
	fileName := fmt.Sprintf("out/user%d.log", userId)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}

	randomNumber := rand.Intn(len(NUMBER_OF_EVENTS))
	totalEvents := NUMBER_OF_EVENTS[randomNumber]

	for j := 0; j < totalEvents; j++ {
		eventType := EVENT_TYPES[rand.Intn(len(EVENT_TYPES))]
		log := fmt.Sprintf(LOG_FORMAT, time.Now().Local().String(), eventType.name, eventType.description, userId)
		file.WriteString(log)
	}
}

type Event struct {
	name        string
	description string
}
