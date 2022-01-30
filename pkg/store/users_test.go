package store

import (
	"context"
	"reflect"
	"testing"

	"github.com/lbrictson/squint/pkg/conf"
	"github.com/lbrictson/squint/pkg/db"
)

func TestUsers(t *testing.T) {
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
	adminUser, err := store.Users.Create(context.Background(), CreateUserInput{
		Email:          "adminuser@admin.com",
		HashedPassword: "hashedpassword",
		Role:           "admin",
	})
	if err != nil {
		t.Errorf("error creating admin user: %v", err)
		return
	}
	if adminUser.Email != "adminuser@admin.com" {
		t.Errorf("admin user created with email %v but got %v", "adminuser@admin.com", adminUser.Email)
		return
	}
	fetched, err := store.Users.GetByID(context.Background(), adminUser.ID)
	if err != nil {
		t.Errorf("error cgetting admin user by id: %v", err)
		return
	}
	if !reflect.DeepEqual(*fetched, *adminUser) {
		t.Errorf("getting admin user by id mismatch, got %v wanted %v", *fetched, *adminUser)
		return
	}
	limit := 1
	offset := 0
	orderBy := "created_at"
	roleEqual := "admin"
	listed, err := store.Users.List(context.Background(), UserListOptions{
		Limit:     &limit,
		OffSet:    &offset,
		OrderBy:   &orderBy,
		RoleEqual: &roleEqual,
	})
	if err != nil {
		t.Errorf("error listing users: %v", err)
		return
	}
	if len(listed) == 0 {
		t.Error("listing users returned 0 results")
		return
	}
	adminUser.Role = "user"
	adminUser.Email = "notadmin@admin.com"
	adminUser.APIKey = "123abc"
	updated, err := store.Users.Update(context.Background(), *adminUser)
	if err != nil {
		t.Errorf("error updating user: %v", err)
		return
	}
	if updated.Role != adminUser.Role {
		t.Errorf("updated role expected %v but got %v", updated.Role, adminUser.Role)
		return
	}
	if updated.Email != adminUser.Email {
		t.Errorf("updated email expected %v but got %v", updated.Email, adminUser.Email)
		return
	}
	if updated.APIKey != adminUser.APIKey {
		t.Errorf("updated APIKey expected %v but got %v", updated.Role, adminUser.Role)
		return
	}
	err = store.Users.Delete(context.Background(), adminUser.ID)
	if err != nil {
		t.Errorf("error deleting user: %v", err)
		return
	}
}
