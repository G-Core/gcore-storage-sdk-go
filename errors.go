package gstorage

// EmptyResultErr customize
type EmptyResultErr string

// Error implemented
func (e EmptyResultErr) Error() string {
	return string(e)
}

// ResponseErr customize
type ResponseErr string

// Error implemented
func (e ResponseErr) Error() string {
	return string(e)
}
