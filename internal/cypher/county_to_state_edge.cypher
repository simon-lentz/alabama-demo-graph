MATCH (c:County), (s:State)
WHERE c.STATE_FIPS = s.STATE_FIPS
CREATE (c)-[:LOCATED_IN]->(s)
