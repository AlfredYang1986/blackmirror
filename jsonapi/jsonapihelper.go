package jsonapi

func IsLeftObjDelim(st string, sv string) bool {
	return st == "json.Delim" && sv == "{"
}

func IsLeftArrayDelim(st string, sv string) bool {
	return st == "json.Delim" && sv == "["
}

func IsRightObjDelim(st string, sv string) bool {
	return st == "json.Delim" && sv == "}"
}

func IsRightArrayDelim(st string, sv string) bool {
	return st == "json.Delim" && sv == "]"
}

func IsMainResult(s *DDStm) bool {
	return false
}

func IsRelationShips(s *DDStm) bool {
	return false
}

func IsIncluded(s *DDStm) bool {
	return false
}
