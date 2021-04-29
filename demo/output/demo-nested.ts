/* eslint:disable */
module GoTSRPC.Demo.Nested {
	// github.com/foomo/gotsrpc/demo/nested.Amount
	export type Amount = number
	// github.com/foomo/gotsrpc/demo/nested.Any
	export type Any = any
	// github.com/foomo/gotsrpc/demo/nested.CustomTypeNested
	export enum CustomTypeNested {
		One = "one",
		Three = "three",
		Two = "two",
	}
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
}