package validation

import "fmt"

func GetFieldErrorFromTag(tag string, field string) string {
	messages := map[string]string {
		"required": fmt.Sprintf("%s is %s", field, tag),
		"min": fmt.Sprintf("%s is too short", field),
		"max": fmt.Sprintf("%s is too big", field),
		"email": fmt.Sprintf("%s is invalid", field),
	}

	message, ok := messages[tag]

	if !ok {
		return fmt.Sprintf("there is no mapped message for %s constraint", tag)
	}

	return message
}
