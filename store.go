package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"
)

var Store *EventStore

func init() {
	Store = &EventStore{
		Events:    make(map[string]Event),
		Responses: make(map[string][]Response),
		mtx:       &sync.Mutex{},
	}
	go SaveData()

	f, err := os.Open("events.json")
	if err != nil {
		return
	}
	json.NewDecoder(f).Decode(&Store.Events)
	fmt.Println(Store.Events)
	r, err := os.Open("responses.json")
	if err != nil {
		return
	}
	json.NewDecoder(r).Decode(&Store.Responses)

}

func SaveData() {
	for tick := range time.Tick(time.Minute) {
		fmt.Println("Save Data:", tick)
		f, err := os.Create("events.json")
		if err != nil {
			return
		}
		json.NewEncoder(f).Encode(Store.Events)
		r, err := os.Create("responses.json")
		if err != nil {
			return
		}
		json.NewEncoder(r).Encode(Store.Responses)
	}
}

type EventStore struct {
	Events    map[string]Event
	Responses map[string][]Response
	mtx       *sync.Mutex
}

func (s *EventStore) NewEvent(e Event) string {
	s.mtx.Lock()
	defer s.mtx.Unlock()

	k := fmt.Sprint(rand.Int63())
	s.Events[k] = e
	fmt.Println(s.Events)
	return k
}

func (s *EventStore) NewResponse(k string, r Response) {
	s.mtx.Lock()
	s.Responses[k] = append(s.Responses[k], r)
	s.mtx.Unlock()
}

func (s *EventStore) GetEvent(k string) (Event, bool) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	e, b := s.Events[k]
	return e, b
}

func (s *EventStore) GetResponses(k string) []Response {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.Responses[k]
}

type Event struct {
	Name         string     `json:"name"`
	EventType    string     `json:"event_type"`
	Date         string     `json:"date"`
	StartTime    string     `json:"startTime"`
	EndTime      string     `json:"endTime"`
	Instructions string     `json:"instructions"`
	Questions    []Question `json:"questions"`
}
type Question struct {
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Required bool     `json:"required"`
	Key      string   `json:"key"`
	Choices  []string `json:"choices"`
}
type Response struct {
	Created time.Time `json:"create"`
	Answers []string  `json:"answers"`
}
