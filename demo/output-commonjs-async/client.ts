/* tslint:disable */
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
	async any(any:any, anyList:any[], anyMap:{[index:string]:any}):Promise<{ret:any; ret_1:any[]; ret_2:{[index:string]:any}}> {
		let response = await this.transport<{0:any; 1:any[]; 2:{[index:string]:any}}>("Any", [any, anyList, anyMap])
		let responseObject = {ret : response[0], ret_1 : response[1], ret_2 : response[2]};
		return responseObject;
	}
	async extractAddress(person:github_com_foomo_gotsrpc_demo.Person):Promise<github_com_foomo_gotsrpc_demo.Address> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_demo.Address; 1:github_com_foomo_gotsrpc_demo.Err}>("ExtractAddress", [person])
		let err = response[1];
		if(err) { throw err }
		return response[0]
	}
	async giveMeAScalar():Promise<{amount:number; wahr:boolean; hier:string}> {
		let response = await this.transport<{0:number; 1:boolean; 2:string}>("GiveMeAScalar", [])
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
	async helloScalarError():Promise<string> {
		return (await this.transport<{0:string}>("HelloScalarError", []))[0]
	}
	async mapCrap():Promise<{[index:string]:number[]}> {
		return (await this.transport<{0:{[index:string]:number[]}}>("MapCrap", []))[0]
	}
	async nest():Promise<github_com_foomo_gotsrpc_demo_nested.Nested[]> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_demo_nested.Nested[]}>("Nest", []))[0]
	}
	async testScalarInPlace():Promise<string> {
		return (await this.transport<{0:string}>("TestScalarInPlace", []))[0]
	}
}
export class BarClient {
	public static defaultEndpoint = "/service/bar";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async hello(number:number):Promise<number> {
		return (await this.transport<{0:number}>("Hello", [number]))[0]
	}
}