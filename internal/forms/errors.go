package forms

type errors map[string][]string

// Adds Error Message For Given Form Field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Gets The First Error Message
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
