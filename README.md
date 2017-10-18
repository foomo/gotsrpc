# Go / TypeScript and Go / Go RPC

## Installation

From source to /usr/local/bin/gotsrpc:

```bash
go get github.com/foomo/gotsrpc
cd $GOPATH/src/github.com/foomo/gotsrpc
make install
```

If you trust us there are precompiled versions:

[/foomo/gotsrpc/releases](/foomo/gotsrpc/releases)

On the mac:

```bash
brew install foomo/gotsrpc/gotsrpc
```


## Usage

```bash
gotsrpc gotsrpc.yml
```

Will generate client and server side go and TypeScript code. Have fun!

## config expamples

### commonjs

[demo/config.yml](demo/config-commonjs.yml)

```yaml
---
modulekind: commonjs
targets:
  demo:
    services:
      /service/foo: Foo
      /service/demo: Demo
    package: github.com/foomo/gotsrpc/demo
    out: /tmp/test.ts
    gorpc:
      - Foo
      - Demo
    tsrpc:
      - Foo
      - Demo

mappings:
  github.com/foomo/gotsrpc/demo:
    out: /tmp/test-files-demo.ts
  github.com/foomo/gotsrpc/demo/nested:
    out: /tmp/test-files-demo-nested.ts
...
```

### oldschool TypeScript

[demo/config.yml](demo/config.yml)

```yaml
---
targets:
  demo:
    module: GoTSRPC.Demo
    services:
      /service/foo: Foo
      /service/demo: Demo
    package: github.com/foomo/gotsrpc/demo
    out: /tmp/test.ts
    gorpc:
      - Foo
      - Demo
    tsrpc:
      - Foo
      - Demo

mappings:
  github.com/foomo/gotsrpc/demo:
    module: GoTSRPC.Demo
    out: /tmp/test-files-demo.ts
  github.com/foomo/gotsrpc/demo/nested:
    module: GoTSRPC.Demo.Nested
    out: /tmp/test-files-demo-nested.ts
...
```