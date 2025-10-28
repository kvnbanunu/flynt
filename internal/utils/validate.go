package utils

// Writes error and returns false if any fields are empty
func ValidateFields(models ...any) bool {
	count := 0
	for _, v := range models {
		switch v := v.(type) {
		case int:
			if v < 0 {
				break
			}
		case string:
			if v == "" {
				break
			}
		default:
			if v == nil {
				break
			}
		}
		count++
	}
	if count < len(models) {
		return false
	}
	return true
}
