package demo

type Address struct {
	City             string `json:"city,omitempty"`
	SecretServerCrap bool   `json:"-"`
}

type Person struct {
	Name    string
	Address *Address `json:"address"`
}

func (s *Service) ExtractAddress(person *Person) (*Address, *Err) {
	if person.Address != nil {
		return person.Address, nil
	}
	return nil, &Err{"there is no address on that person"}
}
