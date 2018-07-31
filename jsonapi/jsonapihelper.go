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

func IsMainResult(s *DDStm, cur string) bool {

	/* rst := true*/
	//for i := 0; i < s.ddsk.Length(); i++ {
	//tmp := s.ddsk.ElemAtIndex(i).(*DDStm)
	//switch tmp.ct {
	//case RELATIONSHIPS:
	//rst = false
	//}
	//}

	//return rst
	return true
}

func IsRelationShips(s *DDStm) bool {

	rst := false
	for i := 0; i < s.ddsk.Length(); i++ {
		tmp := s.ddsk.ElemAtIndex(i).(*DDStm)
		switch tmp.ct {
		case RELATIONSHIPS:
			rst = true
		}
	}

	return rst
}

func IsIncluded(s *DDStm) bool {

	rst := false
	for i := 0; i < s.ddsk.Length(); i++ {
		tmp := s.ddsk.ElemAtIndex(i).(*DDStm)
		switch tmp.ct {
		case INCLUDED:
			rst = true
		}
	}

	return rst
}
