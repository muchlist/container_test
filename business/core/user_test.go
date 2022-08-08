package user_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/muchlist/container_test/business/core/userrepo"
// 	"github.com/muchlist/container_test/business/tests"
// )

// var dbc = tests.DBContainer{
// 	Image: "postgres:13-alpine",
// 	Port:  "5432",
// 	Args:  []string{"-e", "POSTGRES_PASSWORD=postgres"},
// }

// func TestUser(t *testing.T) {
// 	log, db, teardown := tests.NewUnit(t, dbc)
// 	t.Cleanup(teardown)

// 	userStore := userrepo.NewStore(log, db)
// 	_, err := userStore.QueryByID(context.Background(), "1")
// 	if err != nil {
// 		t.Errorf("gagal")
// 	}

// 	t.Log("Given the need to work with User records.")
// 	// {
// 	// 	testID := 0
// 	// 	t.Logf("\tTest %d:\tWhen handling a single User.", testID)
// 	// 	{
// 	// 		// ctx := context.Background()
// 	// 		// now := time.Date(2018, time.October, 1, 0, 0, 0, 0, time.UTC)

// 	// 		// nu := userrepo.User{
// 	// 		// 	Name:  "Bill Kennedy",
// 	// 		// 	Email: "bill@ardanlabs.com",
// 	// 		// 	Roles: []string{auth.RoleAdmin},
// 	// 		// }
// 	// 	}
// 	// }
// }

// // func TestAuthenticate(t *testing.T) {
// // 	log, db, teardown := tests.NewUnit(t, dbc)
// // 	t.Cleanup(teardown)

// // 	store := userrepo.NewStore(log, db)

// // 	t.Log("Given the need to authenticate users")
// // 	{
// // 		testID := 0
// // 		t.Logf("\tTest %d:\tWhen handling a single User.", testID)
// // 		{
// // 			ctx := context.Background()

// // 			nu := userrepo.User{
// // 				Name:  "Anna Walker",
// // 				Email: "anna@ardanlabs.com",
// // 				Roles: []string{auth.RoleAdmin},
// // 			}

// // 			err := store.Create(ctx, nu)
// // 			if err != nil {
// // 				t.Fatalf("\t%s\tTest %d:\tShould be able to create user : %s.", tests.Failed, testID, err)
// // 			}
// // 			t.Logf("\t%s\tTest %d:\tShould be able to create user.", tests.Success, testID)

// // 		}
// // 	}
// // }
