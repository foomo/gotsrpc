import {ServiceClient} from "./service-client.js";
import transport from "./transport.js";

export const init = () => {
	const client = new ServiceClient(transport(ServiceClient.defaultEndpoint))

	client.inlineStruct().then((ret) => {
		console.log(ret.valueA)
	})

	client.inlineStructPtr().then(ret => {
		console.log(ret.valueA)
	})

	client.unionString().then(ret => {
		console.log(ret.One)
		console.log(ret.Two)
		console.log(ret.Three)
		console.log(ret.Four)
	})

	client.unionStruct().then(ret => {
		if (ret) {
			switch (ret.kind) {
				case "UnionStructA":
					console.log(ret.value)
					console.log(ret.bar)
					break
				case "UnionStructB":
					console.log(ret.value)
					console.log(ret.foo)
					break;
				default:
					assertExhaustive(ret, "unhandled")
			}

			if (ret.kind === "UnionStructA") {
				console.log(ret.value)
				console.log(ret.bar)
			}
		}
	})
}

function assertExhaustive(
	value: never,
	message: string = 'msg'
): never {
	throw new Error(message);
}
