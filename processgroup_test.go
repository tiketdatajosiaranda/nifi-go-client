package nifi

import (
	"fmt"
	"testing"
)

func TestProcessGroup_ListProcessGroups(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST JOSI PG: a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1
	_, err = ctx.ProcessGroup.ListProcessGroups("a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1")
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
}

func TestProcessGroup_GetProcessGroup(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST JOSI PG: a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1
	_, err = ctx.ProcessGroup.GetProcessGroup("a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1")
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
}

func TestProcessGroup_ListProcessors(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST JOSI PG: a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1
	_, err = ctx.ProcessGroup.ListProcessors("a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1")
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
}

func TestProcessGroup_GetVariableRegistry(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST JOSI PG: a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1
	_, err = ctx.ProcessGroup.GetVariableRegistry("a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1")
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
}
