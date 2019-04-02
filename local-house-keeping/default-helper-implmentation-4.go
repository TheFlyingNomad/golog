package housekeeping

import (
	"encoding/json"
	"fmt"

	gologC "github.com/brightappsllc/golog/contracts"
)

func (thisRef defaultHelperImplmentation) LogPanicWithFields(fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogPanicWithTagAndLevel("", 0, string(data))
}
func (thisRef defaultHelperImplmentation) LogFatalWithFields(fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogFatalWithTagAndLevel("", 0, string(data))
}
func (thisRef defaultHelperImplmentation) LogErrorWithFields(fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogErrorWithTagAndLevel("", 0, string(data))
}
func (thisRef defaultHelperImplmentation) LogWarningWithFields(fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogWarningWithTagAndLevel("", 0, string(data))
}
func (thisRef defaultHelperImplmentation) LogInfoWithFields(fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogInfoWithTagAndLevel("", 0, string(data))
}
func (thisRef defaultHelperImplmentation) LogDebugWithFields(fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogDebugWithTagAndLevel("", 0, string(data))
}
func (thisRef defaultHelperImplmentation) LogTraceWithFields(fields gologC.Fields) {
	var data, err = json.Marshal(fields)
	if err != nil {
		fmt.Println(fmt.Sprintf("%v", err))
	}

	thisRef.LogTraceWithTagAndLevel("", 0, string(data))
}
