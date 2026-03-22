import { expect, test } from "bun:test";
import transport from "../../common/transport";
import { ServiceClient } from "./client.ts";
import { ACustomType, type Base } from "./vo.ts";

const client = new ServiceClient(
	transport(`${process.env.GOTSRPC_SERVER_URL}${ServiceClient.defaultEndpoint}`)
);

const base: Base = {
	a1: { Foo: "nested" },
	a3: null,
	b1: "hello",
	b3: null,
	c1: "anything",
	c3: null,
	d1: ACustomType.One,
	d3: null,
	e1: null,
	e3: null,
	f1: null,
	f3: null,
	Two: null,
	Two1: null,
	Two2: null,
	Three: null,
	Three1: null,
	Four: null,
	Five: null,
	Six: null,
};

test("variantA", async () => {
	const ret = await client.variantA(base);
	expect(ret.b1).toBe("hello");
	expect(ret.a1).toEqual({ Foo: "nested" });
	expect(ret.d1).toBe(ACustomType.One);
});

test("variantB", async () => {
	expect(await client.variantB("custom")).toBe("custom");
});

test("variantC", async () => {
	expect(await client.variantC(["a", "b"])).toEqual(["a", "b"]);
});

test("variantC null", async () => {
	expect(await client.variantC(null)).toBeNull();
});

test("variantD", async () => {
	const v = { x: "y" };
	expect(await client.variantD(v)).toEqual(v);
});

test("variantD null", async () => {
	expect(await client.variantD(null)).toBeNull();
});

test("variantE", async () => {
	const ret = await client.variantE(base);
	expect(ret).not.toBeNull();
	expect(ret!.b1).toBe("hello");
});

test("variantE null", async () => {
	expect(await client.variantE(null)).toBeNull();
});

test("variantF", async () => {
	const v: (Base | null)[] = [base, null];
	const ret = await client.variantF(v);
	expect(ret).not.toBeNull();
	expect(ret!.length).toBe(2);
	expect(ret![0]!.b1).toBe("hello");
	expect(ret![1]).toBeNull();
});

test("variantG", async () => {
	const v = { k: base };
	const ret = await client.variantG(v);
	expect(ret).not.toBeNull();
	expect(ret!["k"]!.b1).toBe("hello");
});

test("variantH", async () => {
	const i2: Base = { ...base, b1: "two" };
	const i3: (Base | null)[] = [{ ...base, b1: "three" }];
	const i4 = { k: { ...base, b1: "four" } };
	const ret = await client.variantH(base, i2, i3, i4);
	expect(ret.r1.b1).toBe("hello");
	expect(ret.r2).not.toBeNull();
	expect(ret.r2!.b1).toBe("two");
	expect(ret.r3).not.toBeNull();
	expect(ret.r3![0]!.b1).toBe("three");
	expect(ret.r4).not.toBeNull();
	expect(ret.r4!["k"]!.b1).toBe("four");
});
