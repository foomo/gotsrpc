# Configuration

Code generation is driven by a single YAML file, conventionally named
`gotsrpc.yml`. It has three top-level sections — `module`, `targets` and
`mappings` — plus a couple of global options.

## Editor support (JSON schema)

gotsrpc ships a JSON Schema, [`gotsrpc.schema.json`](https://github.com/foomo/gotsrpc/blob/main/gotsrpc.schema.json),
generated directly from the Go config structs. Add this comment as the **first
line** of your config to get autocompletion and inline validation in editors
that run the YAML Language Server (VS Code, Neovim, …):

```yaml
# yaml-language-server: $schema=gotsrpc.schema.json
```

Use a relative path that resolves to the schema file, or reference it remotely:

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/gotsrpc/refs/heads/main/gotsrpc.schema.json
```

::: tip
gotsrpc itself does **not** validate the config beyond parsing the YAML —
unknown keys are silently ignored and there are no required-field checks at
runtime. The schema is your safety net while editing, so keep the
`yaml-language-server` comment in place.
:::

## Full example

```yaml
# yaml-language-server: $schema=gotsrpc.schema.json
module:
  name: github.com/foomo/gotsrpc/v3
  path: ../../

targets:
  basic:
    services:
      /service: Service
    package: github.com/foomo/gotsrpc/v3/example/basic/service
    out: ./client/src/service-client.ts
    gorpc:
      - Service
    tsrpc:
      - Service

mappings:
  github.com/foomo/gotsrpc/v3/example/basic/service:
    out: ./client/src/service-vo.ts
```

## `module`

Identifies the Go module that contains your services.

| Field | Type | Description |
| --- | --- | --- |
| `name` | string | The Go module name, e.g. `github.com/foomo/gotsrpc/v3`. |
| `path` | string | Path to the module root (the directory holding `go.mod`). If relative, it is resolved against the config file's directory. |

When `path` resolves to a directory containing a `go.mod`, gotsrpc parses it to
understand the module. A missing `go.mod` is tolerated.

```yaml
module:
  name: github.com/foomo/gotsrpc/v3
  path: ../../
```

## `targets`

A map of **target name → target settings**. Each target scans one Go package
and produces a set of proxies and clients.

| Field | Type | Description |
| --- | --- | --- |
| `package` | string | Go import path of the package to scan for services. |
| `services` | map | HTTP route → Go service (interface) name, e.g. `/service: Service`. |
| `out` | string | Output path for the generated TypeScript client. |
| `module` | string | Optional TypeScript module name. *(YAML key is `module`; maps to the Go field `TypeScriptModule`.)* |
| `gorpc` | list | Services for which to also generate a Go ↔ Go binary RPC proxy/client. |
| `tsrpc` | list | Services for which to generate a TypeScript client. **Empty means all services.** |
| `skipTSRPCClient` | bool | Skip generating the TypeScript client. |
| `serviceNames` | map | Override the service name used in OpenTelemetry telemetry (span name and the `rpc.method` metric/attribute), keyed by Go service name — e.g. `Service: Monitor`. Defaults to the service name. Affects generated proxies **and** clients; does not change generated Go types or routing. See [Telemetry service names](#telemetry-service-names). |

```yaml
targets:
  basic:
    services:
      /service: Service
    package: github.com/foomo/gotsrpc/v3/example/basic/service
    out: ./client/src/service-client.ts
    gorpc: [Service]
    tsrpc: [Service]
```

### Telemetry service names

Generated proxies and clients emit OpenTelemetry `rpc.*` spans and metrics. The span
name and the `rpc.method` attribute/metric are formed as `<service>/<method>`, where
`<service>` defaults to the Go service type name.

Services are conventionally named `Service`, so telemetry from different packages all
collapses into `Service/<Method>` — indistinguishable across services. Use `serviceNames`
to give a service a distinct telemetry identity **without renaming the Go type**:

```yaml
targets:
  monitor:
    services:
      /service: Service
    serviceNames:
      Service: Monitor   # spans/metrics read "Monitor/Hello" instead of "Service/Hello"
    package: github.com/foomo/gotsrpc/v3/example/monitor/service
    gorpc: [Service]
    tsrpc: [Service]
```

The map is keyed by the Go service name (the value side of `services:`). Services not
listed keep their type name. The override is applied consistently to the generated server
proxy and the generated clients, so their metrics line up under the same name. It only
affects telemetry — generated Go type names and HTTP routing are unchanged.

## `mappings`

A map of **Go import path → mapping settings**. Mappings tell gotsrpc where to
emit the TypeScript type definitions (value objects) for every package your
services reference — including the service package itself, shared/common
packages, and even standard-library packages such as `time` or `encoding/json`.

| Field | Type | Description |
| --- | --- | --- |
| `out` | string | Output path for this package's generated TypeScript types. |
| `structs` | list | Explicit list of Go types to generate (e.g. `github.com/foomo/gotsrpc/v3.Error`). |
| `scalars` | list | Go types to treat as TypeScript scalars. |
| `module` | string | Optional TypeScript module name. |

```yaml
mappings:
  github.com/foomo/gotsrpc/v3/tests/generics/server:
    out: ./client/vo.ts
  github.com/foomo/gotsrpc/v3/tests/common:
    out: ./client/vo-common.ts
  # Pull specific types out of the gotsrpc runtime package
  github.com/foomo/gotsrpc/v3:
    out: ./client/vo-gotsrpc.ts
    structs: [github.com/foomo/gotsrpc/v3.Error]
```

::: warning
Every package referenced by a service's arguments or return values must have a
mapping. Missing mappings cause generation to fail.
:::

## Global options

| Field | Type | Description |
| --- | --- | --- |
| `tsImportJsExtension` | bool | Append a `.js` extension to relative TypeScript import paths. Required for TypeScript projects using `moduleResolution` `nodenext` or `node16`. |

```yaml
tsImportJsExtension: true
```

## Schema reference

For completeness, the committed schema (`gotsrpc.schema.json`) is reproduced
below. It is regenerated from the Go structs by the test suite, so it always
matches the fields gotsrpc understands.

```json
{
  "$schema": "https://json-schema.org/draft/2020-12/schema",
  "$id": "https://raw.githubusercontent.com/foomo/gotsrpc/refs/heads/main/gotsrpc.schema.json",
  "$ref": "#/$defs/Config",
  "$defs": {
    "Config": {
      "properties": {
        "module": { "$ref": "#/$defs/Namespace" },
        "targets": { "additionalProperties": { "$ref": "#/$defs/Target" }, "type": "object" },
        "mappings": { "$ref": "#/$defs/TypeScriptMappings" },
        "tsImportJsExtension": { "type": "boolean" }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Namespace": {
      "properties": {
        "name": { "type": "string" },
        "path": { "type": "string" }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Target": {
      "properties": {
        "package": { "type": "string" },
        "services": { "additionalProperties": { "type": "string" }, "type": "object" },
        "module": { "type": "string" },
        "out": { "type": "string" },
        "gorpc": { "items": { "type": "string" }, "type": "array" },
        "tsrpc": { "items": { "type": "string" }, "type": "array" },
        "skipTSRPCClient": { "type": "boolean" },
        "serviceNames": { "additionalProperties": { "type": "string" }, "type": "object" }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "Mapping": {
      "properties": {
        "out": { "type": "string" },
        "structs": { "items": { "type": "string" }, "type": "array" },
        "scalars": { "items": { "type": "string" }, "type": "array" },
        "module": { "type": "string" }
      },
      "additionalProperties": false,
      "type": "object"
    },
    "TypeScriptMappings": {
      "additionalProperties": { "$ref": "#/$defs/Mapping" },
      "type": "object"
    }
  }
}
```

## Next steps

- [CLI reference](/reference/cli/gotsrpc) — run the generator against your config
- [Writing Services](/guide/writing-services) — the code the config points at
