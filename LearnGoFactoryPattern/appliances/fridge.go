package appliances

// Fridge struct definition
type Fridge struct {
	typeName string
}

// Implementing Start() function
func (fr *Fridge) Start() {
	fr.typeName = " Fridge "
}

// Implementing getPurpose() function
func (fr *Fridge) GetPurpose() string {
	return "I'm a " + fr.typeName + " I cool stuff down!!!"
}
