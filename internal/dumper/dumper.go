package dumper

import (
	"encoding/json"
	"log"
	"os"

	"github.com/s-yakubovskiy/whereami/internal/contracts"
)

type LocationData struct {
	Data []contracts.Location `json:"data"`
}
type DumperJSON struct {
	db DumperInterface
}

func NewDumperJSON(d DumperInterface) (*DumperJSON, error) {
	return &DumperJSON{
		db: d,
	}, nil
}

type DumperInterface interface {
	GetAllLocations() ([]contracts.Location, error)
	ImportLocations([]contracts.Location) error
}

func (d *DumperJSON) Export(exportPath string) {
	data, err := d.db.GetAllLocations()
	if err != nil {
		log.Fatal(err)
	}
	// Convert to JSON
	jsonData, err := json.Marshal(LocationData{Data: data})
	if err != nil {
		log.Fatal(err)
	}

	// Write JSON data to file
	err = os.WriteFile(exportPath, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (d *DumperJSON) Import(importPath string) {
	file, err := os.ReadFile(importPath)
	if err != nil {
		log.Fatal(err)
	}

	var data LocationData
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}
	if err := d.db.ImportLocations(data.Data); err != nil {
		log.Fatal(err)
	}
}
