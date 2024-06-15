package gerrc

import (
	"errors"

	errorsmod "cosmossdk.io/errors"

	grpccodes "google.golang.org/grpc/codes"
)

var DefaultCodespace = "gerrc"

type T struct {
	*errorsmod.Error
}

type Gerr struct{}

func (e Gerr) Error() string {
	return ""
}

func (e Gerr) GrpcCode() grpccodes.Code {
	return 42
}

// use this function only during a program startup phase.
func registerAndWrap(code uint32, err error) T {
	var gerr *Gerr
	if !errors.As(err, gerr) {
		panic("errs as gerr")
	}
	sdkerr := errorsmod.RegisterWithGRPCCode(DefaultCodespace, code, gerr.GrpcCode(), err.Error())
	return T{sdkerr}
}
