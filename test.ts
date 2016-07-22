module GoTSRPC {
    export function call(endPoint:string, method:string, args:any[], success:any, err:any) {
        var request = new XMLHttpRequest();
        request.withCredentials = true;
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.send(JSON.stringify(args));            
        request.onload = function() {
            if (request.status == 200) {
				try {
					var data = JSON.parse(request.responseText);
					success.apply(null, data);
				} catch(e) {
	                err(request);
				}
            } else {
                err(request);
            }
        };            
        request.onerror = function() {
            err(request);
        };
    }
}
module MZG.Services.Order {
    export class ServiceClient {
        static defaultInst = new ServiceClient;
        constructor(public endPoint:string = "/service") {  }
        validateOrder(success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "ValidateOrder", [], success, err);
        }
        confirmOrder(success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "ConfirmOrder", [], success, err);
        }
        flushPaymentInfo(success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "FlushPaymentInfo", [], success, err);
        }
        setPositionSize(oldItemID:string, newItemID:string, success:(order:MZG.Services.Order.Order, err:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetPositionSize", [oldItemID, newItemID], success, err);
        }
        setPositionCount(itemID:string, count:number, success:(order:MZG.Services.Order.Order, err:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetPositionCount", [itemID, count], success, err);
        }
        getOrder(success:(order:MZG.Services.Order.Order, err:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "GetOrder", [], success, err);
        }
        setBillingAddress(addressId:string, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetBillingAddress", [addressId], success, err);
        }
        setShippingAddress(addressId:string, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetShippingAddress", [addressId], success, err);
        }
        setBilling(billing:MZG.Services.Order.Billing, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetBilling", [billing], success, err);
        }
        setShipping(shipping:MZG.Services.Order.Shipping, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetShipping", [shipping], success, err);
        }
        setPayment(payment:MZG.Services.Order.Payment, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetPayment", [payment], success, err);
        }
        setPage(page:string, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetPage", [page], success, err);
        }
        getOrderSummary(success:(summary:MZG.Services.Order.OrderSummary, err:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "GetOrderSummary", [], success, err);
        }
    }
}
// ----------------- MZG.Services.Order --------------------
OrderSummary &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.OrderSummary     export interface OrderSummary {         numberOfItems:number;         subTotal:string;         promotionAmount:string;         hasPromotions:boolean;         vat:string;         total:string;         totalRappen:string;         shippingCosts:string;         hasShippingCosts:boolean;         amountUntilFreeShipping:string;     }] 1}
// ----------------- MZG.Services.Order --------------------
OrderPosition &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.OrderPosition     export interface OrderPosition {         id:string;         name:string;         brand:string;         imageURL:string;         size:string;         style:string;         quantity:number;         pricePerUnit:any;         price:any;         availability:any;         sizes:any[];     }] 1}
// ----------------- MZG.Services.Order --------------------
AddressFormValue &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.AddressFormValue     export interface AddressFormValue {         salutation:string;         firstName:string;         title:string;         lastName:string;         company:string;         street:string;         streetNumber:string;         zip:string;         city:string;     }] 1}
// ----------------- MZG.Services.Order --------------------
Billing &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.Billing     export interface Billing {         newAddress?:MZG.Services.Order.AddressFormValue;         tacAgree:boolean;         email:string;         phone:string;         birthday:string;     }] 1}
// ----------------- MZG.Services.Order --------------------
PaymentInfo &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.PaymentInfo     export interface PaymentInfo {         method:string;         status:string;         feedback:string;         complete:boolean;         amount:string;     }] 1}
// ----------------- MZG.Services.Order --------------------
Session &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.Session     export interface Session {         billing?:MZG.Services.Order.Billing;         shipping?:MZG.Services.Order.Shipping;         payment?:MZG.Services.Order.Payment;         page:string;     }] 1}
// ----------------- MZG.Services.Order --------------------
Order &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.Order     export interface Order {         id:string;         positions:MZG.Services.Order.OrderPosition[];         summary?:MZG.Services.Order.OrderSummary;         paymentInfo:MZG.Services.Order.PaymentInfo[];         session?:MZG.Services.Order.Session;     }] 1}
// ----------------- MZG.Services.Order --------------------
Shipping &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.Shipping     export interface Shipping {         useBilling:boolean;         newAddress?:MZG.Services.Order.AddressFormValue;     }] 1}
// ----------------- MZG.Services.Order --------------------
Payment &{ [    // git.bestbytes.net/Project-Globus-Services/services/order.Payment     export interface Payment {         paymentType:string;         coupons:string[];     }] 1}
// ----------------- MZG.Services --------------------
ServiceError &{ [    // git.bestbytes.net/Project-Globus-Services/services.ServiceError     export interface ServiceError {         error:string;     }] 1}
module GoTSRPC {
    export function call(endPoint:string, method:string, args:any[], success:any, err:any) {
        var request = new XMLHttpRequest();
        request.withCredentials = true;
        request.open('POST', endPoint + "/" + encodeURIComponent(method), true);
        request.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded; charset=UTF-8');
        request.send(JSON.stringify(args));            
        request.onload = function() {
            if (request.status == 200) {
				try {
					var data = JSON.parse(request.responseText);
					success.apply(null, data);
				} catch(e) {
	                err(request);
				}
            } else {
                err(request);
            }
        };            
        request.onerror = function() {
            err(request);
        };
    }
}
module MZG.Services.Profile {
    export class ServiceClient {
        static defaultInst = new ServiceClient;
        constructor(public endPoint:string = "/service") {  }
        validateLogin(email:string, success:(valid:boolean) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "ValidateLogin", [email], success, err);
        }
        newCustomer(email:string, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "NewCustomer", [email], success, err);
        }
        newGuestCustomer(email:string, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "NewGuestCustomer", [email], success, err);
        }
        addShippingAddress(address:MZG.Services.Profile.AddressFormValue, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "AddShippingAddress", [address], success, err);
        }
        addBillingAddress(address:MZG.Services.Profile.AddressFormValue, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "AddBillingAddress", [address], success, err);
        }
        setProfileData(profileData:MZG.Services.Profile.ProfileData, success:(ret:MZG.Services.ServiceError) => void, err:(request:XMLHttpRequest) => void) {
            GoTSRPC.call(this.endPoint, "SetProfileData", [profileData], success, err);
        }
    }
}
// ----------------- MZG.Services --------------------
ServiceError &{ [    // git.bestbytes.net/Project-Globus-Services/services.ServiceError     export interface ServiceError {         error:string;     }] 1}
// ----------------- MZG.Services.Profile --------------------
AddressFormValue &{ [    // git.bestbytes.net/Project-Globus-Services/services/profile.AddressFormValue     export interface AddressFormValue {         salutation:string;         firstName:string;         title:string;         lastName:string;         company:string;         street:string;         streetNumber:string;         zip:string;         city:string;     }] 1}
// ----------------- MZG.Services.Profile --------------------
ProfileData &{ [    // git.bestbytes.net/Project-Globus-Services/services/profile.ProfileData     export interface ProfileData {         tacAgree:boolean;         PhoneMobile:string;         Birthday:string;     }] 1}
