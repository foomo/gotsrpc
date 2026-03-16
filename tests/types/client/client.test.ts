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

test("int8", async () => {
	expect(await client.int8(127)).toBe(127);
});

test("int16", async () => {
	expect(await client.int16(32767)).toBe(32767);
});

test("int32", async () => {
	expect(await client.int32(2147483647)).toBe(2147483647);
});

test("uint", async () => {
	expect(await client.uint(42)).toBe(42);
});

test("uint8", async () => {
	expect(await client.uint8(255)).toBe(255);
});

test("uint16", async () => {
	expect(await client.uint16(65535)).toBe(65535);
});

test("uint32", async () => {
	expect(await client.uint32(4294967295)).toBe(4294967295);
});

test("uint64", async () => {
	expect(await client.uint64(9876543210)).toBe(9876543210);
});

test("float32", async () => {
	expect(await client.float32(3.14)).toBeCloseTo(3.14, 2);
});

test("allScalarsStruct", async () => {
	const v = {
		int8: 127, int16: 32767, int32: 2147483647,
		uint: 42, uint8: 255, uint16: 65535, uint32: 4294967295, uint64: 9876543210,
		float32: 3.14,
	};
	const ret = await client.allScalarsStruct(v);
	expect(ret.int8).toBe(v.int8);
	expect(ret.int16).toBe(v.int16);
	expect(ret.int32).toBe(v.int32);
	expect(ret.uint).toBe(v.uint);
	expect(ret.uint8).toBe(v.uint8);
	expect(ret.uint16).toBe(v.uint16);
	expect(ret.uint32).toBe(v.uint32);
	expect(ret.uint64).toBe(v.uint64);
	expect(ret.float32).toBeCloseTo(v.float32, 2);
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

test("int8Ptr", async () => {
	expect(await client.int8Ptr(127)).toBe(127);
});

test("int16Ptr", async () => {
	expect(await client.int16Ptr(32767)).toBe(32767);
});

test("int32Ptr", async () => {
	expect(await client.int32Ptr(2147483647)).toBe(2147483647);
});

test("uintPtr", async () => {
	expect(await client.uintPtr(42)).toBe(42);
});

test("uint8Ptr", async () => {
	expect(await client.uint8Ptr(255)).toBe(255);
});

test("uint16Ptr", async () => {
	expect(await client.uint16Ptr(65535)).toBe(65535);
});

test("uint32Ptr", async () => {
	expect(await client.uint32Ptr(4294967295)).toBe(4294967295);
});

test("uint64Ptr", async () => {
	expect(await client.uint64Ptr(9876543210)).toBe(9876543210);
});

test("float32Ptr", async () => {
	expect(await client.float32Ptr(3.14)).toBeCloseTo(3.14, 2);
});

test("allScalarPointersStruct", async () => {
	const v = {
		int8Ptr: 127, int16Ptr: 32767, int32Ptr: 2147483647,
		uintPtr: 42, uint8Ptr: 255, uint16Ptr: 65535, uint32Ptr: 4294967295, uint64Ptr: 9876543210,
		float32Ptr: 3.14,
	};
	const ret = await client.allScalarPointersStruct(v);
	expect(ret.int8Ptr).toBe(v.int8Ptr);
	expect(ret.int16Ptr).toBe(v.int16Ptr);
	expect(ret.int32Ptr).toBe(v.int32Ptr);
	expect(ret.uintPtr).toBe(v.uintPtr);
	expect(ret.uint8Ptr).toBe(v.uint8Ptr);
	expect(ret.uint16Ptr).toBe(v.uint16Ptr);
	expect(ret.uint32Ptr).toBe(v.uint32Ptr);
	expect(ret.uint64Ptr).toBe(v.uint64Ptr);
	expect(ret.float32Ptr).toBeCloseTo(v.float32Ptr, 2);
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

test("int8Slice", async () => {
	expect(await client.int8Slice([-128, 0, 127])).toEqual([-128, 0, 127]);
});

test("int16Slice", async () => {
	expect(await client.int16Slice([-32768, 0, 32767])).toEqual([-32768, 0, 32767]);
});

test("int32Slice", async () => {
	expect(await client.int32Slice([-2147483648, 0, 2147483647])).toEqual([-2147483648, 0, 2147483647]);
});

test("uintSlice", async () => {
	expect(await client.uintSlice([0, 42, 100])).toEqual([0, 42, 100]);
});

test("uint16Slice", async () => {
	expect(await client.uint16Slice([0, 1000, 65535])).toEqual([0, 1000, 65535]);
});

test("uint32Slice", async () => {
	expect(await client.uint32Slice([0, 1000, 4294967295])).toEqual([0, 1000, 4294967295]);
});

test("uint64Slice", async () => {
	expect(await client.uint64Slice([0, 1000, 9876543210])).toEqual([0, 1000, 9876543210]);
});

test("float32Slice", async () => {
	const ret = await client.float32Slice([1.1, 2.2, 3.3]);
	expect(ret).toHaveLength(3);
	expect(ret![0]).toBeCloseTo(1.1, 2);
	expect(ret![1]).toBeCloseTo(2.2, 2);
	expect(ret![2]).toBeCloseTo(3.3, 2);
});

test("allScalarSlicesStruct", async () => {
	const v = {
		int8s: [-1, 0, 1], int16s: [-1, 0, 1], int32s: [-1, 0, 1],
		uints: [0, 1, 2], uint16s: [0, 1, 2],
		uint32s: [0, 1, 2], uint64s: [0, 1, 2],
		float32s: [1.1, 2.2],
	};
	const ret = await client.allScalarSlicesStruct(v);
	expect(ret.int8s).toEqual(v.int8s);
	expect(ret.int16s).toEqual(v.int16s);
	expect(ret.int32s).toEqual(v.int32s);
	expect(ret.uints).toEqual(v.uints);
	expect(ret.uint16s).toEqual(v.uint16s);
	expect(ret.uint32s).toEqual(v.uint32s);
	expect(ret.uint64s).toEqual(v.uint64s);
	expect(ret.float32s).toHaveLength(2);
	expect(ret.float32s![0]).toBeCloseTo(1.1, 2);
	expect(ret.float32s![1]).toBeCloseTo(2.2, 2);
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

test("stringInt8Map", async () => {
	const v = { a: 1, b: -1 };
	expect(await client.stringInt8Map(v)).toEqual(v);
});

test("stringInt16Map", async () => {
	const v = { a: 100, b: -100 };
	expect(await client.stringInt16Map(v)).toEqual(v);
});

test("stringInt32Map", async () => {
	const v = { a: 100000, b: -100000 };
	expect(await client.stringInt32Map(v)).toEqual(v);
});

test("stringUintMap", async () => {
	const v = { a: 0, b: 42 };
	expect(await client.stringUintMap(v)).toEqual(v);
});

test("stringUint8Map", async () => {
	const v = { a: 0, b: 255 };
	expect(await client.stringUint8Map(v)).toEqual(v);
});

test("stringUint16Map", async () => {
	const v = { a: 0, b: 65535 };
	expect(await client.stringUint16Map(v)).toEqual(v);
});

test("stringUint32Map", async () => {
	const v = { a: 0, b: 4294967295 };
	expect(await client.stringUint32Map(v)).toEqual(v);
});

test("stringUint64Map", async () => {
	const v = { a: 0, b: 9876543210 };
	expect(await client.stringUint64Map(v)).toEqual(v);
});

test("stringFloat32Map", async () => {
	const v = { a: 1.1, b: 2.2 };
	const ret = await client.stringFloat32Map(v);
	expect(ret!.a).toBeCloseTo(1.1, 2);
	expect(ret!.b).toBeCloseTo(2.2, 2);
});

test("allScalarMapsStruct", async () => {
	const v = {
		int8Map: { x: 1 }, int16Map: { x: 1 }, int32Map: { x: 1 },
		uintMap: { x: 1 }, uint8Map: { x: 1 }, uint16Map: { x: 1 },
		uint32Map: { x: 1 }, uint64Map: { x: 1 }, float32Map: { x: 1.5 },
	};
	const ret = await client.allScalarMapsStruct(v);
	expect(ret.int8Map).toEqual(v.int8Map);
	expect(ret.int16Map).toEqual(v.int16Map);
	expect(ret.int32Map).toEqual(v.int32Map);
	expect(ret.uintMap).toEqual(v.uintMap);
	expect(ret.uint8Map).toEqual(v.uint8Map);
	expect(ret.uint16Map).toEqual(v.uint16Map);
	expect(ret.uint32Map).toEqual(v.uint32Map);
	expect(ret.uint64Map).toEqual(v.uint64Map);
	expect(ret.float32Map!.x).toBeCloseTo(1.5, 2);
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
	const v: (Record<string, string> | null)[] = [{ a: "1" }, { b: "2" }];
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

test("byteSlice", async () => {
	const ret = await client.byteSlice(btoa("hello"))
	expect(ret).toBe(btoa("hello"));
});

test("objectID", async () => {
	const v = btoa(String.fromCharCode(...new Uint8Array(12).fill(42)));
	const ret = await client.objectID(v as unknown as (Uint8Array & { readonly length: 12 }));
	expect(ret).toBe(v as unknown as (Uint8Array & { readonly length: 12 }));
});
