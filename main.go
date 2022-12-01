package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
	"webpersonal/connection"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	connection.DatabaseConnect()

	route.PathPrefix("/public").Handler(http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/formproject", formproject).Methods("GET")
	route.HandleFunc("/detail/{id}", detail).Methods("GET")
	route.HandleFunc("/add-project", addproject).Methods("POST")
	route.HandleFunc("/add-project", formAddProject).Methods("GET")
	route.HandleFunc("/delete-project/{index}", deleteProject).Methods("GET")
	route.HandleFunc("/edit-project/{index}", formEditProject).Methods("GET")
	route.HandleFunc("/edit-project/{index}", editProject).Methods("POST")
	

	fmt.Println("Server running on port 5000")
	http.ListenAndServe("localhost:5000", route)
}


type Blog struct {
	Id 			int
	Title   	string
	StartDate 	time.Time
	EndDate		time.Time
	Description string
	check1 		string
	check2		string
	check3		string
	check4		string
}

type Data struct {
	DataProject []Blog
}

// var blogs = 
var blogs = []Blog{
	{
		Title:   "Dumbways",
		// StartDate: "23 Jan 2022",
		// EndDate: "23 Mart 2022",
		Description: "Hallo semuanya",
		check1: "/public/icon/nodejs.png",
		check2: "/public/icon/reactjs.png",
		check3: "/public/icon/nextjs.png",
		check4: "/public/icon/typescript.png",
	},
}

func editProject(w http.ResponseWriter, r *http.Request){
	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	
	title := r.PostForm.Get("title")
	// starDate := r.PostForm.Get("startDate")
	// endDate := r.PostForm.Get("endDate")
	description := r.PostForm.Get("description")
	check1 := r.PostForm.Get("check1")

	var editBlog = Blog{
		Title:   title,
		// StartDate: starDate,
		// EndDate: endDate,
		Description: description,
		check1: check1,
	}

	blogs[index] = editBlog
	

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
	


func formEditProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/editproject.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	index, _ := strconv.Atoi(mux.Vars(r)["index"])

	var projectEdit = Blog{}

	for i, data := range blogs {
		if i == index {
		 projectEdit = Blog{
				Id: i,
				Title:   data.Title,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Description: data.Description,
			}
		}
	}

	fmt.Println(projectEdit)

	dataNew := map[string]interface{}{
		"Blog": projectEdit,
	}

	tmpt.Execute(w, dataNew)
}


func addproject(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	// starDate := r.PostForm.Get("startDate")
	// endDate := r.PostForm.Get("endDate")
	description := r.PostForm.Get("description")
	check1 := r.PostForm.Get("node")
	check2 := r.PostForm.Get("react")
	check3 := r.PostForm.Get("next")
	check4 := r.PostForm.Get("typescript")

	var pathNode = ""
	var pathReact = ""
	var pathNext = ""
	var pathTypescript = ""

	if check1 == "true" {
		pathNode = "/public/icon/nodejs.png"
	} else {
		pathNode = "d-none"
	}

	if check2 == "true" {
		pathReact = "/public/icon/react.png"
	} else {
		pathReact = "d-none"
	}

	if check3 == "true" {
		pathNext = "/public/icon/nextjs.png"
	} else {
		pathNext = "d-none"
	}

	if check4 == "true" {
		pathTypescript = "/public/icon/typescript.png"
	} else {
		pathTypescript = "d-none"
	}

	var newBlog = Blog{
		Title:   title,
		// StartDate: starDate,
		// EndDate: endDate,
		Description: description,
		check1: pathNode,
		check2: pathReact,
		check3: pathNext,
		check4: pathTypescript,
	}

	// blogs.push(newBlog)
	blogs = append(blogs, newBlog)
	fmt.Println(newBlog)
 
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func formAddProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/formproject.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

func home(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil{
		w.Write([]byte("Mesage :"+ err.Error()))
		return
	}


	dataBlog, errQuery := connection.Conn.Query(context.Background(), "SELECT id, title, description, start_date, end_date FROM tb_project")

	if errQuery != nil {
		fmt.Println("Message : " + errQuery.Error())
		return
	}  

	var result []Blog

	for dataBlog.Next() {
		var eac = Blog{}

		err := dataBlog.Scan(&eac.Id, &eac.Title, &eac.Description, &eac.StartDate, &eac.EndDate)
		if err != nil {
			fmt.Println("Message : " + err.Error())
			return
		}  

		result = append(result, eac)
	}

	fmt.Println(result)

	resData := map[string]interface{}{
	"Blogs": result,
	}

	tmpt.Execute(w, resData)
}
 
func contact(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/form.html")

	if err != nil{
		w.Write([]byte("Mesage :"+ err.Error()))
		return
	}
	tmpt.Execute(w, nil)
}

func formproject(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/formproject.html")

	if err != nil{
		w.Write([]byte("Mesage :"+ err.Error()))
		return
	}
	tmpt.Execute(w, nil)
}

func detail(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Contain type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/detail.html")

	if err != nil{
		w.Write([]byte("Mesage :"+ err.Error()))
		return
	}
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	
	var BlogDetail = Blog{}

	for index, data := range blogs {
		if index == id {
			BlogDetail = Blog{
				Title:   data.Title,
				StartDate: data.StartDate,
				EndDate: data.EndDate,
				Description: data.Description,
			}
		}
	}

	fmt.Println(BlogDetail)

	dataDetail := map[string]interface{}{
		"Blog": BlogDetail,
	}

	tmpt.Execute(w, dataDetail)
}

func deleteProject(w http.ResponseWriter, r *http.Request) {

	index, _ := strconv.Atoi(mux.Vars(r)["index"])
	// fmt.Println(index)

	blogs = append(blogs[:index], blogs[index+1:]...)

	http.Redirect(w, r, "/", http.StatusFound)
}

