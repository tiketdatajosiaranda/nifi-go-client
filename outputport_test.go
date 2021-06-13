package nifi

import (
	"fmt"
	"github.com/tiketdatarisal/nifi-go-client/models"
	"testing"
)

func TestOutputPort(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST Output Port : 6aeb3a9b-ccd6-128c-9404-981b2488269a

	o, err := ctx.OutputPort.GetOutputPort("6aeb3a9b-ccd6-128c-9404-981b2488269a")
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}

	if o.Component.State == "RUNNING" {
		o, err = ctx.OutputPort.PutOutputPortRunStatus(o.Id,&models.PortRunStatusEntity{
			Revision:                     &models.RevisionDto{
				Version: o.Revision.Version,
			},
			State:                        "STOPPED",
			DisconnectedNodeAcknowledged: false,
		})
		if err != nil {
			fmt.Printf("%#v", err)
			t.Error(err)
		}
	}

	o, err = ctx.OutputPort.PutOutputPort(o.Id, &models.PortEntity{
		Revision:  &models.RevisionDto{
			Version: o.Revision.Version,
		},
		Component: &models.PortDto{
			Id:    o.Id,
			Name:  o.Component.Name,
			State: "STOPPED",
		},
	})
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}

	o, err = ctx.OutputPort.PutOutputPortRunStatus(o.Id,&models.PortRunStatusEntity{
		Revision:                     &models.RevisionDto{
			Version: o.Revision.Version,
		},
		State:                        "RUNNING",
		DisconnectedNodeAcknowledged: false,
	})
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}
	o, err = ctx.OutputPort.PutOutputPortRunStatus(o.Id,&models.PortRunStatusEntity{
		Revision:                     &models.RevisionDto{
			Version: o.Revision.Version,
		},
		State:                        "STOPPED",
		DisconnectedNodeAcknowledged: false,
	})
	if err != nil {
		fmt.Printf("%#v", err)
		t.Error(err)
	}



}
