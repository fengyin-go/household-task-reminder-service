# Cancellation loss persists a reminder

## Bug

The reminder request chain replaces the caller context and keeps request state beyond one call.

## Trigger

Cancel one create request, then immediately issue another create request using a fresh context.

## Error

The target test reports a persisted cancelled request or a fresh request that inherited stale cancellation.
