package helper

const (
	BindModelError      = 20200
	NoneParamError      = 20201
	ParamParseError     = 20202
	LoginStatusSQLError = 20319
	LoginStatusError    = 20300
	LoginStatusOK       = 20301
	UserDoesNotExist    = 20302
	SaveStatusOK        = 20400
	SaveStatusError     = 20401
	SaveObjIsNil        = 20402
	DeleteStatusOK      = 20403
	DeleteStatusErr     = 20404
	DeleteObjIsNil      = 20405
	UpdateObjIsNil      = 20406
	ExistSameNameError  = 20501
	ExistSamePhoneError = 20502
	ExistSameEmailError = 20503
	MinThanMaxErr       = 20799
	MaxLessZeroErr      = 20798
	FixLessZeroErr      = 20797
	MediumPasswordErr   = 20801
	StrongPasswordErr   = 20802
	ChineseNameErr      = 20803
	EnglishNameErr      = 20804
)

var statusText = map[int]string{
	BindModelError:      "Model EnClosure Exception",
	NoneParamError:      "No Effective Parameter",
	ParamParseError:     "Invalid Argument",
	LoginStatusOK:       "Login Success",
	LoginStatusSQLError: "Login Error when update database",
	LoginStatusError:    "Invalid username or password",
	UserDoesNotExist:    "User does not exist",
	SaveObjIsNil:        "Object Saved is nil",
	SaveStatusOK:        "Save success",
	DeleteStatusOK:      "Delete success",
	DeleteStatusErr:     "Delete failed",
	UpdateObjIsNil:      "Record does not exist",
	ExistSameNameError:  "Duplicate name",
	ExistSamePhoneError: "Duplicate phone",
	ExistSameEmailError: "Duplicate email",
	DeleteObjIsNil:      "Object does not exist",
	MinThanMaxErr:       "rules error, max is less than min",
	MaxLessZeroErr:      "rules error, max number is less than 0",
	FixLessZeroErr:      "rules error, lenght is less than 0",
	MediumPasswordErr:   "密码为%d-%d位字母、数字，字母数字必须同时存在",
	StrongPasswordErr:   "密码为%d-%d位字母、数字和符号必须,同时存在，符号存在开头和结尾且仅限!@#$%^*",
	ChineseNameErr:      "中文名为%d-%d位中文字符可包含'·'",
	EnglishNameErr:      "英文名为%d-%d英文字符可包含空格",
}

// StatusText return status text
func StatusText(code int) string {
	return statusText[code]
}
