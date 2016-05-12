var Demo;
(function (Demo) {
    var DemoClient = (function () {
        function DemoClient(endpoint) {
            if (endpoint === void 0) { endpoint = "/default"; }
        }
        DemoClient.prototype.foo = function (callback) {
        };
        DemoClient.defaultInst = new DemoClient();
        return DemoClient;
    })();
    Demo.DemoClient = DemoClient;
})(Demo || (Demo = {}));
Demo.DemoClient.defaultInst.foo(function (reply, clientSuccess, request) {
    if (clientSuccess) {
        var sepp, err = reply;
        console.log("success", sepp, err);
    }
    else {
        console.log("upsi", request);
    }
});
