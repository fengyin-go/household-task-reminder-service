# Failed batch item leaves reminder dispatch hanging

## Bug

The failed producer path does not close its channel, leaving the consumer and coordinator waiting forever.

## Trigger

Dispatch a two-item batch with one failed item and one normal item.

## Error

The target test reports a batch dispatch timeout and the normal reminder result is unavailable.
