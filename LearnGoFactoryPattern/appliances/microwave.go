package appliances

// Microwave struct definition
type Microwave struct {
	typeName string
}

// Implementing Start() function
func (mw *Microwave) Start() {
	mw.typeName = " Microwave "
}

// Implementing getPurpose() function
func (sv *Microwave) GetPurpose() string {
	return "I'm a " + sv.typeName + " I heat stuff up!!!"
}
