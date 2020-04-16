package errors_test

import (
	"KubeDesc/src/errors"
	"reflect"

	"testing"
)

func TestHandleHTTPError(t *testing.T) {
	cases := []struct {
		err      error
		expected int
	}{
		{
			nil,
			500,
		},
		{
			errors.NewInvalid("some unknown error"),
			500,
		},
		{
			errors.NewInvalid(errors.MsgDeployNamespaceMismatchError),
			500,
		},
		{
			errors.NewInvalid(errors.MsgLoginUnauthorizedError),
			401,
		},
		{
			errors.NewInvalid(errors.MsgTokenExpiredError),
			401,
		},
		{
			errors.NewInvalid(errors.MsgEncryptionKeyChanged),
			401,
		},
	}
	for _, c := range cases {
		actual := errors.HandleHTTPError(c.err)
		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("HandleHTTPError(%+v) == %+v, expected %+v", c.err, actual, c.expected)
		}
	}
}
