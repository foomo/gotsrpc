// Code generated by gotsrpc https://github.com/foomo/gotsrpc/v2  - DO NOT EDIT.

package service

import (
	tls "crypto/tls"
	gob "encoding/gob"
	fmt "fmt"
	reflect "reflect"
	strings "strings"
	time "time"

	gotsrpc "github.com/foomo/gotsrpc/v2"
	gorpc "github.com/valyala/gorpc"
)

type (
	ServiceGoRPCProxy struct {
		server           *gorpc.Server
		service          Service
		callStatsHandler gotsrpc.GoRPCCallStatsHandlerFun
	}

	ServiceBoolRequest struct {
		V bool
	}
	ServiceBoolResponse struct {
		RetBool_0 bool
	}

	ServiceBoolPtrRequest struct {
		V bool
	}
	ServiceBoolPtrResponse struct {
		RetBoolPtr_0 *bool
	}

	ServiceBoolSliceRequest struct {
		V []bool
	}
	ServiceBoolSliceResponse struct {
		RetBoolSlice_0 []bool
	}

	ServiceFloat32Request struct {
		V float32
	}
	ServiceFloat32Response struct {
		RetFloat32_0 float32
	}

	ServiceFloat32MapRequest struct {
		V map[float32]interface{}
	}
	ServiceFloat32MapResponse struct {
		RetFloat32Map_0 map[float32]interface{}
	}

	ServiceFloat32SliceRequest struct {
		V []float32
	}
	ServiceFloat32SliceResponse struct {
		RetFloat32Slice_0 []float32
	}

	ServiceFloat32TypeRequest struct {
		V Float32Type
	}
	ServiceFloat32TypeResponse struct {
		RetFloat32Type_0 Float32Type
	}

	ServiceFloat32TypeMapRequest struct {
		V map[Float32TypeMapKey]Float32TypeMapValue
	}
	ServiceFloat32TypeMapResponse struct {
		RetFloat32TypeMap_0 map[Float32TypeMapKey]Float32TypeMapValue
	}

	ServiceFloat32TypeMapTypedRequest struct {
		V Float32TypeMapTyped
	}
	ServiceFloat32TypeMapTypedResponse struct {
		RetFloat32TypeMapTyped_0 Float32TypeMapTyped
	}

	ServiceFloat64Request struct {
		V float64
	}
	ServiceFloat64Response struct {
		RetFloat64_0 float64
	}

	ServiceFloat64MapRequest struct {
		V map[float64]interface{}
	}
	ServiceFloat64MapResponse struct {
		RetFloat64Map_0 map[float64]interface{}
	}

	ServiceFloat64SliceRequest struct {
		V []float64
	}
	ServiceFloat64SliceResponse struct {
		RetFloat64Slice_0 []float64
	}

	ServiceFloat64TypeRequest struct {
		V Float64Type
	}
	ServiceFloat64TypeResponse struct {
		RetFloat64Type_0 Float64Type
	}

	ServiceFloat64TypeMapRequest struct {
		V map[Float64TypeMapKey]Float64TypeMapValue
	}
	ServiceFloat64TypeMapResponse struct {
		RetFloat64TypeMap_0 map[Float64TypeMapKey]Float64TypeMapValue
	}

	ServiceFloat64TypeMapTypedRequest struct {
		V Float64TypeMapTyped
	}
	ServiceFloat64TypeMapTypedResponse struct {
		RetFloat64TypeMapTyped_0 Float64TypeMapTyped
	}

	ServiceIntRequest struct {
		V int
	}
	ServiceIntResponse struct {
		RetInt_0 int
	}

	ServiceInt32Request struct {
		V int32
	}
	ServiceInt32Response struct {
		RetInt32_0 int32
	}

	ServiceInt32MapRequest struct {
		V map[int32]interface{}
	}
	ServiceInt32MapResponse struct {
		RetInt32Map_0 map[int32]interface{}
	}

	ServiceInt32SliceRequest struct {
		V []int32
	}
	ServiceInt32SliceResponse struct {
		RetInt32Slice_0 []int32
	}

	ServiceInt32TypeRequest struct {
		V Int32Type
	}
	ServiceInt32TypeResponse struct {
		RetInt32Type_0 Int32Type
	}

	ServiceInt32TypeMapRequest struct {
		V map[Int32TypeMapKey]Int32TypeMapValue
	}
	ServiceInt32TypeMapResponse struct {
		RetInt32TypeMap_0 map[Int32TypeMapKey]Int32TypeMapValue
	}

	ServiceInt32TypeMapTypedRequest struct {
		V Int32TypeMapTyped
	}
	ServiceInt32TypeMapTypedResponse struct {
		RetInt32TypeMapTyped_0 Int32TypeMapTyped
	}

	ServiceInt64Request struct {
		V int64
	}
	ServiceInt64Response struct {
		RetInt64_0 int64
	}

	ServiceInt64MapRequest struct {
		V map[int64]interface{}
	}
	ServiceInt64MapResponse struct {
		RetInt64Map_0 map[int64]interface{}
	}

	ServiceInt64SliceRequest struct {
		V []int64
	}
	ServiceInt64SliceResponse struct {
		RetInt64Slice_0 []int64
	}

	ServiceInt64TypeRequest struct {
		V Int64Type
	}
	ServiceInt64TypeResponse struct {
		RetInt64Type_0 Int64Type
	}

	ServiceInt64TypeMapRequest struct {
		V map[Int64TypeMapKey]Int64TypeMapValue
	}
	ServiceInt64TypeMapResponse struct {
		RetInt64TypeMap_0 map[Int64TypeMapKey]Int64TypeMapValue
	}

	ServiceInt64TypeMapTypedRequest struct {
		V Int64TypeMapTyped
	}
	ServiceInt64TypeMapTypedResponse struct {
		RetInt64TypeMapTyped_0 Int64TypeMapTyped
	}

	ServiceIntMapRequest struct {
		V map[int]interface{}
	}
	ServiceIntMapResponse struct {
		RetIntMap_0 map[int]interface{}
	}

	ServiceIntSliceRequest struct {
		V []int
	}
	ServiceIntSliceResponse struct {
		RetIntSlice_0 []int
	}

	ServiceIntTypeRequest struct {
		V IntType
	}
	ServiceIntTypeResponse struct {
		RetIntType_0 IntType
	}

	ServiceIntTypeMapRequest struct {
		V map[IntTypeMapKey]IntTypeMapValue
	}
	ServiceIntTypeMapResponse struct {
		RetIntTypeMap_0 map[IntTypeMapKey]IntTypeMapValue
	}

	ServiceIntTypeMapTypedRequest struct {
		V IntTypeMapTyped
	}
	ServiceIntTypeMapTypedResponse struct {
		RetIntTypeMapTyped_0 IntTypeMapTyped
	}

	ServiceInterfaceRequest struct {
		V interface{}
	}
	ServiceInterfaceResponse struct {
		RetInterface_0 interface{}
	}

	ServiceInterfaceSliceRequest struct {
		V []interface{}
	}
	ServiceInterfaceSliceResponse struct {
		RetInterfaceSlice_0 []interface{}
	}

	ServiceStringRequest struct {
		V string
	}
	ServiceStringResponse struct {
		RetString_0 string
	}

	ServiceStringMapRequest struct {
		V map[string]interface{}
	}
	ServiceStringMapResponse struct {
		RetStringMap_0 map[string]interface{}
	}

	ServiceStringSliceRequest struct {
		V []string
	}
	ServiceStringSliceResponse struct {
		RetStringSlice_0 []string
	}

	ServiceStringTypeRequest struct {
		V StringType
	}
	ServiceStringTypeResponse struct {
		RetStringType_0 StringType
	}

	ServiceStringTypeMapRequest struct {
		V map[StringTypeMapKey]StringTypeMapValue
	}
	ServiceStringTypeMapResponse struct {
		RetStringTypeMap_0 map[StringTypeMapKey]StringTypeMapValue
	}

	ServiceStringTypeMapTypedRequest struct {
		V StringTypeMapTyped
	}
	ServiceStringTypeMapTypedResponse struct {
		RetStringTypeMapTyped_0 StringTypeMapTyped
	}

	ServiceStructRequest struct {
		V Struct
	}
	ServiceStructResponse struct {
		RetStruct_0 Struct
	}

	ServiceUIntRequest struct {
		V uint
	}
	ServiceUIntResponse struct {
		RetUInt_0 uint
	}

	ServiceUInt32Request struct {
		V uint32
	}
	ServiceUInt32Response struct {
		RetUInt32_0 uint32
	}

	ServiceUInt32MapRequest struct {
		V map[uint32]interface{}
	}
	ServiceUInt32MapResponse struct {
		RetUInt32Map_0 map[uint32]interface{}
	}

	ServiceUInt32SliceRequest struct {
		V []uint32
	}
	ServiceUInt32SliceResponse struct {
		RetUInt32Slice_0 []uint32
	}

	ServiceUInt32TypeRequest struct {
		V UInt32Type
	}
	ServiceUInt32TypeResponse struct {
		RetUInt32Type_0 UInt32Type
	}

	ServiceUInt32TypeMapRequest struct {
		V map[UInt32TypeMapKey]UInt32TypeMapValue
	}
	ServiceUInt32TypeMapResponse struct {
		RetUInt32TypeMap_0 map[UInt32TypeMapKey]UInt32TypeMapValue
	}

	ServiceUInt32TypeMapTypedRequest struct {
		V UInt32TypeMapTyped
	}
	ServiceUInt32TypeMapTypedResponse struct {
		RetUInt32TypeMapTyped_0 UInt32TypeMapTyped
	}

	ServiceUInt64Request struct {
		V uint64
	}
	ServiceUInt64Response struct {
		RetUInt64_0 uint64
	}

	ServiceUInt64MapRequest struct {
		V map[uint64]interface{}
	}
	ServiceUInt64MapResponse struct {
		RetUInt64Map_0 map[uint64]interface{}
	}

	ServiceUInt64SliceRequest struct {
		V []uint64
	}
	ServiceUInt64SliceResponse struct {
		RetUInt64Slice_0 []uint64
	}

	ServiceUInt64TypeRequest struct {
		V UInt64Type
	}
	ServiceUInt64TypeResponse struct {
		RetUInt64Type_0 UInt64Type
	}

	ServiceUInt64TypeMapRequest struct {
		V map[UInt64TypeMapKey]UInt64TypeMapValue
	}
	ServiceUInt64TypeMapResponse struct {
		RetUInt64TypeMap_0 map[UInt64TypeMapKey]UInt64TypeMapValue
	}

	ServiceUInt64TypeMapTypedRequest struct {
		V UInt64TypeMapTyped
	}
	ServiceUInt64TypeMapTypedResponse struct {
		RetUInt64TypeMapTyped_0 UInt64TypeMapTyped
	}

	ServiceUIntMapRequest struct {
		V map[uint]interface{}
	}
	ServiceUIntMapResponse struct {
		RetUIntMap_0 map[uint]interface{}
	}

	ServiceUIntSliceRequest struct {
		V []uint
	}
	ServiceUIntSliceResponse struct {
		RetUIntSlice_0 []uint
	}

	ServiceUIntTypeRequest struct {
		V UIntType
	}
	ServiceUIntTypeResponse struct {
		RetUIntType_0 UIntType
	}

	ServiceUIntTypeMapRequest struct {
		V map[UIntTypeMapKey]UIntTypeMapValue
	}
	ServiceUIntTypeMapResponse struct {
		RetUIntTypeMap_0 map[UIntTypeMapKey]UIntTypeMapValue
	}

	ServiceUIntTypeMapTypedRequest struct {
		V UIntTypeMapTyped
	}
	ServiceUIntTypeMapTypedResponse struct {
		RetUIntTypeMapTyped_0 UIntTypeMapTyped
	}
)

