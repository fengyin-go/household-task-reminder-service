package handler

import (
	"net/http"

	"todolist/internal/model"
	"todolist/pkg/httpx"
)

func (s *Server) registerTaskListRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/task-lists", s.createTaskList)
	mux.HandleFunc("GET /api/task-lists", s.listTaskLists)
	mux.HandleFunc("GET /api/task-lists/{id}", s.getTaskList)
	mux.HandleFunc("PUT /api/task-lists/{id}", s.updateTaskList)
	mux.HandleFunc("DELETE /api/task-lists/{id}", s.deleteTaskList)
	mux.HandleFunc("GET /api/task-lists/{id}/stats", s.getTaskListStats)
}

type createTaskListRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *Server) createTaskList(w http.ResponseWriter, r *http.Request) {
	var req createTaskListRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	tl, err := s.svc.CreateTaskList(model.TaskList{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, tl)
}

func (s *Server) listTaskLists(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.TaskListFilter{
		Keyword: r.URL.Query().Get("keyword"),
	}
	items, total, err := s.svc.ListTaskLists(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getTaskList(w http.ResponseWriter, r *http.Request) {
	tl, err := s.svc.GetTaskList(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, tl)
}

func (s *Server) updateTaskList(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	tl, err := s.svc.UpdateTaskList(r.PathValue("id"), model.TaskList{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, tl)
}

func (s *Server) deleteTaskList(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteTaskList(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

func (s *Server) getTaskListStats(w http.ResponseWriter, r *http.Request) {
	stats, err := s.svc.GetTaskStatsByTaskList(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, stats)
}
