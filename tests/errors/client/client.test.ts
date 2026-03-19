import {expect, test} from "bun:test";
import transport from "../../common/transport"
import {ServiceClient} from "./client.ts";
import {ScalarError, ScalarA} from "./vo.ts";

test("error", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.error();
	expect(ret.m).toBe("error");
});

test("scalar", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.scalar();
	expect(ret).toBe(ScalarError.ScalarOne);
});

test("multiScalar", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.multiScalar();
	expect(JSON.stringify(ret)).toBe('{"ScalarA":"one"}');
});

test("struct", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.struct();
	expect(JSON.stringify(ret)).toBe('{"message":"my custom scalar","data":"hello world"}');
});

test("typedError", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.typedError();
	expect(JSON.stringify(ret)).toBe('{"m":"typed error","p":"github.com/pkg/errors","t":"*errors.fundamental","d":{}}');
});

test("structError", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.structError();
	expect(JSON.stringify(ret)).toBe('{"m":"struct error","p":"github.com/foomo/gotsrpc/v2/tests/errors/server","t":"server.MyStructError","d":{"Msg":"struct error","Map":{"a":"b"},"Slice":["a","b"],"Struct":{"A":"b"}}}');
});

test("scalarError", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.scalarError();
	expect(JSON.stringify(ret)).toBe('{"m":"scalar error one","p":"github.com/foomo/gotsrpc/v2/tests/errors/server","t":"*server.MyScalarError","d":"scalar error one"}');
});

test("customError", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.customError();
	expect(JSON.stringify(ret)).toBe('{"m":"custom error","p":"github.com/foomo/gotsrpc/v2/tests/errors/server","t":"*server.MyCustomError","d":{"Msg":"custom error","Map":{"a":"b"},"Slice":["a","b"],"Struct":{"A":"b"}}}');
});

test("wrappedError", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.wrappedError();
	expect(JSON.stringify(ret)).toBe('{"m":"wrapped","p":"github.com/pkg/errors","t":"*errors.withMessage","d":{},"c":{"m":"error","p":"errors","t":"*errors.errorString","d":{}}}');
});

test("typedError", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.typedError();
	expect(JSON.stringify(ret)).toBe('{"m":"typed error","p":"github.com/pkg/errors","t":"*errors.fundamental","d":{}}');
});

test("typedScalarError", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.typedScalarError();
	expect(JSON.stringify(ret)).toBe('{"m":"scalar error two","p":"github.com/foomo/gotsrpc/v2/tests/errors/server","t":"*server.MyScalarError","d":"scalar error two"}');
});

test("typedCustomError", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.typedCustomError();
	expect(JSON.stringify(ret)).toBe('{"m":"typed custom error","p":"github.com/foomo/gotsrpc/v2/tests/errors/server","t":"*server.MyCustomError","d":{"Msg":"typed custom error","Map":{"a":"b"},"Slice":["a","b"],"Struct":{"A":"b"}}}');
});
