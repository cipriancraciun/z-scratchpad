

package zscratchpad


import "fmt"
import "log"
import "os"




type Error struct {
	Code uint32
	Message string
	Error error
}




func logf (_slug rune, _code uint32, _format string, _arguments ... interface{}) () {
	_pid := os.Getpid ()
	_message := fmt.Sprintf (_format, _arguments ...)
	switch _slug {
		case 's' :
			log.Printf ("[z-scratchpad]  %s\n", _message)
		case '{' :
			log.Printf ("[z-scratchpad:%08d] [%c%c] [%08x]\n\n", _pid, _slug, _slug, _code)
		case '}' :
			log.Printf ("\n[z-scratchpad:%08d] [%c%c] [%08x]\n", _pid, _slug, _slug, _code)
		case '`' :
			_lines, _ := stringSplitLines (_message)
			for _, _line := range _lines {
				log.Printf ("[z-scratchpad:%08d] [%c%c] [%08x]  %s\n", _pid, _slug, _slug, _code, _line)
			}
		default :
			log.Printf ("[z-scratchpad:%08d] [%c%c] [%08x]  %s\n", _pid, _slug, _slug, _code, _message)
	}
}

func logError (_slug rune, _error *Error) () {
	logErrorf (_slug, 0x55d59c80, _error, "unexpected error encountered!")
}

func logErrorf (_slug rune, _code uint32, _error *Error, _format string, _arguments ... interface{}) () {
	_pid := os.Getpid ()
	if (_format != "") || (len (_arguments) != 0) {
		logf (_slug, _code, _format, _arguments ...)
	}
	if _error != nil {
		if _error.Message != "" {
			log.Printf ("[z-scratchpad:%08d] [%c%c] [%08x]  %s\n", _pid, _slug, _slug, _error.Code, _error.Message)
		} else {
			log.Printf ("[z-scratchpad:%08d] [%c%c] [%08x]  %s\n", _pid, _slug, _slug, _error.Code, "unexpected error encountered!")
		}
		if _error.Error != nil {
			_errorMessage := _error.Error.Error ()
			if _errorMessage != "" {
				log.Printf ("[z-scratchpad:%08d] [%c%c] [%08x]  >> %s\n", _pid, _slug, _slug, _error.Code, _errorMessage)
			}
//			log.Printf ("[z-scratchpad:%08d] [%c%c] [%08x]  >> %#v\n", _pid, _slug, _slug, _error.Code, _error.Error)
		}
	}
}


func abortError (_error *Error) (*Error) {
	return abortErrorf (_error, _error.Code, "")
}

func abortErrorf (_error *Error, _code uint32, _format string, _arguments ... interface{}) (*Error) {
	logErrorf ('!', _code, _error, _format, _arguments ...)
	logf ('!', 0xb7a5fb86, "aborting!")
	os.Exit (1)
	panic (0xa235deea)
}

func abortErrorw (_code uint32, _error error) (*Error) {
	return abortError (errorw (_code, _error))
}


func errorf (_code uint32, _format string, _arguments ... interface{}) (*Error) {
	_message := fmt.Sprintf (_format, _arguments ...)
	_error_0 := & Error {
			Code : _code,
			Message : _message,
			Error : nil,
		}
	return returnError (_error_0)
}

func errorw (_code uint32, _error error) (*Error) {
	if _code == 0 {
		panic (0xa4ddfd33)
	}
	_error_0 := & Error {
			Code : _code,
			Message : "",
			Error : _error,
		}
	return returnError (_error_0)
}

func returnError (_error *Error) (*Error) {
	if _error == nil {
		return nil
	} else {
		if true {
			return _error
		} else {
			panic (_error.ToError ())
		}
	}
}


func (_error *Error) ToError () (error) {
	var _message = _error.Message
	if _message == "" {
		_message = "unexpected error encountered"
	}
	if _error.Error != nil {
		return fmt.Errorf ("[%08x]  %s  //  %w", _error.Code, _message, _error.Error)
	} else {
		return fmt.Errorf ("[%08x]  %s", _error.Code, _message)
	}
}

