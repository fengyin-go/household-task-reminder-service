package handler

import (
	"net/http"
	"time"

	"todolist/internal/model"
	"todolist/pkg/httpx"
)

func (s *Server) registerReminderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/reminders", s.createReminder)
	mux.HandleFunc("GET /api/reminders", s.listReminders)
	mux.HandleFunc("GET /api/reminders/{id}", s.getReminder)
	mux.HandleFunc("PUT /api/reminders/{id}", s.updateReminder)
	mux.HandleFunc("DELETE /api/reminders/{id}", s.deleteReminder)
}

type createReminderRequest struct {
	TaskID   string    `json:"task_id"`
	RemindAt time.Time `json:"remind_at"`
	Message  string    `json:"message"`
}

func (s *Server) createReminder(w http.ResponseWriter, r *http.Request) {
	var req createReminderRequest
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	rem, err := s.svc.CreateReminder(model.Reminder{
		TaskID:   req.TaskID,
		RemindAt: req.RemindAt,
		Message:  req.Message,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.Created(w, rem)
}

func (s *Server) listReminders(w http.ResponseWriter, r *http.Request) {
	pp := httpx.ParsePagination(r, 20, s.maxPageSize())
	filter := model.ReminderFilter{
		TaskID: r.URL.Query().Get("task_id"),
	}
	items, total, err := s.svc.ListReminders(filter, pp.Page, pp.Size)
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, httpx.PageResult{
		Items:      items,
		Pagination: httpx.Pagination{Page: pp.Page, Size: pp.Size, Total: total},
	})
}

func (s *Server) getReminder(w http.ResponseWriter, r *http.Request) {
	rem, err := s.svc.GetReminder(r.PathValue("id"))
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, rem)
}

func (s *Server) updateReminder(w http.ResponseWriter, r *http.Request) {
	var req struct {
		TaskID   string    `json:"task_id"`
		RemindAt time.Time `json:"remind_at"`
		Message  string    `json:"message"`
	}
	if err := httpx.Decode(r, &req); err != nil {
		httpx.BadRequest(w, "请求体解析失败: "+err.Error())
		return
	}
	rem, err := s.svc.UpdateReminder(r.PathValue("id"), model.Reminder{
		TaskID:   req.TaskID,
		RemindAt: req.RemindAt,
		Message:  req.Message,
	})
	if err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.OK(w, rem)
}

func (s *Server) deleteReminder(w http.ResponseWriter, r *http.Request) {
	if err := s.svc.DeleteReminder(r.PathValue("id")); err != nil {
		writeServiceError(w, err)
		return
	}
	httpx.NoContent(w)
}
