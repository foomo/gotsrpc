/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_nullable_service from './service-vo'; // ./client/src/service-client.ts to ./client/src/service-vo.ts

export class ServiceClient {
	public static defaultEndpoint = "/service";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async variantA(i1:github_com_foomo_gotsrpc_v2_example_nullable_service.Base):Promise<github_com_foomo_gotsrpc_v2_example_nullable_service.Base> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_nullable_service.Base}>("VariantA", [i1]))[0]
	}
	async variantB(i1:github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomType):Promise<github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomType> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomType}>("VariantB", [i1]))[0]
	}
	async variantC(i1:github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomTypes|null):Promise<github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomTypes|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomTypes|null}>("VariantC", [i1]))[0]
	}
	async variantD(i1:github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomTypesMap|null):Promise<github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomTypesMap|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_nullable_service.BCustomTypesMap|null}>("VariantD", [i1]))[0]
	}
	async variantE(i1:github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null):Promise<github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null}>("VariantE", [i1]))[0]
	}
	async variantF(i1:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null):Promise<Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null> {
		return (await this.transport<{0:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null}>("VariantF", [i1]))[0]
	}
	async variantG(i1:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null):Promise<Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null> {
		return (await this.transport<{0:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null}>("VariantG", [i1]))[0]
	}
	async variantH(i1:github_com_foomo_gotsrpc_v2_example_nullable_service.Base, i2:github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null, i3:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null, i4:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Base>|null):Promise<{r1:github_com_foomo_gotsrpc_v2_example_nullable_service.Base; r2:github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null; r3:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null; r4:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Base>|null}> {
		let response = await this.transport<{0:github_com_foomo_gotsrpc_v2_example_nullable_service.Base; 1:github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null; 2:Array<github_com_foomo_gotsrpc_v2_example_nullable_service.Base|null>|null; 3:Record<string,github_com_foomo_gotsrpc_v2_example_nullable_service.Base>|null}>("VariantH", [i1, i2, i3, i4])
		let responseObject = {r1 : response[0], r2 : response[1], r3 : response[2], r4 : response[3]};
		return responseObject;
	}
}