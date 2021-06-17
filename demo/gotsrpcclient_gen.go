// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package demo

import (
	go_context "context"
	go_net_http "net/http"

	gotsrpc "github.com/foomo/gotsrpc/v2"
	github_com_foomo_gotsrpc_v2_demo_nested "github.com/foomo/gotsrpc/v2/demo/nested"
)

type FooGoTSRPCClient interface {
	Hello(ctx go_context.Context, number int64) (retHello_0 int, clientErr error)
}

type HTTPFooGoTSRPCClient struct {
	URL      string
	EndPoint string
	Client   gotsrpc.Client
}

func NewDefaultFooGoTSRPCClient(url string) *HTTPFooGoTSRPCClient {
	return NewFooGoTSRPCClient(url, "/service/foo")
}

func NewFooGoTSRPCClient(url string, endpoint string) *HTTPFooGoTSRPCClient {
	return NewFooGoTSRPCClientWithClient(url, endpoint, nil)
}

func NewFooGoTSRPCClientWithClient(url string, endpoint string, client *go_net_http.Client) *HTTPFooGoTSRPCClient {
	return &HTTPFooGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
		Client:   gotsrpc.NewClientWithHttpClient(client),
	}
}
func (tsc *HTTPFooGoTSRPCClient) Hello(ctx go_context.Context, number int64) (retHello_0 int, clientErr error) {
	args := []interface{}{number}
	reply := []interface{}{&retHello_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Hello", args, reply)
	return
}

type DemoGoTSRPCClient interface {
	Any(ctx go_context.Context, any github_com_foomo_gotsrpc_v2_demo_nested.Any, anyList []github_com_foomo_gotsrpc_v2_demo_nested.Any, anyMap map[string]github_com_foomo_gotsrpc_v2_demo_nested.Any) (retAny_0 github_com_foomo_gotsrpc_v2_demo_nested.Any, retAny_1 []github_com_foomo_gotsrpc_v2_demo_nested.Any, retAny_2 map[string]github_com_foomo_gotsrpc_v2_demo_nested.Any, clientErr error)
	ExtractAddress(ctx go_context.Context, person *Person) (addr *Address, e *Err, clientErr error)
	GiveMeAScalar(ctx go_context.Context) (amount github_com_foomo_gotsrpc_v2_demo_nested.Amount, wahr github_com_foomo_gotsrpc_v2_demo_nested.True, hier ScalarInPlace, clientErr error)
	Hello(ctx go_context.Context, name string) (retHello_0 string, retHello_1 *Err, clientErr error)
	HelloInterface(ctx go_context.Context, anything interface{}, anythingMap map[string]interface{}, anythingSlice []interface{}) (clientErr error)
	HelloNumberMaps(ctx go_context.Context, intMap map[int]string) (floatMap map[float64]string, clientErr error)
	HelloScalarError(ctx go_context.Context) (err *ScalarError, clientErr error)
	MapCrap(ctx go_context.Context) (crap map[string][]int, clientErr error)
	Nest(ctx go_context.Context) (retNest_0 []*github_com_foomo_gotsrpc_v2_demo_nested.Nested, clientErr error)
	TestScalarInPlace(ctx go_context.Context) (retTestScalarInPlace_0 ScalarInPlace, clientErr error)
}

type HTTPDemoGoTSRPCClient struct {
	URL      string
	EndPoint string
	Client   gotsrpc.Client
}

func NewDefaultDemoGoTSRPCClient(url string) *HTTPDemoGoTSRPCClient {
	return NewDemoGoTSRPCClient(url, "/service/demo")
}

func NewDemoGoTSRPCClient(url string, endpoint string) *HTTPDemoGoTSRPCClient {
	return NewDemoGoTSRPCClientWithClient(url, endpoint, nil)
}

func NewDemoGoTSRPCClientWithClient(url string, endpoint string, client *go_net_http.Client) *HTTPDemoGoTSRPCClient {
	return &HTTPDemoGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
		Client:   gotsrpc.NewClientWithHttpClient(client),
	}
}
func (tsc *HTTPDemoGoTSRPCClient) Any(ctx go_context.Context, any github_com_foomo_gotsrpc_v2_demo_nested.Any, anyList []github_com_foomo_gotsrpc_v2_demo_nested.Any, anyMap map[string]github_com_foomo_gotsrpc_v2_demo_nested.Any) (retAny_0 github_com_foomo_gotsrpc_v2_demo_nested.Any, retAny_1 []github_com_foomo_gotsrpc_v2_demo_nested.Any, retAny_2 map[string]github_com_foomo_gotsrpc_v2_demo_nested.Any, clientErr error) {
	args := []interface{}{any, anyList, anyMap}
	reply := []interface{}{&retAny_0, &retAny_1, &retAny_2}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Any", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) ExtractAddress(ctx go_context.Context, person *Person) (addr *Address, e *Err, clientErr error) {
	args := []interface{}{person}
	reply := []interface{}{&addr, &e}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "ExtractAddress", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) GiveMeAScalar(ctx go_context.Context) (amount github_com_foomo_gotsrpc_v2_demo_nested.Amount, wahr github_com_foomo_gotsrpc_v2_demo_nested.True, hier ScalarInPlace, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&amount, &wahr, &hier}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "GiveMeAScalar", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) Hello(ctx go_context.Context, name string) (retHello_0 string, retHello_1 *Err, clientErr error) {
	args := []interface{}{name}
	reply := []interface{}{&retHello_0, &retHello_1}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Hello", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) HelloInterface(ctx go_context.Context, anything interface{}, anythingMap map[string]interface{}, anythingSlice []interface{}) (clientErr error) {
	args := []interface{}{anything, anythingMap, anythingSlice}
	reply := []interface{}{}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "HelloInterface", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) HelloNumberMaps(ctx go_context.Context, intMap map[int]string) (floatMap map[float64]string, clientErr error) {
	args := []interface{}{intMap}
	reply := []interface{}{&floatMap}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "HelloNumberMaps", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) HelloScalarError(ctx go_context.Context) (err *ScalarError, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&err}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "HelloScalarError", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) MapCrap(ctx go_context.Context) (crap map[string][]int, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&crap}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "MapCrap", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) Nest(ctx go_context.Context) (retNest_0 []*github_com_foomo_gotsrpc_v2_demo_nested.Nested, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&retNest_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Nest", args, reply)
	return
}

