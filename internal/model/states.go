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

type state struct {
	Name      string `csv:"STATE_NAME"`
	StateFIPS string `csv:"STATE_FIPS"` // Primary Key.
}

func loadStates(filepath string) ([]state, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Open(%+v) err = %+v\n", filepath, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("(*csv.Reader).ReadAll(file) err = %v\n", err)
	}
	states := []state{}
	for _, r := range records {
		record := r
		states = append(states, state{
			Name:      record[1],
			StateFIPS: record[0],
		})
	}

	return states, nil
}

func getStateData(node *state) map[string]any {
	stateData := map[string]any{
		"STATE_NAME": node.Name,
		"STATE_FIPS": node.StateFIPS,
	}
	return stateData
}

func CreateStateNodes(filepath string, session neo4j.SessionWithContext, ctx context.Context) {
	states, _ := loadStates(filepath)
	for _, s := range states {
		state := s
		if err := graph.CreateNode(getStateData(&state), "State", session, ctx); err != nil {
			log.Printf("Failed to write %+v to DB, err = %+v\n", state, err)
		}
	}
}
