package demo

type Err struct {
	Message string `json:"message"`
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

func (d *Demo) nothingInNothinOut() {

}
