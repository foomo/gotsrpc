/* eslint:disable */
// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.
import * as github_com_foomo_gotsrpc_v2_example_time_service from './service-vo-service'; // ./client/src/service-client.ts to ./client/src/service-vo-service.ts

export class ServiceClient {
	public static defaultEndpoint = "/service";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async time(v:number):Promise<number> {
		return (await this.transport<{0:number}>("Time", [v]))[0]
	}
	async timeStruct(v:github_com_foomo_gotsrpc_v2_example_time_service.TimeStruct):Promise<github_com_foomo_gotsrpc_v2_example_time_service.TimeStruct> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_time_service.TimeStruct}>("TimeStruct", [v]))[0]
	}
}