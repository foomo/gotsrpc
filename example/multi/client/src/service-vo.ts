/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_multi_service from './service-vo'; // ./client/src/service-vo.ts to ./client/src/service-vo.ts
// github.com/foomo/gotsrpc/v2/example/multi/service.InlineStruct
export interface InlineStruct extends github_com_foomo_gotsrpc_v2_example_multi_service.InlineStructA , github_com_foomo_gotsrpc_v2_example_multi_service.InlineStructB {
	value:string;
}
// github.com/foomo/gotsrpc/v2/example/multi/service.InlineStructA
export interface InlineStructA {
	valueA:string;
}
// github.com/foomo/gotsrpc/v2/example/multi/service.InlineStructB
export interface InlineStructB {
	valueB:string;
}
// github.com/foomo/gotsrpc/v2/example/multi/service.InlineStructPtr
export interface InlineStructPtr extends Partial<github_com_foomo_gotsrpc_v2_example_multi_service.InlineStructA> , Partial<github_com_foomo_gotsrpc_v2_example_multi_service.InlineStructB> {
	bug?:github_com_foomo_gotsrpc_v2_example_multi_service.InlineStructB;
	value:string;
}
// github.com/foomo/gotsrpc/v2/example/multi/service.UnionString
export type UnionString = (typeof github_com_foomo_gotsrpc_v2_example_multi_service.UnionStringA) & (typeof github_com_foomo_gotsrpc_v2_example_multi_service.UnionStringB)
// github.com/foomo/gotsrpc/v2/example/multi/service.UnionStringA
export enum UnionStringA {
	One = "one",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/example/multi/service.UnionStringB
export enum UnionStringB {
	Four = "four",
	Three = "three",
}
// github.com/foomo/gotsrpc/v2/example/multi/service.UnionStruct
export type UnionStruct = github_com_foomo_gotsrpc_v2_example_multi_service.UnionStructA | github_com_foomo_gotsrpc_v2_example_multi_service.UnionStructB | undefined
// github.com/foomo/gotsrpc/v2/example/multi/service.UnionStructA
export interface UnionStructA {
	kind:'UnionStructA';
	value:github_com_foomo_gotsrpc_v2_example_multi_service.UnionStructAValueA;
	bar:string;
}
// github.com/foomo/gotsrpc/v2/example/multi/service.UnionStructAValueA
export enum UnionStructAValueA {
	One = "one",
	Three = "three",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/example/multi/service.UnionStructAValueB
export enum UnionStructAValueB {
	One = "one",
	Three = "three",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/example/multi/service.UnionStructB
export interface UnionStructB {
	kind:'UnionStructB';
	value:github_com_foomo_gotsrpc_v2_example_multi_service.UnionStructAValueB;
	foo:string;
}
// end of common js