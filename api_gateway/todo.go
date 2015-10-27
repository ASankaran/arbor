package main

import (
	"net/http"
)

// type Todo struct {
// 	Id        int       `json:"id"`
// 	Name      string    `json:"name"`
// 	Completed bool      `json:"completed"`
// 	Due       time.Time `json:"due"`
// }
//
// type Todos []Todo

//Location
var TodoURL string = "http://localhost:8080"

//API Interface
var todoRoutes = Routes{
	Route{
		"TodoIndex",
		"GET",
		"/todo",
		todoIndex,
	},
	Route{
		"TodoAll",
		"GET",
		"/todos",
		TodoAll,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
}

//Route handler
func todoIndex(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Welcome!\n")
	//TODO: THIS API CALL IS A SPECIAL CASE
	GetHandler(w, TodoURL, r)
}

func TodoAll(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(todos); err != nil {
	// 	panic(err)
	// }
	//GetHandler(TodoURL + r.URL)
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// var todoId int
	// var err error
	// if todoId, err = strconv.Atoi(vars["todoId"]); err != nil {
	// 	panic(err)
	// }
	// todo := RepoFindTodo(todoId)
	// if todo.Id > 0 {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(http.StatusOK)
	// 	if err := json.NewEncoder(w).Encode(todo); err != nil {
	// 		panic(err)
	// 	}
	// 	return
	// }
	//
	// // If we didn't find it, 404
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusNotFound)
	// if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
	// 	panic(err)
	// }

}

/*
Test with this curl command:

curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos

*/
func TodoCreate(w http.ResponseWriter, r *http.Request) {
	// var todo Todo
	// body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	// if err != nil {
	// 	panic(err)
	// }
	// if err := r.Body.Close(); err != nil {
	// 	panic(err)
	// }
	// if err := json.Unmarshal(body, &todo); err != nil {
	// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.WriteHeader(422) // unprocessable entity
	// 	if err := json.NewEncoder(w).Encode(err); err != nil {
	// 		panic(err)
	// 	}
	// }
	//
	// t := RepoCreateTodo(todo)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(t); err != nil {
	// 	panic(err)
	// }
}
