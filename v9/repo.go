package main

import "fmt"

var currentId int

var todos ContainerRequests

// Give us some seed data
func init() {
	RepoCreateRequest(ContainerRequest{Name: "Write presentation"})
	RepoCreateRequest(ContainerRequest{Name: "Host meetup"})
}

func RepoFindRequest(id int) ContainerRequest {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// return empty ContainerRequest if not found
	return ContainerRequest{}
}

//this is bad, I don't think it passes race condtions
func RepoCreateRequest(t ContainerRequest) ContainerRequest {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoDestroyRequest(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find ContainerRequest with id of %d to delete", id)
}
