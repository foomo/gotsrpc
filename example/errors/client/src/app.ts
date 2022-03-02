import {ServiceClient} from "./service-client.js";
import transport from "./transport.js";

export const init = () => {
	const client = new ServiceClient(transport("/service/frontend"))

	client.simple().then((ret) => {
		console.log(ret)
	})

	client.multiple().then(ret => {
		console.log(ret)
	})
}
