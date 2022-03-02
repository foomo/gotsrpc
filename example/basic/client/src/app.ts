import {ServiceClient} from "./service-client.js";
import transport from "./transport.js";


export const init = () => {
	const client = new ServiceClient(transport("/service"));

	client.boolPtr(true).then((res) => {
		console.log(res);
	});
};
