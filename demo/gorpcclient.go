// Code generated by gotsrpc https://github.com/foomo/gotsrpc  - DO NOT EDIT.

package demo

import (
	tls "crypto/tls"

	github_com_foomo_gotsrpc_demo_nested "github.com/foomo/gotsrpc/demo/nested"
	gorpc "github.com/valyala/gorpc"
)

type FooGoRPCClient struct {
	Client *gorpc.Client
}

func NewFooGoRPCClient(addr string, tlsConfig *tls.Config) *FooGoRPCClient {
	client := &FooGoRPCClient{}
	if tlsConfig == nil {
		client.Client = gorpc.NewTCPClient(addr)
	} else {
		client.Client = gorpc.NewTLSClient(addr, tlsConfig)
	}
	return client
}

func (tsc *FooGoRPCClient) Start() {
	tsc.Client.Start()
}

func (tsc *FooGoRPCClient) Stop() {
	tsc.Client.Stop()
}

func (tsc *FooGoRPCClient) Hello(number int64) (retHello_0 int, clientErr error) {
	req := FooHelloRequest{Number: number}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(FooHelloResponse)
	return response.RetHello_0, nil
}

type DemoGoRPCClient struct {
	Client *gorpc.Client
}

func NewDemoGoRPCClient(addr string, tlsConfig *tls.Config) *DemoGoRPCClient {
	client := &DemoGoRPCClient{}
	if tlsConfig == nil {
		client.Client = gorpc.NewTCPClient(addr)
	} else {
		client.Client = gorpc.NewTLSClient(addr, tlsConfig)
	}
	return client
}

func (tsc *DemoGoRPCClient) Start() {
	tsc.Client.Start()
}

func (tsc *DemoGoRPCClient) Stop() {
	tsc.Client.Stop()
}

func (tsc *DemoGoRPCClient) Any(any github_com_foomo_gotsrpc_demo_nested.Any, anyList []github_com_foomo_gotsrpc_demo_nested.Any, anyMap map[string]github_com_foomo_gotsrpc_demo_nested.Any) (retAny_0 github_com_foomo_gotsrpc_demo_nested.Any, retAny_1 []github_com_foomo_gotsrpc_demo_nested.Any, retAny_2 map[string]github_com_foomo_gotsrpc_demo_nested.Any, clientErr error) {
	req := DemoAnyRequest{Any: any, AnyList: anyList, AnyMap: anyMap}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoAnyResponse)
	return response.RetAny_0, response.RetAny_1, response.RetAny_2, nil
}

func (tsc *DemoGoRPCClient) ExtractAddress(person *Person) (addr *Address, e *Err, clientErr error) {
	req := DemoExtractAddressRequest{Person: person}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoExtractAddressResponse)
	return response.Addr, response.E, nil
}

func (tsc *DemoGoRPCClient) GiveMeAScalar() (amount github_com_foomo_gotsrpc_demo_nested.Amount, wahr github_com_foomo_gotsrpc_demo_nested.True, hier ScalarInPlace, clientErr error) {
	req := DemoGiveMeAScalarRequest{}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoGiveMeAScalarResponse)
	return response.Amount, response.Wahr, response.Hier, nil
}

func (tsc *DemoGoRPCClient) Hello(name string) (retHello_0 string, retHello_1 *Err, clientErr error) {
	req := DemoHelloRequest{Name: name}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoHelloResponse)
	return response.RetHello_0, response.RetHello_1, nil
}

func (tsc *DemoGoRPCClient) HelloInterface(anything interface{}, anythingMap map[string]interface{}, anythingSlice []interface{}) (clientErr error) {
	req := DemoHelloInterfaceRequest{Anything: anything, AnythingMap: anythingMap, AnythingSlice: anythingSlice}
	_, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	return nil
}

func (tsc *DemoGoRPCClient) HelloNumberMaps(intMap map[int]string) (floatMap map[float64]string, clientErr error) {
	req := DemoHelloNumberMapsRequest{IntMap: intMap}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoHelloNumberMapsResponse)
	return response.FloatMap, nil
}

func (tsc *DemoGoRPCClient) HelloScalarError() (err *ScalarError, clientErr error) {
	req := DemoHelloScalarErrorRequest{}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoHelloScalarErrorResponse)
	return response.Err, nil
}

func (tsc *DemoGoRPCClient) MapCrap() (crap map[string][]int, clientErr error) {
	req := DemoMapCrapRequest{}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoMapCrapResponse)
	return response.Crap, nil
}

func (tsc *DemoGoRPCClient) Nest() (retNest_0 []*github_com_foomo_gotsrpc_demo_nested.Nested, clientErr error) {
	req := DemoNestRequest{}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoNestResponse)
	return response.RetNest_0, nil
}

func (tsc *DemoGoRPCClient) TestScalarInPlace() (retTestScalarInPlace_0 ScalarInPlace, clientErr error) {
	req := DemoTestScalarInPlaceRequest{}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(DemoTestScalarInPlaceResponse)
	return response.RetTestScalarInPlace_0, nil
}

type BarGoRPCClient struct {
	Client *gorpc.Client
}

func NewBarGoRPCClient(addr string, tlsConfig *tls.Config) *BarGoRPCClient {
	client := &BarGoRPCClient{}
	if tlsConfig == nil {
		client.Client = gorpc.NewTCPClient(addr)
	} else {
		client.Client = gorpc.NewTLSClient(addr, tlsConfig)
	}
	return client
}

func (tsc *BarGoRPCClient) Start() {
	tsc.Client.Start()
}

func (tsc *BarGoRPCClient) Stop() {
	tsc.Client.Stop()
}

func (tsc *BarGoRPCClient) CustomType(customTypeInt CustomTypeInt, customTypeString CustomTypeString, CustomTypeStruct CustomTypeStruct) (retCustomType_0 *CustomTypeInt, retCustomType_1 *CustomTypeString, retCustomType_2 CustomTypeStruct, clientErr error) {
	req := BarCustomTypeRequest{CustomTypeInt: customTypeInt, CustomTypeString: customTypeString, CustomTypeStruct: CustomTypeStruct}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(BarCustomTypeResponse)
	return response.RetCustomType_0, response.RetCustomType_1, response.RetCustomType_2, nil
}
