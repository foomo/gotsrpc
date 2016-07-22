var GoTSRPC;
(function (GoTSRPC) {
    function call(endPoint, method, args, success, err) {
        var request = new XMLHttpRequest();
        request.withCredentials = true;
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.send(JSON.stringify(args));
        request.onload = function () {
            if (request.status == 200) {
                try {
                    var data = JSON.parse(request.responseText);
                    success.apply(null, data);
                }
                catch (e) {
                    err(request);
                }
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
var MZG;
(function (MZG) {
    var Services;
    (function (Services) {
        var Order;
        (function (Order) {
            var ServiceClient = (function () {
                function ServiceClient(endPoint) {
                    if (endPoint === void 0) { endPoint = "/service"; }
                    this.endPoint = endPoint;
                }
                ServiceClient.prototype.validateOrder = function (success, err) {
                    GoTSRPC.call(this.endPoint, "ValidateOrder", [], success, err);
                };
                ServiceClient.prototype.confirmOrder = function (success, err) {
                    GoTSRPC.call(this.endPoint, "ConfirmOrder", [], success, err);
                };
                ServiceClient.prototype.flushPaymentInfo = function (success, err) {
                    GoTSRPC.call(this.endPoint, "FlushPaymentInfo", [], success, err);
                };
                ServiceClient.prototype.setPositionSize = function (oldItemID, newItemID, success, err) {
                    GoTSRPC.call(this.endPoint, "SetPositionSize", [oldItemID, newItemID], success, err);
                };
                ServiceClient.prototype.setPositionCount = function (itemID, count, success, err) {
                    GoTSRPC.call(this.endPoint, "SetPositionCount", [itemID, count], success, err);
                };
                ServiceClient.prototype.getOrder = function (success, err) {
                    GoTSRPC.call(this.endPoint, "GetOrder", [], success, err);
                };
                ServiceClient.prototype.setBillingAddress = function (addressId, success, err) {
                    GoTSRPC.call(this.endPoint, "SetBillingAddress", [addressId], success, err);
                };
                ServiceClient.prototype.setShippingAddress = function (addressId, success, err) {
                    GoTSRPC.call(this.endPoint, "SetShippingAddress", [addressId], success, err);
                };
                ServiceClient.prototype.setBilling = function (billing, success, err) {
                    GoTSRPC.call(this.endPoint, "SetBilling", [billing], success, err);
                };
                ServiceClient.prototype.setShipping = function (shipping, success, err) {
                    GoTSRPC.call(this.endPoint, "SetShipping", [shipping], success, err);
                };
                ServiceClient.prototype.setPayment = function (payment, success, err) {
                    GoTSRPC.call(this.endPoint, "SetPayment", [payment], success, err);
                };
                ServiceClient.prototype.setPage = function (page, success, err) {
                    GoTSRPC.call(this.endPoint, "SetPage", [page], success, err);
                };
                ServiceClient.prototype.getOrderSummary = function (success, err) {
                    GoTSRPC.call(this.endPoint, "GetOrderSummary", [], success, err);
                };
                ServiceClient.defaultInst = new ServiceClient;
                return ServiceClient;
            }());
            Order.ServiceClient = ServiceClient;
        })(Order = Services.Order || (Services.Order = {}));
    })(Services = MZG.Services || (MZG.Services = {}));
})(MZG || (MZG = {}));
// ----------------- MZG.Services --------------------
var MZG;
(function (MZG) {
    var Services;
    (function (Services) {
        var Profile;
        (function (Profile) {
            var ServiceClient = (function () {
                function ServiceClient(endPoint) {
                    if (endPoint === void 0) { endPoint = "/service"; }
                    this.endPoint = endPoint;
                }
                ServiceClient.prototype.validateLogin = function (email, success, err) {
                    GoTSRPC.call(this.endPoint, "ValidateLogin", [email], success, err);
                };
                ServiceClient.prototype.newCustomer = function (email, success, err) {
                    GoTSRPC.call(this.endPoint, "NewCustomer", [email], success, err);
                };
                ServiceClient.prototype.newGuestCustomer = function (email, success, err) {
                    GoTSRPC.call(this.endPoint, "NewGuestCustomer", [email], success, err);
                };
                ServiceClient.prototype.addShippingAddress = function (address, success, err) {
                    GoTSRPC.call(this.endPoint, "AddShippingAddress", [address], success, err);
                };
                ServiceClient.prototype.addBillingAddress = function (address, success, err) {
                    GoTSRPC.call(this.endPoint, "AddBillingAddress", [address], success, err);
                };
                ServiceClient.prototype.setProfileData = function (profileData, success, err) {
                    GoTSRPC.call(this.endPoint, "SetProfileData", [profileData], success, err);
                };
                ServiceClient.defaultInst = new ServiceClient;
                return ServiceClient;
            }());
            Profile.ServiceClient = ServiceClient;
        })(Profile = Services.Profile || (Services.Profile = {}));
    })(Services = MZG.Services || (MZG.Services = {}));
})(MZG || (MZG = {}));
