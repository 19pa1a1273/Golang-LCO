package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var courses []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == "" // you can create more such methods
	return c.CourseName == ""
}

func main() {

	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	// seeding

	courses = append(courses, Course{CourseId: "2", CourseName: "React JS", CoursePrice: 299,
		Author: &Author{Fullname: "Bhanu Lingolu", Website: "lco.in"}})

	courses = append(courses, Course{CourseId: "6", CourseName: "MERN", CoursePrice: 199,
		Author: &Author{Fullname: "Bhanu Lingolu", Website: "go.dev"}})

	// routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) { // Check HERE...
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request

	params := mux.Vars(r)

	// fmt.Println(params)

	// loop through courses, find matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			// fmt.Println(courses)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id") // Exercise --> we have params write id value here
	// return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if: body is empty

	// var course Course
	// _ = json.NewDecoder(r.Body).Decode(&course)

	// fmt.Println(r.Body)
	// if r.Body == nil { // Check HERE...
	// 	json.NewEncoder(w).Encode("Please send some data")
	// 	return
	// }

	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		fmt.Println(err)
		json.NewEncoder(w).Encode("Please Send Some Data")
		return
	}

	// what about - {}

	// var course Course
	// _ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// TODO: check only if title is Duplicate
	// loop, title matches with course.coursename, JSON

	//generate unique id, string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	// return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req

	params := mux.Vars(r)

	// loop, id, remove, add with my id (same id)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	//TODO: Send a response when id is not found

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {

	// Exercise: Delete All Courses

	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// Exercise: Send json response(We are able to delete this successfully)
			// TODO: Send a Confirm or deny Response
			break
		}
	}

}
