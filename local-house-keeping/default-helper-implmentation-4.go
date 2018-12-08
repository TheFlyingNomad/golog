package housekeeping

import (
	"encoding/json"
	"fmt"

	gologC "github.com/TheFlyingNomad/golog/contracts"
)

func (thisRef defaultHelperImplmentation) LogPanicWithFields(message string, fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogPanicWithTagAndLevel("", 0, fmt.Sprintf("%s -> %s", message, string(data)))
}
func (thisRef defaultHelperImplmentation) LogFatalWithFields(message string, fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogFatalWithTagAndLevel("", 0, fmt.Sprintf("%s -> %s", message, string(data)))
}
func (thisRef defaultHelperImplmentation) LogErrorWithFields(message string, fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogErrorWithTagAndLevel("", 0, fmt.Sprintf("%s -> %s", message, string(data)))
}
func (thisRef defaultHelperImplmentation) LogWarningWithFields(message string, fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogWarningWithTagAndLevel("", 0, fmt.Sprintf("%s -> %s", message, string(data)))
}
func (thisRef defaultHelperImplmentation) LogInfoWithFields(message string, fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogInfoWithTagAndLevel("", 0, fmt.Sprintf("%s -> %s", message, string(data)))

}
func (thisRef defaultHelperImplmentation) LogDebugWithFields(message string, fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogDebugWithTagAndLevel("", 0, fmt.Sprintf("%s -> %s", message, string(data)))
}
