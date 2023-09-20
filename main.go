package main

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	cfg "github.com/simon-lentz/oz_cdfi_model/internal/config"
	cql "github.com/simon-lentz/oz_cdfi_model/internal/cypher"
	graph "github.com/simon-lentz/oz_cdfi_model/internal/graph"
	model "github.com/simon-lentz/oz_cdfi_model/internal/model"
)

const (
	statesPath   string = "./data/states.csv"
	countiesPath string = "./data/counties.csv"
	oppzonesPath string = "./data/opportunityzones.csv"
	testdataPath string = "./data/cejdata.csv"
)

func main() {
	ctx := context.Background()
	// Load config from .env and use credentials to connect to DB.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("godotenv.Load() err = %+v\n", err)
	}
	cfg := cfg.LoadFromEnv()
	driver, err := neo4j.NewDriverWithContext(
		cfg.Neo4jURI,
		neo4j.BasicAuth(cfg.Neo4jUserName,
			cfg.Neo4jPassword,
			""))
	if err != nil {
		log.Printf("Failed to connect to Neo4j DB: %+v\n", err)
	}
	defer driver.Close(ctx)
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	// Get the state and county nodes, write them to DB,
	// then link the county nodes to their state nodes.
	model.CreateStateNodes(statesPath, session, ctx)
	trace("model.CreateStateNodes()")()
	model.CreateCountyNodes(countiesPath, session, ctx)
	trace("model.CreateCountyNodes()")()
	_ = graph.CreateEdges(cql.CountyToState, session, ctx)
	trace("graph.CreateEdges(cql.CountyToState)")

	// Get the opportunity zone nodes, write them to the DB,
	// then link the opportunity zone nodes to their counties.
	model.CreateOppZoneNodes(oppzonesPath, session, ctx)
	trace("model.CreateOppZoneNodes()")()
	_ = graph.CreateEdges(cql.OppZoneToCounty, session, ctx)
	trace("graph.CreateEdges(cql.OppZoneToCounty)")()

	// Get the data nodes, write them to the DB,
	// then link them to their opportunity zones.
	model.CreateDataNodes(testdataPath, session, ctx)
	trace("model.CreateDataNodes()")()

	_ = graph.CreateEdges(cql.DataToOppZone, session, ctx)
	trace("graph.CreateEdges(cql.DataToOppZone)")()

}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("TRACE: %s", msg)
	return func() {
		log.Printf("%s ELAPSED: %s", msg, time.Since(start))
	}
}
