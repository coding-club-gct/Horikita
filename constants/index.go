package constants

type Constants struct {
	Strings map[string] string
}

var C Constants

func (*Constants) GetDeptRole(dept string) string {
	if role, ok := deptKeys[dept]; ok {
		return role
	} else {
		return ""
	}
}

func (*Constants) GetBatchRole(batch int) string {
	if role, ok := batchKeys[batch]; ok {
		return role
	} else {
		return ""
	}
}

func (*Constants) GetGenderRole(gender string) string {
	if role, ok := genderKeys[gender]; ok {
		return role
	} else {
		return ""
	}
}

func (*Constants) GetVerifiedRole() string {
	return "1088406695404896306"
}
