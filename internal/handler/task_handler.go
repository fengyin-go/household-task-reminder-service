package handler

import (
	"net/http"
	"time"

	"todolist/internal/model"
	"todolist/pkg/httpx"
)

func (s *Server) registerTaskRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/tasks", s.createTask)
	mux.HandleFunc("GET /api/tasks", s.listTasks)
	mux.HandleFunc("GET /api/tasks/{id}", s.getTask)
	mux.HandleFunc("PUT /api/tasks/{id}", s.updateTask)
	mux.HandleFunc("DELETE /api/tasks/{id}", s.deleteTask)
	mux.HandleFunc("POST /api/tasks/{id}/transition", s.transitionTask)
}

type createTaskRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	TaskListID  string    `json:"task_list_id"`
	Priority    int       `json:"priority"`
	DueDate     time.Time `json:"due_date"`
}

func (s *Server) createTask(w http.ResponseWriter, r *http.Request) {
	var req createTaskRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	t, err := s.svc.CreateTask(model.Task{
		Title:       req.Title,
		Description: req.Description,
		TaskListID:  req.TaskListID,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, t)
}

func (s *Server) listTasks(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.TaskFilter{
		TaskListID: r.URL.Query().Get("task_list_id"),
		Status:     r.URL.Query().Get("status"),
		Keyword:    r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListTasks(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getTask(w http.ResponseWriter, r *http.Request) {
	t, err := s.svc.GetTask(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, t)
}

func (s *Server) updateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string    `json:"title"`
		Description string    `json:"description"`
		TaskListID  string    `json:"task_list_id"`
		Priority    int       `json:"priority"`
		DueDate     time.Time `json:"due_date"`
	}
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	t, err := s.svc.UpdateTask(r.PathValue("id"), model.Task{
		Title:       req.Title,
		Description: req.Description,
		TaskListID:  req.TaskListID,
		Priority:    req.Priority,
		DueDate:     req.DueDate,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, t)
}

func (s *Server) deleteTask(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteTask(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

type transitionTaskRequest struct {
	Status string `json:"status"`
}

func (s *Server) transitionTask(w http.ResponseWriter, r *http.Request) {
	var req transitionTaskRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	t, err := s.svc.TransitionTaskStatus(r.PathValue("id"), req.Status)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, t)
}
