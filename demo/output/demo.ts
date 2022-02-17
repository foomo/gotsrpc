/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_demo from './demo'; // demo/output/demo.ts to demo/output/demo.ts
import * as github_com_foomo_gotsrpc_v2_demo_nested from './demo-nested'; // demo/output/demo.ts to demo/output/demo-nested.ts
// github.com/foomo/gotsrpc/v2/demo.Address
export interface Address {
	city?:string;
	signs?:Array<string>|null;
	PeoplePtr:Array<github_com_foomo_gotsrpc_v2_demo.Person|null>|null;
	ArrayOfMaps:Array<Record<string,boolean>|null>|null;
	ArrayArrayAddress:Array<Array<github_com_foomo_gotsrpc_v2_demo.Address|null>|null>|null;
	People:Array<github_com_foomo_gotsrpc_v2_demo.Person>|null;
	MapCrap:Record<string,Record<number,boolean>|null>|null;
	NestedPtr:github_com_foomo_gotsrpc_v2_demo_nested.Nested|null;
	NestedStruct:github_com_foomo_gotsrpc_v2_demo_nested.Nested;
}
// github.com/foomo/gotsrpc/v2/demo.AttributeDefinition
export interface AttributeDefinition {
	Key:string;
	Value:string;
}
// github.com/foomo/gotsrpc/v2/demo.AttributeID
export type AttributeID = string
// github.com/foomo/gotsrpc/v2/demo.AttributeMapping
export type AttributeMapping = Record<github_com_foomo_gotsrpc_v2_demo.AttributeID,github_com_foomo_gotsrpc_v2_demo.AttributeDefinition|null>
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
// github.com/foomo/gotsrpc/v2/demo.LocalKey
export type LocalKey = string
// github.com/foomo/gotsrpc/v2/demo.MapOfOtherStuff
export type MapOfOtherStuff = Record<github_com_foomo_gotsrpc_v2_demo_nested.JustAnotherStingType,number>
// github.com/foomo/gotsrpc/v2/demo.MapWithLocalStuff
export type MapWithLocalStuff = Record<github_com_foomo_gotsrpc_v2_demo.LocalKey,number>
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
	address:github_com_foomo_gotsrpc_v2_demo.Address|null;
	AddressStruct:github_com_foomo_gotsrpc_v2_demo.Address;
	Addresses:Record<string,github_com_foomo_gotsrpc_v2_demo.Address|null>|null;
	InlinePtr:{
		Foo:boolean;
	}|null;
	InlineStruct:{
		Bar:string;
	};
	DNA:string|null;
}
// github.com/foomo/gotsrpc/v2/demo.RemoteScalarStruct
export interface RemoteScalarStruct {
	Foo:github_com_foomo_gotsrpc_v2_demo.RemoteScalarsStrings|null;
	Bar:github_com_foomo_gotsrpc_v2_demo.RemoteScalarsStrings|null;
}
// github.com/foomo/gotsrpc/v2/demo.RemoteScalarsStrings
export type RemoteScalarsStrings = Array<github_com_foomo_gotsrpc_v2_demo_nested.JustAnotherStingType>
// github.com/foomo/gotsrpc/v2/demo.ScalarError
export type ScalarError = string
// github.com/foomo/gotsrpc/v2/demo.ScalarInPlace
export type ScalarInPlace = string
// end of common js