import {ServiceClient} from "./service-client.js";
import transport from "./transport.js";


export const init = () => {
	const client = new ServiceClient(transport("/service"));

	client.timeStruct({
		time: 631148400000,
		timePtr: 631148400000,
	}).then((res) => {
		console.log(new Date(res.time).toString());
	});
};
