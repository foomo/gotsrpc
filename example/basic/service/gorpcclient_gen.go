// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package service

import (
	tls "crypto/tls"

	gorpc "github.com/valyala/gorpc"
)

type ServiceGoRPCClient struct {
	Client *gorpc.Client
}

func NewServiceGoRPCClient(addr string, tlsConfig *tls.Config) *ServiceGoRPCClient {
	client := &ServiceGoRPCClient{}
	if tlsConfig == nil {
		client.Client = gorpc.NewTCPClient(addr)
	} else {
		client.Client = gorpc.NewTLSClient(addr, tlsConfig)
	}
	return client
}

func (tsc *ServiceGoRPCClient) Start() {
	tsc.Client.Start()
}

func (tsc *ServiceGoRPCClient) Stop() {
	tsc.Client.Stop()
}

func (tsc *ServiceGoRPCClient) Bool(v bool) (retBool_0 bool, clientErr error) {
	req := ServiceBoolRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceBoolResponse)
	return response.RetBool_0, nil
}

func (tsc *ServiceGoRPCClient) BoolPtr(v bool) (retBoolPtr_0 *bool, clientErr error) {
	req := ServiceBoolPtrRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceBoolPtrResponse)
	return response.RetBoolPtr_0, nil
}

func (tsc *ServiceGoRPCClient) BoolSlice(v []bool) (retBoolSlice_0 []bool, clientErr error) {
	req := ServiceBoolSliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceBoolSliceResponse)
	return response.RetBoolSlice_0, nil
}

func (tsc *ServiceGoRPCClient) Empty() (clientErr error) {
	req := ServiceEmptyRequest{}
	_, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	return nil
}

func (tsc *ServiceGoRPCClient) Float32(v float32) (retFloat32_0 float32, clientErr error) {
	req := ServiceFloat32Request{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat32Response)
	return response.RetFloat32_0, nil
}

func (tsc *ServiceGoRPCClient) Float32Map(v map[float32]interface{}) (retFloat32Map_0 map[float32]interface{}, clientErr error) {
	req := ServiceFloat32MapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat32MapResponse)
	return response.RetFloat32Map_0, nil
}

func (tsc *ServiceGoRPCClient) Float32Slice(v []float32) (retFloat32Slice_0 []float32, clientErr error) {
	req := ServiceFloat32SliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat32SliceResponse)
	return response.RetFloat32Slice_0, nil
}

func (tsc *ServiceGoRPCClient) Float32Type(v Float32Type) (retFloat32Type_0 Float32Type, clientErr error) {
	req := ServiceFloat32TypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat32TypeResponse)
	return response.RetFloat32Type_0, nil
}

func (tsc *ServiceGoRPCClient) Float32TypeMap(v map[Float32TypeMapKey]Float32TypeMapValue) (retFloat32TypeMap_0 map[Float32TypeMapKey]Float32TypeMapValue, clientErr error) {
	req := ServiceFloat32TypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat32TypeMapResponse)
	return response.RetFloat32TypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) Float32TypeMapTyped(v Float32TypeMapTyped) (retFloat32TypeMapTyped_0 Float32TypeMapTyped, clientErr error) {
	req := ServiceFloat32TypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat32TypeMapTypedResponse)
	return response.RetFloat32TypeMapTyped_0, nil
}

func (tsc *ServiceGoRPCClient) Float64(v float64) (retFloat64_0 float64, clientErr error) {
	req := ServiceFloat64Request{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat64Response)
	return response.RetFloat64_0, nil
}

func (tsc *ServiceGoRPCClient) Float64Map(v map[float64]interface{}) (retFloat64Map_0 map[float64]interface{}, clientErr error) {
	req := ServiceFloat64MapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat64MapResponse)
	return response.RetFloat64Map_0, nil
}

func (tsc *ServiceGoRPCClient) Float64Slice(v []float64) (retFloat64Slice_0 []float64, clientErr error) {
	req := ServiceFloat64SliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat64SliceResponse)
	return response.RetFloat64Slice_0, nil
}

func (tsc *ServiceGoRPCClient) Float64Type(v Float64Type) (retFloat64Type_0 Float64Type, clientErr error) {
	req := ServiceFloat64TypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat64TypeResponse)
	return response.RetFloat64Type_0, nil
}

