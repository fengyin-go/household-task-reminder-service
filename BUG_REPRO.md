# Stale reminder events reopen completed tasks

## Bug

An older reminder event can overwrite a newer completed task view.

## Trigger

Apply a version 2 completed event, then replay a version 1 in-progress event for the same task.

## Error

The targeted test reports `stale event reopened task: state=doing completed=0`.
