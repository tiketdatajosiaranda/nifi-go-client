package nifi

import (
	"fmt"
	"github.com/tiketdatarisal/nifi-go-client/models"
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

func TestProcessGroup_CreateDeleteProcessGroup(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST JOSI PG: a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1
	pg, err := ctx.ProcessGroup.CreateProcessGroup("a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1", &models.ProcessGroupEntity{
		Component: &models.ProcessGroupDto{
			Position: &models.PositionDto{X: new(float64), Y: new(float64)},
			Name: "Test Folder",
		},
		Revision: &models.RevisionDto{Version: new(int64)},
	})
	if err != nil {
		t.Error(err)
	}

	_, err = ctx.ProcessGroup.DeleteProcessGroup(pg.Component.Id, *pg.Revision.Version)
	if err != nil {
		t.Error(err)
	}

	//raw, _ := json.MarshalIndent(pg, "", "  ")
	//fmt.Println(string(raw))
}