func init() {
	gob.Register(ServiceBoolRequest{})
	gob.Register(ServiceBoolResponse{})
	gob.Register(ServiceBoolPtrRequest{})
	gob.Register(ServiceBoolPtrResponse{})
	gob.Register(ServiceBoolSliceRequest{})
	gob.Register(ServiceBoolSliceResponse{})
	gob.Register(ServiceFloat32Request{})
	gob.Register(ServiceFloat32Response{})
	gob.Register(ServiceFloat32MapRequest{})
	gob.Register(ServiceFloat32MapResponse{})
	gob.Register(ServiceFloat32SliceRequest{})
	gob.Register(ServiceFloat32SliceResponse{})
	gob.Register(ServiceFloat32TypeRequest{})
	gob.Register(ServiceFloat32TypeResponse{})
	gob.Register(ServiceFloat32TypeMapRequest{})
	gob.Register(ServiceFloat32TypeMapResponse{})
	gob.Register(ServiceFloat32TypeMapTypedRequest{})
	gob.Register(ServiceFloat32TypeMapTypedResponse{})
	gob.Register(ServiceFloat64Request{})
	gob.Register(ServiceFloat64Response{})
	gob.Register(ServiceFloat64MapRequest{})
	gob.Register(ServiceFloat64MapResponse{})
	gob.Register(ServiceFloat64SliceRequest{})
	gob.Register(ServiceFloat64SliceResponse{})
	gob.Register(ServiceFloat64TypeRequest{})
	gob.Register(ServiceFloat64TypeResponse{})
	gob.Register(ServiceFloat64TypeMapRequest{})
	gob.Register(ServiceFloat64TypeMapResponse{})
	gob.Register(ServiceFloat64TypeMapTypedRequest{})
	gob.Register(ServiceFloat64TypeMapTypedResponse{})
	gob.Register(ServiceIntRequest{})
	gob.Register(ServiceIntResponse{})
	gob.Register(ServiceInt32Request{})
	gob.Register(ServiceInt32Response{})
	gob.Register(ServiceInt32MapRequest{})
	gob.Register(ServiceInt32MapResponse{})
	gob.Register(ServiceInt32SliceRequest{})
	gob.Register(ServiceInt32SliceResponse{})
	gob.Register(ServiceInt32TypeRequest{})
	gob.Register(ServiceInt32TypeResponse{})
	gob.Register(ServiceInt32TypeMapRequest{})
	gob.Register(ServiceInt32TypeMapResponse{})
	gob.Register(ServiceInt32TypeMapTypedRequest{})
	gob.Register(ServiceInt32TypeMapTypedResponse{})
	gob.Register(ServiceInt64Request{})
	gob.Register(ServiceInt64Response{})
	gob.Register(ServiceInt64MapRequest{})
	gob.Register(ServiceInt64MapResponse{})
	gob.Register(ServiceInt64SliceRequest{})
	gob.Register(ServiceInt64SliceResponse{})
	gob.Register(ServiceInt64TypeRequest{})
	gob.Register(ServiceInt64TypeResponse{})
	gob.Register(ServiceInt64TypeMapRequest{})
	gob.Register(ServiceInt64TypeMapResponse{})
	gob.Register(ServiceInt64TypeMapTypedRequest{})
	gob.Register(ServiceInt64TypeMapTypedResponse{})
	gob.Register(ServiceIntMapRequest{})
	gob.Register(ServiceIntMapResponse{})
	gob.Register(ServiceIntSliceRequest{})
	gob.Register(ServiceIntSliceResponse{})
	gob.Register(ServiceIntTypeRequest{})
	gob.Register(ServiceIntTypeResponse{})
	gob.Register(ServiceIntTypeMapRequest{})
	gob.Register(ServiceIntTypeMapResponse{})
	gob.Register(ServiceIntTypeMapTypedRequest{})
	gob.Register(ServiceIntTypeMapTypedResponse{})
	gob.Register(ServiceInterfaceRequest{})
	gob.Register(ServiceInterfaceResponse{})
	gob.Register(ServiceInterfaceSliceRequest{})
	gob.Register(ServiceInterfaceSliceResponse{})
	gob.Register(ServiceStringRequest{})
	gob.Register(ServiceStringResponse{})
	gob.Register(ServiceStringMapRequest{})
	gob.Register(ServiceStringMapResponse{})
	gob.Register(ServiceStringSliceRequest{})
	gob.Register(ServiceStringSliceResponse{})
	gob.Register(ServiceStringTypeRequest{})
	gob.Register(ServiceStringTypeResponse{})
	gob.Register(ServiceStringTypeMapRequest{})
	gob.Register(ServiceStringTypeMapResponse{})
	gob.Register(ServiceStringTypeMapTypedRequest{})
	gob.Register(ServiceStringTypeMapTypedResponse{})
	gob.Register(ServiceStructRequest{})
	gob.Register(ServiceStructResponse{})
	gob.Register(ServiceUIntRequest{})
	gob.Register(ServiceUIntResponse{})
	gob.Register(ServiceUInt32Request{})
	gob.Register(ServiceUInt32Response{})
	gob.Register(ServiceUInt32MapRequest{})
	gob.Register(ServiceUInt32MapResponse{})
	gob.Register(ServiceUInt32SliceRequest{})
	gob.Register(ServiceUInt32SliceResponse{})
	gob.Register(ServiceUInt32TypeRequest{})
	gob.Register(ServiceUInt32TypeResponse{})
	gob.Register(ServiceUInt32TypeMapRequest{})
	gob.Register(ServiceUInt32TypeMapResponse{})
	gob.Register(ServiceUInt32TypeMapTypedRequest{})
	gob.Register(ServiceUInt32TypeMapTypedResponse{})
	gob.Register(ServiceUInt64Request{})
	gob.Register(ServiceUInt64Response{})
	gob.Register(ServiceUInt64MapRequest{})
	gob.Register(ServiceUInt64MapResponse{})
	gob.Register(ServiceUInt64SliceRequest{})
	gob.Register(ServiceUInt64SliceResponse{})
	gob.Register(ServiceUInt64TypeRequest{})
	gob.Register(ServiceUInt64TypeResponse{})
	gob.Register(ServiceUInt64TypeMapRequest{})
	gob.Register(ServiceUInt64TypeMapResponse{})
	gob.Register(ServiceUInt64TypeMapTypedRequest{})
	gob.Register(ServiceUInt64TypeMapTypedResponse{})
	gob.Register(ServiceUIntMapRequest{})
	gob.Register(ServiceUIntMapResponse{})
	gob.Register(ServiceUIntSliceRequest{})
	gob.Register(ServiceUIntSliceResponse{})
	gob.Register(ServiceUIntTypeRequest{})
	gob.Register(ServiceUIntTypeResponse{})
	gob.Register(ServiceUIntTypeMapRequest{})
	gob.Register(ServiceUIntTypeMapResponse{})
	gob.Register(ServiceUIntTypeMapTypedRequest{})
	gob.Register(ServiceUIntTypeMapTypedResponse{})
}

