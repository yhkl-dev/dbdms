package helper

const (
	BindModelError      = 20200
	NoneParamError      = 20201
	LoginStatusSQLError = 20319
	LoginStatusError    = 20300
	LoginStatusOK       = 20301
	SaveStatusOK        = 20400
	SaveStatusError     = 20401
	SaveObjIsNil        = 20402
	UpdateObjIsNil      = 20406
	ExistSameNameError  = 20501
	ExistSamePhoneError = 20502
	ExistSameEmailError = 20503
)

var statusText = map[int]string{
	BindModelError:      "Model EnClosure Exception",
	NoneParamError:      "No Effective Parameter",
	LoginStatusOK:       "Login Success",
	LoginStatusSQLError: "Login Error when update database",
	LoginStatusError:    "Invalied username or password",
	SaveObjIsNil:        "Object Saved is nil",
	UpdateObjIsNil:      "Record does not exist",
	ExistSameNameError:  "Duplicate name",
	ExistSamePhoneError: "Duplicate phone",
	ExistSameEmailError: "Duplicate email",
}

// StatusText return status text
func StatusText(code int) string {
	return statusText[code]
}
