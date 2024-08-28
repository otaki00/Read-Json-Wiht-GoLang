package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var FileLocation string = "./data.json"

func main() {
	data, err := os.ReadFile(FileLocation)
	if err != nil {
		fmt.Println(err)
	}

	var usersWrapper UsersWrapper

	err = json.Unmarshal(data, &usersWrapper)
	if err != nil {
		fmt.Println(err)
	}

	users := usersWrapper.Users

	printDataInFormatedWay(users, FileLocation)

}

type Address struct {
	City       string 
	Street     string 
	PostalCode string 
	Country    string 
}

type SocialMedia struct {
	LinkedIn string 
	Twitter  string 
}

type ContactInfo struct {
	Email       string      
	Phone       string      
	SocialMedia SocialMedia 
}

type Education struct {
	Degree         string 
	Major          string 
	University     string 
	GraduationYear int    
}

type Job struct {
	Title            string   
	Company          string   
	Salary           float64  
	Department       string   
	StartYear        int      
	Responsibilities []string 
}

type EmploymentHistory struct {
	Company          string   
	Title            string   
	Duration         string   
	Responsibilities []string 
}

type User struct {
	Name              string              
	Age               int                 
	Address           Address             
	ContactInfo       ContactInfo         
	Hobbies           []string            
	Education         Education         
	Job               Job               
	EmploymentHistory []EmploymentHistory 
}

type UsersWrapper struct {
	Users []User `json:"users"`
}



func printDataInFormatedWay(users []User, fileLoc string) {
	fmt.Printf("File %s contains the following data: \n", fileLoc)

	fields := []string{"Name", "Age", "Address", "ContactInfo", "Hobbies", "Education", "Job", "EmploymentHistory"}

	fmt.Println("==============================================")
	
	// show table in console
	for _, user := range users {
		fmt.Println("==============================================")
		for _, field := range fields {
			fmt.Printf("%s: ", field)
			switch field {
			case "Name":
				fmt.Println(user.Name)
			case "Age":
				fmt.Println(user.Age)
			case "Address":
				fmt.Println(user.Address)
			case "ContactInfo":
				fmt.Println(user.ContactInfo)
			case "Hobbies":
				fmt.Println(user.Hobbies)
			case "Education":
				fmt.Println(user.Education)
			case "Job":
				fmt.Println(user.Job)
			case "EmploymentHistory":
				fmt.Println(user.EmploymentHistory)
			}
		}
		fmt.Println("==============================================")
	}
}