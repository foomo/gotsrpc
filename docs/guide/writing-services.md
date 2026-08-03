# Writing Services

A gotsrpc service is a plain Go **interface** together with a struct that
implements it. The interface is the contract gotsrpc reads to generate proxies
and clients; the struct is your actual implementation.

```go
package service

// The contract gotsrpc generates code from.
type Service interface {
	Hello(name string) string
}

// Your implementation.
type Handler struct{}

func (h *Handler) Hello(name string) string {
	return name
}
```

The interface name is what you reference in `gotsrpc.yml` under a target's
`services` map (route → service name). See [Configuration](/guide/configuration).

## Serializable arguments and returns

Method arguments and return values must be **serializable** — they are
marshalled over the wire (JSON by default). Supported types include:

- scalars: `bool`, all `int`/`uint` widths, `float32`/`float64`, `string`
- pointers to any supported type (become nullable — see below)
- slices and maps of supported types
- named/defined types (`type IntType int`) and `const` groups (become enums)
- structs of supported fields, including nested structs
- `any` / `interface{}`
- `time.Time` and generics

Struct field names, JSON tags and pointer/`omitempty` combinations control the
generated TypeScript shape and nullability.

## Method signatures

gotsrpc is flexible about which "special" parameters and returns you include.
All of the following are valid.

### Plain values

```go
Bool(v bool) bool
String(v string) string
Struct(v Struct) Struct
Empty()                       // no args, no returns
```

### `context.Context`

Add `context.Context` as the first parameter to propagate cancellation,
deadlines and request-scoped values:

```go
Hello(ctx context.Context, msg string) string
Error(ctx context.Context, msg string) error
```

### Direct HTTP access

Include `http.ResponseWriter` and `*http.Request` when you need the raw request —
for example to read headers (auth tokens), set cookies, or reach
`r.Context()`:

```go
Context(w http.ResponseWriter, r *http.Request)
Error(w http.ResponseWriter, r *http.Request) (e error)
```

These parameters are handled by the proxy and **do not appear** in the generated
client signatures — callers never pass `w`/`r`. Header, cookie and context
handling belong in your server code and in the client [transport](/guide/getting-started#_5-call-it-from-typescript).

### Multiple and named returns, errors

Methods may return multiple values, use named returns, and return `error` or a
**typed** error:

```go
Errors(w http.ResponseWriter, r *http.Request) (e1 error, e2 error)
Scalar(w http.ResponseWriter, r *http.Request) (e *ScalarError)
```

Errors are serialized with their cause chain. You can return the built-in
`error` interface, a custom error type, a scalar error (`type MyError string`
implementing `Error()`), a struct error carrying data, or wrapped errors.

### Pointers → nullable

Pointer arguments and returns map to nullable types on the client. A Go
`*bool` return becomes `Promise<boolean | null>` in TypeScript:

```go
BoolPtr(v bool) *bool
```

```ts
async boolPtr(v: boolean): Promise<boolean | null> { /* ... */ }
```

## What the generated clients look like

**TypeScript** — method names are camelCased and results come back positionally:

```ts
async hello(name: string): Promise<string> {
	return (await this.transport<{0: string}>("Hello", [name]))[0];
}
```

**Go** — regardless of the server signature, every generated client method takes
a leading `ctx context.Context` and returns a trailing `clientErr error`:

```go
type ServiceGoTSRPCClient interface {
	Hello(ctx context.Context, name string) (retHello_0 string, clientErr error)
	BoolPtr(ctx context.Context, v bool) (retBoolPtr_0 *bool, clientErr error)
	Context(ctx context.Context) (clientErr error)
}
```

## Feature examples

The repository's `tests/` directory contains focused, runnable examples for
specific type features. Use them as references when you hit one of these cases:

| Feature | Where |
| --- | --- |
| Enums (named scalar types + `const` groups) | `example/basic/service/vo.go` |
| Nullable / pointer & `omitempty` matrix | `tests/nullable/server/vo.go` |
| Unions (discriminated & registered) | `tests/union/server/vo.go` |
| `time.Time` handling | `tests/time/server/` |
| Generics | `tests/generics/server/`, `tests/nested-generics/server/` |
| Type aliases | `tests/aliases/server/` |
| Errors (scalar, struct, typed, wrapped) | `tests/errors/server/` |
| Context-aware errors | `tests/context/server/` |

## Next steps

- [Configuration](/guide/configuration) — wire your service into `gotsrpc.yml`
- [CLI reference](/reference/cli/gotsrpc) — run the generator
