package gstorage

import "fmt"

//SwaggerResponseError customizing
type SwaggerResponseError string

//Error implemented
func (r SwaggerResponseError) Error() string {
	return fmt.Sprintf("response reading error: %s", r)
}
