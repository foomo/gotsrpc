package demo

import (
	"fmt"
	"math/rand"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
	"time"

	"github.com/foomo/gotsrpc/v2/demo/nested"
	"github.com/stretchr/testify/assert"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

var (
	client DemoGoTSRPCClient
	server *httptest.Server
)

func setup() {
	server = httptest.NewServer(NewDefaultDemoGoTSRPCProxy(&Demo{}))
	serverUrl, _ := url.Parse(server.URL)
	client = NewDefaultDemoGoTSRPCClient(serverUrl.String())
}

func teardown() {
	server.Close()
}

func TestDefault(t *testing.T) {
	setup()
	defer teardown()

	resp, errServer, errClient := client.Hello("stefan")
	assert.NoError(t, errClient)
	assert.Nil(t, errServer)
	fmt.Println(resp)
}

func TestHelloNumberMaps(t *testing.T) {
	setup()
	defer teardown()
	intMap := map[int]string{1: "one", 2: "two", 3: "three"}
	floatMap, errClient := client.HelloNumberMaps(intMap)
	assert.NoError(t, errClient)
	for f, fstr := range floatMap {
		i := int(f)
		assert.Equal(t, fstr, intMap[i])
	}
}

func benchmarkRequests(b *testing.B, count int) {
	setup()
	defer teardown()

	person := GeneratePerson(count)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		client.ExtractAddress(person)
	}
}

func BenchmarkRequest1(b *testing.B)      { benchmarkRequests(b, 1) }
func BenchmarkRequest10(b *testing.B)     { benchmarkRequests(b, 10) }
func BenchmarkRequest100(b *testing.B)    { benchmarkRequests(b, 100) }
func BenchmarkRequest1000(b *testing.B)   { benchmarkRequests(b, 1000) }
func BenchmarkRequest10000(b *testing.B)  { benchmarkRequests(b, 10000) }
func BenchmarkRequest100000(b *testing.B) { benchmarkRequests(b, 100000) }

func GeneratePerson(count int) *Person {
	person := &Person{}
	person.AddressPtr = GenerateAddress()
	person.Addresses = map[string]*Address{}
	for i := 0; i < count; i++ {
		person.Addresses[strconv.Itoa(i)] = GenerateAddress()
	}
	return person
}

func GenerateAddress() *Address {
	gen := func() string {
		return RandStringRunes(32)
	}
	genarr := func(count int) (ret []string) {
		ret = make([]string, count)
		for i := 0; i < count; i++ {
			ret[i] = gen()
		}
		return
	}
	return &Address{
		City:              gen(),
		Signs:             genarr(100),
		SecretServerCrap:  false,
		PeoplePtr:         nil,
		ArrayOfMaps:       nil,
		ArrayArrayAddress: nil,
		People:            nil,
		MapCrap:           nil,
		NestedPtr: &nested.Nested{
			Name: gen(),
			SuperNestedString: struct {
				Ha int64
			}{
				Ha: 011,
			},
			SuperNestedPtr: &struct {
				Bla string
			}{
				Bla: gen(),
			},
		},
		NestedStruct: nested.Nested{
			Name: gen(),
			SuperNestedString: struct {
				Ha int64
			}{
				Ha: 0,
			},
			SuperNestedPtr: &struct {
				Bla string
			}{
				Bla: gen(),
			},
		},
	}
}
