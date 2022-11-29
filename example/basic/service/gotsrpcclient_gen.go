// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package service

import (
	go_context "context"
	go_net_http "net/http"

	gotsrpc "github.com/foomo/gotsrpc/v2"
	pkg_errors "github.com/pkg/errors"
)

type ServiceGoTSRPCClient interface {
	Bool(ctx go_context.Context, v bool) (retBool_0 bool, clientErr error)
	BoolPtr(ctx go_context.Context, v bool) (retBoolPtr_0 *bool, clientErr error)
	BoolSlice(ctx go_context.Context, v []bool) (retBoolSlice_0 []bool, clientErr error)
	Empty(ctx go_context.Context) (clientErr error)
	Float32(ctx go_context.Context, v float32) (retFloat32_0 float32, clientErr error)
	Float32Map(ctx go_context.Context, v map[float32]interface{}) (retFloat32Map_0 map[float32]interface{}, clientErr error)
	Float32Slice(ctx go_context.Context, v []float32) (retFloat32Slice_0 []float32, clientErr error)
	Float32Type(ctx go_context.Context, v Float32Type) (retFloat32Type_0 Float32Type, clientErr error)
	Float32TypeMap(ctx go_context.Context, v map[Float32TypeMapKey]Float32TypeMapValue) (retFloat32TypeMap_0 map[Float32TypeMapKey]Float32TypeMapValue, clientErr error)
	Float32TypeMapTyped(ctx go_context.Context, v Float32TypeMapTyped) (retFloat32TypeMapTyped_0 Float32TypeMapTyped, clientErr error)
	Float64(ctx go_context.Context, v float64) (retFloat64_0 float64, clientErr error)
	Float64Map(ctx go_context.Context, v map[float64]interface{}) (retFloat64Map_0 map[float64]interface{}, clientErr error)
	Float64Slice(ctx go_context.Context, v []float64) (retFloat64Slice_0 []float64, clientErr error)
	Float64Type(ctx go_context.Context, v Float64Type) (retFloat64Type_0 Float64Type, clientErr error)
	Float64TypeMap(ctx go_context.Context, v map[Float64TypeMapKey]Float64TypeMapValue) (retFloat64TypeMap_0 map[Float64TypeMapKey]Float64TypeMapValue, clientErr error)
	Float64TypeMapTyped(ctx go_context.Context, v Float64TypeMapTyped) (retFloat64TypeMapTyped_0 Float64TypeMapTyped, clientErr error)
	Int(ctx go_context.Context, v int) (retInt_0 int, clientErr error)
	Int32(ctx go_context.Context, v int32) (retInt32_0 int32, clientErr error)
	Int32Map(ctx go_context.Context, v map[int32]interface{}) (retInt32Map_0 map[int32]interface{}, clientErr error)
	Int32Slice(ctx go_context.Context, v []int32) (retInt32Slice_0 []int32, clientErr error)
	Int32Type(ctx go_context.Context, v Int32Type) (retInt32Type_0 Int32Type, clientErr error)
	Int32TypeMap(ctx go_context.Context, v map[Int32TypeMapKey]Int32TypeMapValue) (retInt32TypeMap_0 map[Int32TypeMapKey]Int32TypeMapValue, clientErr error)
	Int32TypeMapTyped(ctx go_context.Context, v Int32TypeMapTyped) (retInt32TypeMapTyped_0 Int32TypeMapTyped, clientErr error)
	Int64(ctx go_context.Context, v int64) (retInt64_0 int64, clientErr error)
	Int64Map(ctx go_context.Context, v map[int64]interface{}) (retInt64Map_0 map[int64]interface{}, clientErr error)
	Int64Slice(ctx go_context.Context, v []int64) (retInt64Slice_0 []int64, clientErr error)
	Int64Type(ctx go_context.Context, v Int64Type) (retInt64Type_0 Int64Type, clientErr error)
	Int64TypeMap(ctx go_context.Context, v map[Int64TypeMapKey]Int64TypeMapValue) (retInt64TypeMap_0 map[Int64TypeMapKey]Int64TypeMapValue, clientErr error)
	Int64TypeMapTyped(ctx go_context.Context, v Int64TypeMapTyped) (retInt64TypeMapTyped_0 Int64TypeMapTyped, clientErr error)
	IntMap(ctx go_context.Context, v map[int]interface{}) (retIntMap_0 map[int]interface{}, clientErr error)
	IntSlice(ctx go_context.Context, v []int) (retIntSlice_0 []int, clientErr error)
	IntType(ctx go_context.Context, v IntType) (retIntType_0 IntType, clientErr error)
	IntTypeMap(ctx go_context.Context, v map[IntTypeMapKey]IntTypeMapValue) (retIntTypeMap_0 map[IntTypeMapKey]IntTypeMapValue, clientErr error)
	IntTypeMapTyped(ctx go_context.Context, v IntTypeMapTyped) (retIntTypeMapTyped_0 IntTypeMapTyped, clientErr error)
	Interface(ctx go_context.Context, v interface{}) (retInterface_0 interface{}, clientErr error)
	InterfaceSlice(ctx go_context.Context, v []interface{}) (retInterfaceSlice_0 []interface{}, clientErr error)
	NestedType(ctx go_context.Context) (retNestedType_0 NestedType, clientErr error)
	String(ctx go_context.Context, v string) (retString_0 string, clientErr error)
	StringMap(ctx go_context.Context, v map[string]interface{}) (retStringMap_0 map[string]interface{}, clientErr error)
	StringSlice(ctx go_context.Context, v []string) (retStringSlice_0 []string, clientErr error)
	StringType(ctx go_context.Context, v StringType) (retStringType_0 StringType, clientErr error)
	StringTypeMap(ctx go_context.Context, v map[StringTypeMapKey]StringTypeMapValue) (retStringTypeMap_0 map[StringTypeMapKey]StringTypeMapValue, clientErr error)
	StringTypeMapTyped(ctx go_context.Context, v StringTypeMapTyped) (retStringTypeMapTyped_0 StringTypeMapTyped, clientErr error)
	Struct(ctx go_context.Context, v Struct) (retStruct_0 Struct, clientErr error)
	UInt(ctx go_context.Context, v uint) (retUInt_0 uint, clientErr error)
	UInt32(ctx go_context.Context, v uint32) (retUInt32_0 uint32, clientErr error)
	UInt32Map(ctx go_context.Context, v map[uint32]interface{}) (retUInt32Map_0 map[uint32]interface{}, clientErr error)
	UInt32Slice(ctx go_context.Context, v []uint32) (retUInt32Slice_0 []uint32, clientErr error)
	UInt32Type(ctx go_context.Context, v UInt32Type) (retUInt32Type_0 UInt32Type, clientErr error)
	UInt32TypeMap(ctx go_context.Context, v map[UInt32TypeMapKey]UInt32TypeMapValue) (retUInt32TypeMap_0 map[UInt32TypeMapKey]UInt32TypeMapValue, clientErr error)
	UInt32TypeMapTyped(ctx go_context.Context, v UInt32TypeMapTyped) (retUInt32TypeMapTyped_0 UInt32TypeMapTyped, clientErr error)
	UInt64(ctx go_context.Context, v uint64) (retUInt64_0 uint64, clientErr error)
	UInt64Map(ctx go_context.Context, v map[uint64]interface{}) (retUInt64Map_0 map[uint64]interface{}, clientErr error)
	UInt64Slice(ctx go_context.Context, v []uint64) (retUInt64Slice_0 []uint64, clientErr error)
	UInt64Type(ctx go_context.Context, v UInt64Type) (retUInt64Type_0 UInt64Type, clientErr error)
	UInt64TypeMap(ctx go_context.Context, v map[UInt64TypeMapKey]UInt64TypeMapValue) (retUInt64TypeMap_0 map[UInt64TypeMapKey]UInt64TypeMapValue, clientErr error)
	UInt64TypeMapTyped(ctx go_context.Context, v UInt64TypeMapTyped) (retUInt64TypeMapTyped_0 UInt64TypeMapTyped, clientErr error)
	UIntMap(ctx go_context.Context, v map[uint]interface{}) (retUIntMap_0 map[uint]interface{}, clientErr error)
	UIntSlice(ctx go_context.Context, v []uint) (retUIntSlice_0 []uint, clientErr error)
	UIntType(ctx go_context.Context, v UIntType) (retUIntType_0 UIntType, clientErr error)
	UIntTypeMap(ctx go_context.Context, v map[UIntTypeMapKey]UIntTypeMapValue) (retUIntTypeMap_0 map[UIntTypeMapKey]UIntTypeMapValue, clientErr error)
	UIntTypeMapTyped(ctx go_context.Context, v UIntTypeMapTyped) (retUIntTypeMapTyped_0 UIntTypeMapTyped, clientErr error)
}

