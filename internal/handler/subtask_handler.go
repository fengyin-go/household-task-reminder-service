package handler

import (
	"net/http"

	"todolist/internal/model"
	"todolist/pkg/httpx"
)

func (s *Server) registerSubtaskRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/subtasks", s.createSubtask)
	mux.HandleFunc("GET /api/subtasks", s.listSubtasks)
	mux.HandleFunc("POST /api/subtasks/{id}/toggle", s.toggleSubtask)
	mux.HandleFunc("DELETE /api/subtasks/{id}", s.deleteSubtask)
	mux.HandleFunc("GET /api/tasks/{id}/subtasks/stats", s.subtaskStats)
}

type subtaskRequest struct {
	TaskID string `json:"task_id"`
	Title  string `json:"title"`
}

func (s *Server) createSubtask(w http.ResponseWriter, r *http.Request) {
	var req subtaskRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	st, err := s.svc.CreateSubtask(model.Subtask{TaskID: req.TaskID, Title: req.Title})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, st)
}

func (s *Server) listSubtasks(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.SubtaskFilter{TaskID: r.URL.Query().Get("task_id")}
	if v := r.URL.Query().Get("done"); v != "" {
		done := v == "true"
		filter.Done = &done
	}
	items, total, err := s.svc.ListSubtasks(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) toggleSubtask(w http.ResponseWriter, r *http.Request) {
	st, err := s.svc.ToggleSubtask(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, st)
}

func (s *Server) deleteSubtask(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteSubtask(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}

func (s *Server) subtaskStats(w http.ResponseWriter, r *http.Request) {
	stats, err := s.svc.GetSubtaskStats(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, stats)
}
