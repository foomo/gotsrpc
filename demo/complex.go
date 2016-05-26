package demo

type Address struct {
	City              string `json:"city,omitempty"`
	SecretServerCrap  bool   `json:"-"`
	PeoplePtr         []*Person
	ArrayOfMaps       []map[string]bool
	ArrayArrayAddress [][]*Address
	People            []Person
	MapCrap           map[string]map[int]bool
}

type Person struct {
	Name          string
	AddressPtr    *Address `json:"address"`
	AddressStruct Address
	Addresses     map[string]*Address

	InlinePtr *struct {
		Foo bool
	}
	iAmPrivate string
}

func (s *Service) ExtractAddress(person *Person) (*Address, *Err) {
	if person.AddressPtr != nil {
		return person.AddressPtr, nil
	}
	return nil, &Err{"there is no address on that person"}
}
