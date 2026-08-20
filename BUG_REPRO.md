# Default preferences crash and bypass validation

## Bug

The empty preference path exposes both an uninitialized labels map and a typed-nil checker.

## Trigger

Load defaults, save a delivery label, then validate an empty reminder using the default checker path.

## Error

The target test reports that the default write panicked or invalid reminder input bypassed validation.
