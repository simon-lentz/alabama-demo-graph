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

type county struct {
	StateFIPS  string `csv:"STATE_FIPS"`
	CountyFIPS string `csv:"COUNTY_FIPS"` // Primary Key.
	CountyName string `csv:"COUNTY_NAME"`
}

func loadCounties(filepath string) ([]county, error) {
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
	counties := []county{}
	for _, r := range records {
		record := r
		counties = append(counties, county{
			StateFIPS:  record[0],
			CountyFIPS: record[1],
			CountyName: record[2],
		})
	}

	return counties, nil
}

func countyData(node *county) map[string]any {
	countyData := map[string]any{
		"COUNTY_NAME": node.CountyName,
		"COUNTY_FIPS": node.CountyFIPS,
		"STATE_FIPS":  node.StateFIPS,
	}

	return countyData
}

func CreateCountyNodes(filepath string, session neo4j.SessionWithContext, ctx context.Context) {
	counties, _ := loadCounties(filepath)
	for _, c := range counties {
		county := c
		if err := graph.CreateNode(countyData(&county), "County", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", county, err)
		}
	}
}
