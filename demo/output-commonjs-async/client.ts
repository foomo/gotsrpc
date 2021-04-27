/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_demo from './demo'; // demo/output-commonjs-async/client.ts to demo/output-commonjs-async/demo.ts
import * as github_com_foomo_gotsrpc_demo_nested from './demo-nested'; // demo/output-commonjs-async/client.ts to demo/output-commonjs-async/demo-nested.ts

export class FooClient {
	public static defaultEndpoint = "/service/foo";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async hello(number:number):Promise<number> {
		return (await this.transport<{0:number}>("Hello", [number]))[0]
	}
}
export class DemoClient {
	public static defaultEndpoint = "/service/demo";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async any(any:github_com_foomo_gotsrpc_demo_nested.Any, anyList:github_com_foomo_gotsrpc_demo_nested.Any[], anyMap:{[index:string]:github_com_foomo_gotsrpc_demo_nested.Any}):Promise<{ret:github_com_foomo_gotsrpc_demo_nested.Any; ret_1:github_com_foomo_gotsrpc_demo_nested.Any[]; ret_2:{[index:string]:github_com_foomo_gotsrpc_demo_nested.Any}}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_demo_nested.Any; 1:github_com_foomo_gotsrpc_demo_nested.Any[]; 2:{[index:string]:github_com_foomo_gotsrpc_demo_nested.Any}}>("Any", [any, anyList, anyMap])
		let responseObject = {ret : response[0], ret_1 : response[1], ret_2 : response[2]};
		return responseObject;
	}
	async extractAddress(person:github_com_foomo_gotsrpc_demo.Person):Promise<{addr:github_com_foomo_gotsrpc_demo.Address; e:github_com_foomo_gotsrpc_demo.Err}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_demo.Address; 1:github_com_foomo_gotsrpc_demo.Err}>("ExtractAddress", [person])
		let responseObject = {addr : response[0], e : response[1]};
		return responseObject;
	}
	async giveMeAScalar():Promise<{amount:github_com_foomo_gotsrpc_demo_nested.Amount; wahr:github_com_foomo_gotsrpc_demo_nested.True; hier:github_com_foomo_gotsrpc_demo.ScalarInPlace}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_demo_nested.Amount; 1:github_com_foomo_gotsrpc_demo_nested.True; 2:github_com_foomo_gotsrpc_demo.ScalarInPlace}>("GiveMeAScalar", [])
		let responseObject = {amount : response[0], wahr : response[1], hier : response[2]};
		return responseObject;
	}
	async hello(name:string):Promise<string> {
		let response = await this.transport<{0:string; 1:github_com_foomo_gotsrpc_demo.Err}>("Hello", [name])
		let err = response[1];
		if(err) { throw err }
		return response[0]
	}
	async helloInterface(anything:any, anythingMap:{[index:string]:any}, anythingSlice:any[]):Promise<void> {
		await this.transport<void>("HelloInterface", [anything, anythingMap, anythingSlice])
	}
	async helloNumberMaps(intMap:{[index:number]:string}):Promise<{[index:number]:string}> {
		return (await this.transport<{0:{[index:number]:string}}>("HelloNumberMaps", [intMap]))[0]
	}
	async helloScalarError():Promise<github_com_foomo_gotsrpc_demo.ScalarError> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_demo.ScalarError}>("HelloScalarError", []))[0]
	}
	async mapCrap():Promise<{[index:string]:number[]}> {
		return (await this.transport<{0:{[index:string]:number[]}}>("MapCrap", []))[0]
	}
	async nest():Promise<github_com_foomo_gotsrpc_demo_nested.Nested[]> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_demo_nested.Nested[]}>("Nest", []))[0]
	}
	async testScalarInPlace():Promise<github_com_foomo_gotsrpc_demo.ScalarInPlace> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_demo.ScalarInPlace}>("TestScalarInPlace", []))[0]
	}
}
export class BarClient {
	public static defaultEndpoint = "/service/bar";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async customType(customTypeInt:github_com_foomo_gotsrpc_demo.CustomTypeInt, customTypeString:github_com_foomo_gotsrpc_demo.CustomTypeString, CustomTypeStruct:github_com_foomo_gotsrpc_demo.CustomTypeStruct):Promise<{ret:github_com_foomo_gotsrpc_demo.CustomTypeInt; ret_1:github_com_foomo_gotsrpc_demo.CustomTypeString; ret_2:github_com_foomo_gotsrpc_demo.CustomTypeStruct}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_demo.CustomTypeInt; 1:github_com_foomo_gotsrpc_demo.CustomTypeString; 2:github_com_foomo_gotsrpc_demo.CustomTypeStruct}>("CustomType", [customTypeInt, customTypeString, CustomTypeStruct])
		let responseObject = {ret : response[0], ret_1 : response[1], ret_2 : response[2]};
		return responseObject;
	}
}