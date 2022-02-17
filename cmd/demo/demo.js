/// <reference path="output/demo.ts"/>
/// <reference path="output/demo-nested.ts"/>
/// <reference path="output/client.ts"/>
var handleCrap = function (err, request) {
    if (err) {
        console.log("fuckit logic");
    }
    else if (request) {
        console.warn("request crap", request);
    }
    else {
        console.log("no crap", err);
    }
};
GoTSRPC.Demo.DemoClient.defaultInst.hello("Hansi", function (reply, err) {
    console.log("server says hello to Hansi", reply, err);
    handleCrap(err, null);
}, function (request) {
    console.log("wtf", request);
    handleCrap(null, request);
});
GoTSRPC.Demo.DemoClient.defaultInst.hello("Peter", function (reply, err) {
    console.log("server should not like Peter, sorry Peter ;)", reply, err);
    handleCrap(err, null);
}, function (request) {
    console.log("wtf", request);
    handleCrap(null, request);
});
