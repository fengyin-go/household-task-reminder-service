package service_test

import (
	"testing"

	"todolist/internal/service"
	"todolist/internal/transaction"
)

func TestFailedReminderDoesNotLeakTransactionOrCommit(t *testing.T) {
	repo := transaction.NewRepository()
	svc := service.NewReminderTransactionService(transaction.NewRunner(repo))
	if err := svc.Send("r1", true); err == nil {
		t.Fatal("failed delivery returned success")
	}
	if repo.Active() != 0 || repo.Saved("r1") {
		t.Fatalf("failed reminder leaked transaction or committed: active=%d saved=%v", repo.Active(), repo.Saved("r1"))
	}
}
