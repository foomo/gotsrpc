/* eslint:disable */
// hello commonjs - we need some imports - sorted in alphabetical order, by go package
import * as github_com_foomo_gotsrpc_v2_example_time_service from './service-vo-service'; // ./client/src/service-vo-service.ts to ./client/src/service-vo-service.ts
import * as time from './service-vo-time'; // ./client/src/service-vo-service.ts to ./client/src/service-vo-time.ts
// github.com/foomo/gotsrpc/v2/example/time/service.TimeStruct
export interface TimeStruct {
	time:number;
	timePtr:number|null;
	timePtrOmit?:number;
}
// end of common js