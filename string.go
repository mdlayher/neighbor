// generated by stringer -output=string.go -type=Family,Flags,State; DO NOT EDIT

package neighbor

import "fmt"

const (
	_Family_name_0 = "FamilyUnspec"
	_Family_name_1 = "FamilyIPv4"
	_Family_name_2 = "FamilyIPv6"
)

var (
	_Family_index_0 = [...]uint8{0, 12}
	_Family_index_1 = [...]uint8{0, 10}
	_Family_index_2 = [...]uint8{0, 10}
)

func (i Family) String() string {
	switch {
	case i == 0:
		return _Family_name_0
	case i == 2:
		return _Family_name_1
	case i == 10:
		return _Family_name_2
	default:
		return fmt.Sprintf("Family(%d)", i)
	}
}

const (
	_Flags_name_0 = "FlagsUseFlagsSelf"
	_Flags_name_1 = "FlagsMaster"
	_Flags_name_2 = "FlagsProxy"
	_Flags_name_3 = "FlagsRouter"
)

var (
	_Flags_index_0 = [...]uint8{0, 8, 17}
	_Flags_index_1 = [...]uint8{0, 11}
	_Flags_index_2 = [...]uint8{0, 10}
	_Flags_index_3 = [...]uint8{0, 11}
)

func (i Flags) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _Flags_name_0[_Flags_index_0[i]:_Flags_index_0[i+1]]
	case i == 4:
		return _Flags_name_1
	case i == 8:
		return _Flags_name_2
	case i == 128:
		return _Flags_name_3
	default:
		return fmt.Sprintf("Flags(%d)", i)
	}
}

const (
	_State_name_0 = "StateNoneStateIncompleteStateReachable"
	_State_name_1 = "StateStale"
	_State_name_2 = "StateDelay"
	_State_name_3 = "StateProbe"
	_State_name_4 = "StateFailed"
	_State_name_5 = "StateNoARP"
	_State_name_6 = "StatePermanent"
)

var (
	_State_index_0 = [...]uint8{0, 9, 24, 38}
	_State_index_1 = [...]uint8{0, 10}
	_State_index_2 = [...]uint8{0, 10}
	_State_index_3 = [...]uint8{0, 10}
	_State_index_4 = [...]uint8{0, 11}
	_State_index_5 = [...]uint8{0, 10}
	_State_index_6 = [...]uint8{0, 14}
)

func (i State) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _State_name_0[_State_index_0[i]:_State_index_0[i+1]]
	case i == 4:
		return _State_name_1
	case i == 8:
		return _State_name_2
	case i == 16:
		return _State_name_3
	case i == 32:
		return _State_name_4
	case i == 64:
		return _State_name_5
	case i == 128:
		return _State_name_6
	default:
		return fmt.Sprintf("State(%d)", i)
	}
}