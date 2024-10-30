package main

import "time"

type CrewMember struct {
	Crewmemberid int32  `json:"id"`
	Name         string `json:"name"`
	Role         string `json:"role"`
}

type Ship struct {
	Shipnr       int32        `json:"id"`
	Name         string       `json:"name"`
	Owner        int32        `json:"owner"`
	ShipType     string       `json:"type"`
	Image        string       `json:"image"`
	Currentvalue string       `json:"currentValue"`
	Year         string       `json:"year"`
	Crewmembers  []CrewMember `json:"crewmembers"`
}

type Benchmark struct {
	Size      int32         `json:"size"`
	Objective string        `json:"objective"`
	Mysql     time.Duration `json:"mysql"`
	Mongodb   time.Duration `json:"mongodb"`
}
