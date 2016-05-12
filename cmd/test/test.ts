
module Demo {
    
    export interface ClientErr {
        
    }
    
    export interface Sepp {
        foo:number;
    }
    
    export interface Err {
        Message:string;
    }
    
    export class DemoClient {
        static defaultInst = new DemoClient()
        constructor(endpoint:string = "/default") {
            
        }
        foo(args:[boolean, number], success:(reply:[Sepp,Err]) => void, err:(request:XMLHttpRequest) => void) {
            
        }
    }
}


Demo.DemoClient.defaultInst.foo(
    [true, 23.4], 
    (reply:[Demo.Sepp,Demo.Err]) => {
        var sepp, err = reply;
        console.log("success", sepp, err);
    },
    (request:XMLHttpRequest) => {
        console.log("wtf", request);
    }
);


