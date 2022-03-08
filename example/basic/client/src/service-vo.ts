/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_basic_service from './service-vo'; // ./client/src/service-vo.ts to ./client/src/service-vo.ts
// github.com/foomo/gotsrpc/v2/example/basic/service.Float32Type
export enum Float32Type {
	Float32AType = 1,
	Float32BType = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Float32TypeMapKey
export enum Float32TypeMapKey {
	Float32ATypeMapKey = 1,
	Float32BTypeMapKey = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Float32TypeMapTyped
export type Float32TypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.Float32TypeMapValue
export enum Float32TypeMapValue {
	Float32ATypeMapValue = 1,
	Float32BTypeMapValue = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Float64Type
export enum Float64Type {
	Float64AType = 1,
	Float64BType = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Float64TypeMapKey
export enum Float64TypeMapKey {
	Float64ATypeMapKey = 1,
	Float64BTypeMapKey = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Float64TypeMapTyped
export type Float64TypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.Float64TypeMapValue
export enum Float64TypeMapValue {
	Float64ATypeMapValue = 1,
	Float64BTypeMapValue = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Int32Type
export enum Int32Type {
	Int32AType = 1,
	Int32BType = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Int32TypeMapKey
export enum Int32TypeMapKey {
	Int32ATypeMapKey = 1,
	Int32BTypeMapKey = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Int32TypeMapTyped
export type Int32TypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.Int32TypeMapValue
export enum Int32TypeMapValue {
	Int32ATypeMapValue = 1,
	Int32BTypeMapValue = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Int64Type
export enum Int64Type {
	Int64AType = 1,
	Int64BType = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Int64TypeMapKey
export enum Int64TypeMapKey {
	Int64ATypeMapKey = 1,
	Int64BTypeMapKey = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Int64TypeMapTyped
export type Int64TypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.Int64TypeMapValue
export enum Int64TypeMapValue {
	Int64ATypeMapValue = 1,
	Int64BTypeMapValue = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.IntType
export enum IntType {
	IntAType = 1,
	IntBType = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.IntTypeMapKey
export enum IntTypeMapKey {
	IntATypeMapKey = 1,
	IntBTypeMapKey = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.IntTypeMapTyped
export type IntTypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.IntTypeMapValue
export enum IntTypeMapValue {
	IntATypeMapValue = 1,
	IntBTypeMapValue = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.StringType
export enum StringType {
	StringAType = "A",
	StringBType = "B",
}
// github.com/foomo/gotsrpc/v2/example/basic/service.StringTypeMapKey
export enum StringTypeMapKey {
	StringATypeMapKey = "A",
	StringBTypeMapKey = "B",
}
// github.com/foomo/gotsrpc/v2/example/basic/service.StringTypeMapTyped
export type StringTypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.StringTypeMapValue
export enum StringTypeMapValue {
	StringATypeMapValue = "A",
	StringBTypeMapValue = "B",
}
// github.com/foomo/gotsrpc/v2/example/basic/service.Struct
export interface Struct {
	Int:number;
	Int32:number;
	Int64:number;
	UInt:number;
	UInt32:number;
	UInt64:number;
	Float32:number;
	Float64:number;
	String:string;
	Interface:any;
	IntTypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapValue>|null;
	Int32TypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapValue>|null;
	Int64TypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapValue>|null;
	UIntTypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapValue>|null;
	UInt32TypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapValue>|null;
	UInt64TypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapValue>|null;
	Float32TypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapValue>|null;
	Float64TypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapValue>|null;
	StringTypeMapTyped:Record<github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapValue>|null;
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UInt32Type
export enum UInt32Type {
	UInt32AType = 1,
	UInt32BType = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UInt32TypeMapKey
export enum UInt32TypeMapKey {
	UInt32ATypeMapKey = 1,
	UInt32BTypeMapKey = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UInt32TypeMapTyped
export type UInt32TypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.UInt32TypeMapValue
export enum UInt32TypeMapValue {
	UInt32ATypeMapValue = 1,
	UInt32BTypeMapValue = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UInt64Type
export enum UInt64Type {
	UInt64AType = 1,
	UInt64BType = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UInt64TypeMapKey
export enum UInt64TypeMapKey {
	UInt64ATypeMapKey = 1,
	UInt64BTypeMapKey = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UInt64TypeMapTyped
export type UInt64TypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.UInt64TypeMapValue
export enum UInt64TypeMapValue {
	UInt64ATypeMapValue = 1,
	UInt64BTypeMapValue = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UIntType
export enum UIntType {
	UIntAType = 1,
	UIntBType = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UIntTypeMapKey
export enum UIntTypeMapKey {
	UIntATypeMapKey = 1,
	UIntBTypeMapKey = 2,
}
// github.com/foomo/gotsrpc/v2/example/basic/service.UIntTypeMapTyped
export type UIntTypeMapTyped = Record<github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapValue>
// github.com/foomo/gotsrpc/v2/example/basic/service.UIntTypeMapValue
export enum UIntTypeMapValue {
	UIntATypeMapValue = 1,
	UIntBTypeMapValue = 2,
}
// end of common js