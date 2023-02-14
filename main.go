package main

import (
	"errors"
	"fmt"
	"os"

	// "reflect"

	"github.com/CIDgravity/deal-filter/storagemarket"
	"github.com/urfave/cli"

	// "github.com/fatih/structtag"
	"github.com/invopop/jsonschema"
)

// func addMissingTags() {
// 	tag := reflect.TypeOf(storagemarket.StorageDeal{}).Field(1).Tag
// 	// ... and start using structtag by parsing the tag
// 	_, err := structtag.Parse(string(tag))
// 	if err != nil {
// 		panic(err)
// 	}
// }

func generateStorageJsonSchema() ([]byte, error) {
	schema := jsonschema.Reflect(&storagemarket.StorageDeal{})
	return schema.MarshalJSON()
}

func generateRetrievalJsonSchema() ([]byte, error) {
	schema := jsonschema.Reflect(&storagemarket.StorageDeal{}) // TODO change to Retrieval
	return schema.MarshalJSON()
}

func StorageAction(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return errors.New("no arguments is expected")
	}

	schemaBinary, err := generateStorageJsonSchema()
	if err != nil {
		return err
	}
	fmt.Println(string(schemaBinary))
	return nil
}

func RetrievalAction(c *cli.Context) error {
	if len(c.Args()) > 0 {
		return errors.New("no arguments is expected")
	}

	schemaBinary, err := generateRetrievalJsonSchema()
	if err != nil {
		return err
	}
	fmt.Println(string(schemaBinary))
	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "Generate Filecoin deal-filter JSON Schema"
	app.Usage = "Generate Storage or Retrieval deal-filter JSON Schema"
	app.Commands = []cli.Command{
		{
			Name:        "storage",
			HelpName:    "storage",
			Action:      StorageAction,
			ArgsUsage:   ` `,
			Usage:       `Generate Filecoin Storage deal-filter JSON Schema`,
			Description: `Generate Filecoin Storage deal-filter JSON Schema`,
		},
		{
			Name:        "retrieval",
			HelpName:    "retrieval",
			Action:      RetrievalAction,
			ArgsUsage:   ` `,
			Usage:       `Generate Filecoin Retrieval deal-filter JSON Schema`,
			Description: `Generate Filecoin Retrieval deal-filter JSON Schema`,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
