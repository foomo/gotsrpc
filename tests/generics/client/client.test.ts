import { expect, test } from "bun:test";
import transport from "../../common/transport";
import { ServiceClient } from "./client.ts";

const client = new ServiceClient(
	transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`)
);

test("getStringResponse", async () => {
	const ret = await client.getStringResponse();
	expect(ret.data).toBe("hello");
	expect(ret.error).toBeUndefined();
});

test("getItemResponse", async () => {
	const ret = await client.getItemResponse();
	expect(ret.data.id).toBe("1");
	expect(ret.data.name).toBe("test");
});

test("setItemResponse", async () => {
	const ret = await client.setItemResponse({ data: { id: "1", name: "x" }, error: "" });
	expect(ret).toBe(true);
});

test("getPair", async () => {
	const ret = await client.getPair();
	expect(ret.first).toBe("hello");
	expect(ret.second).toBe(42);
});

test("getPagedItems", async () => {
	const ret = await client.getPagedItems(1);
	expect(ret.items).toHaveLength(1);
	expect(ret.items![0]!.id).toBe("1");
	expect(ret.total).toBe(1);
});

test("getNestedGeneric", async () => {
	const ret = await client.getNestedGeneric();
	expect(ret.items).toHaveLength(1);
	expect(ret.items![0]!.first).toBe("key");
	expect(ret.items![0]!.second.id).toBe("1");
});

test("getResult", async () => {
	const ret = await client.getResult();
	expect(ret.value).toBeDefined();
	expect(ret.value!.id).toBe("1");
});

test("getContainer", async () => {
	const ret = await client.getContainer();
	expect(ret.data!["key"]!.id).toBe("1");
	expect(ret.default.name).toBe("default");
});
