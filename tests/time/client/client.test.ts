import { expect, test } from "bun:test";
import transport from "../../lib/transport"
import {ServiceClient} from "./client.ts";

const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));

test("time", async () => {
	const now = new Date().getTime()
	const ret = await client.time(now);
	expect(ret).toBe(now);
});

test("timeStruct", async () => {
	const now = new Date().getTime()
	const ret = await client.timeStruct({
		time: now,
		timePtr: now,
	});
	expect(ret.time).toBe(now);
	expect(ret.timePtr).toBe(now);
});
