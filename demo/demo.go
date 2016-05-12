package demo

type Err struct {
	Message string `json:"message"`
}

type Service struct {
}

func (s *Service) Hello(name string) (reply string, err *Err) {
	if name == "Peter" {
		return "", &Err{"fuck you Peter I do not like you"}
	}
	return "Hello from the server: " + name, nil
}

type Address struct {
	City string `json:"city"`
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
