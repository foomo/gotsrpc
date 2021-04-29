/* eslint:disable */
module GoTSRPC.Demo {
	// github.com/foomo/gotsrpc/demo.Address
	export interface Address {
		city?:string;
		signs?:string[];
		PeoplePtr:GoTSRPC.Demo.Person[];
		ArrayOfMaps:{[index:string]:boolean}[];
		ArrayArrayAddress:GoTSRPC.Demo.Address[][];
		People:GoTSRPC.Demo.Person[];
		MapCrap:{[index:string]:{[index:number]:boolean}};
		NestedPtr?:GoTSRPC.Demo.Nested.Nested;
		NestedStruct:GoTSRPC.Demo.Nested.Nested;
	}
	// github.com/foomo/gotsrpc/demo.Bar
	export type Bar = any
	// github.com/foomo/gotsrpc/demo.Check
	export interface Check {
		Foo:string;
	}
	// github.com/foomo/gotsrpc/demo.CustomTypeFoo
	export type CustomTypeFoo = string
	// github.com/foomo/gotsrpc/demo.CustomTypeInt
	export type CustomTypeInt = 1 | 2 | 3
	// github.com/foomo/gotsrpc/demo.CustomTypeString
	export enum CustomTypeString {
		Regular = "regular",
		CamelCase = "camelCase",
		SnakeCase = "snake_case",
		SlugCase = "slug-case",
		ConstCase = "CONST_CASE",
		SlugCase = "SLUG-CASE",
		DotCase = "dot.case",
	}
	// github.com/foomo/gotsrpc/demo.CustomTypeStruct
	export interface CustomTypeStruct {
		CustomTypeFoo:GoTSRPC.Demo.CustomTypeFoo;
		CustomTypeInt:GoTSRPC.Demo.CustomTypeInt;
		CustomTypeString:GoTSRPC.Demo.CustomTypeString;
		CustomTypeNested:GoTSRPC.Demo.Nested.CustomTypeNested;
		Check:GoTSRPC.Demo.Check;
	}
	// github.com/foomo/gotsrpc/demo.Err
	export interface Err {
		message:string;
	}
	// github.com/foomo/gotsrpc/demo.Inner
	export interface Inner {
		one:string;
	}
	// github.com/foomo/gotsrpc/demo.OuterInline
	export interface OuterInline {
		one:string;
		two:string;
	}
	// github.com/foomo/gotsrpc/demo.OuterNested
	export interface OuterNested {
		inner:GoTSRPC.Demo.Inner;
		two:string;
	}
	// github.com/foomo/gotsrpc/demo.Person
	export interface Person {
		Name:string;
		address?:GoTSRPC.Demo.Address;
		AddressStruct:GoTSRPC.Demo.Address;
		Addresses:{[index:string]:GoTSRPC.Demo.Address};
		InlinePtr?:{
			Foo:boolean;
		};
		InlineStruct:{
			Bar:string;
		};
		DNA:string;
	}
	// github.com/foomo/gotsrpc/demo.ScalarError
	export type ScalarError = string
	// github.com/foomo/gotsrpc/demo.ScalarInPlace
	export type ScalarInPlace = string
}