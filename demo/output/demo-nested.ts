/* eslint:disable */
module GoTSRPC.Demo.Nested {
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
}