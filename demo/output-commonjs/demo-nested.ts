/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_demo from './demo'; // demo/output-commonjs/demo-nested.ts to demo/output-commonjs/demo.ts
import * as github_com_foomo_gotsrpc_v2_demo_nested from './demo-nested'; // demo/output-commonjs/demo-nested.ts to demo/output-commonjs/demo-nested.ts
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
// github.com/foomo/gotsrpc/v2/demo/nested.Nested
export interface Nested {
	Name:string;
	Any:any;
	AnyMap:Record<string, any>;
	AnyList:any[];
	SuperNestedString:{
		Ha:number;
	};
	SuperNestedPtr?:{
		Bla:string;
	};
}
// github.com/foomo/gotsrpc/v2/demo/nested.True
export type True = boolean
// end of common js