package errors

import (
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"net/http"
)

var _ error = &errors.StatusError{}

func NewUnauthorized(reason string) *errors.StatusError {
	return errors.NewUnauthorized(reason)
}

func NewTokenExpired(reason string) *errors.StatusError {
	return &errors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusUnauthorized,
			Reason:  metav1.StatusReasonExpired,
			Message: reason,
		},
	}
}

func NewBadRequest(reason string) *errors.StatusError {
	return errors.NewBadRequest(reason)
}

func NewInvalid(reason string) *errors.StatusError {
	return &errors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusInternalServerError,
			Reason:  metav1.StatusReasonInvalid,
			Message: reason,
		},
	}
}

func NewNotFound(reason string) *errors.StatusError {
	return &errors.StatusError{
		ErrStatus: metav1.Status{
			Status:  metav1.StatusFailure,
			Code:    http.StatusNotFound,
			Reason:  metav1.StatusReasonNotFound,
			Message: reason,
		},
	}
}
