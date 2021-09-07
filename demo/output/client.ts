/* eslint:disable */
module GoTSRPC {
export const call = (endPoint:string, method:string, args:any[], success:any, err:any) => {
        var request = new XMLHttpRequest();
        request.withCredentials = true;
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
		// this causes problems, when the browser decides to do a cors OPTIONS request
        // request.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
        request.send(JSON.stringify(args));
        request.onload = function() {
            if (request.status == 200) {
				try {
					var data = JSON.parse(request.responseText);
				} catch(e) {
	                err(request, e);
				}
				success.apply(null, data);
            } else {
                err(request);
            }
        };
        request.onerror = function() {
            err(request);
        };
    }

} // close
module GoTSRPC.Demo {
	export class FooClient {
		static defaultInst = new FooClient;
		constructor(public endPoint:string = "/service/foo", public transport = GoTSRPC.call) {  }
		hello(number:number, success:(ret:number) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Hello", [number], success, err);
		}
	}
	export class DemoClient {
		static defaultInst = new DemoClient;
		constructor(public endPoint:string = "/service/demo", public transport = GoTSRPC.call) {  }
		any(any:GoTSRPC.Demo.Nested.Any, anyList:GoTSRPC.Demo.Nested.Any[], anyMap:Record<string,GoTSRPC.Demo.Nested.Any>, success:(ret:GoTSRPC.Demo.Nested.Any, ret_1:GoTSRPC.Demo.Nested.Any[], ret_2:Record<string,GoTSRPC.Demo.Nested.Any>) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Any", [any, anyList, anyMap], success, err);
		}
		arrayOfRemoteScalars(success:(arrayOfRemoteScalars:GoTSRPC.Demo.Nested.JustAnotherStingType) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "ArrayOfRemoteScalars", [], success, err);
		}
		arrayOfRemoteScalarsInAStruct(success:(strct:GoTSRPC.Demo.RemoteScalarStruct) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "ArrayOfRemoteScalarsInAStruct", [], success, err);
		}
		extractAddress(person:GoTSRPC.Demo.Person, success:(addr:GoTSRPC.Demo.Address, e:GoTSRPC.Demo.Err) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "ExtractAddress", [person], success, err);
		}
		giveMeAScalar(success:(amount:GoTSRPC.Demo.Nested.Amount, wahr:GoTSRPC.Demo.Nested.True, hier:GoTSRPC.Demo.ScalarInPlace) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "GiveMeAScalar", [], success, err);
		}
		hello(name:string, success:(ret:string, ret_1:GoTSRPC.Demo.Err) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Hello", [name], success, err);
		}
		helloInterface(anything:any, anythingMap:Record<string,any>, anythingSlice:any[], success:() => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "HelloInterface", [anything, anythingMap, anythingSlice], success, err);
		}
		helloLocalMapType(success:(localStuff:GoTSRPC.Demo.MapWithLocalStuff) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "HelloLocalMapType", [], success, err);
		}
		helloMapType(success:(otherStuff:GoTSRPC.Demo.MapOfOtherStuff) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "HelloMapType", [], success, err);
		}
		helloNumberMaps(intMap:Record<number,string>, success:(floatMap:Record<number,string>) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "HelloNumberMaps", [intMap], success, err);
		}
		helloScalarError(success:(err:GoTSRPC.Demo.ScalarError) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "HelloScalarError", [], success, err);
		}
		mapCrap(success:(crap:Record<string,number[]>) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "MapCrap", [], success, err);
		}
		nest(success:(ret:GoTSRPC.Demo.Nested.Nested[]) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Nest", [], success, err);
		}
		testScalarInPlace(success:(ret:GoTSRPC.Demo.ScalarInPlace) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "TestScalarInPlace", [], success, err);
		}
	}
	export class BarClient {
		static defaultInst = new BarClient;
		constructor(public endPoint:string = "/service/bar", public transport = GoTSRPC.call) {  }
		attributeMapping(success:(ret:GoTSRPC.Demo.AttributeMapping) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "AttributeMapping", [], success, err);
		}
		customError(one:GoTSRPC.Demo.CustomError, two:GoTSRPC.Demo.CustomError, success:(three:GoTSRPC.Demo.CustomError, four:GoTSRPC.Demo.CustomError) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "CustomError", [one, two], success, err);
		}
		customType(customTypeInt:GoTSRPC.Demo.CustomTypeInt, customTypeString:GoTSRPC.Demo.CustomTypeString, CustomTypeStruct:GoTSRPC.Demo.CustomTypeStruct, success:(ret:GoTSRPC.Demo.CustomTypeInt, ret_1:GoTSRPC.Demo.CustomTypeString, ret_2:GoTSRPC.Demo.CustomTypeStruct) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "CustomType", [customTypeInt, customTypeString, CustomTypeStruct], success, err);
		}
		hello(number:number, success:(ret:number) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Hello", [number], success, err);
		}
		inheritance(inner:GoTSRPC.Demo.Inner, nested:GoTSRPC.Demo.OuterNested, inline:GoTSRPC.Demo.OuterInline, success:(ret:GoTSRPC.Demo.Inner, ret_1:GoTSRPC.Demo.OuterNested, ret_2:GoTSRPC.Demo.OuterInline) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Inheritance", [inner, nested, inline], success, err);
		}
		repeat(one:string, two:string, success:(three:boolean, four:boolean) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Repeat", [one, two], success, err);
		}
	}
}