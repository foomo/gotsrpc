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
