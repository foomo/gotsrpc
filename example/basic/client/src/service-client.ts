/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_basic_service from './service-vo'; // ./client/src/service-client.ts to ./client/src/service-vo.ts

export class ServiceClient {
	public static defaultEndpoint = "/service";
	constructor(
		public transport:<T>(method: string, data?: any[]) => Promise<T>
	) {}
	async bool(v:boolean):Promise<boolean> {
		return (await this.transport<{0:boolean}>("Bool", [v]))[0]
	}
	async boolSlice(v:Array<boolean>|null):Promise<Array<boolean>|null> {
		return (await this.transport<{0:Array<boolean>|null}>("BoolSlice", [v]))[0]
	}
	async float32(v:number):Promise<number> {
		return (await this.transport<{0:number}>("Float32", [v]))[0]
	}
	async float32Map(v:Record<number,any>|null):Promise<Record<number,any>|null> {
		return (await this.transport<{0:Record<number,any>|null}>("Float32Map", [v]))[0]
	}
	async float32Slice(v:Array<number>|null):Promise<Array<number>|null> {
		return (await this.transport<{0:Array<number>|null}>("Float32Slice", [v]))[0]
	}
	async float32Type(v:github_com_foomo_gotsrpc_v2_example_basic_service.Float32Type):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Float32Type> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Float32Type}>("Float32Type", [v]))[0]
	}
	async float32TypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapValue>|null}>("Float32TypeMap", [v]))[0]
	}
	async float32TypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Float32TypeMapTyped|null}>("Float32TypeMapTyped", [v]))[0]
	}
	async float64(v:number):Promise<number> {
		return (await this.transport<{0:number}>("Float64", [v]))[0]
	}
	async float64Map(v:Record<number,any>|null):Promise<Record<number,any>|null> {
		return (await this.transport<{0:Record<number,any>|null}>("Float64Map", [v]))[0]
	}
	async float64Slice(v:Array<number>|null):Promise<Array<number>|null> {
		return (await this.transport<{0:Array<number>|null}>("Float64Slice", [v]))[0]
	}
	async float64Type(v:github_com_foomo_gotsrpc_v2_example_basic_service.Float64Type):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Float64Type> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Float64Type}>("Float64Type", [v]))[0]
	}
	async float64TypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapValue>|null}>("Float64TypeMap", [v]))[0]
	}
	async float64TypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Float64TypeMapTyped|null}>("Float64TypeMapTyped", [v]))[0]
	}
	async int(v:number):Promise<number> {
		return (await this.transport<{0:number}>("Int", [v]))[0]
	}
	async int32(v:number):Promise<number> {
		return (await this.transport<{0:number}>("Int32", [v]))[0]
	}
	async int32Map(v:Record<number,any>|null):Promise<Record<number,any>|null> {
		return (await this.transport<{0:Record<number,any>|null}>("Int32Map", [v]))[0]
	}
	async int32Slice(v:Array<number>|null):Promise<Array<number>|null> {
		return (await this.transport<{0:Array<number>|null}>("Int32Slice", [v]))[0]
	}
	async int32Type(v:github_com_foomo_gotsrpc_v2_example_basic_service.Int32Type):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Int32Type> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Int32Type}>("Int32Type", [v]))[0]
	}
	async int32TypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapValue>|null}>("Int32TypeMap", [v]))[0]
	}
	async int32TypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Int32TypeMapTyped|null}>("Int32TypeMapTyped", [v]))[0]
	}
	async int64(v:number):Promise<number> {
		return (await this.transport<{0:number}>("Int64", [v]))[0]
	}
	async int64Map(v:Record<number,any>|null):Promise<Record<number,any>|null> {
		return (await this.transport<{0:Record<number,any>|null}>("Int64Map", [v]))[0]
	}
	async int64Slice(v:Array<number>|null):Promise<Array<number>|null> {
		return (await this.transport<{0:Array<number>|null}>("Int64Slice", [v]))[0]
	}
	async int64Type(v:github_com_foomo_gotsrpc_v2_example_basic_service.Int64Type):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Int64Type> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Int64Type}>("Int64Type", [v]))[0]
	}
	async int64TypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapValue>|null}>("Int64TypeMap", [v]))[0]
	}
	async int64TypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Int64TypeMapTyped|null}>("Int64TypeMapTyped", [v]))[0]
	}
	async intMap(v:Record<number,any>|null):Promise<Record<number,any>|null> {
		return (await this.transport<{0:Record<number,any>|null}>("IntMap", [v]))[0]
	}
	async intSlice(v:Array<number>|null):Promise<Array<number>|null> {
		return (await this.transport<{0:Array<number>|null}>("IntSlice", [v]))[0]
	}
	async intType(v:github_com_foomo_gotsrpc_v2_example_basic_service.IntType):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.IntType> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.IntType}>("IntType", [v]))[0]
	}
	async intTypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapValue>|null}>("IntTypeMap", [v]))[0]
	}
	async intTypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.IntTypeMapTyped|null}>("IntTypeMapTyped", [v]))[0]
	}
	async interface(v:any):Promise<any> {
		return (await this.transport<{0:any}>("Interface", [v]))[0]
	}
	async interfaceSlice(v:Array<any>|null):Promise<Array<any>|null> {
		return (await this.transport<{0:Array<any>|null}>("InterfaceSlice", [v]))[0]
	}
	async string(v:string):Promise<string> {
		return (await this.transport<{0:string}>("String", [v]))[0]
	}
	async stringMap(v:Record<string,any>|null):Promise<Record<string,any>|null> {
		return (await this.transport<{0:Record<string,any>|null}>("StringMap", [v]))[0]
	}
	async stringSlice(v:Array<string>|null):Promise<Array<string>|null> {
		return (await this.transport<{0:Array<string>|null}>("StringSlice", [v]))[0]
	}
	async stringType(v:github_com_foomo_gotsrpc_v2_example_basic_service.StringType):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.StringType> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.StringType}>("StringType", [v]))[0]
	}
	async stringTypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapValue>|null}>("StringTypeMap", [v]))[0]
	}
	async stringTypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.StringTypeMapTyped|null}>("StringTypeMapTyped", [v]))[0]
	}
	async struct(v:github_com_foomo_gotsrpc_v2_example_basic_service.Struct):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.Struct> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.Struct}>("Struct", [v]))[0]
	}
	async uInt(v:number):Promise<number> {
		return (await this.transport<{0:number}>("UInt", [v]))[0]
	}
	async uInt32(v:number):Promise<number> {
		return (await this.transport<{0:number}>("UInt32", [v]))[0]
	}
	async uInt32Map(v:Record<number,any>|null):Promise<Record<number,any>|null> {
		return (await this.transport<{0:Record<number,any>|null}>("UInt32Map", [v]))[0]
	}
	async uInt32Slice(v:Array<number>|null):Promise<Array<number>|null> {
		return (await this.transport<{0:Array<number>|null}>("UInt32Slice", [v]))[0]
	}
	async uInt32Type(v:github_com_foomo_gotsrpc_v2_example_basic_service.UInt32Type):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.UInt32Type> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.UInt32Type}>("UInt32Type", [v]))[0]
	}
	async uInt32TypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapValue>|null}>("UInt32TypeMap", [v]))[0]
	}
	async uInt32TypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.UInt32TypeMapTyped|null}>("UInt32TypeMapTyped", [v]))[0]
	}
	async uInt64(v:number):Promise<number> {
		return (await this.transport<{0:number}>("UInt64", [v]))[0]
	}
	async uInt64Map(v:Record<number,any>|null):Promise<Record<number,any>|null> {
		return (await this.transport<{0:Record<number,any>|null}>("UInt64Map", [v]))[0]
	}
	async uInt64Slice(v:Array<number>|null):Promise<Array<number>|null> {
		return (await this.transport<{0:Array<number>|null}>("UInt64Slice", [v]))[0]
	}
	async uInt64Type(v:github_com_foomo_gotsrpc_v2_example_basic_service.UInt64Type):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.UInt64Type> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.UInt64Type}>("UInt64Type", [v]))[0]
	}
	async uInt64TypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapValue>|null}>("UInt64TypeMap", [v]))[0]
	}
	async uInt64TypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.UInt64TypeMapTyped|null}>("UInt64TypeMapTyped", [v]))[0]
	}
	async uIntMap(v:Record<number,any>|null):Promise<Record<number,any>|null> {
		return (await this.transport<{0:Record<number,any>|null}>("UIntMap", [v]))[0]
	}
	async uIntSlice(v:Array<number>|null):Promise<Array<number>|null> {
		return (await this.transport<{0:Array<number>|null}>("UIntSlice", [v]))[0]
	}
	async uIntType(v:github_com_foomo_gotsrpc_v2_example_basic_service.UIntType):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.UIntType> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.UIntType}>("UIntType", [v]))[0]
	}
	async uIntTypeMap(v:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapValue>|null):Promise<Record<github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapValue>|null> {
		return (await this.transport<{0:Record<github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapKey,github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapValue>|null}>("UIntTypeMap", [v]))[0]
	}
	async uIntTypeMapTyped(v:github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapTyped|null):Promise<github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapTyped|null> {
		return (await this.transport<{0:github_com_foomo_gotsrpc_v2_example_basic_service.UIntTypeMapTyped|null}>("UIntTypeMapTyped", [v]))[0]
	}
}