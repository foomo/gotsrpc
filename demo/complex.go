package demo

import nstd "github.com/foomo/gotsrpc/demo/nested"

type Address struct {
	City              string   `json:"city,omitempty"`
	Signs             []string `json:"signs,omitempty"`
	SecretServerCrap  bool     `json:"-"`
	PeoplePtr         []*Person
	ArrayOfMaps       []map[string]bool
	ArrayArrayAddress [][]*Address
	People            []Person
	MapCrap           map[string]map[int]bool
	NestedPtr         *nstd.Nested
	NestedStruct      nstd.Nested
}

type Person struct {
	Name          string
	AddressPtr    *Address `json:"address"`
	AddressStruct Address
	Addresses     map[string]*Address

	InlinePtr *struct {
		Foo bool
	}
	InlineStruct struct {
		Bar string
	}
	iAmPrivate string
}

func (s *Service) ExtractAddress(person *Person) (addr *Address, e *Err) {
	if person.AddressPtr != nil {
		return person.AddressPtr, nil
	}
	return nil, &Err{"there is no address on that person"}
}