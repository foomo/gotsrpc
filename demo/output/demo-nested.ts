/* eslint:disable */
module GoTSRPC.Demo.Nested {
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
}