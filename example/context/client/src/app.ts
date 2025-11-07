import {ServiceClient} from "./service-client.js";
import transport from "./transport.js";


export const init = () => {
	const client = new ServiceClient(transport("/service"));

	client.hello("foomo").then((err) => {
		console.log(err);
	});

	// {"m":"something went wrong: ups","p":"errors","t":"*errors.errorString","d":{}}
	client.standardError("ups").then((err) => {
		console.log(JSON.stringify(err));
	});

	// {"m":"something","p":"github.com/pkg/errors","t":"*errors.fundamental","d":{}}
	client.typedError("ups").then((err) => {
		console.log(JSON.stringify(err));
	});

	// {"m":"ups","p":"github.com/pkg/errors","t":"*errors.withMessage","d":{},"c":{"m":"something","p":"github.com/pkg/errors","t":"*errors.fundamental","d":{}}}
	client.wrappedError("ups").then((err) => {
		console.log(JSON.stringify(err));
	});

	// {"m":"ups","p":"github.com/foomo/gotsrpc/v2/example/context/service","t":"*service.MyError","d":{"payload":"ups"},"c":{"m":"something","p":"github.com/pkg/errors","t":"*errors.fundamental","d":{}}}
	client.customError("ups").then((err) => {
		console.log(JSON.stringify(err));
	});
};
