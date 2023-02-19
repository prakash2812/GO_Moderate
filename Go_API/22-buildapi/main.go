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

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// facke DB
var courses []Course

// helper -- middleware
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}
func main() {
	fmt.Println("build api concept")
	r := mux.NewRouter()

	// seeding data
	courses = append(courses, Course{"1", "react", 32, &Author{"arjun", "google"}})
	courses = append(courses, Course{"2", "nextjs", 12, &Author{"prakash", "apple"}})
	courses = append(courses, Course{"3", "GO", 2812, &Author{"krishna", "amazon"}})
	courses = append(courses, Course{"4", "Nodejs", 28, &Author{"shiva", "microsoft"}})
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUST")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers

// serve home
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api build</h1>"))
}
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get  one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course using post")
	w.Header().Set("Content-Type", "application/json")

	var course Course
	// if body is nil
	if r.Body == nil {
		json.NewEncoder(w).Encode("Forgot to send body")
	}
	// if body is {}
	json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("Forgot to send properties inside body")
		return
	}

	// TODO: if title or coursename is same then response its akreay exist no ned to create duplicate data
	//generate unique id
	// append cour in to courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update course using put")
	w.Header().Set("Content-Type", "application/json")
	// take id then iterate couse
	params := mux.Vars(r)

	// pull out that and append the new user given body in to the data courses
	for index, course := range courses {
		if params["id"] == course.CourseId {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: send a response when id is invalid and body is not valid so on conditions
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete course using delete method")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, course := range courses {
		if params["id"] == course.CourseId {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm after delete or any resposne
			break
		}
	}
}
