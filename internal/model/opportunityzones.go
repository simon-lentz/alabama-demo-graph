package internal

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	graph "github.com/simon-lentz/oz_cdfi_model/internal/graph"
)

type oppZone struct {
	NodeType    string
	CountyFIPS  string `csv:"COUNTY_FIPS"`
	OppZoneFIPS string `csv:"OPPORTUNITY_ZONE_FIPS"` // Primary Key.
}

func loadOppZones(filepath string) ([]oppZone, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%+v) err = %+v\n", filepath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("(*csv.Reader).ReadAll(file) err = %+v\n", err)
	}
	oppZones := []oppZone{}
	for _, r := range records {
		record := r
		oppZones = append(oppZones, oppZone{
			NodeType:    "Opportunity_Zone",
			CountyFIPS:  record[0],
			OppZoneFIPS: record[1],
		})
	}

	return oppZones, nil
}

func getOppZoneData(node *oppZone) map[string]any {
	oppZoneData := map[string]any{
		"COUNTY_FIPS":           node.CountyFIPS,
		"OPPORTUNITY_ZONE_FIPS": node.OppZoneFIPS,
	}
	return oppZoneData
}

func CreateOppZoneNodes(filepath string, session neo4j.SessionWithContext, ctx context.Context) {
	oppZones, _ := loadOppZones(filepath)
	for _, oz := range oppZones {
		oppZone := oz
		if err := graph.CreateNode(getOppZoneData(&oppZone), "OpportunityZone", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", oppZone, err)
		}
	}
}
