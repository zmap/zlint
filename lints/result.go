package lints

type ResultEnum int

const (
	NA    ResultEnum = iota // 0
	NE                      // 1
	Pass                    // 2
	Info                    // 3
	Warn                    // 4
	Error                   // 5
	Fatal                   // 6
)

type ResultStruct struct {
	Result ResultEnum `json:"result"` //this is the ResultEnum enumeration and uses the values found there
	//Details string     `json:"details,omitempty"`
}

type FinalResult struct {
	Result string `json:"result"`
	//Details string `json:"details,omitempty"`
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