func (tsc *HTTPDemoGoTSRPCClient) TestScalarInPlace(ctx go_context.Context) (retTestScalarInPlace_0 ScalarInPlace, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&retTestScalarInPlace_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "TestScalarInPlace", args, reply)
	return
}

type BarGoTSRPCClient interface {
	AttributeMapping(ctx go_context.Context) (retAttributeMapping_0 AttributeMapping, clientErr error)
	CustomError(ctx go_context.Context, one CustomError, two *CustomError) (three CustomError, four *CustomError, clientErr error)
	CustomType(ctx go_context.Context, customTypeInt CustomTypeInt, customTypeString CustomTypeString, CustomTypeStruct CustomTypeStruct) (retCustomType_0 *CustomTypeInt, retCustomType_1 *CustomTypeString, retCustomType_2 CustomTypeStruct, clientErr error)
	Hello(ctx go_context.Context, number int64) (retHello_0 int, clientErr error)
	Inheritance(ctx go_context.Context, inner Inner, nested OuterNested, inline OuterInline) (retInheritance_0 Inner, retInheritance_1 OuterNested, retInheritance_2 OuterInline, clientErr error)
	Repeat(ctx go_context.Context, one string, two string) (three bool, four bool, clientErr error)
}

type HTTPBarGoTSRPCClient struct {
	URL      string
	EndPoint string
	Client   gotsrpc.Client
}

func NewDefaultBarGoTSRPCClient(url string) *HTTPBarGoTSRPCClient {
	return NewBarGoTSRPCClient(url, "/service/bar")
}

func NewBarGoTSRPCClient(url string, endpoint string) *HTTPBarGoTSRPCClient {
	return NewBarGoTSRPCClientWithClient(url, endpoint, nil)
}

func NewBarGoTSRPCClientWithClient(url string, endpoint string, client *go_net_http.Client) *HTTPBarGoTSRPCClient {
	return &HTTPBarGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
		Client:   gotsrpc.NewClientWithHttpClient(client),
	}
}
func (tsc *HTTPBarGoTSRPCClient) AttributeMapping(ctx go_context.Context) (retAttributeMapping_0 AttributeMapping, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&retAttributeMapping_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "AttributeMapping", args, reply)
	return
}

func (tsc *HTTPBarGoTSRPCClient) CustomError(ctx go_context.Context, one CustomError, two *CustomError) (three CustomError, four *CustomError, clientErr error) {
	args := []interface{}{one, two}
	reply := []interface{}{&three, &four}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "CustomError", args, reply)
	return
}

func (tsc *HTTPBarGoTSRPCClient) CustomType(ctx go_context.Context, customTypeInt CustomTypeInt, customTypeString CustomTypeString, CustomTypeStruct CustomTypeStruct) (retCustomType_0 *CustomTypeInt, retCustomType_1 *CustomTypeString, retCustomType_2 CustomTypeStruct, clientErr error) {
	args := []interface{}{customTypeInt, customTypeString, CustomTypeStruct}
	reply := []interface{}{&retCustomType_0, &retCustomType_1, &retCustomType_2}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "CustomType", args, reply)
	return
}

func (tsc *HTTPBarGoTSRPCClient) Hello(ctx go_context.Context, number int64) (retHello_0 int, clientErr error) {
	args := []interface{}{number}
	reply := []interface{}{&retHello_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Hello", args, reply)
	return
}

func (tsc *HTTPBarGoTSRPCClient) Inheritance(ctx go_context.Context, inner Inner, nested OuterNested, inline OuterInline) (retInheritance_0 Inner, retInheritance_1 OuterNested, retInheritance_2 OuterInline, clientErr error) {
	args := []interface{}{inner, nested, inline}
	reply := []interface{}{&retInheritance_0, &retInheritance_1, &retInheritance_2}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Inheritance", args, reply)
	return
}

func (tsc *HTTPBarGoTSRPCClient) Repeat(ctx go_context.Context, one string, two string) (three bool, four bool, clientErr error) {
	args := []interface{}{one, two}
	reply := []interface{}{&three, &four}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Repeat", args, reply)
	return
}
