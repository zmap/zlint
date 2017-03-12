package lints

type ResultEnum int

const (
	Reserved ResultEnum = iota //0
	NA    			// 1
	NE                      // 2
	Pass                    // 3
	Info                    // 4
	Warn                    // 5
	Error                   // 6
	Fatal                   // 7
)

type ResultStruct struct {
	Result ResultEnum `json:"result"` //this is the ResultEnum enumeration and uses the values found there
	//Details string     `json:"details,omitempty"`
}

type FinalResult struct {
	Result 	ResultEnum 	`json:"result"`
	Details string 		`json:"details,omitempty"`
}

func EnumToString(e ResultEnum) string {
	switch e {
	case NA:
		return "NA"
	case NE:
		return "NE"
	case Pass:
		return "pass"
	case Info:
		return "info"
	case Warn:
		return "warn"
	case Error:
		return "error"
	case Fatal:
		return "fatal"
	default:
		return ""
	}
}
