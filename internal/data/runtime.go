package data 
import (
	"fmt"
	"strconv"
	"errors"
	"strings"
)


var ErrInvalidRuntimeFormat = errors.New("Invalid runtime format")


// Declare a custom Runtime type, which has the underlying type int32 (the same as our Movie struct field).
type Runtime int32 

func (r Runtime) MarshalJSON() ([]byte, error){
	jsonValue := fmt.Sprintf("%d mins", r) 
	quotedJSONValue := strconv.Quote(jsonValue)

	return []byte(quotedJSONValue), nil
}


func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	parts := strings.Split(unquotedJSONValue, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidRuntimeFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat  
	} 

	*r = Runtime(i)
	return nil 
}