---
layout: home

hero:
  name: gotsrpc
  text: Type-safe RPC for Go & TypeScript
  tagline: Generate Go service proxies and idiomatic Go and TypeScript clients from your Go interfaces — code-first, no schema language to learn.
  image:
    src: /logo.png
    alt: gotsrpc
  actions:
    - theme: brand
      text: Get Started
      link: /guide/getting-started
    - theme: alt
      text: Writing Services
      link: /guide/writing-services
    - theme: alt
      text: View on GitHub
      link: https://github.com/foomo/gotsrpc

features:
  - title: Code-first
    details: Define plain Go interfaces. gotsrpc parses your source via the AST and generates everything else — no IDL, no annotations required.
  - title: Type-safe across the wire
    details: Go structs, enums, slices, maps, pointers and generics map to matching TypeScript types, so the compiler catches mismatches on both sides.
  - title: Go ↔ TypeScript and Go ↔ Go
    details: Generate an HTTP client for TypeScript frontends and, optionally, a binary Go-to-Go RPC client for service-to-service calls.
  - title: Idiomatic clients
    details: Generated clients read like hand-written code — camelCased TypeScript methods, context-aware Go methods, and a transport you fully control.
---

## What is gotsrpc?

**gotsrpc** is a code generator that produces type-safe RPC bindings between Go
services and TypeScript clients, with optional Go-to-Go RPC support. You write a
Go interface, describe it in a small YAML config (`gotsrpc.yml`), and gotsrpc
generates:

- **Go HTTP service proxies** (`gotsrpc_gen.go`) that serve your implementation over HTTP
- **Go HTTP clients** (`gotsrpcclient_gen.go`) for calling those services from Go
- **Go binary-protocol proxies/clients** (`gorpc_gen.go`, `gorpcclient_gen.go`) — opt-in
- **TypeScript clients and type definitions** for the frontend

## Who is it for?

- **Go developers** who want to expose Go interfaces as HTTP RPC services and
  consume them from Go without hand-writing boilerplate.
- **TypeScript developers** who want to call a Go API through a fully typed,
  idiomatic TypeScript client instead of untyped `fetch` calls.

## Why not REST or GraphQL?

gotsrpc is deliberately **not RESTful**. It targets teams where the same people
(or closely collaborating teams) own the Go backend and the TypeScript frontend.
In that setting a REST layer adds ceremony without buying much, and GraphQL's
value — letting frontend teams query independently of backend developers — is
less relevant. gotsrpc instead makes the Go interface the single source of
truth and keeps both ends in lockstep through generated, type-checked code.

Ready to try it? Head to the [Getting Started](/guide/getting-started) guide.
