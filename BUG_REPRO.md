# Reused reminder payload overwrites earlier reminders

## Bug

Payload bytes from the buffer pool are retained after the pool has reclaimed them.

## Trigger

Queue a first reminder, queue a second reminder, then inspect both the first returned payload and its cached value.

## Error

The target test reports that the first payload was overwritten with the second reminder text.
