# gerr-cosmos

## What?

A set of concrete error types to use as a basis for errors across your cosmos sdk application.

## Why?

In modern idiomatic golang it is normal to almost mimic a stack trace by wrapping errors at appropriate moments. Errors should be handled at caller level using `errors.Is()` and related utilities from the standard library.

If done wrong, this leads to a proliferation of error types, and bug prone code.

This library follows Google [API design guidelines](https://cloud.google.com/apis/design/errors#handling_errors)

> Individual APIs must avoid defining additional error codes, since developers are very unlikely to write logic to handle a large number of error codes. For reference, handling an average of three error codes per API call would mean most application logic would just be for error handling, which would not be a good developer experience.

The philosophy is to use a small set of well thought out errors as first class citizens in your program.

This library extends the selection of Google errors with a few ubiquitous errors encountered in Cosmos/Blockchain applications.

## Contributing

Suggestions, PRs, pointing-out-of-mistakes, all welcome 🙌

## Basis

This lib is actually a wrapper around a more [general purpose library](https://github.com/danwt/gerr).
