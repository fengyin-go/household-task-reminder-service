# Failed reminder commits transaction state

## Bug

The send-failure path commits the reminder instead of aborting the transaction lifecycle.

## Trigger

Send a reminder with the deterministic failure switch enabled and then inspect transaction and saved state.

## Error

The target test reports that a failed reminder was committed or transaction state was leaked.
