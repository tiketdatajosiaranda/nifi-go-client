package nifi

import (
	"fmt"
	"testing"
)

func TestProcessor_GetProcessor(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// GenerateFlowFile: 38f738bb-5301-12c3-0000-0000025bcb9a
	_, err = ctx.Processor.GetProcessor("38f738bb-5301-12c3-0000-0000025bcb9a")
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
}
