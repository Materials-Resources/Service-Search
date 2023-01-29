package schemas

import (
	"github.com/typesense/typesense-go/typesense"
	"log"
)

func RegisterSchemas(client *typesense.Client) {
	if err := registerProductSchema(client); err != nil {
		log.Println(err)
	}
	if err := registerProductGroupSchema(client); err != nil {
		log.Println(err)
	}
}
