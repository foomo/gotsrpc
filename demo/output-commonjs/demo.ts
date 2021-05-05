/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_demo from './demo'; // demo/output-commonjs/demo.ts to demo/output-commonjs/demo.ts
import * as github_com_foomo_gotsrpc_v2_demo_nested from './demo-nested'; // demo/output-commonjs/demo.ts to demo/output-commonjs/demo-nested.ts
// github.com/foomo/gotsrpc/v2/demo.Address
export interface Address {
	city?:string;
	signs?:string[];
	PeoplePtr:github_com_foomo_gotsrpc_v2_demo.Person[];
	ArrayOfMaps:{[index:string]:boolean}[];
	ArrayArrayAddress:github_com_foomo_gotsrpc_v2_demo.Address[][];
	People:github_com_foomo_gotsrpc_v2_demo.Person[];
	MapCrap:{[index:string]:{[index:number]:boolean}};
	NestedPtr?:github_com_foomo_gotsrpc_v2_demo_nested.Nested;
	NestedStruct:github_com_foomo_gotsrpc_v2_demo_nested.Nested;
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
	CustomTypeFoo:github_com_foomo_gotsrpc_v2_demo.CustomTypeFoo;
	CustomTypeInt:github_com_foomo_gotsrpc_v2_demo.CustomTypeInt;
	CustomTypeString:github_com_foomo_gotsrpc_v2_demo.CustomTypeString;
	CustomTypeNested:github_com_foomo_gotsrpc_v2_demo_nested.CustomTypeNested;
	Check:github_com_foomo_gotsrpc_v2_demo.Check;
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
	inner:github_com_foomo_gotsrpc_v2_demo.Inner;
	two:string;
}
// github.com/foomo/gotsrpc/v2/demo.Person
export interface Person {
	Name:string;
	address?:github_com_foomo_gotsrpc_v2_demo.Address;
	AddressStruct:github_com_foomo_gotsrpc_v2_demo.Address;
	Addresses:{[index:string]:github_com_foomo_gotsrpc_v2_demo.Address};
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
// end of common js