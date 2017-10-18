/* tslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_demo from './demo'; // demo/output-commonjs/demo.ts to demo/output-commonjs/demo.ts
import * as github_com_foomo_gotsrpc_demo_nested from './demo-nested'; // demo/output-commonjs/demo.ts to demo/output-commonjs/demo-nested.ts
	// github.com/foomo/gotsrpc/demo.Address
	export interface Address {
		city?:string;
		signs?:string[];
		PeoplePtr:github_com_foomo_gotsrpc_demo.Person[];
		ArrayOfMaps:{[index:string]:boolean}[];
		ArrayArrayAddress:github_com_foomo_gotsrpc_demo.Address[][];
		People:github_com_foomo_gotsrpc_demo.Person[];
		MapCrap:{[index:string]:{[index:number]:boolean}};
		NestedPtr?:github_com_foomo_gotsrpc_demo_nested.Nested;
		NestedStruct:github_com_foomo_gotsrpc_demo_nested.Nested;
	}
	// github.com/foomo/gotsrpc/demo.Err
	export interface Err {
		message:string;
	}
	// github.com/foomo/gotsrpc/demo.Person
	export interface Person {
		Name:string;
		address?:github_com_foomo_gotsrpc_demo.Address;
		AddressStruct:github_com_foomo_gotsrpc_demo.Address;
		Addresses:{[index:string]:github_com_foomo_gotsrpc_demo.Address};
		InlinePtr?:{
			Foo:boolean;
		};
		InlineStruct:{
			Bar:string;
		};
		DNA:string;
	}
// end of common js