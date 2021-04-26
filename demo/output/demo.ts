/* eslint:disable */
module GoTSRPC.Demo {
	// github.com/foomo/gotsrpc/demo.Address
	export interface Address {
		city?:string;
		signs?:string[];
		PeoplePtr:GoTSRPC.Demo.Person[];
		ArrayOfMaps:{[index:string]:boolean}[];
		ArrayArrayAddress:GoTSRPC.Demo.Address[][];
		People:GoTSRPC.Demo.Person[];
		MapCrap:{[index:string]:{[index:number]:boolean}};
		NestedPtr?:GoTSRPC.Demo.Nested.Nested;
		NestedStruct:GoTSRPC.Demo.Nested.Nested;
	}
	// github.com/foomo/gotsrpc/demo.Err
	export interface Err {
		message:string;
	}
	// github.com/foomo/gotsrpc/demo.Inner
	export interface Inner {
		one:string;
	}
	// github.com/foomo/gotsrpc/demo.OuterInline
	export interface OuterInline {
		one:string;
		two:string;
	}
	// github.com/foomo/gotsrpc/demo.OuterNested
	export interface OuterNested {
		inner:GoTSRPC.Demo.Inner;
		two:string;
	}
	// github.com/foomo/gotsrpc/demo.Person
	export interface Person {
		Name:string;
		address?:GoTSRPC.Demo.Address;
		AddressStruct:GoTSRPC.Demo.Address;
		Addresses:{[index:string]:GoTSRPC.Demo.Address};
		InlinePtr?:{
			Foo:boolean;
		};
		InlineStruct:{
			Bar:string;
		};
		DNA:string;
	}
}