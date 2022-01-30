package store

import (
	"context"
	"testing"

	"github.com/lbrictson/squint/pkg/conf"
	"github.com/lbrictson/squint/pkg/db"
)

func TestGroupsStore(t *testing.T) {
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
	g, err := store.Groups.Create(ctx, CreateGroupInput{
		Name:        "testgroup1",
		Description: "testing group!  Anything can go here",
		Expanded:    true,
	})
	if err != nil {
		t.Errorf("error creating group: %v", err)
		return
	}
	if g.Name != "testgroup1" {
		t.Errorf("expected group name %v got %v", "testgroup1", g.Name)
		return
	}
	if g.Description != "testing group!  Anything can go here" {
		t.Errorf("expected group Description %v got %v", "testing group!  Anything can go here", g.Description)
		return
	}
	if g.Expanded != true {
		t.Error("expected expanded to be true but got false")
		return
	}
	g.Description = "updated"
	g.Name = "updated-name"
	g.Expanded = false
	updated, err := store.Groups.Update(ctx, *g)
	if err != nil {
		t.Errorf("error updating group: %v", err)
		return
	}
	if updated.Name != "updated-name" {
		t.Errorf("expected group name %v got %v", "updated-name", updated.Name)
		return
	}
	if updated.Description != "updated" {
		t.Errorf("expected group Description %v got %v", "updated", updated.Description)
		return
	}
	if updated.Expanded != false {
		t.Error("expected expanded to be false but got true")
		return
	}
	fetched, err := store.Groups.GetByID(ctx, updated.ID)
	if err != nil {
		t.Errorf("error getting group by id: %v", err)
		return
	}
	if fetched.Name != updated.Name {
		t.Errorf("got %v but expected name to be %v", fetched.Name, updated.Name)
		return
	}
	limit := 5
	offset := 0
	order := "created_at"
	listed, err := store.Groups.List(ctx, GroupListOptions{
		OffSet:  &offset,
		Limit:   &limit,
		OrderBy: &order,
	})
	if err != nil {
		t.Errorf("error listing groups: %v", err)
		return
	}
	if len(listed) == 0 {
		t.Error("got 0 groups in list group operation, at least 1 is expected")
		return
	}
	err = store.Groups.Delete(ctx, updated.ID)
	if err != nil {
		t.Errorf("error deleting group: %v", err)
		return
	}
	return
}
