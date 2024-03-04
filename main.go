package main

import (
	"fmt"
	"os"
)

type Friend struct {
	Id string
	Name string
	Address string
	Occupation string
	Reason string
}

var friends []Friend = []Friend{
	{
		Id: "F1",
		Name: "John Doe",
		Address: "123 Main Street, Anytown, USA",
		Occupation: "Software Engineer",
		Reason: "I love the simplicity and efficiency of Golang, and I believe it will help me build scalable and reliable applications.",
	},
	{
		Id: "F2",
		Name: "Jane Smith",
		Address: "456 Elm Avenue, Othertown, USA",
		Occupation: "Web Developer",
		Reason: "I'm interested in Golang's concurrency model and its ability to handle high-performance tasks.",
	},
	{
		Id: "F3",
		Name: "Alice Johnson",
		Address: "789 Oak Drive, Anothercity, USA",
		Occupation: "Data Scientist",
		Reason: "I want to leverage Golang's speed and simplicity for building data processing pipelines.",
	},
	{
		Id: "F4",
		Name: "Bob Williams",
		Address: "101 Pine Road, Somewhere, USA",
		Occupation: "Systems Administrator",
		Reason: "Golang's built-in testing support and cross-platform compatibility make it ideal for building robust and scalable infrastructure.",
	},	
	{
		Id: "F5",
		Name: "Emily Brown",
		Address: "111 Cedar Lane, Nowhere, USA",
		Occupation: "Mobile App Developer",
		Reason: "I see a growing demand for Golang in mobile development, especially for building backend services for mobile apps.",
	},	
	{
		Id: "F6",
		Name: "Michael Wilson",
		Address: "222 Maple Street, Anywhere, USA",
		Occupation: "DevOps Engineer",
		Reason: "Golang's lightweight goroutines and strong standard library make it a great choice for building scalable and reliable microservices.",
	},	
	{
		Id: "F7",
		Name: "Sarah Martinez",
		Address: "333 Birch Avenue, Elsewhere, USA",
		Occupation: "Security Analyst",
		Reason: "I'm impressed by Golang's focus on security and its ability to write secure code with minimal vulnerabilities.",
	},	
	{
		Id: "F8",
		Name: "David Garcia",
		Address: "444 Ash Street, Anytown, USA",
		Occupation: "Backend Developer",
		Reason: "I believe Golang's simplicity and performance will help me develop high-performance backend services for web applications.",
	},	
	{
		Id: "F9",
		Name: "Jessica Rodriguez",
		Address: "555 Pine Lane, Nowhere, USA",
		Occupation: "Full Stack Developer",
		Reason: "I'm excited to learn Golang for its powerful standard library and its suitability for building both backend and frontend components of web applications.",
	},	
	{
		Id: "F10",
		Name: "Ryan Thompson",
		Address: "666 Elm Road, Somewhere, USA",
		Occupation: "Cloud Engineer",
		Reason: "As a cloud engineer, I see Golang as a valuable tool for building scalable and efficient cloud-native applications.",
	},
}

func searchFriendById(id string) {
	for _, v := range friends {
		if v.Id == id {
			fmt.Println("Data Teman Berdasarkan Id")
			fmt.Println("-------------------------")
			fmt.Printf("Id: %s\n", v.Id)
			fmt.Printf("Name: %s\n", v.Name)
			fmt.Printf("Address: %s\n", v.Address)
			fmt.Printf("Occupation: %s\n", v.Occupation)
			fmt.Printf("Reason: %s\n", v.Reason)
			return
		}
	}

	fmt.Printf("Friend with Id %s can not be found\n", id)
}

func main() {
	var arguments []string = os.Args
	
	if len(arguments) < 2 {
		fmt.Printf("Masukkan id teman terlebih dahulu!\n")
		return;
	} else if len(arguments) > 2 {
		fmt.Printf("Cukup masukkan 1 id teman yang ingin dicari!\n")
		return;
	} else {
		searchFriendById(arguments[1])
	}
}