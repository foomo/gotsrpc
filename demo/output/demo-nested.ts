module GoTSRPC.Demo.Nested {
	// github.com/foomo/gotsrpc/demo/nested.Nested
	export interface Nested {
		Name:string;
		SuperNestedString:{
			Ha:number;
		};
		SuperNestedPtr?:{
			Bla:string;
		};
	}
}