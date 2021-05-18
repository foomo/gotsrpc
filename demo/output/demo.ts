/* eslint:disable */
module GoTSRPC.Demo {
	// github.com/foomo/gotsrpc/v2/demo.Address
	export interface Address {
		city?:string;
		signs?:string[];
		PeoplePtr:GoTSRPC.Demo.Person[];
		ArrayOfMaps:Record<string, boolean>[];
		ArrayArrayAddress:GoTSRPC.Demo.Address[][];
		People:GoTSRPC.Demo.Person[];
		MapCrap:Record<string, Record<number, boolean>>;
		NestedPtr?:GoTSRPC.Demo.Nested.Nested;
		NestedStruct:GoTSRPC.Demo.Nested.Nested;
	}
	// github.com/foomo/gotsrpc/v2/demo.Bar
	export type Bar = any
	// github.com/foomo/gotsrpc/v2/demo.Check
	export interface Check {
		Foo:string;
	}
	// github.com/foomo/gotsrpc/v2/demo.CustomError
	export enum CustomError {
		Demo = "demo",
	}
	// github.com/foomo/gotsrpc/v2/demo.CustomTypeFoo
	export type CustomTypeFoo = string
	// github.com/foomo/gotsrpc/v2/demo.CustomTypeInt
	export enum CustomTypeInt {
		One = 1,
		Three = 3,
		Two = 2,
	}
	// github.com/foomo/gotsrpc/v2/demo.CustomTypeString
	export enum CustomTypeString {
		Five = "CONST_CASE",
		Four = "slug-case",
		One = "regular",
		Seven = "dot.case",
		Six = "SLUG-CASE-UPPER",
		Three = "snake_case",
		Two = "camelCase",
	}
	// github.com/foomo/gotsrpc/v2/demo.CustomTypeStruct
	export interface CustomTypeStruct {
		CustomTypeFoo:GoTSRPC.Demo.CustomTypeFoo;
		CustomTypeInt:GoTSRPC.Demo.CustomTypeInt;
		CustomTypeString:GoTSRPC.Demo.CustomTypeString;
		CustomTypeNested:GoTSRPC.Demo.Nested.CustomTypeNested;
		Check:GoTSRPC.Demo.Check;
	}
	// github.com/foomo/gotsrpc/v2/demo.Err
	export interface Err {
		message:string;
	}
	// github.com/foomo/gotsrpc/v2/demo.Inner
	export interface Inner {
		one:string;
	}
	// github.com/foomo/gotsrpc/v2/demo.OuterInline
	export interface OuterInline {
		one:string;
		two:string;
	}
	// github.com/foomo/gotsrpc/v2/demo.OuterNested
	export interface OuterNested {
		inner:GoTSRPC.Demo.Inner;
		two:string;
	}
	// github.com/foomo/gotsrpc/v2/demo.Person
	export interface Person {
		Name:string;
		address?:GoTSRPC.Demo.Address;
		AddressStruct:GoTSRPC.Demo.Address;
		Addresses:Record<string, GoTSRPC.Demo.Address>;
		InlinePtr?:{
			Foo:boolean;
		};
		InlineStruct:{
			Bar:string;
		};
		DNA:string;
	}
	// github.com/foomo/gotsrpc/v2/demo.ScalarError
	export type ScalarError = string
	// github.com/foomo/gotsrpc/v2/demo.ScalarInPlace
	export type ScalarInPlace = string
}