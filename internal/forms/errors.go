package forms

//map of errors slice of strings, indexed by string
type errors map[string][]string

//Adds an error message for a given form field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

//See whether a given field has an error
func (e errors) Get(field string) string {
	es := e[field] 
	if len(es) == 0 {
		return ""
	}

	return es[0]
}