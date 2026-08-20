# Tag summary retains mutable request state

## Bug

The summary pipeline stores references to a caller-owned tag slice.

## Trigger

Refresh a summary from one tag, then reuse the same backing array for a later tag update.

## Error

The target test reports that the earlier summary observed the later tag mutation.
