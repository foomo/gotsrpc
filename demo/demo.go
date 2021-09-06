package demo

import "github.com/foomo/gotsrpc/v2/demo/nested"

type Err struct {
	Message string `json:"message"`
}

type LocalKey string

type MapWithLocalStuff map[LocalKey]int

type MapOfOtherStuff map[nested.JustAnotherStingType]int

func (e *Err) Error() string {
	return e.Message
}

type ScalarError string

func (se *ScalarError) Error() string {
	return string(*se)
}

type ScalarInPlace string

type Demo struct {
	Bla bool
}

func (d *Demo) Hello(name string) (string, *Err) {
	if name == "Peter" {
		return "", &Err{"fuck you Peter I do not like you"}
	}
	return "Hello from the server: " + name, nil
}

func (d *Demo) HelloInterface(anything interface{}, anythingMap map[string]interface{}, anythingSlice []interface{}) {

}

func (d *Demo) HelloNumberMaps(intMap map[int]string) (floatMap map[float64]string) {
	floatMap = map[float64]string{}
	for i, str := range intMap {
		floatMap[float64(i)] = str
	}
	return
}

func (d *Demo) HelloScalarError() (err *ScalarError) {
	return
}

func (d *Demo) HelloMapType() (otherStuff MapOfOtherStuff) {
	return MapOfOtherStuff{
		nested.JustAnotherStingType("foo"): 1,
	}
}

func (d *Demo) HelloLocalMapType() (localStuff MapWithLocalStuff) {
	return MapWithLocalStuff{
		"foo": 1,
	}
}

type RemoteScalarsStrings []nested.JustAnotherStingType

type RemoteScalarStruct struct {
	Foo RemoteScalarsStrings
	Bar RemoteScalarsStrings
}

func (d *Demo) ArrayOfRemoteScalars() (arrayOfRemoteScalars RemoteScalarsStrings) {
	return []nested.JustAnotherStingType{"foo", "bar"}
}

func (d *Demo) ArrayOfRemoteScalarsInAStruct() (strct RemoteScalarStruct) {
	return RemoteScalarStruct{
		Foo: RemoteScalarsStrings{"Foo"},
		Bar: RemoteScalarsStrings{"Bar"},
	}
}

func (d *Demo) nothingInNothinOut() {

}
