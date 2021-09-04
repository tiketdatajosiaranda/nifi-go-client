package nifi

import (
	"fmt"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"testing"
)

func TestConnection(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST JOSI PG: a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1
	rootPgId := "a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1"
	c,err := ctx.ProcessGroup.CreateConnection(rootPgId,&models.ConnectionEntity{
		Revision:                     &models.RevisionDto{
			Version:      new(int64),
		},
		DisconnectedNodeAcknowledged: false,
		Component:                    &models.ConnectionDto{
			Destination: &models.ConnectableDto{
				Id:                   "b4083e21-5a78-11a5-841c-290f0518d056",
				Type:                 "PROCESSOR",
				GroupId:              rootPgId,
			},
			Source: &models.ConnectableDto{
				Id:                   "38f738bb-5301-12c3-0000-0000025bcb9a",
				Type:                 "PROCESSOR",
				GroupId:              rootPgId,
			},
			SelectedRelationships: []string{"success"},
		},
	})

	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}

	c, err = ctx.Connection.GetConnection(c.Id)
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}

	c, err = ctx.Connection.PutConnection(c.Id,c)
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
	c, err = ctx.Connection.DeleteConnection(c.Id,*c.Revision.Version)
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
}
