package onepassword

// Credentials defines git credentials.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type response struct {
	Details struct {
		Fields []field `json:"fields"`
	} `json:"details"`
}

type login struct {
	Fields []field `json:"fields"`
	Notes  string  `json:"notesPlain"`
}

type field struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	Type        string `json:"type"`
	Designation string `json:"designation"`
}
