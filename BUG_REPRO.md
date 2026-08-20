# Partial task-list cleanup retries duplicate work

## Bug

A cleanup that already ran is treated as an ordinary retryable delete failure.

## Trigger

Delete a task list in the failure mode that reports an error after the cleanup side effect.

## Error

The target test reports a hidden partial cleanup and observes the cleanup operation run twice.
