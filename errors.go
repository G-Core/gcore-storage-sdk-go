package gstorage

type EmptyResultErr string

func (e EmptyResultErr) Error() string {
	return string(e)
}
