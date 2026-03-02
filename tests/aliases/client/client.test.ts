import { expect, test } from "bun:test";
import transport from "../../lib/transport";
import { ServiceClient } from "./client.ts";
import { Status, Priority, Category } from "./vo.ts";

const client = new ServiceClient(
	transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`)
);

// Scalar aliases

test("statusValue", async () => {
	expect(await client.statusValue(Status.Active)).toBe(Status.Active);
});

test("categoryValue", async () => {
	expect(await client.categoryValue(Category.A)).toBe(Category.A);
});

test("priorityValue", async () => {
	expect(await client.priorityValue(Priority.High)).toBe(Priority.High);
});

test("ratingValue", async () => {
	expect(await client.ratingValue(4.5)).toBeCloseTo(4.5);
});

// Typed collections

test("tagsValue", async () => {
	expect(await client.tagsValue(["go", "typescript", "rpc"])).toEqual(["go", "typescript", "rpc"]);
});

test("entriesValue", async () => {
	const v = [{ id: "1", status: Status.Active, priority: Priority.High, rating: 4.5, tags: ["a"] }];
	const ret = await client.entriesValue(v);
	expect(ret).toHaveLength(1);
	expect(ret![0]!.id).toBe("1");
	expect(ret![0]!.status).toBe(Status.Active);
	expect(ret![0]!.priority).toBe(Priority.High);
});

test("registryValue", async () => {
	const v = { first: { id: "1", status: Status.Active, priority: Priority.Low, rating: 1.0 } };
	const ret = await client.registryValue(v);
	expect(ret!["first"].id).toBe("1");
	expect(ret!["first"].status).toBe(Status.Active);
});

test("indexValue", async () => {
	const v = { [Category.A]: [{ id: "1", status: Status.Active, priority: Priority.Low, rating: 1.0 }] };
	const ret = await client.indexValue(v);
	expect(ret![Category.A]).toHaveLength(1);
	expect(ret![Category.A]![0].id).toBe("1");
});

test("labelMapValue", async () => {
	const v = { key: "val", env: "prod" };
	expect(await client.labelMapValue(v)).toEqual(v);
});

// Structs with enums

test("entryValue", async () => {
	const v = { id: "1", status: Status.Active, priority: Priority.High, rating: 4.5, tags: ["a", "b"] };
	expect(await client.entryValue(v)).toEqual(v);
});

test("detailValue", async () => {
	const v = {
		name: "test",
		description: "desc",
		entry: { id: "1", status: Status.Active, priority: Priority.High, rating: 4.5 },
		labels: { key: "val" },
	};
	const ret = await client.detailValue(v);
	expect(ret.name).toBe("test");
	expect(ret.description).toBe("desc");
	expect(ret.entry.status).toBe(Status.Active);
	expect(ret.labels).toEqual({ key: "val" });
});

test("dataRecordValue", async () => {
	const v = {
		id: "1",
		title: "test",
		status: Status.Pending,
		amount: { value: 100, currency: "USD" },
		items: [{ id: "i1", status: Status.Active, priority: Priority.Low, rating: 1.0 }],
		metadata: { createdBy: "admin", note: "some note" },
		categories: [Category.A, Category.B],
	};
	const ret = await client.dataRecordValue(v);
	expect(ret.id).toBe("1");
	expect(ret.status).toBe(Status.Pending);
	expect(ret.amount?.value).toBe(100);
	expect(ret.amount?.currency).toBe("USD");
	expect(ret.items).toHaveLength(1);
	expect(ret.metadata?.createdBy).toBe("admin");
	expect(ret.metadata?.note).toBe("some note");
	expect(ret.categories).toEqual([Category.A, Category.B]);
});

// Complex nesting

test("mapOfEntries", async () => {
	const v = {
		[Category.A]: [{ id: "1", status: Status.Active, priority: Priority.Low, rating: 1.0 }],
	};
	const ret = await client.mapOfEntries(v);
	expect(ret![Category.A]).toHaveLength(1);
	expect(ret![Category.A]![0].id).toBe("1");
});

// Nil optionals

test("dataRecordNil", async () => {
	const v = { id: "1", title: "minimal", status: Status.Closed };
	const ret = await client.dataRecordNil(v);
	expect(ret.id).toBe("1");
	expect(ret.status).toBe(Status.Closed);
});