func (tsc *ServiceGoRPCClient) Float64TypeMap(v map[Float64TypeMapKey]Float64TypeMapValue) (retFloat64TypeMap_0 map[Float64TypeMapKey]Float64TypeMapValue, clientErr error) {
	req := ServiceFloat64TypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat64TypeMapResponse)
	return response.RetFloat64TypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) Float64TypeMapTyped(v Float64TypeMapTyped) (retFloat64TypeMapTyped_0 Float64TypeMapTyped, clientErr error) {
	req := ServiceFloat64TypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceFloat64TypeMapTypedResponse)
	return response.RetFloat64TypeMapTyped_0, nil
}

func (tsc *ServiceGoRPCClient) Int(v int) (retInt_0 int, clientErr error) {
	req := ServiceIntRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceIntResponse)
	return response.RetInt_0, nil
}

func (tsc *ServiceGoRPCClient) Int32(v int32) (retInt32_0 int32, clientErr error) {
	req := ServiceInt32Request{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt32Response)
	return response.RetInt32_0, nil
}

func (tsc *ServiceGoRPCClient) Int32Map(v map[int32]interface{}) (retInt32Map_0 map[int32]interface{}, clientErr error) {
	req := ServiceInt32MapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt32MapResponse)
	return response.RetInt32Map_0, nil
}

func (tsc *ServiceGoRPCClient) Int32Slice(v []int32) (retInt32Slice_0 []int32, clientErr error) {
	req := ServiceInt32SliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt32SliceResponse)
	return response.RetInt32Slice_0, nil
}

func (tsc *ServiceGoRPCClient) Int32Type(v Int32Type) (retInt32Type_0 Int32Type, clientErr error) {
	req := ServiceInt32TypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt32TypeResponse)
	return response.RetInt32Type_0, nil
}

func (tsc *ServiceGoRPCClient) Int32TypeMap(v map[Int32TypeMapKey]Int32TypeMapValue) (retInt32TypeMap_0 map[Int32TypeMapKey]Int32TypeMapValue, clientErr error) {
	req := ServiceInt32TypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt32TypeMapResponse)
	return response.RetInt32TypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) Int32TypeMapTyped(v Int32TypeMapTyped) (retInt32TypeMapTyped_0 Int32TypeMapTyped, clientErr error) {
	req := ServiceInt32TypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt32TypeMapTypedResponse)
	return response.RetInt32TypeMapTyped_0, nil
}

func (tsc *ServiceGoRPCClient) Int64(v int64) (retInt64_0 int64, clientErr error) {
	req := ServiceInt64Request{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt64Response)
	return response.RetInt64_0, nil
}

func (tsc *ServiceGoRPCClient) Int64Map(v map[int64]interface{}) (retInt64Map_0 map[int64]interface{}, clientErr error) {
	req := ServiceInt64MapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt64MapResponse)
	return response.RetInt64Map_0, nil
}

func (tsc *ServiceGoRPCClient) Int64Slice(v []int64) (retInt64Slice_0 []int64, clientErr error) {
	req := ServiceInt64SliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt64SliceResponse)
	return response.RetInt64Slice_0, nil
}

func (tsc *ServiceGoRPCClient) Int64Type(v Int64Type) (retInt64Type_0 Int64Type, clientErr error) {
	req := ServiceInt64TypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt64TypeResponse)
	return response.RetInt64Type_0, nil
}

func (tsc *ServiceGoRPCClient) Int64TypeMap(v map[Int64TypeMapKey]Int64TypeMapValue) (retInt64TypeMap_0 map[Int64TypeMapKey]Int64TypeMapValue, clientErr error) {
	req := ServiceInt64TypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt64TypeMapResponse)
	return response.RetInt64TypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) Int64TypeMapTyped(v Int64TypeMapTyped) (retInt64TypeMapTyped_0 Int64TypeMapTyped, clientErr error) {
	req := ServiceInt64TypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInt64TypeMapTypedResponse)
	return response.RetInt64TypeMapTyped_0, nil
}

func (tsc *ServiceGoRPCClient) IntMap(v map[int]interface{}) (retIntMap_0 map[int]interface{}, clientErr error) {
	req := ServiceIntMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceIntMapResponse)
	return response.RetIntMap_0, nil
}

