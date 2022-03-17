package ptrhelpers

func String(s string) *string {
	return &s
}

func StringValue(s *string) string {
	return StringValueDefault(s, "")
}

func StringValueDefault(s *string, defaultValue string) string {
	if s == nil {
		return defaultValue
	}

	return *s
}

func Bool(b bool) *bool {
	return &b
}

func BoolValue(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}

func Int(i int) *int {
	return &i
}

func IntValue(i *int) int {
	return IntValueDefault(i, 0)
}

func IntValueDefault(i *int, defaultValue int) int {
	if i == nil {
		return defaultValue
	}

	return *i
}
