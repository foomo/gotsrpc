/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_union_service from './service-vo'; // ./client/src/service-vo.ts to ./client/src/service-vo.ts
// github.com/foomo/gotsrpc/v2/example/union/service.InlineStruct
export interface InlineStruct extends github_com_foomo_gotsrpc_v2_example_union_service.InlineStructA , github_com_foomo_gotsrpc_v2_example_union_service.InlineStructB {
	value:string;
}
// github.com/foomo/gotsrpc/v2/example/union/service.InlineStructA
export interface InlineStructA {
	valueA:string;
}
// github.com/foomo/gotsrpc/v2/example/union/service.InlineStructB
export interface InlineStructB {
	valueB:string;
}
// github.com/foomo/gotsrpc/v2/example/union/service.InlineStructPtr
export interface InlineStructPtr extends Partial<github_com_foomo_gotsrpc_v2_example_union_service.InlineStructA> , Partial<github_com_foomo_gotsrpc_v2_example_union_service.InlineStructB> {
	bug?:github_com_foomo_gotsrpc_v2_example_union_service.InlineStructB;
	value:string;
}
// github.com/foomo/gotsrpc/v2/example/union/service.UnionString
export type UnionString = (typeof github_com_foomo_gotsrpc_v2_example_union_service.UnionStringA) & (typeof github_com_foomo_gotsrpc_v2_example_union_service.UnionStringB)
// github.com/foomo/gotsrpc/v2/example/union/service.UnionStringA
export enum UnionStringA {
	One = "one",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/example/union/service.UnionStringB
export enum UnionStringB {
	Four = "four",
	Three = "three",
}
// github.com/foomo/gotsrpc/v2/example/union/service.UnionStruct
export type UnionStruct = github_com_foomo_gotsrpc_v2_example_union_service.UnionStructA | github_com_foomo_gotsrpc_v2_example_union_service.UnionStructB | undefined
// github.com/foomo/gotsrpc/v2/example/union/service.UnionStructA
export interface UnionStructA {
	kind:'UnionStructA';
	value:github_com_foomo_gotsrpc_v2_example_union_service.UnionStructAValueA;
	bar:string;
}
// github.com/foomo/gotsrpc/v2/example/union/service.UnionStructAValueA
export enum UnionStructAValueA {
	One = "one",
	Three = "three",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/example/union/service.UnionStructAValueB
export enum UnionStructAValueB {
	One = "one",
	Three = "three",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/example/union/service.UnionStructB
export interface UnionStructB {
	kind:'UnionStructB';
	value:github_com_foomo_gotsrpc_v2_example_union_service.UnionStructAValueB;
	foo:string;
}
// end of common js