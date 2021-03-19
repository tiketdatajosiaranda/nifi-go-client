package nifi

import (
	"testing"
)

func TestFlow_ListTemplates(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	_, err = ctx.Flow.ListTemplates()
	if err != nil {
		t.Error(err)
	}
}

func TestFlow_GetProcessGroup(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	_, err = ctx.Flow.GetProcessGroup("root")
	if err != nil {
		t.Error(err)
	}
}

func TestFlow_ListRegistries(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	_, err = ctx.Flow.ListRegistries()
	if err != nil {
		t.Error(err)
	}
}

func TestFlow_ListBuckets(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	r, err := ctx.Flow.ListRegistries()
	if err != nil {
		t.Error(err)
	}

	_, err = ctx.Flow.ListBuckets(r.Registries[0].Registry.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestFlow_ListFlows(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	r, err := ctx.Flow.ListRegistries()
	if err != nil {
		t.Error(err)
	}

	b, err := ctx.Flow.ListBuckets(r.Registries[0].Registry.Id)
	if err != nil {
		t.Error(err)
	}

	_, err = ctx.Flow.ListFlows(r.Registries[0].Registry.Id, b.Buckets[0].Id)
	if err != nil {
		t.Error(err)
	}
}

func TestFlow_ListFlowVersions(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	r, err := ctx.Flow.ListRegistries()
	if err != nil {
		t.Error(err)
	}

	b, err := ctx.Flow.ListBuckets(r.Registries[0].Registry.Id)
	if err != nil {
		t.Error(err)
	}

	f, err := ctx.Flow.ListFlows(r.Registries[0].Registry.Id, b.Buckets[0].Id)
	if err != nil {
		t.Error(err)
	}

	_, err = ctx.Flow.ListFlowVersions(f.VersionedFlows[0].VersionedFlow.RegistryId, f.VersionedFlows[0].VersionedFlow.BucketId, f.VersionedFlows[0].VersionedFlow.FlowId)
	if err != nil {
		t.Error(err)
	}
}

func TestFlow_ListProcessGroupProcessors(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST JOSI PG: a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1
	_, err = ctx.Flow.ListProcessGroupProcessors("a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1")
	if err != nil {
		t.Error(err)
	}
}

func TestFlow_ListControllerServices(t *testing.T) {
	ctx, err := NewContext(nifiHost)
	if err != nil {
		t.Error(err)
	}

	err = ctx.Access.NewLogin(nifiUsername, nifiPassword)
	if err != nil {
		t.Error(err)
	}

	// TEST JOSI PG: a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1
	_, err = ctx.Flow.ListControllerServices("a4e53fa7-ace9-10d2-ffff-ffffe93bb3a1")
	if err != nil {
		t.Error(err)
	}

	//raw, _ := json.MarshalIndent(p, "", "  ")
	//fmt.Println(string(raw))
}
