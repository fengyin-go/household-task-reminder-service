package service

import "todolist/internal/transaction"

type ReminderTransactionService struct{ runner *transaction.Runner }

func NewReminderTransactionService(runner *transaction.Runner) *ReminderTransactionService {
	return &ReminderTransactionService{runner: runner}
}

func (s *ReminderTransactionService) Send(id string, fail bool) error {
	return s.runner.Deliver(&transaction.Reminder{ID: id}, fail)
}
