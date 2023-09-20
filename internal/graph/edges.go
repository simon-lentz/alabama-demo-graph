package internal

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func CreateEdges(cypherQuery string, session neo4j.SessionWithContext, ctx context.Context) error {
	_, err := session.ExecuteWrite(
		ctx,
		func(transaction neo4j.ManagedTransaction) (any, error) {
			_, err := transaction.Run(
				ctx,
				cypherQuery,
				nil) // The LOCATED_IN edge does not have properties as the moment.

			if err != nil {
				return nil, fmt.Errorf("session.ExecuteWrite() err = %+v\n", err)
			}

			return nil, nil
		})

	if err != nil {
		return fmt.Errorf("session.ExecuteWrite() err = %+v\n", err)
	}

	return nil
}
