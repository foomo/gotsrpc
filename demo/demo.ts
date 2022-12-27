
/// <reference path="output/demo.ts"/>
/// <reference path="output/demo-nested.ts"/>
/// <reference path="output/client.ts"/>

var handleCrap = (err:GoTSRPC.Demo.Err, request:XMLHttpRequest) => {
    if(err) {
        console.log("fuckit logic");        
    } else if(request) {
        console.warn("request crap", request);
    } else {
        console.log("no crap", err);
    }
}

GoTSRPC.Demo.DemoClient.defaultInst.hello(
    "Hansi",
    (reply:string, err:GoTSRPC.Demo.Err) => {
        console.log("server says hello to Hansi", reply, err);
        handleCrap(err, null);
    },
    (request:XMLHttpRequest) => {
        console.log("wtf", request);
        handleCrap(null, request);        
    }
);    


GoTSRPC.Demo.DemoClient.defaultInst.hello(
    "Peter",
    (reply:string, err:GoTSRPC.Demo.Err) => {
        console.log("I particularly like you Peter ;)", reply, err);
        handleCrap(err, null);
    },
    (request:XMLHttpRequest) => {
        console.log("wtf", request);
        handleCrap(null, request);
    }
);    