func (tsc *ServiceGoRPCClient) IntSlice(v []int) (retIntSlice_0 []int, clientErr error) {
	req := ServiceIntSliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceIntSliceResponse)
	return response.RetIntSlice_0, nil
}

func (tsc *ServiceGoRPCClient) IntType(v IntType) (retIntType_0 IntType, clientErr error) {
	req := ServiceIntTypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceIntTypeResponse)
	return response.RetIntType_0, nil
}

func (tsc *ServiceGoRPCClient) IntTypeMap(v map[IntTypeMapKey]IntTypeMapValue) (retIntTypeMap_0 map[IntTypeMapKey]IntTypeMapValue, clientErr error) {
	req := ServiceIntTypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceIntTypeMapResponse)
	return response.RetIntTypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) IntTypeMapTyped(v IntTypeMapTyped) (retIntTypeMapTyped_0 IntTypeMapTyped, clientErr error) {
	req := ServiceIntTypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceIntTypeMapTypedResponse)
	return response.RetIntTypeMapTyped_0, nil
}

func (tsc *ServiceGoRPCClient) Interface(v interface{}) (retInterface_0 interface{}, clientErr error) {
	req := ServiceInterfaceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInterfaceResponse)
	return response.RetInterface_0, nil
}

func (tsc *ServiceGoRPCClient) InterfaceSlice(v []interface{}) (retInterfaceSlice_0 []interface{}, clientErr error) {
	req := ServiceInterfaceSliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceInterfaceSliceResponse)
	return response.RetInterfaceSlice_0, nil
}

func (tsc *ServiceGoRPCClient) NestedType() (retNestedType_0 NestedType, clientErr error) {
	req := ServiceNestedTypeRequest{}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceNestedTypeResponse)
	return response.RetNestedType_0, nil
}

func (tsc *ServiceGoRPCClient) String(v string) (retString_0 string, clientErr error) {
	req := ServiceStringRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceStringResponse)
	return response.RetString_0, nil
}

func (tsc *ServiceGoRPCClient) StringMap(v map[string]interface{}) (retStringMap_0 map[string]interface{}, clientErr error) {
	req := ServiceStringMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceStringMapResponse)
	return response.RetStringMap_0, nil
}

func (tsc *ServiceGoRPCClient) StringSlice(v []string) (retStringSlice_0 []string, clientErr error) {
	req := ServiceStringSliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceStringSliceResponse)
	return response.RetStringSlice_0, nil
}

func (tsc *ServiceGoRPCClient) StringType(v StringType) (retStringType_0 StringType, clientErr error) {
	req := ServiceStringTypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceStringTypeResponse)
	return response.RetStringType_0, nil
}

func (tsc *ServiceGoRPCClient) StringTypeMap(v map[StringTypeMapKey]StringTypeMapValue) (retStringTypeMap_0 map[StringTypeMapKey]StringTypeMapValue, clientErr error) {
	req := ServiceStringTypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceStringTypeMapResponse)
	return response.RetStringTypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) StringTypeMapTyped(v StringTypeMapTyped) (retStringTypeMapTyped_0 StringTypeMapTyped, clientErr error) {
	req := ServiceStringTypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceStringTypeMapTypedResponse)
	return response.RetStringTypeMapTyped_0, nil
}

func (tsc *ServiceGoRPCClient) Struct(v Struct) (retStruct_0 Struct, clientErr error) {
	req := ServiceStructRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceStructResponse)
	return response.RetStruct_0, nil
}

func (tsc *ServiceGoRPCClient) UInt(v uint) (retUInt_0 uint, clientErr error) {
	req := ServiceUIntRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUIntResponse)
	return response.RetUInt_0, nil
}

func (tsc *ServiceGoRPCClient) UInt32(v uint32) (retUInt32_0 uint32, clientErr error) {
	req := ServiceUInt32Request{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt32Response)
	return response.RetUInt32_0, nil
}

func (tsc *ServiceGoRPCClient) UInt32Map(v map[uint32]interface{}) (retUInt32Map_0 map[uint32]interface{}, clientErr error) {
	req := ServiceUInt32MapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt32MapResponse)
	return response.RetUInt32Map_0, nil
}

func (tsc *ServiceGoRPCClient) UInt32Slice(v []uint32) (retUInt32Slice_0 []uint32, clientErr error) {
	req := ServiceUInt32SliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt32SliceResponse)
	return response.RetUInt32Slice_0, nil
}

