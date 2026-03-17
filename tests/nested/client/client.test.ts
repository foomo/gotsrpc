import { expect, test } from "bun:test";
import transport from "../../lib/transport";
import { ExtendedServiceClient } from "./client.ts";

const client = new ExtendedServiceClient(
	transport(`${process.env.GOTSRPC_SERVER_URL}${ExtendedServiceClient.defaultEndpoint}`)
);

test("getFirstName", async () => {
	const ret = await client.getFirstName();
	expect(ret).toBe("John");
});

test("getMiddleName", async () => {
	const ret = await client.getMiddleName();
	expect(ret).toBe("Michael");
});

test("getAge", async () => {
	const ret = await client.getAge();
	expect(ret).toBe(30);
});

test("getLastName", async () => {
	const ret = await client.getLastName();
	expect(ret).toBe("Doe");
});

test("getPerson", async () => {
	const ret = await client.getPerson();
	expect(ret.firstName).toBe("John");
	expect(ret.middleName).toBe("Michael");
	expect(ret.lastName).toBe("Doe");
	expect(ret.age).toBe(30);
});
