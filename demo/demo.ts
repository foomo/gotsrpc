
module GoTSRPC {
    export function call(endPoint:string, method:string, args:any[], success:any, err:any) {
        var request = new XMLHttpRequest();
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.send(JSON.stringify(args));            
        request.onload = function() {
            if (request.status == 200) {
                var data = JSON.parse(request.responseText);
                success.apply(null, data);
            } else {
                err(request)
            }
        };            
        request.onerror = function() {
            err(request);
        };
    }
}

module Demo {
    export interface Err {
        Message:string;
    }
    export class ServiceClient {
        static defaultInst = new ServiceClient()
        constructor(public endPoint:string = "/service") {  }
        hello(name:string, success:(reply:string, err:Err) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "Hello", [name], success, err);
        }
    }
}

var handleCrap = (err:Demo.Err, request:XMLHttpRequest) => {
    if(err) {
        console.log("fuckit logic");        
    } else if(request) {
        console.warn("request crap", request);
    } else {
        console.log("no crap", err);
    }
}

Demo.ServiceClient.defaultInst.hello(
    "Hansi",
    (reply:string, err:Demo.Err) => {
        console.log("server says hello to Hansi", reply, err);
        handleCrap(err, null);
    },
    (request:XMLHttpRequest) => {
        console.log("wtf", request);
        handleCrap(null, request);        
    }
);    


Demo.ServiceClient.defaultInst.hello(
    "Peter",
    (reply:string, err:Demo.Err) => {
        console.log("server should not like Peter", reply, err);
        handleCrap(err, null);
    },
    (request:XMLHttpRequest) => {
        console.log("wtf", request);
        handleCrap(null, request);
    }
);    
