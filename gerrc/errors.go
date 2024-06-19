package gerrc

import (
	"errors"

	errorsmod "cosmossdk.io/errors"
	"github.com/danwt/gerr/gerr"
)

var DefaultCodespace = "gerrc"

var _ error = Error{}

// necessary to make Error() pass through to the underlying impl
type alias = errorsmod.Error

type Error struct {
	*alias
}

var (
	/*
		Google errors
	*/

	ErrCancelled          = registerAndWrap(0, gerr.ErrCancelled)
	ErrUnknown            = registerAndWrap(1, gerr.ErrUnknown)
	ErrInvalidArgument    = registerAndWrap(2, gerr.ErrInvalidArgument)
	ErrDeadlineExceeded   = registerAndWrap(3, gerr.ErrDeadlineExceeded)
	ErrNotFound           = registerAndWrap(4, gerr.ErrNotFound)
	ErrAlreadyExist       = registerAndWrap(5, gerr.ErrAlreadyExist)
	ErrPermissionDenied   = registerAndWrap(6, gerr.ErrPermissionDenied)
	ErrUnauthenticated    = registerAndWrap(7, gerr.ErrUnauthenticated)
	ErrResourceExhausted  = registerAndWrap(8, gerr.ErrResourceExhausted)
	ErrFailedPrecondition = registerAndWrap(9, gerr.ErrFailedPrecondition)
	ErrAborted            = registerAndWrap(10, gerr.ErrAborted)
	ErrOutOfRange         = registerAndWrap(11, gerr.ErrOutOfRange)
	ErrUnimplemented      = registerAndWrap(12, gerr.ErrUnimplemented)
	ErrInternal           = registerAndWrap(13, gerr.ErrInternal)
	ErrUnavailable        = registerAndWrap(14, gerr.ErrUnavailable)
	ErrDataLoss           = registerAndWrap(15, gerr.ErrDataLoss)

	/*
		Blockchain errors
	*/

	// ErrFault is a failed precondition error for malicious actor faults or frauds.
	ErrFault = errorsmod.Wrap(ErrFailedPrecondition, "fraud or fault")
)

// use this function only during a program startup phase.
func registerAndWrap(code uint32, err error) Error {
	var gErr gerr.Error
	if !errors.As(err, &gErr) {
		panic("errs as gerr")
	}
	sdkerr := errorsmod.RegisterWithGRPCCode(DefaultCodespace, code, gErr.GrpcCode(), err.Error())
	return Error{sdkerr}
}