func NewServiceGoRPCProxy(addr string, service Service, tlsConfig *tls.Config) *ServiceGoRPCProxy {
	proxy := &ServiceGoRPCProxy{
		service: service,
	}

	if tlsConfig != nil {
		proxy.server = gorpc.NewTLSServer(addr, proxy.handler, tlsConfig)
	} else {
		proxy.server = gorpc.NewTCPServer(addr, proxy.handler)
	}

	return proxy
}

func (p *ServiceGoRPCProxy) Start() error {
	return p.server.Start()
}

func (p *ServiceGoRPCProxy) Serve() error {
	return p.server.Serve()
}

func (p *ServiceGoRPCProxy) Stop() {
	p.server.Stop()
}

func (p *ServiceGoRPCProxy) SetCallStatsHandler(handler gotsrpc.GoRPCCallStatsHandlerFun) {
	p.callStatsHandler = handler
}

func (p *ServiceGoRPCProxy) handler(clientAddr string, request interface{}) (response interface{}) {
	start := time.Now()

	reqType := reflect.TypeOf(request).String()
	funcNameParts := strings.Split(reqType, ".")
	funcName := funcNameParts[len(funcNameParts)-1]

	switch funcName {
	case "ServiceBoolRequest":
		req := request.(ServiceBoolRequest)
		retBool_0 := p.service.Bool(req.V)
		response = ServiceBoolResponse{RetBool_0: retBool_0}
	case "ServiceBoolPtrRequest":
		req := request.(ServiceBoolPtrRequest)
		retBoolPtr_0 := p.service.BoolPtr(req.V)
		response = ServiceBoolPtrResponse{RetBoolPtr_0: retBoolPtr_0}
	case "ServiceBoolSliceRequest":
		req := request.(ServiceBoolSliceRequest)
		retBoolSlice_0 := p.service.BoolSlice(req.V)
		response = ServiceBoolSliceResponse{RetBoolSlice_0: retBoolSlice_0}
	case "ServiceFloat32Request":
		req := request.(ServiceFloat32Request)
		retFloat32_0 := p.service.Float32(req.V)
		response = ServiceFloat32Response{RetFloat32_0: retFloat32_0}
	case "ServiceFloat32MapRequest":
		req := request.(ServiceFloat32MapRequest)
		retFloat32Map_0 := p.service.Float32Map(req.V)
		response = ServiceFloat32MapResponse{RetFloat32Map_0: retFloat32Map_0}
	case "ServiceFloat32SliceRequest":
		req := request.(ServiceFloat32SliceRequest)
		retFloat32Slice_0 := p.service.Float32Slice(req.V)
		response = ServiceFloat32SliceResponse{RetFloat32Slice_0: retFloat32Slice_0}
	case "ServiceFloat32TypeRequest":
		req := request.(ServiceFloat32TypeRequest)
		retFloat32Type_0 := p.service.Float32Type(req.V)
		response = ServiceFloat32TypeResponse{RetFloat32Type_0: retFloat32Type_0}
	case "ServiceFloat32TypeMapRequest":
		req := request.(ServiceFloat32TypeMapRequest)
		retFloat32TypeMap_0 := p.service.Float32TypeMap(req.V)
		response = ServiceFloat32TypeMapResponse{RetFloat32TypeMap_0: retFloat32TypeMap_0}
	case "ServiceFloat32TypeMapTypedRequest":
		req := request.(ServiceFloat32TypeMapTypedRequest)
		retFloat32TypeMapTyped_0 := p.service.Float32TypeMapTyped(req.V)
		response = ServiceFloat32TypeMapTypedResponse{RetFloat32TypeMapTyped_0: retFloat32TypeMapTyped_0}
	case "ServiceFloat64Request":
		req := request.(ServiceFloat64Request)
		retFloat64_0 := p.service.Float64(req.V)
		response = ServiceFloat64Response{RetFloat64_0: retFloat64_0}
	case "ServiceFloat64MapRequest":
		req := request.(ServiceFloat64MapRequest)
		retFloat64Map_0 := p.service.Float64Map(req.V)
		response = ServiceFloat64MapResponse{RetFloat64Map_0: retFloat64Map_0}
	case "ServiceFloat64SliceRequest":
		req := request.(ServiceFloat64SliceRequest)
		retFloat64Slice_0 := p.service.Float64Slice(req.V)
		response = ServiceFloat64SliceResponse{RetFloat64Slice_0: retFloat64Slice_0}
	case "ServiceFloat64TypeRequest":
		req := request.(ServiceFloat64TypeRequest)
		retFloat64Type_0 := p.service.Float64Type(req.V)
		response = ServiceFloat64TypeResponse{RetFloat64Type_0: retFloat64Type_0}
	case "ServiceFloat64TypeMapRequest":
		req := request.(ServiceFloat64TypeMapRequest)
		retFloat64TypeMap_0 := p.service.Float64TypeMap(req.V)
		response = ServiceFloat64TypeMapResponse{RetFloat64TypeMap_0: retFloat64TypeMap_0}
	case "ServiceFloat64TypeMapTypedRequest":
		req := request.(ServiceFloat64TypeMapTypedRequest)
		retFloat64TypeMapTyped_0 := p.service.Float64TypeMapTyped(req.V)
		response = ServiceFloat64TypeMapTypedResponse{RetFloat64TypeMapTyped_0: retFloat64TypeMapTyped_0}
	case "ServiceIntRequest":
		req := request.(ServiceIntRequest)
		retInt_0 := p.service.Int(req.V)
		response = ServiceIntResponse{RetInt_0: retInt_0}
	case "ServiceInt32Request":
		req := request.(ServiceInt32Request)
		retInt32_0 := p.service.Int32(req.V)
		response = ServiceInt32Response{RetInt32_0: retInt32_0}
	case "ServiceInt32MapRequest":
		req := request.(ServiceInt32MapRequest)
		retInt32Map_0 := p.service.Int32Map(req.V)
		response = ServiceInt32MapResponse{RetInt32Map_0: retInt32Map_0}
	case "ServiceInt32SliceRequest":
		req := request.(ServiceInt32SliceRequest)
		retInt32Slice_0 := p.service.Int32Slice(req.V)
		response = ServiceInt32SliceResponse{RetInt32Slice_0: retInt32Slice_0}
	case "ServiceInt32TypeRequest":
		req := request.(ServiceInt32TypeRequest)
		retInt32Type_0 := p.service.Int32Type(req.V)
		response = ServiceInt32TypeResponse{RetInt32Type_0: retInt32Type_0}
	case "ServiceInt32TypeMapRequest":
		req := request.(ServiceInt32TypeMapRequest)
		retInt32TypeMap_0 := p.service.Int32TypeMap(req.V)
		response = ServiceInt32TypeMapResponse{RetInt32TypeMap_0: retInt32TypeMap_0}
	case "ServiceInt32TypeMapTypedRequest":
		req := request.(ServiceInt32TypeMapTypedRequest)
		retInt32TypeMapTyped_0 := p.service.Int32TypeMapTyped(req.V)
		response = ServiceInt32TypeMapTypedResponse{RetInt32TypeMapTyped_0: retInt32TypeMapTyped_0}
	case "ServiceInt64Request":
		req := request.(ServiceInt64Request)
		retInt64_0 := p.service.Int64(req.V)
		response = ServiceInt64Response{RetInt64_0: retInt64_0}
	case "ServiceInt64MapRequest":
		req := request.(ServiceInt64MapRequest)
		retInt64Map_0 := p.service.Int64Map(req.V)
		response = ServiceInt64MapResponse{RetInt64Map_0: retInt64Map_0}
	case "ServiceInt64SliceRequest":
		req := request.(ServiceInt64SliceRequest)
		retInt64Slice_0 := p.service.Int64Slice(req.V)
		response = ServiceInt64SliceResponse{RetInt64Slice_0: retInt64Slice_0}
	case "ServiceInt64TypeRequest":
		req := request.(ServiceInt64TypeRequest)
		retInt64Type_0 := p.service.Int64Type(req.V)
		response = ServiceInt64TypeResponse{RetInt64Type_0: retInt64Type_0}
	case "ServiceInt64TypeMapRequest":
		req := request.(ServiceInt64TypeMapRequest)
		retInt64TypeMap_0 := p.service.Int64TypeMap(req.V)
		response = ServiceInt64TypeMapResponse{RetInt64TypeMap_0: retInt64TypeMap_0}
	case "ServiceInt64TypeMapTypedRequest":
		req := request.(ServiceInt64TypeMapTypedRequest)
		retInt64TypeMapTyped_0 := p.service.Int64TypeMapTyped(req.V)
		response = ServiceInt64TypeMapTypedResponse{RetInt64TypeMapTyped_0: retInt64TypeMapTyped_0}
	case "ServiceIntMapRequest":
		req := request.(ServiceIntMapRequest)
		retIntMap_0 := p.service.IntMap(req.V)
		response = ServiceIntMapResponse{RetIntMap_0: retIntMap_0}
	case "ServiceIntSliceRequest":
		req := request.(ServiceIntSliceRequest)
		retIntSlice_0 := p.service.IntSlice(req.V)
		response = ServiceIntSliceResponse{RetIntSlice_0: retIntSlice_0}
	case "ServiceIntTypeRequest":
		req := request.(ServiceIntTypeRequest)
		retIntType_0 := p.service.IntType(req.V)
		response = ServiceIntTypeResponse{RetIntType_0: retIntType_0}
	case "ServiceIntTypeMapRequest":
		req := request.(ServiceIntTypeMapRequest)
		retIntTypeMap_0 := p.service.IntTypeMap(req.V)
		response = ServiceIntTypeMapResponse{RetIntTypeMap_0: retIntTypeMap_0}
	case "ServiceIntTypeMapTypedRequest":
		req := request.(ServiceIntTypeMapTypedRequest)
		retIntTypeMapTyped_0 := p.service.IntTypeMapTyped(req.V)
		response = ServiceIntTypeMapTypedResponse{RetIntTypeMapTyped_0: retIntTypeMapTyped_0}
	case "ServiceInterfaceRequest":
		req := request.(ServiceInterfaceRequest)
		retInterface_0 := p.service.Interface(req.V)
		response = ServiceInterfaceResponse{RetInterface_0: retInterface_0}
	case "ServiceInterfaceSliceRequest":
		req := request.(ServiceInterfaceSliceRequest)
		retInterfaceSlice_0 := p.service.InterfaceSlice(req.V)
		response = ServiceInterfaceSliceResponse{RetInterfaceSlice_0: retInterfaceSlice_0}
	case "ServiceStringRequest":
		req := request.(ServiceStringRequest)
		retString_0 := p.service.String(req.V)
		response = ServiceStringResponse{RetString_0: retString_0}
	case "ServiceStringMapRequest":
		req := request.(ServiceStringMapRequest)
		retStringMap_0 := p.service.StringMap(req.V)
		response = ServiceStringMapResponse{RetStringMap_0: retStringMap_0}
	case "ServiceStringSliceRequest":
		req := request.(ServiceStringSliceRequest)
		retStringSlice_0 := p.service.StringSlice(req.V)
		response = ServiceStringSliceResponse{RetStringSlice_0: retStringSlice_0}
	case "ServiceStringTypeRequest":
		req := request.(ServiceStringTypeRequest)
		retStringType_0 := p.service.StringType(req.V)
		response = ServiceStringTypeResponse{RetStringType_0: retStringType_0}
	case "ServiceStringTypeMapRequest":
		req := request.(ServiceStringTypeMapRequest)
		retStringTypeMap_0 := p.service.StringTypeMap(req.V)
		response = ServiceStringTypeMapResponse{RetStringTypeMap_0: retStringTypeMap_0}
	case "ServiceStringTypeMapTypedRequest":
		req := request.(ServiceStringTypeMapTypedRequest)
		retStringTypeMapTyped_0 := p.service.StringTypeMapTyped(req.V)
		response = ServiceStringTypeMapTypedResponse{RetStringTypeMapTyped_0: retStringTypeMapTyped_0}
	case "ServiceStructRequest":
		req := request.(ServiceStructRequest)
		retStruct_0 := p.service.Struct(req.V)
		response = ServiceStructResponse{RetStruct_0: retStruct_0}
	case "ServiceUIntRequest":
		req := request.(ServiceUIntRequest)
		retUInt_0 := p.service.UInt(req.V)
		response = ServiceUIntResponse{RetUInt_0: retUInt_0}
	case "ServiceUInt32Request":
		req := request.(ServiceUInt32Request)
		retUInt32_0 := p.service.UInt32(req.V)
		response = ServiceUInt32Response{RetUInt32_0: retUInt32_0}
	case "ServiceUInt32MapRequest":
		req := request.(ServiceUInt32MapRequest)
		retUInt32Map_0 := p.service.UInt32Map(req.V)
		response = ServiceUInt32MapResponse{RetUInt32Map_0: retUInt32Map_0}
	case "ServiceUInt32SliceRequest":
		req := request.(ServiceUInt32SliceRequest)
		retUInt32Slice_0 := p.service.UInt32Slice(req.V)
		response = ServiceUInt32SliceResponse{RetUInt32Slice_0: retUInt32Slice_0}
	case "ServiceUInt32TypeRequest":
		req := request.(ServiceUInt32TypeRequest)
		retUInt32Type_0 := p.service.UInt32Type(req.V)
		response = ServiceUInt32TypeResponse{RetUInt32Type_0: retUInt32Type_0}
	case "ServiceUInt32TypeMapRequest":
		req := request.(ServiceUInt32TypeMapRequest)
		retUInt32TypeMap_0 := p.service.UInt32TypeMap(req.V)
		response = ServiceUInt32TypeMapResponse{RetUInt32TypeMap_0: retUInt32TypeMap_0}
	case "ServiceUInt32TypeMapTypedRequest":
		req := request.(ServiceUInt32TypeMapTypedRequest)
		retUInt32TypeMapTyped_0 := p.service.UInt32TypeMapTyped(req.V)
		response = ServiceUInt32TypeMapTypedResponse{RetUInt32TypeMapTyped_0: retUInt32TypeMapTyped_0}
	case "ServiceUInt64Request":
		req := request.(ServiceUInt64Request)
		retUInt64_0 := p.service.UInt64(req.V)
		response = ServiceUInt64Response{RetUInt64_0: retUInt64_0}
	case "ServiceUInt64MapRequest":
		req := request.(ServiceUInt64MapRequest)
		retUInt64Map_0 := p.service.UInt64Map(req.V)
		response = ServiceUInt64MapResponse{RetUInt64Map_0: retUInt64Map_0}
	case "ServiceUInt64SliceRequest":
		req := request.(ServiceUInt64SliceRequest)
		retUInt64Slice_0 := p.service.UInt64Slice(req.V)
		response = ServiceUInt64SliceResponse{RetUInt64Slice_0: retUInt64Slice_0}
	case "ServiceUInt64TypeRequest":
		req := request.(ServiceUInt64TypeRequest)
		retUInt64Type_0 := p.service.UInt64Type(req.V)
		response = ServiceUInt64TypeResponse{RetUInt64Type_0: retUInt64Type_0}
	case "ServiceUInt64TypeMapRequest":
		req := request.(ServiceUInt64TypeMapRequest)
		retUInt64TypeMap_0 := p.service.UInt64TypeMap(req.V)
		response = ServiceUInt64TypeMapResponse{RetUInt64TypeMap_0: retUInt64TypeMap_0}
	case "ServiceUInt64TypeMapTypedRequest":
		req := request.(ServiceUInt64TypeMapTypedRequest)
		retUInt64TypeMapTyped_0 := p.service.UInt64TypeMapTyped(req.V)
		response = ServiceUInt64TypeMapTypedResponse{RetUInt64TypeMapTyped_0: retUInt64TypeMapTyped_0}
	case "ServiceUIntMapRequest":
		req := request.(ServiceUIntMapRequest)
		retUIntMap_0 := p.service.UIntMap(req.V)
		response = ServiceUIntMapResponse{RetUIntMap_0: retUIntMap_0}
	case "ServiceUIntSliceRequest":
		req := request.(ServiceUIntSliceRequest)
		retUIntSlice_0 := p.service.UIntSlice(req.V)
		response = ServiceUIntSliceResponse{RetUIntSlice_0: retUIntSlice_0}
	case "ServiceUIntTypeRequest":
		req := request.(ServiceUIntTypeRequest)
		retUIntType_0 := p.service.UIntType(req.V)
		response = ServiceUIntTypeResponse{RetUIntType_0: retUIntType_0}
	case "ServiceUIntTypeMapRequest":
		req := request.(ServiceUIntTypeMapRequest)
		retUIntTypeMap_0 := p.service.UIntTypeMap(req.V)
		response = ServiceUIntTypeMapResponse{RetUIntTypeMap_0: retUIntTypeMap_0}
	case "ServiceUIntTypeMapTypedRequest":
		req := request.(ServiceUIntTypeMapTypedRequest)
		retUIntTypeMapTyped_0 := p.service.UIntTypeMapTyped(req.V)
		response = ServiceUIntTypeMapTypedResponse{RetUIntTypeMapTyped_0: retUIntTypeMapTyped_0}
	default:
		fmt.Println("Unkown request type", reflect.TypeOf(request).String())
	}

	if p.callStatsHandler != nil {
		p.callStatsHandler(&gotsrpc.CallStats{
			Func:      funcName,
			Package:   "github.com/foomo/gotsrpc/v2/example/basic/service",
			Service:   "Service",
			Execution: time.Since(start),
		})
	}

	return
}