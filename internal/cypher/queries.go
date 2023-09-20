package internal

const (
	CountyToState string = `
	MATCH (c:County)
	WITH c
	MATCH (s:State)
	WHERE s.STATE_FIPS = c.STATE_FIPS
	CREATE (s)<-[:LOCATED_IN]-(c);
	`
	OppZoneToCounty string = `
	MATCH (oz:OpportunityZone)
	WITH oz
	MATCH (co:County)
	WHERE co.COUNTY_FIPS = oz.COUNTY_FIPS
	CREATE (co)<-[:LOCATED_IN]-(oz);
	`
	DataToOppZone string = `
	MATCH (td:TestData)
	WITH td
	MATCH (oz:OpportunityZone)
	WHERE oz.OPPORTUNITY_ZONE_FIPS = td.OPPORTUNITY_ZONE_FIPS
	CREATE (td)-[r:DESCRIBES]->(oz)
	`
)

/*
Need to rewrite the edge queries to
use an index of the foreign keys
to write to relationship on to reduce
the number of db comparisons performed.
Right now it is forming a cartesian product
between disjoint pairs, not great!

temp:
OPTIONAL MATCH (c:County)
WITH c
MATCH (s:State)
WHERE s.STATE_FIPS = c.STATE_FIPS
CREATE (s)-[:LOCATED_IN]->(c);

Not sure if the query should be using
optional match for the first or second
match only, or for both matches...
*/
