var GoTSRPC;
(function (GoTSRPC) {
    function call(endPoint, method, args, success, err) {
        var request = new XMLHttpRequest();
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.send(JSON.stringify(args));
        request.onload = function () {
            if (request.status == 200) {
                var data = JSON.parse(request.responseText);
                success.apply(null, data);
            }
            else {
                err(request);
            }
        };
        request.onerror = function () {
            err(request);
        };
    }
    GoTSRPC.call = call;
})(GoTSRPC || (GoTSRPC = {}));
var Demo;
(function (Demo) {
    var ServiceClient = (function () {
        function ServiceClient(endPoint) {
            if (endPoint === void 0) { endPoint = "/service"; }
            this.endPoint = endPoint;
        }
        ServiceClient.prototype.hello = function (name, success, err) {
            GoTSRPC.call(this.endPoint, "Hello", [name], success, err);
        };
        ServiceClient.defaultInst = new ServiceClient();
        return ServiceClient;
    })();
    Demo.ServiceClient = ServiceClient;
})(Demo || (Demo = {}));
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
Demo.ServiceClient.defaultInst.hello("Hansi", function (reply, err) {
    console.log("server says hello to Hansi", reply, err);
    handleCrap(err, null);
}, function (request) {
    console.log("wtf", request);
    handleCrap(null, request);
});
Demo.ServiceClient.defaultInst.hello("Peter", function (reply, err) {
    console.log("server should not like Peter", reply, err);
    handleCrap(err, null);
}, function (request) {
    console.log("wtf", request);
    handleCrap(null, request);
});
