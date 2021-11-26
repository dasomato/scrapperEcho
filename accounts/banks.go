package accounts

type Boolean bool

func NewBoolean(flag bool) Boolean {
	returnFlag := Boolean(flag)
	return returnFlag
}

func (b Boolean) SystemFlag() Boolean {
	return b
}
