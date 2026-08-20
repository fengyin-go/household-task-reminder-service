# Reminder cancellation is persisted

## Bug

A reminder scheduling request that has already been cancelled is accepted as if it were still active.

## Trigger

Cancel a request before scheduling its reminder, then submit that cancelled job through the reminder scheduler.

## Error

The targeted test reports `cancelled request error = <nil>, want context cancellation`; the cancelled job is also stored instead of being discarded.
