# Go TypeScript RPC

Where to define this ?!

```bash
gotsrpc -ts-target path/to/local/whatever.ts -ts-module My.Module.Foo my/package/path ServiceStructA [ServiceStructB ...]
gotsrpc build gotsrpc.yml
```

```yaml
---
# gotsrpc.yml
targets:
  -
    package: my/package/path
    expose:
      - ServiceStructA
      - ServiceStructB
    ts:
        module: My.Module.Foo
        target: path/to/local/whatever.ts
    
```