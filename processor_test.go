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

	// GenerateFlowFile: c92537a3-55a6-3666-a07b-a1dce921d986
	_, err = ctx.Processor.GetProcessor("c92537a3-55a6-3666-a07b-a1dce921d986")
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
}

