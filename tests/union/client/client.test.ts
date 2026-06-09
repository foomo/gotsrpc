import {expect, test} from "bun:test";
import transport from "../../common/transport"
import {ServiceClient} from "./client.ts";
import {UnionString, UnionStructAValueB} from "./vo.ts";
import {UnionStructAValueB as PrivateUnionStructAValueB} from "./vo-private.ts";

test("inlineStruct", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.inlineStruct();
	expect(ret.valueA).toBe("a");
});

test("inlineStructPtr", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.inlineStructPtr();
	expect(ret.valueA).toBe("a");
});

test("unionString", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.unionString();
	expect(ret).toEqual(UnionString.Three);
});

test("unionStruct", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.unionStruct();
	expect(ret?.kind).toBe("UnionStructB");
	expect(ret?.value).toBe(UnionStructAValueB.One);
});

test("privateInlineStruct", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.privateInlineStruct();
	expect(ret?.valueA).toBe("a");
});

test("privateInlineStructPtr", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.privateInlineStructPtr();
	expect(ret?.valueA).toBe("a");
});

test("privateUnionString", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.privateUnionString();
	expect(ret).toEqual(UnionString.Three);
});

test("privateUnionStruct", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.privateUnionStruct();
	expect(ret?.kind).toBe("UnionStructB");
	expect(ret?.value).toBe(PrivateUnionStructAValueB.One);
});
