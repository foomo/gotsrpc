/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_nullable_service from './service-vo'; // ./client/src/service-vo.ts to ./client/src/service-vo.ts
// github.com/foomo/gotsrpc/v2/example/nullable/service.ACustomType
export enum ACustomType {
	One = "one",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/example/nullable/service.ACustomTypes
export type ACustomTypes = Array<github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomType>
// github.com/foomo/gotsrpc/v2/example/nullable/service.ACustomTypesMap
export type ACustomTypesMap = Record<github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomType,github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomType>
// github.com/foomo/gotsrpc/v2/example/nullable/service.BCustomType
export type BCustomType = string
// github.com/foomo/gotsrpc/v2/example/nullable/service.BCustomTypes
export type BCustomTypes = Array<github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomType>
// github.com/foomo/gotsrpc/v2/example/nullable/service.BCustomTypesMap
export type BCustomTypesMap = Record<github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomType,github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomType>
// github.com/foomo/gotsrpc/v2/example/nullable/service.Base
export interface Base {
	a1:github_com_foomo_gotsrpc_v2_example_nullable_service.Nested;
	a2?:github_com_foomo_gotsrpc_v2_example_nullable_service.Nested;
	a3:github_com_foomo_gotsrpc_v2_example_nullable_service.Nested|null;
	b1:string;
	b2?:string;
	b3:string|null;
	c1:any;
	c2?:any;
	c3:any;
	d1:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomType;
	d2?:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomType;
	d3:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomType|null;
	e1:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomTypes|null;
	e2?:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomTypes;
	e3:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomTypes|null;
	f1:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomTypesMap|null;
	f2?:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomTypesMap;
	f3:github_com_foomo_gotsrpc_v2_example_nullable_service.ACustomTypesMap|null;
	Two:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Nested>|null;
	Two1:Array<Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Nested>|null>|null;
	Two2:Array<Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Nested>|null>|null;
	Three:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Nested|null>|null;
	Three1:Array<string|null>|null;
	Four:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Nested>|null;
	Five:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Nested|null>|null;
	Six:{
		Foo:string;
	}|null;
}
// github.com/foomo/gotsrpc/v2/example/nullable/service.Nested
export interface Nested {
	Foo:string;
}
// end of common js