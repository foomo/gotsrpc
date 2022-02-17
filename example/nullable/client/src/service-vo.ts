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
	two:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Nested>|null;
	two1:Array<Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Nested>|null>|null;
	two2:Array<Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Nested>|null>|null;
	three:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Nested|null>|null;
	three1:Array<string|null>|null;
	four:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Nested>|null;
	five:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Nested|null>|null;
	six:{
		foo:string;
	}|null;
}
// github.com/foomo/gotsrpc/v2/example/nullable/service.Nested
export interface Nested {
	foo:string;
}
// end of common js