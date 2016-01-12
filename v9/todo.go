package main

import "time"

type ContainerRequest struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
	Image     string    `json:"image"`
	Command   string    `json:"command"`
	Cpus      string    `json:"cpus"`
	Mem       string    `json:"mem"`
	Status    string    `json:"status"`
}

type ContainerRequests []ContainerRequest
