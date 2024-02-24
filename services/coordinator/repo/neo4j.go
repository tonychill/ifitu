package repo

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type graphClient struct {
}

func (r *repoImpl) getGraphDbAuth() neo4j.AuthToken {
	return neo4j.BasicAuth("neo4j", "letmein!", "")

}

// func (r *repoImpl) createGuestInGraph(ctx context.Context, guest *global.Guest) error {
// 	driver, err := neo4j.NewDriverWithContext(r.GraphDbAddress, r.getGraphDbAuth())
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Starting with 5.0, you can control the execution of most driver APIs
// 	// To keep things simple, we create here a never-cancelling context
// 	// Read https://pkg.go.dev/context to learn more about contexts

// 	// Handle driver lifetime based on your application lifetime requirements.
// 	// driver's lifetime is usually bound by the application lifetime, which usually implies one driver instance per
// 	// application

// 	defer driver.Close(ctx) // Make sure to handle errors during deferred calls
// 	gmap := utils.ConvertStructToMap(guest, true)
// 	fmt.Printf("***** TESTING: converted guest map: %+v ", gmap)
// 	result, err := neo4j.ExecuteQuery(ctx, driver,
// 		`MERGE (g:Guest {
// 			id: $id,
// 			first_name: $first_name,
// 			last_name: $last_name
// 			}
// 		)
// 		RETURN g`,
// 		gmap, neo4j.EagerResultTransformer)
// 	if err != nil {
// 		return err
// 	}

// 	guestNode, _, err := neo4j.GetRecordValue[neo4j.Node](result.Records[0], "n")
// 	if err != nil {
// 		return fmt.Errorf("could not find node n")
// 	}

// 	_, err = neo4j.GetProperty[int64](guestNode, "id")
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *repoImpl) getGuest(ctx context.Context, guest any) error {}
