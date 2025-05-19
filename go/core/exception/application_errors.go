package exception

import "net/http"

type ApplicationErrors struct {
	Debug        bool
	MemberErrors *MemberErrors
}

func NewApplicationErrors() *ApplicationErrors {
	return &ApplicationErrors{
		Debug:        false,
		MemberErrors: NewMemberErrors(),
	}
}

// Application Errors Interface
type CommonApplicationErrors interface {
	ThrowUnauthorized() *ExceptionError
	ThrowPermissionDenied() *ExceptionError
	ThrowInvalidRequest() *ExceptionError
}

type MemberErrors struct {
	CommonApplicationErrors
	ErrUnauthorized             *ExceptionError
	ErrPermissionDenied         *ExceptionError
	ErrNotFound                 *ExceptionError
	ErrValidation               *ExceptionError
	ErrInvalidEmailFormat       *ExceptionError
	ErrEffectiveDate            *ExceptionError
	ErrLoginFailed              *ExceptionError
	ErrMemberStatusLock         *ExceptionError
	ErrOTPLimitExceed           *ExceptionError
	ErrInvalidOTP               *ExceptionError
	ErrOTPExpired               *ExceptionError
	ErrOrAccountSuspended       *ExceptionError
	ErrEmailAlreadyExist        *ExceptionError
	ErrTooManyFrequentRequests  *ExceptionError
	ErrUrlUpdatePasswordExpired *ExceptionError
	ErrGatewayTimeout           *ExceptionError
	ErrUnableToProceed          *ExceptionError
}

func NewMemberErrors() *MemberErrors {
	return &MemberErrors{
		ErrUnauthorized:             NewExceptionError(4000, 200000, "Unauthorized", http.StatusUnauthorized),
		ErrPermissionDenied:         NewExceptionError(4000, 200001, "Permission Denied (Forbidden error)", http.StatusForbidden),
		ErrNotFound:                 NewExceptionError(4000, 200002, "Not found", http.StatusNotFound),
		ErrValidation:               NewExceptionError(4000, 200003, "Invalid Request", http.StatusBadRequest),
		ErrInvalidEmailFormat:       NewExceptionError(4000, 200004, "Invalid Email Format", http.StatusBadRequest),
		ErrEffectiveDate:            NewExceptionError(4000, 200005, "Effective Date must be from the current date onwards", http.StatusBadRequest),
		ErrLoginFailed:              NewExceptionError(4000, 200006, "Login Failed", http.StatusBadRequest),
		ErrMemberStatusLock:         NewExceptionError(4000, 200007, "Member status Lock", http.StatusBadRequest),
		ErrOTPLimitExceed:           NewExceptionError(4000, 200008, "OTP limit exceed", http.StatusBadRequest),
		ErrInvalidOTP:               NewExceptionError(4000, 200009, "Invalid OTP", http.StatusBadRequest),
		ErrOTPExpired:               NewExceptionError(4000, 200010, "OTP expired", http.StatusBadRequest),
		ErrOrAccountSuspended:       NewExceptionError(4000, 200011, "OR Account suspended", http.StatusBadRequest),
		ErrEmailAlreadyExist:        NewExceptionError(4000, 200012, "Email Already Exist", http.StatusBadRequest),
		ErrTooManyFrequentRequests:  NewExceptionError(4000, 200013, "Requests are too many frequent", http.StatusBadRequest),
		ErrUrlUpdatePasswordExpired: NewExceptionError(4000, 200014, "Url update password expired", http.StatusBadRequest),
		ErrGatewayTimeout:           NewExceptionError(5000, 209998, "Gateway Timeout", http.StatusInternalServerError),
		ErrUnableToProceed:          NewExceptionError(5000, 209999, "Unable to proceed", http.StatusInternalServerError),
	}
}
