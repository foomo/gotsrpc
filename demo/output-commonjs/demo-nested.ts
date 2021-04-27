/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_demo from './demo'; // demo/output-commonjs/demo-nested.ts to demo/output-commonjs/demo.ts
import * as github_com_foomo_gotsrpc_demo_nested from './demo-nested'; // demo/output-commonjs/demo-nested.ts to demo/output-commonjs/demo-nested.ts
// github.com/foomo/gotsrpc/demo/nested.Amount
export type Amount = number
// github.com/foomo/gotsrpc/demo/nested.Any
export type Any = any
// github.com/foomo/gotsrpc/demo/nested.CustomTypeNested
export type CustomTypeNested = "one" | "two" | "three"
// github.com/foomo/gotsrpc/demo/nested.Nested
export interface Nested {
	Name:string;
	Any:any;
	AnyMap:{[index:string]:any};
	AnyList:any[];
	SuperNestedString:{
		Ha:number;
	};
	SuperNestedPtr?:{
		Bla:string;
	};
}
// github.com/foomo/gotsrpc/demo/nested.True
export type True = boolean
// constants from github.com/foomo/gotsrpc/demo/nested
export const GoConst = {
	CustomTypeNestedOne : "one",
	CustomTypeNestedThree : "three",
	CustomTypeNestedTwo : "two",
}
// end of common js