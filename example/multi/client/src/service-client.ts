/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_multi_service from './service-vo'; // ./client/src/service-client.ts to ./client/src/service-vo.ts

export class ServiceClient {
	public static defaultEndpoint = "/service";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async inlineStruct():Promise<github_com_foomo_gotsrpc_v2_example_multi_service.InlineStruct> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_multi_service.InlineStruct}>("InlineStruct", []))[0]
	}
	async inlineStructPtr():Promise<github_com_foomo_gotsrpc_v2_example_multi_service.InlineStructPtr> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_multi_service.InlineStructPtr}>("InlineStructPtr", []))[0]
	}
	async unionString():Promise<github_com_foomo_gotsrpc_v2_example_multi_service.UnionString> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_multi_service.UnionString}>("UnionString", []))[0]
	}
	async unionStruct():Promise<github_com_foomo_gotsrpc_v2_example_multi_service.UnionStruct> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_multi_service.UnionStruct}>("UnionStruct", []))[0]
	}
}