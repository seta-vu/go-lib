package utilities

//Convert bool to string
func Btos(bVal bool) string {
	if bVal == true {
		return "true"
	} else if bVal == false {
		return "false"
	}
	return "true"
}