package gstorage

// EmptyResultErr customize
type EmptyResultErr string

// Error implemented
func (e EmptyResultErr) Error() string {
	return string(e)
}
