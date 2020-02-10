/* tslint:disable */
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
		extractAddress(person:GoTSRPC.Demo.Person, success:(addr:GoTSRPC.Demo.Address, e:GoTSRPC.Demo.Err) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "ExtractAddress", [person], success, err);
		}
		giveMeAScalar(success:(amount:number, wahr:boolean, hier:string) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "GiveMeAScalar", [], success, err);
		}
		hello(name:string, success:(ret:string, ret_1:GoTSRPC.Demo.Err) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Hello", [name], success, err);
		}
		helloInterface(anything:any, anythingMap:{[index:string]:any}, anythingSlice:any[], success:() => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "HelloInterface", [anything, anythingMap, anythingSlice], success, err);
		}
		helloNumberMaps(intMap:{[index:number]:string}, success:(floatMap:{[index:number]:string}) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "HelloNumberMaps", [intMap], success, err);
		}
		helloScalarError(success:(err:string) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "HelloScalarError", [], success, err);
		}
		mapCrap(success:(crap:{[index:string]:number[]}) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "MapCrap", [], success, err);
		}
		nest(success:(ret:GoTSRPC.Demo.Nested.Nested[]) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "Nest", [], success, err);
		}
		testScalarInPlace(success:(ret:string) => void, err:(request:XMLHttpRequest, e?:Error) => void) {
			this.transport(this.endPoint, "TestScalarInPlace", [], success, err);
		}
	}
}