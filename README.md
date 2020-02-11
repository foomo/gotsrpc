# Go / TypeScript and Go / Go RPC

[![Build Status](https://travis-ci.org/foomo/gotsrpc.svg?branch=master)](https://travis-ci.org/foomo/gotsrpc)

## Installation

From source to /usr/local/bin/gotsrpc:

```bash
go get github.com/foomo/gotsrpc
cd $GOPATH/src/github.com/foomo/gotsrpc
make install
```

If you trust us there are precompiled versions:

[releases](https://github.com/foomo/gotsrpc/releases)

On the mac:

```bash
brew install foomo/gotsrpc/gotsrpc
```


## Usage

```bash
gotsrpc gotsrpc.yml
```

Will generate client and server side go and TypeScript code. Have fun!

## Configuration Examples

### Standard Example

[demo/config.yml](demo/config-commonjs.yml)

```yaml

modulekind: commonjs
# if you want an async api vs classic callbacks - here you are
tsclientflavor: async
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

#### Async Example

How to use async clients in this case with axios:

```TypeScript
import axios, { AxiosResponse } from "axios";
import { ServiceClient as ExampleClient } from "./some/generated/client";

// axios transport
let getTransport = endpoint => async <T>(method, args = []) => {
	return new Promise<T>(async (resolve, reject) => {
		try {
			let axiosPromise: AxiosResponse<T> = await axios.post<T>(
				endpoint + "/" + encodeURIComponent(method),
				JSON.stringify(args),
			);
			return resolve(axiosPromise.data);
		} catch (e) {
			return reject(e);
		}
	});
};

let client = new ExampleClient(getTransport(ExampleClient.defaultEndpoint));

export async function test() {
	try {
		let result = await client.getResult();
		console.log("here is the result", result);
	} catch (e) {
		// e => network?
		// e => json
		// e => domain error type
		console.error("something went wrong ...", e);
	}
}
```


### Oldschool Typescript

[demo/config.yml](demo/config.yml)

```yaml

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

## GOModule Support

To support go modules add 

```yaml

module:
  name: github.com/foomo/gotsrpc
  path: ../ # Relative Or Absolute Path where the package was checked out (root of the package)

```
