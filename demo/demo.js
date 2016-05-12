var Demo;
(function (Demo) {
    var ServiceClient = (function () {
        function ServiceClient(endpoint) {
            if (endpoint === void 0) { endpoint = "/default"; }
        }
        ServiceClient.prototype.hello = function (msg, success, err) {
        };
        ServiceClient.defaultInst = new ServiceClient();
        return ServiceClient;
    })();
    Demo.ServiceClient = ServiceClient;
})(Demo || (Demo = {}));
Demo.ServiceClient.defaultInst.hello("Hansi", function (reply, err) {
    console.log("server says hello", reply, err);
}, function (request) {
    console.log("wtf", request);
});
