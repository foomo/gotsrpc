import {ServiceClient} from "./service-client.js";
import {Base} from "./service-vo.js";
import transport from "./transport.js";

export const init = () => {
	const client = new ServiceClient(transport("/service"))

	client.variantA({} as Base).then(ret => {
		console.log("Variant A", ret)
	})
}
