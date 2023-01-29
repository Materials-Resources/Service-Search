package index

import (
	"github.com/materials-resources/Service-Search/core/external-services/prophet"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"log"
)

func Product(client *typesense.Client) error {

	nextInvMastUid := new(int32)

	for true {
		requestedProductGroups := []string{
			"IS05",
			"HD70",
			"HD30",
			"IS55",
			"IS35",
		}

		productsByGroup := prophet.GetProductsByGroup(requestedProductGroups, nextInvMastUid)

		*nextInvMastUid = productsByGroup.Cursor.NextInvMastUid

		doc := convertProphetProductsToDoc(productsByGroup.Products)
		if doc == nil {
			break
		}

		action := "upsert"
		params := &api.ImportDocumentsParams{
			Action:      &action,
			BatchSize:   nil,
			DirtyValues: nil,
		}

		res, _ := client.Collection("Products").Documents().Import(doc, params)
		log.Printf("Success: %t, %s", res[0].Success, res[0].Error)

		if *nextInvMastUid == 0 {
			break
		}

	}
	return nil
}

//func importDocuments(documents a) error {
//	var documents []interface{}
//
//	action := "upsert"
//	params := &api.ImportDocumentsParams{
//		Action:      &action,
//		BatchSize:   nil,
//		DirtyValues: nil,
//	}
//
//	res, err := client.Collection("Products").Documents().Import(documents, params)
//	if err != nil {
//		return fmt.Errorf("error processing products %w", err)
//	}
//	log.Printf("Success: %t, %s", res[0].Success, res[0].Error)
//}
