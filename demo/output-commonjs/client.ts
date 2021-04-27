/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_demo from './demo'; // demo/output-commonjs/client.ts to demo/output-commonjs/demo.ts
import * as github_com_foomo_gotsrpc_demo_nested from './demo-nested'; // demo/output-commonjs/client.ts to demo/output-commonjs/demo-nested.ts

export class FooClient {
	constructor(public endPoint:string = "/service/foo", public transport:(endPoint:string, method:string, args:any[], success:any, err:any) => void) {  }
	hello(number:number, success:(ret:number) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "Hello", [number], success, err);
	}
}
export class DemoClient {
	constructor(public endPoint:string = "/service/demo", public transport:(endPoint:string, method:string, args:any[], success:any, err:any) => void) {  }
	any(any:github_com_foomo_gotsrpc_demo_nested.Any, anyList:github_com_foomo_gotsrpc_demo_nested.Any[], anyMap:{[index:string]:github_com_foomo_gotsrpc_demo_nested.Any}, success:(ret:github_com_foomo_gotsrpc_demo_nested.Any, ret_1:github_com_foomo_gotsrpc_demo_nested.Any[], ret_2:{[index:string]:github_com_foomo_gotsrpc_demo_nested.Any}) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "Any", [any, anyList, anyMap], success, err);
	}
	extractAddress(person:github_com_foomo_gotsrpc_demo.Person, success:(addr:github_com_foomo_gotsrpc_demo.Address, e:github_com_foomo_gotsrpc_demo.Err) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "ExtractAddress", [person], success, err);
	}
	giveMeAScalar(success:(amount:github_com_foomo_gotsrpc_demo_nested.Amount, wahr:github_com_foomo_gotsrpc_demo_nested.True, hier:github_com_foomo_gotsrpc_demo.ScalarInPlace) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "GiveMeAScalar", [], success, err);
	}
	hello(name:string, success:(ret:string, ret_1:github_com_foomo_gotsrpc_demo.Err) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "Hello", [name], success, err);
	}
	helloInterface(anything:any, anythingMap:{[index:string]:any}, anythingSlice:any[], success:() => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "HelloInterface", [anything, anythingMap, anythingSlice], success, err);
	}
	helloNumberMaps(intMap:{[index:number]:string}, success:(floatMap:{[index:number]:string}) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "HelloNumberMaps", [intMap], success, err);
	}
	helloScalarError(success:(err:github_com_foomo_gotsrpc_demo.ScalarError) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "HelloScalarError", [], success, err);
	}
	mapCrap(success:(crap:{[index:string]:number[]}) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "MapCrap", [], success, err);
	}
	nest(success:(ret:github_com_foomo_gotsrpc_demo_nested.Nested[]) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "Nest", [], success, err);
	}
	testScalarInPlace(success:(ret:github_com_foomo_gotsrpc_demo.ScalarInPlace) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "TestScalarInPlace", [], success, err);
	}
}
export class BarClient {
	constructor(public endPoint:string = "/service/bar", public transport:(endPoint:string, method:string, args:any[], success:any, err:any) => void) {  }
	customType(customTypeInt:github_com_foomo_gotsrpc_demo.CustomTypeInt, customTypeString:github_com_foomo_gotsrpc_demo.CustomTypeString, CustomTypeStruct:github_com_foomo_gotsrpc_demo.CustomTypeStruct, success:(ret:github_com_foomo_gotsrpc_demo.CustomTypeInt, ret_1:github_com_foomo_gotsrpc_demo.CustomTypeString, ret_2:github_com_foomo_gotsrpc_demo.CustomTypeStruct) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
		this.transport(this.endPoint, "CustomType", [customTypeInt, customTypeString, CustomTypeStruct], success, err);
	}
}