type HTTPServiceGoTSRPCClient struct {
	URL      string
	EndPoint string
	Client   gotsrpc.Client
}

func NewDefaultServiceGoTSRPCClient(url string) *HTTPServiceGoTSRPCClient {
	return NewServiceGoTSRPCClient(url, "/service")
}

func NewServiceGoTSRPCClient(url string, endpoint string) *HTTPServiceGoTSRPCClient {
	return NewServiceGoTSRPCClientWithClient(url, endpoint, nil)
}

func NewServiceGoTSRPCClientWithClient(url string, endpoint string, client *go_net_http.Client) *HTTPServiceGoTSRPCClient {
	return &HTTPServiceGoTSRPCClient{
		URL:      url,
		EndPoint: endpoint,
		Client:   gotsrpc.NewClientWithHttpClient(client),
	}
}
func (tsc *HTTPServiceGoTSRPCClient) Bool(ctx go_context.Context, v bool) (retBool_0 bool, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retBool_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Bool", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Bool")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) BoolPtr(ctx go_context.Context, v bool) (retBoolPtr_0 *bool, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retBoolPtr_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "BoolPtr", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy BoolPtr")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) BoolSlice(ctx go_context.Context, v []bool) (retBoolSlice_0 []bool, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retBoolSlice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "BoolSlice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy BoolSlice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Empty(ctx go_context.Context) (clientErr error) {
	args := []interface{}{}
	reply := []interface{}{}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Empty", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Empty")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float32(ctx go_context.Context, v float32) (retFloat32_0 float32, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat32_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float32", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float32")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float32Map(ctx go_context.Context, v map[float32]interface{}) (retFloat32Map_0 map[float32]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat32Map_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float32Map", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float32Map")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float32Slice(ctx go_context.Context, v []float32) (retFloat32Slice_0 []float32, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat32Slice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float32Slice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float32Slice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float32Type(ctx go_context.Context, v Float32Type) (retFloat32Type_0 Float32Type, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat32Type_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float32Type", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float32Type")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float32TypeMap(ctx go_context.Context, v map[Float32TypeMapKey]Float32TypeMapValue) (retFloat32TypeMap_0 map[Float32TypeMapKey]Float32TypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat32TypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float32TypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float32TypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float32TypeMapTyped(ctx go_context.Context, v Float32TypeMapTyped) (retFloat32TypeMapTyped_0 Float32TypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat32TypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float32TypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float32TypeMapTyped")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float64(ctx go_context.Context, v float64) (retFloat64_0 float64, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat64_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float64", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float64")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float64Map(ctx go_context.Context, v map[float64]interface{}) (retFloat64Map_0 map[float64]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat64Map_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float64Map", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float64Map")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float64Slice(ctx go_context.Context, v []float64) (retFloat64Slice_0 []float64, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat64Slice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float64Slice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float64Slice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float64Type(ctx go_context.Context, v Float64Type) (retFloat64Type_0 Float64Type, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat64Type_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float64Type", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float64Type")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float64TypeMap(ctx go_context.Context, v map[Float64TypeMapKey]Float64TypeMapValue) (retFloat64TypeMap_0 map[Float64TypeMapKey]Float64TypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat64TypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float64TypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float64TypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Float64TypeMapTyped(ctx go_context.Context, v Float64TypeMapTyped) (retFloat64TypeMapTyped_0 Float64TypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retFloat64TypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Float64TypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Float64TypeMapTyped")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int(ctx go_context.Context, v int) (retInt_0 int, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int32(ctx go_context.Context, v int32) (retInt32_0 int32, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt32_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int32", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int32")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int32Map(ctx go_context.Context, v map[int32]interface{}) (retInt32Map_0 map[int32]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt32Map_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int32Map", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int32Map")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int32Slice(ctx go_context.Context, v []int32) (retInt32Slice_0 []int32, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt32Slice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int32Slice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int32Slice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int32Type(ctx go_context.Context, v Int32Type) (retInt32Type_0 Int32Type, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt32Type_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int32Type", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int32Type")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int32TypeMap(ctx go_context.Context, v map[Int32TypeMapKey]Int32TypeMapValue) (retInt32TypeMap_0 map[Int32TypeMapKey]Int32TypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt32TypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int32TypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int32TypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int32TypeMapTyped(ctx go_context.Context, v Int32TypeMapTyped) (retInt32TypeMapTyped_0 Int32TypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt32TypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int32TypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int32TypeMapTyped")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int64(ctx go_context.Context, v int64) (retInt64_0 int64, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt64_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int64", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int64")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int64Map(ctx go_context.Context, v map[int64]interface{}) (retInt64Map_0 map[int64]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt64Map_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int64Map", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int64Map")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int64Slice(ctx go_context.Context, v []int64) (retInt64Slice_0 []int64, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt64Slice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int64Slice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int64Slice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int64Type(ctx go_context.Context, v Int64Type) (retInt64Type_0 Int64Type, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt64Type_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int64Type", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int64Type")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int64TypeMap(ctx go_context.Context, v map[Int64TypeMapKey]Int64TypeMapValue) (retInt64TypeMap_0 map[Int64TypeMapKey]Int64TypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt64TypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int64TypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int64TypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Int64TypeMapTyped(ctx go_context.Context, v Int64TypeMapTyped) (retInt64TypeMapTyped_0 Int64TypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInt64TypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Int64TypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Int64TypeMapTyped")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) IntMap(ctx go_context.Context, v map[int]interface{}) (retIntMap_0 map[int]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retIntMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "IntMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy IntMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) IntSlice(ctx go_context.Context, v []int) (retIntSlice_0 []int, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retIntSlice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "IntSlice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy IntSlice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) IntType(ctx go_context.Context, v IntType) (retIntType_0 IntType, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retIntType_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "IntType", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy IntType")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) IntTypeMap(ctx go_context.Context, v map[IntTypeMapKey]IntTypeMapValue) (retIntTypeMap_0 map[IntTypeMapKey]IntTypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retIntTypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "IntTypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy IntTypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) IntTypeMapTyped(ctx go_context.Context, v IntTypeMapTyped) (retIntTypeMapTyped_0 IntTypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retIntTypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "IntTypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy IntTypeMapTyped")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Interface(ctx go_context.Context, v interface{}) (retInterface_0 interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInterface_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Interface", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Interface")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) InterfaceSlice(ctx go_context.Context, v []interface{}) (retInterfaceSlice_0 []interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retInterfaceSlice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "InterfaceSlice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy InterfaceSlice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) NestedType(ctx go_context.Context) (retNestedType_0 NestedType, clientErr error) {
	args := []interface{}{}
	reply := []interface{}{&retNestedType_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "NestedType", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy NestedType")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) String(ctx go_context.Context, v string) (retString_0 string, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retString_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "String", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy String")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) StringMap(ctx go_context.Context, v map[string]interface{}) (retStringMap_0 map[string]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retStringMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "StringMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy StringMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) StringSlice(ctx go_context.Context, v []string) (retStringSlice_0 []string, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retStringSlice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "StringSlice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy StringSlice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) StringType(ctx go_context.Context, v StringType) (retStringType_0 StringType, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retStringType_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "StringType", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy StringType")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) StringTypeMap(ctx go_context.Context, v map[StringTypeMapKey]StringTypeMapValue) (retStringTypeMap_0 map[StringTypeMapKey]StringTypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retStringTypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "StringTypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy StringTypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) StringTypeMapTyped(ctx go_context.Context, v StringTypeMapTyped) (retStringTypeMapTyped_0 StringTypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retStringTypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "StringTypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy StringTypeMapTyped")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) Struct(ctx go_context.Context, v Struct) (retStruct_0 Struct, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retStruct_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "Struct", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy Struct")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt(ctx go_context.Context, v uint) (retUInt_0 uint, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt32(ctx go_context.Context, v uint32) (retUInt32_0 uint32, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt32_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt32", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt32")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt32Map(ctx go_context.Context, v map[uint32]interface{}) (retUInt32Map_0 map[uint32]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt32Map_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt32Map", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt32Map")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt32Slice(ctx go_context.Context, v []uint32) (retUInt32Slice_0 []uint32, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt32Slice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt32Slice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt32Slice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt32Type(ctx go_context.Context, v UInt32Type) (retUInt32Type_0 UInt32Type, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt32Type_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt32Type", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt32Type")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt32TypeMap(ctx go_context.Context, v map[UInt32TypeMapKey]UInt32TypeMapValue) (retUInt32TypeMap_0 map[UInt32TypeMapKey]UInt32TypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt32TypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt32TypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt32TypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt32TypeMapTyped(ctx go_context.Context, v UInt32TypeMapTyped) (retUInt32TypeMapTyped_0 UInt32TypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt32TypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt32TypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt32TypeMapTyped")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt64(ctx go_context.Context, v uint64) (retUInt64_0 uint64, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt64_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt64", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt64")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt64Map(ctx go_context.Context, v map[uint64]interface{}) (retUInt64Map_0 map[uint64]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt64Map_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt64Map", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt64Map")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt64Slice(ctx go_context.Context, v []uint64) (retUInt64Slice_0 []uint64, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt64Slice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt64Slice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt64Slice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt64Type(ctx go_context.Context, v UInt64Type) (retUInt64Type_0 UInt64Type, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt64Type_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt64Type", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt64Type")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt64TypeMap(ctx go_context.Context, v map[UInt64TypeMapKey]UInt64TypeMapValue) (retUInt64TypeMap_0 map[UInt64TypeMapKey]UInt64TypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt64TypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt64TypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt64TypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UInt64TypeMapTyped(ctx go_context.Context, v UInt64TypeMapTyped) (retUInt64TypeMapTyped_0 UInt64TypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUInt64TypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UInt64TypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UInt64TypeMapTyped")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UIntMap(ctx go_context.Context, v map[uint]interface{}) (retUIntMap_0 map[uint]interface{}, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUIntMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UIntMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UIntMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UIntSlice(ctx go_context.Context, v []uint) (retUIntSlice_0 []uint, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUIntSlice_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UIntSlice", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UIntSlice")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UIntType(ctx go_context.Context, v UIntType) (retUIntType_0 UIntType, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUIntType_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UIntType", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UIntType")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UIntTypeMap(ctx go_context.Context, v map[UIntTypeMapKey]UIntTypeMapValue) (retUIntTypeMap_0 map[UIntTypeMapKey]UIntTypeMapValue, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUIntTypeMap_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UIntTypeMap", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UIntTypeMap")
	}
	return
}

func (tsc *HTTPServiceGoTSRPCClient) UIntTypeMapTyped(ctx go_context.Context, v UIntTypeMapTyped) (retUIntTypeMapTyped_0 UIntTypeMapTyped, clientErr error) {
	args := []interface{}{v}
	reply := []interface{}{&retUIntTypeMapTyped_0}
	clientErr = tsc.Client.Call(ctx, tsc.URL, tsc.EndPoint, "UIntTypeMapTyped", args, reply)
	if clientErr != nil {
		clientErr = pkg_errors.WithMessage(clientErr, "failed to call service.ServiceGoTSRPCProxy UIntTypeMapTyped")
	}
	return
}
