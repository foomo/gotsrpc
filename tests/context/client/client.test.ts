import { expect, test } from "bun:test";
import transport from "../../lib/transport"
import {ServiceClient} from "./client.ts";

test("hello", async () => {
	const client = new ServiceClient(transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`));
	const ret = await client.hello("world");
	expect(ret).toBe("Hello world");
});