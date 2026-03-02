import { expect, test } from "bun:test";
import transport from "../../lib/transport";
import { ServiceClient } from "./client.ts";

const client = new ServiceClient(
	transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`)
);

// Scalars

test("bool", async () => {
	expect(await client.bool(true)).toBe(true);
});

test("int", async () => {
	expect(await client.int(42)).toBe(42);
});

test("int64", async () => {
	expect(await client.int64(9876543210)).toBe(9876543210);
});

test("float64", async () => {
	expect(await client.float64(3.14159)).toBeCloseTo(3.14159);
});

test("string", async () => {
	expect(await client.string("hello world")).toBe("hello world");
});

// Pointers

test("stringPtr", async () => {
	expect(await client.stringPtr("test")).toBe("test");
});

test("int64Ptr", async () => {
	expect(await client.int64Ptr(42)).toBe(42);
});

test("boolPtr", async () => {
	expect(await client.boolPtr(true)).toBe(true);
});

// Structs

test("simpleStruct", async () => {
	const v = { bool: true, int: 42, int64: 100, float64: 2.718, string: "test" };
	expect(await client.simpleStruct(v)).toEqual(v);
});

test("nestedStruct", async () => {
	const v = { name: "parent", child: { bool: true, int: 1, int64: 2, float64: 3.0, string: "child" } };
	expect(await client.nestedStruct(v)).toEqual(v);
});

test("structWithPointers", async () => {
	const v = { strPtr: "hello", int64Ptr: 42, boolPtr: true, child: { bool: true, int: 1, int64: 2, float64: 3.0, string: "child" } };
	expect(await client.structWithPointers(v)).toEqual(v);
});

test("structWithCollections", async () => {
	const v = {
		strings: ["a", "b"],
		int64s: [1, 2, 3],
		items: [{ bool: true, int: 1, int64: 2, float64: 3.0, string: "item" }],
		itemPtrs: [{ bool: false, int: 10, int64: 20, float64: 30.0, string: "ptr" }],
		stringMap: { key: "val" },
		structMap: { s: { bool: true, int: 5, int64: 6, float64: 7.0, string: "map" } },
	};
	expect(await client.structWithCollections(v)).toEqual(v);
});

// Slices

test("stringSlice", async () => {
	expect(await client.stringSlice(["a", "b", "c"])).toEqual(["a", "b", "c"]);
});

test("int64Slice", async () => {
	expect(await client.int64Slice([10, 20, 30])).toEqual([10, 20, 30]);
});

test("simpleSlice", async () => {
	const v = [
		{ bool: true, int: 1, int64: 2, float64: 3.0, string: "one" },
		{ bool: false, int: 4, int64: 5, float64: 6.0, string: "two" },
	];
	expect(await client.simpleSlice(v)).toEqual(v);
});

test("simplePtrSlice", async () => {
	const v = [
		{ bool: true, int: 1, int64: 2, float64: 3.0, string: "one" },
		{ bool: false, int: 4, int64: 5, float64: 6.0, string: "two" },
	];
	expect(await client.simplePtrSlice(v)).toEqual(v);
});

test("stringSlice2D", async () => {
	const v = [["a", "b"], ["c", "d"]];
	expect(await client.stringSlice2D(v)).toEqual(v);
});

// Maps

test("stringStringMap", async () => {
	const v = { a: "1", b: "2" };
	expect(await client.stringStringMap(v)).toEqual(v);
});

test("stringInt64Map", async () => {
	const v = { x: 10, y: 20 };
	expect(await client.stringInt64Map(v)).toEqual(v);
});

test("stringSimpleMap", async () => {
	const v = { one: { bool: true, int: 1, int64: 2, float64: 3.0, string: "one" } };
	expect(await client.stringSimpleMap(v)).toEqual(v);
});

test("stringSimplePtrMap", async () => {
	const v = { k: { bool: true, int: 1, int64: 2, float64: 3.0, string: "val" } };
	expect(await client.stringSimplePtrMap(v)).toEqual(v);
});

test("stringStringSliceMap", async () => {
	const v = { colors: ["red", "blue"], sizes: ["s", "m"] };
	expect(await client.stringStringSliceMap(v)).toEqual(v);
});

// Complex nested

test("mapOfMaps", async () => {
	const v = { outer: { inner: "val" } };
	expect(await client.mapOfMaps(v)).toEqual(v);
});

test("mapOfSimpleSlice", async () => {
	const v = { group: [{ bool: true, int: 1, int64: 2, float64: 3.0, string: "item" }] };
	expect(await client.mapOfSimpleSlice(v)).toEqual(v);
});

test("sliceOfMaps", async () => {
	const v = [{ a: "1" }, { b: "2" }];
	expect(await client.sliceOfMaps(v)).toEqual(v);
});

// Multi-args

test("multiArgs", async () => {
	const ret = await client.multiArgs("hello", 42, true);
	expect(ret.ret).toBe("hello");
	expect(ret.ret_1).toBe(42);
	expect(ret.ret_2).toBe(true);
});

test("mixedArgs", async () => {
	const s = { bool: true, int: 1, int64: 2, float64: 3.0, string: "mix" };
	const items = ["a", "b"];
	const m = { x: 10 };
	const ret = await client.mixedArgs(s, items, m);
	expect(ret.ret).toEqual(s);
	expect(ret.ret_1).toEqual(items);
	expect(ret.ret_2).toEqual(m);
});

// Edge cases

test("empty", async () => {
	const ret = await client.empty();
	expect(ret).toBe(true);
});
