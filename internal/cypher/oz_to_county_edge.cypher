MATCH (oz:Opportunity_Zone), (co:County)
WHERE oz.COUNTY_FIPS = co.COUNTY_FIPS
CREATE (oz)-[:LOCATED_IN]->(co)
