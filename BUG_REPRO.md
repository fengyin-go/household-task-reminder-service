# Reminder retry duplicates a committed delivery

## Bug

When the gateway fails after accepting a reminder, the application loses the fact that the delivery was committed and retries it.

## Trigger

Submit a reminder using the repository failure mode that returns an error immediately after its first commit.

## Error

The target test reports that the partial commit was hidden; the duplicate side-effect count becomes two.
