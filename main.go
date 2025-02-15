package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

var NUMBER_OF_EVENTS = []int{1000000, 10000, 5000, 3000}
var EVENT_TYPES = []Event{
	{name: "click", description: "User clicked on a button"},
	{name: "view", description: "User viewed a page"},
	{name: "purchase", description: "User purchased an item"},
	{name: "search", description: "User searched for an item"},
}
var LOG_FORMAT = "%s | EVENT_TYPE=%s | DESCRIPTION=\"%s\" | USER_ID=%d\n"

func main() {

	var logGenerator = &cobra.Command{
		Use:   "logg",
		Short: "Generate random log files",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	logGenerator.AddCommand(Generate())

	if err := logGenerator.Execute(); err != nil {
		log.Println(err)
	}
}

func Generate() *cobra.Command {
	var dirLocation string
	var numberofFiles int

	generateLogs := &cobra.Command{
		Use:   "generate",
		Short: "Generate logs",
		Run: func(cmd *cobra.Command, args []string) {
			var wg sync.WaitGroup

			for i := 0; i < numberofFiles; i++ {
				wg.Add(1)
				go generateFile(dirLocation, i, &wg)
			}

			wg.Wait()
		},
	}

	generateLogs.Flags().StringVarP(&dirLocation, "location", "l", "", "Generated file output location")
	generateLogs.Flags().IntVarP(&numberofFiles, "count", "c", 5, "Number of files")

	generateLogs.MarkFlagRequired("location")

	return generateLogs
}

func generateFile(location string, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := os.ReadDir(location)
	if err != nil {
		err = os.MkdirAll(location, os.ModeAppend)
		if err != nil {
			panic("unable to create directory")
		}
	}

	fileName := fmt.Sprintf("%s/file%d.log", location, id+1)
	file, err := os.Create(fileName)

	randomNumber := rand.Intn(len(NUMBER_OF_EVENTS))
	totalEvents := NUMBER_OF_EVENTS[randomNumber]

	for j := 1; j < totalEvents; j++ {
		if err != nil {
			log.Fatal(err)
		}
		eventType := EVENT_TYPES[rand.Intn(len(EVENT_TYPES))]
		log := fmt.Sprintf(LOG_FORMAT, time.Now().Local().String(), eventType.name, eventType.description, id)
		file.WriteString(log)
	}
}

type Event struct {
	name        string
	description string
}
