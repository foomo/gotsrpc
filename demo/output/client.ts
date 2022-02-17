/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_demo from './demo'; // demo/output/client.ts to demo/output/demo.ts
import * as github_com_foomo_gotsrpc_v2_demo_nested from './demo-nested'; // demo/output/client.ts to demo/output/demo-nested.ts

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
	async any(any:github_com_foomo_gotsrpc_v2_demo_nested.Any, anyList:Array<github_com_foomo_gotsrpc_v2_demo_nested.Any>|null, anyMap:Record<string,github_com_foomo_gotsrpc_v2_demo_nested.Any>|null):Promise<{ret:github_com_foomo_gotsrpc_v2_demo_nested.Any; ret_1:Array<github_com_foomo_gotsrpc_v2_demo_nested.Any>|null; ret_2:Record<string,github_com_foomo_gotsrpc_v2_demo_nested.Any>|null}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_v2_demo_nested.Any; 1:Array<github_com_foomo_gotsrpc_v2_demo_nested.Any>|null; 2:Record<string,github_com_foomo_gotsrpc_v2_demo_nested.Any>|null}>("Any", [any, anyList, anyMap])
		let responseObject = {ret : response[0], ret_1 : response[1], ret_2 : response[2]};
		return responseObject;
	}
	async arrayOfRemoteScalars():Promise<github_com_foomo_gotsrpc_v2_demo.RemoteScalarsStrings|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.RemoteScalarsStrings|null}>("ArrayOfRemoteScalars", []))[0]
	}
	async arrayOfRemoteScalarsInAStruct():Promise<github_com_foomo_gotsrpc_v2_demo.RemoteScalarStruct> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.RemoteScalarStruct}>("ArrayOfRemoteScalarsInAStruct", []))[0]
	}
	async extractAddress(person:github_com_foomo_gotsrpc_v2_demo.Person|null):Promise<github_com_foomo_gotsrpc_v2_demo.Address|null> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.Address|null; 1:github_com_foomo_gotsrpc_v2_demo.Err|null}>("ExtractAddress", [person])
		let err = response[1];
		if(err) { throw err }
		return response[0]
	}
	async giveMeAScalar():Promise<{amount:github_com_foomo_gotsrpc_v2_demo_nested.Amount; wahr:github_com_foomo_gotsrpc_v2_demo_nested.True; hier:github_com_foomo_gotsrpc_v2_demo.ScalarInPlace}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_v2_demo_nested.Amount; 1:github_com_foomo_gotsrpc_v2_demo_nested.True; 2:github_com_foomo_gotsrpc_v2_demo.ScalarInPlace}>("GiveMeAScalar", [])
		let responseObject = {amount : response[0], wahr : response[1], hier : response[2]};
		return responseObject;
	}
	async hello(name:string):Promise<string> {
		let response = await this.transport<{0:string; 1:github_com_foomo_gotsrpc_v2_demo.Err|null}>("Hello", [name])
		let err = response[1];
		if(err) { throw err }
		return response[0]
	}
	async helloInterface(anything:any, anythingMap:Record<string,any>|null, anythingSlice:Array<any>|null):Promise<void> {
		await this.transport<void>("HelloInterface", [anything, anythingMap, anythingSlice])
	}
	async helloLocalMapType():Promise<github_com_foomo_gotsrpc_v2_demo.MapWithLocalStuff|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.MapWithLocalStuff|null}>("HelloLocalMapType", []))[0]
	}
	async helloMapType():Promise<github_com_foomo_gotsrpc_v2_demo.MapOfOtherStuff|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.MapOfOtherStuff|null}>("HelloMapType", []))[0]
	}
	async helloNumberMaps(intMap:Record<number,string>|null):Promise<Record<number,string>|null> {
		return (await this.transport<{0:Record<number,string>|null}>("HelloNumberMaps", [intMap]))[0]
	}
	async helloScalarError():Promise<github_com_foomo_gotsrpc_v2_demo.ScalarError|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.ScalarError|null}>("HelloScalarError", []))[0]
	}
	async mapCrap():Promise<Record<string,Array<number>|null>|null> {
		return (await this.transport<{0:Record<string,Array<number>|null>|null}>("MapCrap", []))[0]
	}
	async nest():Promise<Array<github_com_foomo_gotsrpc_v2_demo_nested.Nested|null>|null> {
		return (await this.transport<{0:Array<github_com_foomo_gotsrpc_v2_demo_nested.Nested|null>|null}>("Nest", []))[0]
	}
	async testScalarInPlace():Promise<github_com_foomo_gotsrpc_v2_demo.ScalarInPlace> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.ScalarInPlace}>("TestScalarInPlace", []))[0]
	}
}
export class BarClient {
	public static defaultEndpoint = "/service/bar";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async attributeMapping():Promise<github_com_foomo_gotsrpc_v2_demo.AttributeMapping|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.AttributeMapping|null}>("AttributeMapping", []))[0]
	}
	async customError(one:github_com_foomo_gotsrpc_v2_demo.CustomError, two:github_com_foomo_gotsrpc_v2_demo.CustomError|null):Promise<{three:github_com_foomo_gotsrpc_v2_demo.CustomError; four:github_com_foomo_gotsrpc_v2_demo.CustomError|null}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.CustomError; 1:github_com_foomo_gotsrpc_v2_demo.CustomError|null}>("CustomError", [one, two])
		let responseObject = {three : response[0], four : response[1]};
		return responseObject;
	}
	async customType(customTypeInt:github_com_foomo_gotsrpc_v2_demo.CustomTypeInt, customTypeString:github_com_foomo_gotsrpc_v2_demo.CustomTypeString, CustomTypeStruct:github_com_foomo_gotsrpc_v2_demo.CustomTypeStruct):Promise<{ret:github_com_foomo_gotsrpc_v2_demo.CustomTypeInt|null; ret_1:github_com_foomo_gotsrpc_v2_demo.CustomTypeString|null; ret_2:github_com_foomo_gotsrpc_v2_demo.CustomTypeStruct}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.CustomTypeInt|null; 1:github_com_foomo_gotsrpc_v2_demo.CustomTypeString|null; 2:github_com_foomo_gotsrpc_v2_demo.CustomTypeStruct}>("CustomType", [customTypeInt, customTypeString, CustomTypeStruct])
		let responseObject = {ret : response[0], ret_1 : response[1], ret_2 : response[2]};
		return responseObject;
	}
	async hello(number:number):Promise<number> {
		return (await this.transport<{0:number}>("Hello", [number]))[0]
	}
	async inheritance(inner:github_com_foomo_gotsrpc_v2_demo.Inner, nested:github_com_foomo_gotsrpc_v2_demo.OuterNested, inline:github_com_foomo_gotsrpc_v2_demo.OuterInline):Promise<{ret:github_com_foomo_gotsrpc_v2_demo.Inner; ret_1:github_com_foomo_gotsrpc_v2_demo.OuterNested; ret_2:github_com_foomo_gotsrpc_v2_demo.OuterInline}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_v2_demo.Inner; 1:github_com_foomo_gotsrpc_v2_demo.OuterNested; 2:github_com_foomo_gotsrpc_v2_demo.OuterInline}>("Inheritance", [inner, nested, inline])
		let responseObject = {ret : response[0], ret_1 : response[1], ret_2 : response[2]};
		return responseObject;
	}
	async repeat(one:string, two:string):Promise<{three:boolean; four:boolean}> {
		let response = await this.transport<{0:boolean; 1:boolean}>("Repeat", [one, two])
		let responseObject = {three : response[0], four : response[1]};
		return responseObject;
	}
}