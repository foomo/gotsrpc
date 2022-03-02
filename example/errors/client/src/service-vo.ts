/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_errors_service_frontend from './service-vo'; // ./client/src/service-vo.ts to ./client/src/service-vo.ts
// github.com/foomo/gotsrpc/v2/example/errors/service/frontend.ErrMulti
export type ErrMulti = (typeof github_com_foomo_gotsrpc_v2_example_errors_service_frontend.ErrMultiA) & (typeof github_com_foomo_gotsrpc_v2_example_errors_service_frontend.ErrMultiB)
// github.com/foomo/gotsrpc/v2/example/errors/service/frontend.ErrMultiA
export enum ErrMultiA {
	One = "one",
	Two = "two",
}
// github.com/foomo/gotsrpc/v2/example/errors/service/frontend.ErrMultiB
export enum ErrMultiB {
	Four = "four",
	Three = "three",
}
// github.com/foomo/gotsrpc/v2/example/errors/service/frontend.ErrSimple
export enum ErrSimple {
	One = "one",
	Two = "two",
}
// end of common js