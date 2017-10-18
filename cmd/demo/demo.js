/* tslint:disable */
var GoTSRPC;
(function (GoTSRPC) {
    GoTSRPC.call = function (endPoint, method, args, success, err) {
        var request = new XMLHttpRequest();
        request.withCredentials = true;
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
        // this causes problems, when the browser decides to do a cors OPTIONS request
        // request.setRequestHeader('Content-Type', 'application/json; charset=UTF-8');
        request.send(JSON.stringify(args));
        request.onload = function () {
            if (request.status == 200) {
                try {
                    var data = JSON.parse(request.responseText);
                }
                catch (e) {
                    err(request, e);
                }
                success.apply(null, data);
            }
            else {
                err(request);
            }
        };
        request.onerror = function () {
            err(request);
        };
    };
})(GoTSRPC || (GoTSRPC = {})); // close
var GoTSRPC;
(function (GoTSRPC) {
    var Demo;
    (function (Demo) {
        var FooClient = (function () {
            function FooClient(endPoint, transport) {
                if (endPoint === void 0) { endPoint = "/service/foo"; }
                if (transport === void 0) { transport = GoTSRPC.call; }
                this.endPoint = endPoint;
                this.transport = transport;
            }
            FooClient.prototype.hello = function (number, success, err) {
                this.transport(this.endPoint, "Hello", [number], success, err);
            };
            FooClient.defaultInst = new FooClient;
            return FooClient;
        }());
        Demo.FooClient = FooClient;
        var DemoClient = (function () {
            function DemoClient(endPoint, transport) {
                if (endPoint === void 0) { endPoint = "/service/demo"; }
                if (transport === void 0) { transport = GoTSRPC.call; }
                this.endPoint = endPoint;
                this.transport = transport;
            }
            DemoClient.prototype.extractAddress = function (person, success, err) {
                this.transport(this.endPoint, "ExtractAddress", [person], success, err);
            };
            DemoClient.prototype.giveMeAScalar = function (success, err) {
                this.transport(this.endPoint, "GiveMeAScalar", [], success, err);
            };
            DemoClient.prototype.hello = function (name, success, err) {
                this.transport(this.endPoint, "Hello", [name], success, err);
            };
            DemoClient.prototype.helloInterface = function (anything, anythingMap, anythingSlice, success, err) {
                this.transport(this.endPoint, "HelloInterface", [anything, anythingMap, anythingSlice], success, err);
            };
            DemoClient.prototype.mapCrap = function (success, err) {
                this.transport(this.endPoint, "MapCrap", [], success, err);
            };
            DemoClient.prototype.nest = function (success, err) {
                this.transport(this.endPoint, "Nest", [], success, err);
            };
            DemoClient.prototype.testScalarInPlace = function (success, err) {
                this.transport(this.endPoint, "TestScalarInPlace", [], success, err);
            };
            DemoClient.defaultInst = new DemoClient;
            return DemoClient;
        }());
        Demo.DemoClient = DemoClient;
    })(Demo = GoTSRPC.Demo || (GoTSRPC.Demo = {}));
})(GoTSRPC || (GoTSRPC = {}));
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
