package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"todo/structures"

	"github.com/gorilla/mux"
)

type requestBody struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}
type errorPrint struct {
	Message string
	Time    time.Time
}

type HTTPHandlers struct {
	todoList *structures.TodoList
}

func NewHttpHandlers(todoList *structures.TodoList) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func sendError(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	response := errorPrint{
		Message: err.Error(),
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

func (h *HTTPHandlers) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var req requestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
	tasks, err := h.todoList.CreateTask(req.Author, req.Text)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func (h *HTTPHandlers) ListAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tasks := h.todoList.ListTasks()
	json.NewEncoder(w).Encode(tasks)
}
func (h *HTTPHandlers) GetTaskById(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		sendError(w, fmt.Errorf("invalid id: %v", idStr), http.StatusBadRequest)
		return
	}
	task, err := h.todoList.GetTaskById(id)
	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}
	b, err := json.MarshalIndent(task, "", "     ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failder to write in write response!")
		return
	}
}
