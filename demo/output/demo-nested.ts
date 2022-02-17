/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_demo from './demo'; // demo/output/demo-nested.ts to demo/output/demo.ts
import * as github_com_foomo_gotsrpc_v2_demo_nested from './demo-nested'; // demo/output/demo-nested.ts to demo/output/demo-nested.ts
// github.com/foomo/gotsrpc/v2/demo/nested.Amount
export type Amount = number
// github.com/foomo/gotsrpc/v2/demo/nested.Any
export type Any = any
// github.com/foomo/gotsrpc/v2/demo/nested.CustomTypeNested
export enum CustomTypeNested {
	One = "one",
	Three = "three",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/demo/nested.JustAnotherStingType
export type JustAnotherStingType = string
// github.com/foomo/gotsrpc/v2/demo/nested.Nested
export interface Nested {
	Name:string;
	Any:github_com_foomo_gotsrpc_v2_demo_nested.Any;
	AnyMap:Record<string,github_com_foomo_gotsrpc_v2_demo_nested.Any>|null;
	AnyList:Array<github_com_foomo_gotsrpc_v2_demo_nested.Any>|null;
	SuperNestedString:{
		Ha:number;
	};
	SuperNestedPtr:{
		Bla:string;
	}|null;
}
// github.com/foomo/gotsrpc/v2/demo/nested.True
export type True = boolean
// end of common js