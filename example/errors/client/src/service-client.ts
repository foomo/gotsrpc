/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_errors_service_frontend from './service-vo'; // ./client/src/service-client.ts to ./client/src/service-vo.ts

export class ServiceClient {
	public static defaultEndpoint = "/service/frontend";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async multiple():Promise<github_com_foomo_gotsrpc_v2_example_errors_service_frontend.ErrMulti|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_errors_service_frontend.ErrMulti|null}>("Multiple", []))[0]
	}
	async simple():Promise<github_com_foomo_gotsrpc_v2_example_errors_service_frontend.ErrSimple|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_errors_service_frontend.ErrSimple|null}>("Simple", []))[0]
	}
}