package nifi

import (
	"os"
	"testing"
)

var (
	nifiHost     string
	nifiUsername string
	nifiPassword string
)

func init() {
	nifiHost = os.Getenv("NIFI_HOST")
	nifiUsername = os.Getenv("NIFI_USERNAME")
	nifiPassword = os.Getenv("NIFI_PASSWORD")
}

func TestAccess_NewLogin(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	if ctx.token == "" {
		t.Error("token must not empty")
	}
}

func TestAccess_GetStatus(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	_, err = ctx.Access.GetStatus()
	if err == nil {
		t.Error("status should return error")
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	res, err := ctx.Access.GetStatus()
	if err != nil {
		t.Error(err)
	} else if res == nil {
		t.Errorf("status result should not nil")
	}
}
