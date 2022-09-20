package appliances

// Stove struct definition
type Stove struct {
	typeName string
}

// Implementing Start() function
func (fr *Stove) Start() {
	fr.typeName = " Stove "
}

// Implementing getPurpose() function
func (sv *Stove) GetPurpose() string {
	return "I'm a " + sv.typeName + " I cook food!!!"
}
