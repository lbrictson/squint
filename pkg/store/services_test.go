package store

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/lbrictson/squint/pkg/conf"
	"github.com/lbrictson/squint/pkg/db"
)

func TestServicesStore(t *testing.T) {
	pool, err := db.New(conf.Config{
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "squinttest",
		DBHost:     "localhost",
		DBPort:     5432,
	})
	if err != nil {
		t.Errorf("error setting up test DB: %v", err)
		return
	}
	store, err := New(pool)
	if err != nil {
		t.Errorf("error setting up test store: %v", err)
		return
	}
	ctx := context.Background()
	fakeGroupUUID := uuid.NewString()
	serviceOne, err := store.Services.Create(ctx, CreateServiceInput{
		Name:        "serviceOne",
		Description: "service one description",
		Group:       "aaaabbbbcccceeeedddd",
	})
	if err != nil {
		t.Errorf("error creating test service %v", err)
		return
	}
	serviceTwo, err := store.Services.Create(ctx, CreateServiceInput{
		Name:        "serviceTwo",
		Description: "service two description",
		Group:       "aaaabbbbcccceeeedddd",
	})
	if err != nil {
		t.Errorf("error creating test service %v", err)
		return
	}
	serviceThree, err := store.Services.Create(ctx, CreateServiceInput{
		Name:        "serviceThree",
		Description: "service three description",
		Group:       fakeGroupUUID,
	})
	if err != nil {
		t.Errorf("error creating test service %v", err)
		return
	}
	if serviceOne.Name != "serviceOne" {
		t.Errorf("got name %v instead of name %v", serviceOne.Name, "serviceOne")
		return
	}
	if serviceTwo.Name != "serviceTwo" {
		t.Errorf("got name %v instead of name %v", serviceOne.Name, "serviceTwo")
		return
	}
	if serviceThree.Name != "serviceThree" {
		t.Errorf("got name %v instead of name %v", serviceOne.Name, "serviceThree")
		return
	}
	// add a page
	serviceOne.Pages = []string{"aaa", "bbb"}
	serviceTwo.Pages = []string{"aaa"}
	serviceTwo.Status = "Major Outage"
	serviceTwo.Group = fakeGroupUUID
	serviceOne, err = store.Services.Update(ctx, *serviceOne)

	if err != nil {
		t.Errorf("error updating service %v", err)
		return
	}
	serviceTwo, err = store.Services.Update(ctx, *serviceTwo)
	if err != nil {
		t.Errorf("error updating service %v", err)
		return
	}
	limit := 5
	offset := 0
	order := "name"
	page := "aaa"
	pageServices, err := store.Services.List(ctx, ServiceListOptions{
		Limit:       &limit,
		OffSet:      &offset,
		OrderBy:     &order,
		StatusEqual: nil,
		GroupEqual:  nil,
		PageEqual:   &page,
	})
	if err != nil {
		t.Errorf("error getting services list %v", err)
		return
	}
	if len(pageServices) != 2 {
		t.Errorf("expected 2 services but got %v", len(pageServices))
		return
	}
	groupServices, err := store.Services.List(ctx, ServiceListOptions{
		Limit:       nil,
		OffSet:      nil,
		OrderBy:     nil,
		StatusEqual: nil,
		GroupEqual:  &fakeGroupUUID,
		PageEqual:   nil,
	})
	if err != nil {
		t.Errorf("error getting services list %v", err)
		return
	}
	if len(groupServices) == 0 {
		t.Error("expected at least 1 service but got 0")
		return
	}
	err = store.Services.Delete(ctx, serviceOne.ID)
	if err != nil {
		t.Errorf("error deleting service %v", err)
		return
	}
	return
}