func (tsc *ServiceGoRPCClient) UInt32Type(v UInt32Type) (retUInt32Type_0 UInt32Type, clientErr error) {
	req := ServiceUInt32TypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt32TypeResponse)
	return response.RetUInt32Type_0, nil
}

func (tsc *ServiceGoRPCClient) UInt32TypeMap(v map[UInt32TypeMapKey]UInt32TypeMapValue) (retUInt32TypeMap_0 map[UInt32TypeMapKey]UInt32TypeMapValue, clientErr error) {
	req := ServiceUInt32TypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt32TypeMapResponse)
	return response.RetUInt32TypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) UInt32TypeMapTyped(v UInt32TypeMapTyped) (retUInt32TypeMapTyped_0 UInt32TypeMapTyped, clientErr error) {
	req := ServiceUInt32TypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt32TypeMapTypedResponse)
	return response.RetUInt32TypeMapTyped_0, nil
}

func (tsc *ServiceGoRPCClient) UInt64(v uint64) (retUInt64_0 uint64, clientErr error) {
	req := ServiceUInt64Request{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt64Response)
	return response.RetUInt64_0, nil
}

func (tsc *ServiceGoRPCClient) UInt64Map(v map[uint64]interface{}) (retUInt64Map_0 map[uint64]interface{}, clientErr error) {
	req := ServiceUInt64MapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt64MapResponse)
	return response.RetUInt64Map_0, nil
}

func (tsc *ServiceGoRPCClient) UInt64Slice(v []uint64) (retUInt64Slice_0 []uint64, clientErr error) {
	req := ServiceUInt64SliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt64SliceResponse)
	return response.RetUInt64Slice_0, nil
}

func (tsc *ServiceGoRPCClient) UInt64Type(v UInt64Type) (retUInt64Type_0 UInt64Type, clientErr error) {
	req := ServiceUInt64TypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt64TypeResponse)
	return response.RetUInt64Type_0, nil
}

func (tsc *ServiceGoRPCClient) UInt64TypeMap(v map[UInt64TypeMapKey]UInt64TypeMapValue) (retUInt64TypeMap_0 map[UInt64TypeMapKey]UInt64TypeMapValue, clientErr error) {
	req := ServiceUInt64TypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt64TypeMapResponse)
	return response.RetUInt64TypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) UInt64TypeMapTyped(v UInt64TypeMapTyped) (retUInt64TypeMapTyped_0 UInt64TypeMapTyped, clientErr error) {
	req := ServiceUInt64TypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUInt64TypeMapTypedResponse)
	return response.RetUInt64TypeMapTyped_0, nil
}

func (tsc *ServiceGoRPCClient) UIntMap(v map[uint]interface{}) (retUIntMap_0 map[uint]interface{}, clientErr error) {
	req := ServiceUIntMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUIntMapResponse)
	return response.RetUIntMap_0, nil
}

func (tsc *ServiceGoRPCClient) UIntSlice(v []uint) (retUIntSlice_0 []uint, clientErr error) {
	req := ServiceUIntSliceRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUIntSliceResponse)
	return response.RetUIntSlice_0, nil
}

func (tsc *ServiceGoRPCClient) UIntType(v UIntType) (retUIntType_0 UIntType, clientErr error) {
	req := ServiceUIntTypeRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUIntTypeResponse)
	return response.RetUIntType_0, nil
}

func (tsc *ServiceGoRPCClient) UIntTypeMap(v map[UIntTypeMapKey]UIntTypeMapValue) (retUIntTypeMap_0 map[UIntTypeMapKey]UIntTypeMapValue, clientErr error) {
	req := ServiceUIntTypeMapRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUIntTypeMapResponse)
	return response.RetUIntTypeMap_0, nil
}

func (tsc *ServiceGoRPCClient) UIntTypeMapTyped(v UIntTypeMapTyped) (retUIntTypeMapTyped_0 UIntTypeMapTyped, clientErr error) {
	req := ServiceUIntTypeMapTypedRequest{V: v}
	rpcCallRes, rpcCallErr := tsc.Client.Call(req)
	if rpcCallErr != nil {
		clientErr = rpcCallErr
		return
	}
	response := rpcCallRes.(ServiceUIntTypeMapTypedResponse)
	return response.RetUIntTypeMapTyped_0, nil
}
