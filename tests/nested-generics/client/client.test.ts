import { expect, test } from "bun:test";
import transport from "../../lib/transport";
import { ServiceClient } from "./client.ts";

const client = new ServiceClient(
	transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`)
);

test("getValue", async () => {
	const ret = await client.getValue();
	expect(ret.id).toBe("1");
	expect(ret.name).toBe("test");
});

test("getWrapped", async () => {
	const ret = await client.getWrapped();
	expect(ret.data.id).toBe("1");
	expect(ret.data.name).toBe("test");
});

test("getByKey", async () => {
	const ret = await client.getByKey("hello");
	expect(ret).toBe(5);
});

test("getName", async () => {
	const ret = await client.getName();
	expect(ret).toBe("service");
});
