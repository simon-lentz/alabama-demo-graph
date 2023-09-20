package internal

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	graph "github.com/simon-lentz/oz_cdfi_model/internal/graph"
)

type data struct {
	OPPORTUNITY_ZONE_FIPS                                                                          string `csv:"OPPORTUNITY_ZONE_FIPS"`
	LOCALE_FIPS_CODE                                                                               string `csv:"LOCALE_FIPS_CODE"`
	ZONE_TYPE                                                                                      string `csv:"ZONE_TYPE"`
	PERCENT_BLACK_OR_AFRICAN_AMERICAN                                                              string `csv:"PERCENT_BLACK_OR_AFRICAN_AMERICAN"`
	PERCENT_AMERICAN_INDIAN_OR_ALASKA_NATIVE                                                       string `csv:"PERCENT_AMERICAN_INDIAN_OR_ALASKA_NATIVE"`
	PERCENT_ASIAN                                                                                  string `csv:"PERCENT_ASIAN"`
	PERCENT_NATIVE_HAWAIIAN_OR_PACIFIC                                                             string `csv:"PERCENT_NATIVE_HAWAIIAN_OR_PACIFIC"`
	PERCENT_TWO_OR_MORE_RACES                                                                      string `csv:"PERCENT_TWO_OR_MORE_RACES"`
	PERCENT_WHITE                                                                                  string `csv:"PERCENT_WHITE"`
	PERCENT_HISPANIC_OR_LATINO                                                                     string `csv:"PERCENT_HISPANIC_OR_LATINO"`
	PERCENT_OTHER_RACES                                                                            string `csv:"PERCENT_OTHER_RACES"`
	PERCENT_BELOW_10_YOA                                                                           string `csv:"PERCENT_BELOW_10_YOA"`
	PERCENT_10_TO_64_YOA                                                                           string `csv:"PERCENT_10_TO_64_YOA"`
	PERCENT_64_PLUS_YOA                                                                            string `csv:"PERCENT_64_PLUS_YOA"`
	TOTAL_THRESHOLD_CRITERIA_EXCEEDED                                                              string `csv:"TOTAL_THRESHOLD_CRITERIA_EXCEEDED"`
	TOTAL_CATEGORIES_EXCEEDED                                                                      string `csv:"TOTAL_CATEGORIES_EXCEEDED"`
	DISADVANTAGED_WITHOUT_CONSIDERING_NEIGHBORS                                                    string `csv:"DISADVANTAGED_WITHOUT_CONSIDERING_NEIGHBORS"`
	DISADVANTAGED_BASED_ON_NEIGHBORS_AND_RELAXED_LOW_INCOME_THRESHOLD_ONLY                         string `csv:"DISADVANTAGED_BASED_ON_NEIGHBORS_AND_RELAXED_LOW_INCOME_THRESHOLD_ONLY"`
	DISADVANTAGED_DUE_TO_TRIBAL_OVERLAP                                                            string `csv:"DISADVANTAGED_DUE_TO_TRIBAL_OVERLAP"`
	DISADVANTAGED                                                                                  string `csv:"DISADVANTAGED"`
	PERCENTAGE_DISADVANTAGED_BY_AREA                                                               string `csv:"PERCENTAGE_DISADVANTAGED_BY_AREA"`
	PERCENT_NEIGHBORS_DISADVANTAGED                                                                string `csv:"PERCENT_NEIGHBORS_DISADVANTAGED"`
	TOTAL_POPULATION                                                                               string `csv:"TOTAL_POPULATION"`
	PERCENTILE_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE                                              string `csv:"PERCENTILE_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE"`
	PERCENT_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE                                                 string `csv:"PERCENT_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE"`
	LOW_INCOME                                                                                     string `csv:"LOW_INCOME"`
	INCOME_ESTIMATED_ON_NEIGHBOR_INCOME                                                            string `csv:"INCOME_ESTIMATED_ON_NEIGHBOR_INCOME"`
	GTET_90_PERCENTILE_EXPECTED_AGRICULTURE_LOSS_RATE_AND_LOW_INCOME                               string `csv:"GTET_90_PERCENTILE_EXPECTED_AGRICULTURE_LOSS_RATE_AND_LOW_INCOME"`
	PERCENTILE_EXPECTED_AGRICULTURAL_LOSS_RATE                                                     string `csv:"PERCENTILE_EXPECTED_AGRICULTURAL_LOSS_RATE"`
	EXPECTED_AGRICULTURAL_LOSS_RATE                                                                string `csv:"EXPECTED_AGRICULTURAL_LOSS_RATE"`
	GTET_90_PERCENTILE_EXPECTED_BUILDING_LOSS_RATE_AND_LOW_INCOME                                  string `csv:"GTET_90_PERCENTILE_EXPECTED_BUILDING_LOSS_RATE_AND_LOW_INCOME"`
	PERCENTILE_EXPECTED_BUILDING_LOSS_RATE                                                         string `csv:"PERCENTILE_EXPECTED_BUILDING_LOSS_RATE"`
	EXPECTED_BUILDING_LOSS_RATE                                                                    string `csv:"EXPECTED_BUILDING_LOSS_RATE"`
	GTET_90_PERCENTILE_EXPECTED_POPULATION_LOSS_RATE_AND_LOW_INCOME                                string `csv:"GTET_90_PERCENTILE_EXPECTED_POPULATION_LOSS_RATE_AND_LOW_INCOME"`
	PERCENTILE_EXPECTED_POPULATION_LOSS_RATE                                                       string `csv:"PERCENTILE_EXPECTED_POPULATION_LOSS_RATE"`
	EXPECTED_POPULATION_LOSS_RATE                                                                  string `csv:"EXPECTED_POPULATION_LOSS_RATE"`
	PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS                                             string `csv:"PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS"`
	PERCENT_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS                                                string `csv:"PERCENT_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS"`
	GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS                                     string `csv:"GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS"`
	GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS_AND_LOW_INCOME                      string `csv:"GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS_AND_LOW_INCOME"`
	PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS                                              string `csv:"PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS"`
	PERCENT_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS                                                 string `csv:"PERCENT_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS"`
	GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS                                      string `csv:"GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS"`
	GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS_AND_LOW_INCOME                       string `csv:"GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS_AND_LOW_INCOME"`
	GTET_90_PERCENTILE_ENERGY_BURDEN_AND_LOW_INCOME                                                string `csv:"GTET_90_PERCENTILE_ENERGY_BURDEN_AND_LOW_INCOME"`
	PERCENTILE_ENERGY_BURDEN                                                                       string `csv:"PERCENTILE_ENERGY_BURDEN"`
	ENERGY_BURDEN                                                                                  string `csv:"ENERGY_BURDEN"`
	GTET_90_PERCENTILE_PM25_EXPOSURE_AND_LOW_INCOME                                                string `csv:"GTET_90_PERCENTILE_PM25_EXPOSURE_AND_LOW_INCOME"`
	PERCENTILE_AIRBORNE_PM25                                                                       string `csv:"PERCENTILE_AIRBORNE_PM25"`
	AIRBORNE_PM25                                                                                  string `csv:"AIRBORNE_PM25"`
	GTET_90_PERCENTILE_DIESEL_PARTICULATE_MATTER_AND_LOW_INCOME                                    string `csv:"GTET_90_PERCENTILE_DIESEL_PARTICULATE_MATTER_AND_LOW_INCOME"`
	PERCENTILE_DIESEL_PARTICULATE_MATTER_EXPOSURE                                                  string `csv:"PERCENTILE_DIESEL_PARTICULATE_MATTER_EXPOSURE"`
	DIESEL_PARTICULATE_MATTER_EXPOSURE                                                             string `csv:"DIESEL_PARTICULATE_MATTER_EXPOSURE"`
	GTET_90_PERCENTILE_TRAFFIC_PROXIMITY_AND_LOW_INCOME                                            string `csv:"GTET_90_PERCENTILE_TRAFFIC_PROXIMITY_AND_LOW_INCOME"`
	PERCENTILE_TRAFFIC_PROXIMITY_AND_VOLUME                                                        string `csv:"PERCENTILE_TRAFFIC_PROXIMITY_AND_VOLUME"`
	TRAFFIC_PROXIMITY_AND_VOLUME                                                                   string `csv:"TRAFFIC_PROXIMITY_AND_VOLUME"`
	GTET_90_PERCENTILE_FOR_DOT_TRANSIT_BARRIERS_AND_LOW_INCOME                                     string `csv:"GTET_90_PERCENTILE_FOR_DOT_TRANSIT_BARRIERS_AND_LOW_INCOME"`
	PERCENTILE_DOT_TRAVEL_BARRIERS_SCORE                                                           string `csv:"PERCENTILE_DOT_TRAVEL_BARRIERS_SCORE"`
	GTET_90_PERCENTILE_HOUSING_BURDEN_AND_LOW_INCOME                                               string `csv:"GTET_90_PERCENTILE_HOUSING_BURDEN_AND_LOW_INCOME"`
	PERCENTILE_HOUSING_BURDEN                                                                      string `csv:"PERCENTILE_HOUSING_BURDEN"`
	PERCENT_HOUSING_BURDEN                                                                         string `csv:"PERCENT_HOUSING_BURDEN"`
	GTET_90_PERCENTILE_FOR_LEAD_PAINT                                                              string `csv:"GTET_90_PERCENTILE_FOR_LEAD_PAINT"`
	MEDIAN_HOUSE_VALUE_BELOW_90_PERCENTILE_AND_LOW_INCOME                                          string `csv:"MEDIAN_HOUSE_VALUE_BELOW_90_PERCENTILE_AND_LOW_INCOME"`
	PERCENTILE_PRE_1960S_HOUSING                                                                   string `csv:"PERCENTILE_PRE_1960S_HOUSING"`
	PERCENT_PRE_1960S_HOUSING                                                                      string `csv:"PERCENT_PRE_1960S_HOUSING"`
	PERCENTILE_MEDIAN_VALUE_OWNER_OCCUPIED_HOUSING_UNITS                                           string `csv:"PERCENTILE_MEDIAN_VALUE_OWNER_OCCUPIED_HOUSING_UNITS"`
	MEDIAN_DOLLAR_VALUE_OWNER_OCCUPIED_HOUSING_UNITS                                               string `csv:"MEDIAN_DOLLAR_VALUE_OWNER_OCCUPIED_HOUSING_UNITS"`
	GTET_90_PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND_AND_LOW_INCOME                     string `csv:"GTET_90_PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND_AND_LOW_INCOME"`
	GTET_90_PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND                                    string `csv:"GTET_90_PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND"`
	PERCENT_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND                                               string `csv:"PERCENT_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND"`
	PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND                                            string `csv:"PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND"`
	AREA_GTET_35_ACRES                                                                             string `csv:"AREA_GTET_35_ACRES"`
	HISTORIC_UNDERINVESTMENT_AND_LOW_INCOME                                                        string `csv:"HISTORIC_UNDERINVESTMENT_AND_LOW_INCOME"`
	HISTORIC_UNDERINVESTMENT                                                                       string `csv:"HISTORIC_UNDERINVESTMENT"`
	PERCENTILE_HOMES_NO_KITCHEN_NO_INDOOR_PLUMBING                                                 string `csv:"PERCENTILE_HOMES_NO_KITCHEN_NO_INDOOR_PLUMBING"`
	PERCENT_HOMES_NO_KITCHEN_NO_INDOOR_PLUMBING                                                    string `csv:"PERCENT_HOMES_NO_KITCHEN_NO_INDOOR_PLUMBING"`
	GTET_90_PERCENTILE_PROXIMITY_TO_HAZARDOUS_WASTE_FACILITIES_AND_LOW_INCOME                      string `csv:"GTET_90_PERCENTILE_PROXIMITY_TO_HAZARDOUS_WASTE_FACILITIES_AND_LOW_INCOME"`
	PERCENTILE_PROXIMITY_TO_HAZARDOUS_WASTE_SITES                                                  string `csv:"PERCENTILE_PROXIMITY_TO_HAZARDOUS_WASTE_SITES"`
	PROXIMITY_TO_HAZARDOUS_WASTE_SITES                                                             string `csv:"PROXIMITY_TO_HAZARDOUS_WASTE_SITES"`
	GTET_90_PERCENTILE_PROXIMITY_TO_SUPERFUND_SITES_AND_LOW_INCOME                                 string `csv:"GTET_90_PERCENTILE_PROXIMITY_TO_SUPERFUND_SITES_AND_LOW_INCOME"`
	PERCENTILE_PROXIMITY_TO_SUPERFUND_SITES                                                        string `csv:"PERCENTILE_PROXIMITY_TO_SUPERFUND_SITES"`
	PROXIMITY_TO_SUPERFUND_SITES                                                                   string `csv:"PROXIMITY_TO_SUPERFUND_SITES"`
	GTET_90_PERCENTILE_PROXIMITY_TO_RISK_MANAGEMENT_PLAN_SITES_AND_LOW_INCOME                      string `csv:"GTET_90_PERCENTILE_PROXIMITY_TO_RISK_MANAGEMENT_PLAN_SITES_AND_LOW_INCOME"`
	PERCENTILE_PROXIMITY_TO_RISK_MANAGEMENT_PLAN_FACILITIES                                        string `csv:"PERCENTILE_PROXIMITY_TO_RISK_MANAGEMENT_PLAN_FACILITIES"`
	PROXIMITY_TO_RISK_MANAGEMENT_PLAN_FACILITIES                                                   string `csv:"PROXIMITY_TO_RISK_MANAGEMENT_PLAN_FACILITIES"`
	FORMERLY_USED_DEFENSE_SITE_IN_TRACT                                                            string `csv:"FORMERLY_USED_DEFENSE_SITE_IN_TRACT"`
	ABANDONED_MINE_IN_TRACT                                                                        string `csv:"ABANDONED_MINE_IN_TRACT"`
	ABANDONED_MINE_IN_TRACT_AND_LOW_INCOME                                                         string `csv:"ABANDONED_MINE_IN_TRACT_AND_LOW_INCOME"`
	FORMERLY_USED_DEFENSE_SITE_IN_TRACT_AND_LOW_INCOME                                             string `csv:"FORMERLY_USED_DEFENSE_SITE_IN_TRACT_AND_LOW_INCOME"`
	FORMERLY_USED_DEFENSE_SITE_IN_TRACT_MISSING_DATA_EQUALS_FALSE                                  string `csv:"FORMERLY_USED_DEFENSE_SITE_IN_TRACT_MISSING_DATA_EQUALS_FALSE"`
	ABANDONED_MINE_IN_TRACT_MISSING_DATA_EQUALS_FALSE                                              string `csv:"ABANDONED_MINE_IN_TRACT_MISSING_DATA_EQUALS_FALSE"`
	GTET_90_PERCENTILE_WASTEWATER_DISCHARGE_AND_LOW_INCOME                                         string `csv:"GTET_90_PERCENTILE_WASTEWATER_DISCHARGE_AND_LOW_INCOME"`
	PERCENTILE_WASTEWATER_DISCHARGE                                                                string `csv:"PERCENTILE_WASTEWATER_DISCHARGE"`
	WASTEWATER_DISCHARGE                                                                           string `csv:"WASTEWATER_DISCHARGE"`
	GTET_90_PERCENTILE_LEAKY_UNDERGROUND_STORAGE_TANKS_AND_LOW_INCOME                              string `csv:"GTET_90_PERCENTILE_LEAKY_UNDERGROUND_STORAGE_TANKS_AND_LOW_INCOME"`
	PERCENTILE_LEAKY_UNDERGROUND_STORAGE_TANKS                                                     string `csv:"PERCENTILE_LEAKY_UNDERGROUND_STORAGE_TANKS"`
	LEAKY_UNDERGROUND_STORAGE_TANKS                                                                string `csv:"LEAKY_UNDERGROUND_STORAGE_TANKS"`
	GTET_90_PERCENTILE_ASTHMA_AND_LOW_INCOME                                                       string `csv:"GTET_90_PERCENTILE_ASTHMA_AND_LOW_INCOME"`
	PERCENTILE_ASTHMA_GTET_18_YOA                                                                  string `csv:"PERCENTILE_ASTHMA_GTET_18_YOA"`
	PERCENT_ASTHMA_GTET_18_YOA                                                                     string `csv:"PERCENT_ASTHMA_GTET_18_YOA"`
	GTET_90_PERCENTILE_FOR_DIABETES_AND_LOW_INCOME                                                 string `csv:"GTET_90_PERCENTILE_FOR_DIABETES_AND_LOW_INCOME"`
	PERCENTILE_DIAGNOSED_DIABETES_GTET_18_YOA                                                      string `csv:"PERCENTILE_DIAGNOSED_DIABETES_GTET_18_YOA"`
	DIAGNOSED_DIABETES_GTET_18_YOA                                                                 string `csv:"DIAGNOSED_DIABETES_GTET_18_YOA"`
	GTET_90_PERCENTILE_HEART_DISEASE_AND_LOW_INCOME                                                string `csv:"GTET_90_PERCENTILE_HEART_DISEASE_AND_LOW_INCOME"`
	PERCENTILE_CORONARY_HEART_DISEASE_GTET_18_YOA                                                  string `csv:"PERCENTILE_CORONARY_HEART_DISEASE_GTET_18_YOA"`
	PERCENT_CORONARY_HEART_DISEASE_GTET_18_YOA                                                     string `csv:"PERCENT_CORONARY_HEART_DISEASE_GTET_18_YOA"`
	GTET_90_PERCENTILE_LOW_LIFE_EXPECTANCY_AND_LOW_INCOME                                          string `csv:"GTET_90_PERCENTILE_LOW_LIFE_EXPECTANCY_AND_LOW_INCOME"`
	PERCENTILE_LOW_LIFE_EXPECTANCY                                                                 string `csv:"PERCENTILE_LOW_LIFE_EXPECTANCY"`
	LIFE_EXPECTANCY_YEARS                                                                          string `csv:"LIFE_EXPECTANCY_YEARS"`
	GTET_THE_90_PERCENTILE_FOR_LOW_MEDIAN_INCOME_AS_PERCENT_OF_AREA_MEDIAN_AND_LOW_HS_ATTAINMENT   string `csv:"GTET_THE_90_PERCENTILE_FOR_LOW_MEDIAN_INCOME_AS_PERCENT_OF_AREA_MEDIAN_AND_LOW_HS_ATTAINMENT"`
	PERCENTILE_LOW_MEDIAN_INCOME_AS_PERCENT_AREA_MEDIAN_INCOME                                     string `csv:"PERCENTILE_LOW_MEDIAN_INCOME_AS_PERCENT_AREA_MEDIAN_INCOME"`
	MEDIAN_INCOME_AS_PERCENT_OF_AREA_MEDIAN                                                        string `csv:"MEDIAN_INCOME_AS_PERCENT_OF_AREA_MEDIAN"`
	GTET_THE_90_PERCENTILE_HOUSEHOLDS_IN_LINGUISTIC_ISOLATION_AND_LOW_HS_ATTAINMENT                string `csv:"GTET_THE_90_PERCENTILE_HOUSEHOLDS_IN_LINGUISTIC_ISOLATION_AND_LOW_HS_ATTAINMENT"`
	PERCENTILE_LINGUISTIC_ISOLATION                                                                string `csv:"PERCENTILE_LINGUISTIC_ISOLATION"`
	PERCENT_LINGUISTIC_ISOLATION                                                                   string `csv:"PERCENT_LINGUISTIC_ISOLATION"`
	GTET_90_PERCENTILE_UNEMPLOYMENT_AND_LOW_HS                                                     string `csv:"GTET_90_PERCENTILE_UNEMPLOYMENT_AND_LOW_HS"`
	PERCENTILE_UNEMPLOYMENT                                                                        string `csv:"PERCENTILE_UNEMPLOYMENT"`
	PERCENT_UNEMPLOYMENT                                                                           string `csv:"PERCENT_UNEMPLOYMENT"`
	GTET_90_PERCENTILE_HOUSEHOLDS_BELOW_100_PERCENT_FEDERAL_POVERTY_LEVEL_AND_LOW_HS_ATTAINMENT    string `csv:"GTET_90_PERCENTILE_HOUSEHOLDS_BELOW_100_PERCENT_FEDERAL_POVERTY_LEVEL_AND_LOW_HS_ATTAINMENT"`
	PERCENTILE_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE_TWO                                          string `csv:"PERCENTILE_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE_TWO"`
	PERCENT_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE_TWO                                             string `csv:"PERCENT_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE_TWO"`
	PERCENTILE_BELOW_100_PERCENT_FEDERAL_POVERTY_LINE                                              string `csv:"PERCENTILE_BELOW_100_PERCENT_FEDERAL_POVERTY_LINE"`
	PERCENT_BELOW_100_PERCENT_FEDERAL_POVERTY_LINE                                                 string `csv:"PERCENT_BELOW_100_PERCENT_FEDERAL_POVERTY_LINE"`
	PERCENTILE_AGE_25_PLUS_LESS_THAN_HIGH_SCHOOL                                                   string `csv:"PERCENTILE_AGE_25_PLUS_LESS_THAN_HIGH_SCHOOL"`
	PERCENT_AGE_25_PLUS_LESS_THAN_HIGH_SCHOOL                                                      string `csv:"PERCENT_AGE_25_PLUS_LESS_THAN_HIGH_SCHOOL"`
	PERCENT_RESIDENTS_NOT_CURRENTLY_ENROLLED_IN_HIGHER_ED                                          string `csv:"PERCENT_RESIDENTS_NOT_CURRENTLY_ENROLLED_IN_HIGHER_ED"`
	UNEMPLOYMENT_PERCENT_ISLAND_AREAS_AND_STATES_AND_PR                                            string `csv:"UNEMPLOYMENT_PERCENT_ISLAND_AREAS_AND_STATES_AND_PR"`
	PERCENTAGE_HOUSEHOLDS_BELOW_100_PERCENT_OF_FEDERAL_POVERTY_LINE_ISLAND_AREAS_AND_STATES_AND_PR string `csv:"PERCENTAGE_HOUSEHOLDS_BELOW_100_PERCENT_OF_FEDERAL_POVERTY_LINE_ISLAND_AREAS_AND_STATES_AND_PR"`
	GTET_90_PERCENTILE_FOR_UNEMPLOYMENT_AND_LOW_HS_EDUCATION_ISLAND_AREAS                          string `csv:"GTET_90_PERCENTILE_FOR_UNEMPLOYMENT_AND_LOW_HS_EDUCATION_ISLAND_AREAS"`
	GTET_90_PERCENTILE_BELOW_100_PERCENT_FEDERAL_POVERTY_LEVEL_AND_LOW_HS_EDUCATION_ISLAND_AREAS   string `csv:"GTET_90_PERCENTILE_BELOW_100_PERCENT_FEDERAL_POVERTY_LEVEL_AND_LOW_HS_EDUCATION_ISLAND_AREAS"`
	GTET_90_PERCENTILE_LOW_AREA_MEDIAN_HOUSEHOLD_INCOME_AND_LOW_HS_EDUCATION_ISLAND_AREAS          string `csv:"GTET_90_PERCENTILE_LOW_AREA_MEDIAN_HOUSEHOLD_INCOME_AND_LOW_HS_EDUCATION_ISLAND_AREAS"`
	NUMBER_OF_TRIBAL_AREAS_ALASKA                                                                  string `csv:"NUMBER_OF_TRIBAL_AREAS_ALASKA"`
	NAMES_TRIBAL_AREAS                                                                             string `csv:"NAMES_TRIBAL_AREAS"`
	PERCENT_WITHIN_TRIBAL_AREAS                                                                    string `csv:"PERCENT_WITHIN_TRIBAL_AREAS"`
}

func loadData(filepath string) ([]data, error) {
	f, err := os.Open("./data/cejdata.csv")
	if err != nil {
		return nil, fmt.Errorf("\nos.Open(%+v) err = %+v\n", "./data/cejdata.csv", err)
	}
	defer f.Close()

	var out []data
	if err := gocsv.UnmarshalFile(f, &out); err != nil {
		panic(err)
	}

	return out, nil
}

func allData(node *data) map[string]any {
	testData := map[string]any{
		"OPPORTUNITY_ZONE_FIPS":                       node.OPPORTUNITY_ZONE_FIPS,
		"LOCALE_FIPS_CODE":                            node.LOCALE_FIPS_CODE,
		"ZONE_TYPE":                                   node.ZONE_TYPE,
		"PERCENT_BLACK_OR_AFRICAN_AMERICAN":           node.PERCENT_BLACK_OR_AFRICAN_AMERICAN,
		"PERCENT_AMERICAN_INDIAN_OR_ALASKA_NATIVE":    node.PERCENT_AMERICAN_INDIAN_OR_ALASKA_NATIVE,
		"PERCENT_ASIAN":                               node.PERCENT_ASIAN,
		"PERCENT_NATIVE_HAWAIIAN_OR_PACIFIC":          node.PERCENT_NATIVE_HAWAIIAN_OR_PACIFIC,
		"PERCENT_TWO_OR_MORE_RACES":                   node.PERCENT_TWO_OR_MORE_RACES,
		"PERCENT_WHITE":                               node.PERCENT_WHITE,
		"PERCENT_HISPANIC_OR_LATINO":                  node.PERCENT_HISPANIC_OR_LATINO,
		"PERCENT_OTHER_RACES":                         node.PERCENT_OTHER_RACES,
		"PERCENT_BELOW_10_YOA":                        node.PERCENT_BELOW_10_YOA,
		"PERCENT_10_TO_64_YOA":                        node.PERCENT_10_TO_64_YOA,
		"PERCENT_64_PLUS_YOA":                         node.PERCENT_64_PLUS_YOA,
		"TOTAL_THRESHOLD_CRITERIA_EXCEEDED":           node.TOTAL_THRESHOLD_CRITERIA_EXCEEDED,
		"TOTAL_CATEGORIES_EXCEEDED":                   node.TOTAL_CATEGORIES_EXCEEDED,
		"DISADVANTAGED_WITHOUT_CONSIDERING_NEIGHBORS": node.DISADVANTAGED_WITHOUT_CONSIDERING_NEIGHBORS,
		"DISADVANTAGED_BASED_ON_NEIGHBORS_AND_RELAXED_LOW_INCOME_THRESHOLD_ONLY": node.DISADVANTAGED_BASED_ON_NEIGHBORS_AND_RELAXED_LOW_INCOME_THRESHOLD_ONLY,
		"DISADVANTAGED_DUE_TO_TRIBAL_OVERLAP":                                    node.DISADVANTAGED_DUE_TO_TRIBAL_OVERLAP,
		"DISADVANTAGED":                                                          node.DISADVANTAGED,
		"PERCENTAGE_DISADVANTAGED_BY_AREA":                                       node.PERCENTAGE_DISADVANTAGED_BY_AREA,
		"PERCENT_NEIGHBORS_DISADVANTAGED":                                        node.PERCENT_NEIGHBORS_DISADVANTAGED,
		"TOTAL_POPULATION":                                                       node.TOTAL_POPULATION,
		"PERCENTILE_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE":                      node.PERCENTILE_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE,
		"PERCENT_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE":                         node.PERCENT_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE,
		"LOW_INCOME":                          node.LOW_INCOME,
		"INCOME_ESTIMATED_ON_NEIGHBOR_INCOME": node.INCOME_ESTIMATED_ON_NEIGHBOR_INCOME,
		"GTET_90_PERCENTILE_EXPECTED_AGRICULTURE_LOSS_RATE_AND_LOW_INCOME":                               node.GTET_90_PERCENTILE_EXPECTED_AGRICULTURE_LOSS_RATE_AND_LOW_INCOME,
		"PERCENTILE_EXPECTED_AGRICULTURAL_LOSS_RATE":                                                     node.PERCENTILE_EXPECTED_AGRICULTURAL_LOSS_RATE,
		"EXPECTED_AGRICULTURAL_LOSS_RATE":                                                                node.EXPECTED_AGRICULTURAL_LOSS_RATE,
		"GTET_90_PERCENTILE_EXPECTED_BUILDING_LOSS_RATE_AND_LOW_INCOME":                                  node.GTET_90_PERCENTILE_EXPECTED_BUILDING_LOSS_RATE_AND_LOW_INCOME,
		"PERCENTILE_EXPECTED_BUILDING_LOSS_RATE":                                                         node.PERCENTILE_EXPECTED_BUILDING_LOSS_RATE,
		"EXPECTED_BUILDING_LOSS_RATE":                                                                    node.EXPECTED_BUILDING_LOSS_RATE,
		"GTET_90_PERCENTILE_EXPECTED_POPULATION_LOSS_RATE_AND_LOW_INCOME":                                node.GTET_90_PERCENTILE_EXPECTED_POPULATION_LOSS_RATE_AND_LOW_INCOME,
		"PERCENTILE_EXPECTED_POPULATION_LOSS_RATE":                                                       node.PERCENTILE_EXPECTED_POPULATION_LOSS_RATE,
		"EXPECTED_POPULATION_LOSS_RATE":                                                                  node.EXPECTED_POPULATION_LOSS_RATE,
		"PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS":                                             node.PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS,
		"PERCENT_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS":                                                node.PERCENT_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS,
		"GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS":                                     node.GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS,
		"GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS_AND_LOW_INCOME":                      node.GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FLOOD_IN_30_YEARS_AND_LOW_INCOME,
		"PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS":                                              node.PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS,
		"PERCENT_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS":                                                 node.PERCENT_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS,
		"GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS":                                      node.GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS,
		"GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS_AND_LOW_INCOME":                       node.GTET_90_PERCENTILE_PROPERTIES_AT_RISK_OF_FIRE_IN_30_YEARS_AND_LOW_INCOME,
		"GTET_90_PERCENTILE_ENERGY_BURDEN_AND_LOW_INCOME":                                                node.GTET_90_PERCENTILE_ENERGY_BURDEN_AND_LOW_INCOME,
		"PERCENTILE_ENERGY_BURDEN":                                                                       node.PERCENTILE_ENERGY_BURDEN,
		"ENERGY_BURDEN":                                                                                  node.ENERGY_BURDEN,
		"GTET_90_PERCENTILE_PM25_EXPOSURE_AND_LOW_INCOME":                                                node.GTET_90_PERCENTILE_PM25_EXPOSURE_AND_LOW_INCOME,
		"PERCENTILE_AIRBORNE_PM25":                                                                       node.PERCENTILE_AIRBORNE_PM25,
		"AIRBORNE_PM25":                                                                                  node.AIRBORNE_PM25,
		"GTET_90_PERCENTILE_DIESEL_PARTICULATE_MATTER_AND_LOW_INCOME":                                    node.GTET_90_PERCENTILE_DIESEL_PARTICULATE_MATTER_AND_LOW_INCOME,
		"PERCENTILE_DIESEL_PARTICULATE_MATTER_EXPOSURE":                                                  node.PERCENTILE_DIESEL_PARTICULATE_MATTER_EXPOSURE,
		"DIESEL_PARTICULATE_MATTER_EXPOSURE":                                                             node.DIESEL_PARTICULATE_MATTER_EXPOSURE,
		"GTET_90_PERCENTILE_TRAFFIC_PROXIMITY_AND_LOW_INCOME":                                            node.GTET_90_PERCENTILE_TRAFFIC_PROXIMITY_AND_LOW_INCOME,
		"PERCENTILE_TRAFFIC_PROXIMITY_AND_VOLUME":                                                        node.PERCENTILE_TRAFFIC_PROXIMITY_AND_VOLUME,
		"TRAFFIC_PROXIMITY_AND_VOLUME":                                                                   node.TRAFFIC_PROXIMITY_AND_VOLUME,
		"GTET_90_PERCENTILE_FOR_DOT_TRANSIT_BARRIERS_AND_LOW_INCOME":                                     node.GTET_90_PERCENTILE_FOR_DOT_TRANSIT_BARRIERS_AND_LOW_INCOME,
		"PERCENTILE_DOT_TRAVEL_BARRIERS_SCORE":                                                           node.PERCENTILE_DOT_TRAVEL_BARRIERS_SCORE,
		"GTET_90_PERCENTILE_HOUSING_BURDEN_AND_LOW_INCOME":                                               node.GTET_90_PERCENTILE_HOUSING_BURDEN_AND_LOW_INCOME,
		"PERCENTILE_HOUSING_BURDEN":                                                                      node.PERCENTILE_HOUSING_BURDEN,
		"PERCENT_HOUSING_BURDEN":                                                                         node.PERCENT_HOUSING_BURDEN,
		"GTET_90_PERCENTILE_FOR_LEAD_PAINT":                                                              node.GTET_90_PERCENTILE_FOR_LEAD_PAINT,
		"MEDIAN_HOUSE_VALUE_BELOW_90_PERCENTILE_AND_LOW_INCOME":                                          node.MEDIAN_HOUSE_VALUE_BELOW_90_PERCENTILE_AND_LOW_INCOME,
		"PERCENTILE_PRE_1960S_HOUSING":                                                                   node.PERCENTILE_PRE_1960S_HOUSING,
		"PERCENT_PRE_1960S_HOUSING":                                                                      node.PERCENT_PRE_1960S_HOUSING,
		"PERCENTILE_MEDIAN_VALUE_OWNER_OCCUPIED_HOUSING_UNITS":                                           node.PERCENTILE_MEDIAN_VALUE_OWNER_OCCUPIED_HOUSING_UNITS,
		"MEDIAN_DOLLAR_VALUE_OWNER_OCCUPIED_HOUSING_UNITS":                                               node.MEDIAN_DOLLAR_VALUE_OWNER_OCCUPIED_HOUSING_UNITS,
		"GTET_90_PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND_AND_LOW_INCOME":                     node.GTET_90_PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND_AND_LOW_INCOME,
		"GTET_90_PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND":                                    node.GTET_90_PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND,
		"PERCENT_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND":                                               node.PERCENT_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND,
		"PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND":                                            node.PERCENTILE_LAND_AREA_IMPERVIOUS_SURFACE_OR_CROPLAND,
		"AREA_GTET_35_ACRES":                                                                             node.AREA_GTET_35_ACRES,
		"HISTORIC_UNDERINVESTMENT_AND_LOW_INCOME":                                                        node.HISTORIC_UNDERINVESTMENT_AND_LOW_INCOME,
		"HISTORIC_UNDERINVESTMENT":                                                                       node.HISTORIC_UNDERINVESTMENT,
		"PERCENTILE_HOMES_NO_KITCHEN_NO_INDOOR_PLUMBING":                                                 node.PERCENTILE_HOMES_NO_KITCHEN_NO_INDOOR_PLUMBING,
		"PERCENT_HOMES_NO_KITCHEN_NO_INDOOR_PLUMBING":                                                    node.PERCENT_HOMES_NO_KITCHEN_NO_INDOOR_PLUMBING,
		"GTET_90_PERCENTILE_PROXIMITY_TO_HAZARDOUS_WASTE_FACILITIES_AND_LOW_INCOME":                      node.GTET_90_PERCENTILE_PROXIMITY_TO_HAZARDOUS_WASTE_FACILITIES_AND_LOW_INCOME,
		"PERCENTILE_PROXIMITY_TO_HAZARDOUS_WASTE_SITES":                                                  node.PERCENTILE_PROXIMITY_TO_HAZARDOUS_WASTE_SITES,
		"PROXIMITY_TO_HAZARDOUS_WASTE_SITES":                                                             node.PROXIMITY_TO_HAZARDOUS_WASTE_SITES,
		"GTET_90_PERCENTILE_PROXIMITY_TO_SUPERFUND_SITES_AND_LOW_INCOME":                                 node.GTET_90_PERCENTILE_PROXIMITY_TO_SUPERFUND_SITES_AND_LOW_INCOME,
		"PERCENTILE_PROXIMITY_TO_SUPERFUND_SITES":                                                        node.PERCENTILE_PROXIMITY_TO_SUPERFUND_SITES,
		"PROXIMITY_TO_SUPERFUND_SITES":                                                                   node.PROXIMITY_TO_SUPERFUND_SITES,
		"GTET_90_PERCENTILE_PROXIMITY_TO_RISK_MANAGEMENT_PLAN_SITES_AND_LOW_INCOME":                      node.GTET_90_PERCENTILE_PROXIMITY_TO_RISK_MANAGEMENT_PLAN_SITES_AND_LOW_INCOME,
		"PERCENTILE_PROXIMITY_TO_RISK_MANAGEMENT_PLAN_FACILITIES":                                        node.PERCENTILE_PROXIMITY_TO_RISK_MANAGEMENT_PLAN_FACILITIES,
		"PROXIMITY_TO_RISK_MANAGEMENT_PLAN_FACILITIES":                                                   node.PROXIMITY_TO_RISK_MANAGEMENT_PLAN_FACILITIES,
		"FORMERLY_USED_DEFENSE_SITE_IN_TRACT":                                                            node.FORMERLY_USED_DEFENSE_SITE_IN_TRACT,
		"ABANDONED_MINE_IN_TRACT":                                                                        node.ABANDONED_MINE_IN_TRACT,
		"ABANDONED_MINE_IN_TRACT_AND_LOW_INCOME":                                                         node.ABANDONED_MINE_IN_TRACT_AND_LOW_INCOME,
		"FORMERLY_USED_DEFENSE_SITE_IN_TRACT_AND_LOW_INCOME":                                             node.FORMERLY_USED_DEFENSE_SITE_IN_TRACT_AND_LOW_INCOME,
		"FORMERLY_USED_DEFENSE_SITE_IN_TRACT_MISSING_DATA_EQUALS_FALSE":                                  node.FORMERLY_USED_DEFENSE_SITE_IN_TRACT_MISSING_DATA_EQUALS_FALSE,
		"ABANDONED_MINE_IN_TRACT_MISSING_DATA_EQUALS_FALSE":                                              node.ABANDONED_MINE_IN_TRACT_MISSING_DATA_EQUALS_FALSE,
		"GTET_90_PERCENTILE_WASTEWATER_DISCHARGE_AND_LOW_INCOME":                                         node.GTET_90_PERCENTILE_WASTEWATER_DISCHARGE_AND_LOW_INCOME,
		"PERCENTILE_WASTEWATER_DISCHARGE":                                                                node.PERCENTILE_WASTEWATER_DISCHARGE,
		"WASTEWATER_DISCHARGE":                                                                           node.WASTEWATER_DISCHARGE,
		"GTET_90_PERCENTILE_LEAKY_UNDERGROUND_STORAGE_TANKS_AND_LOW_INCOME":                              node.GTET_90_PERCENTILE_LEAKY_UNDERGROUND_STORAGE_TANKS_AND_LOW_INCOME,
		"PERCENTILE_LEAKY_UNDERGROUND_STORAGE_TANKS":                                                     node.PERCENTILE_LEAKY_UNDERGROUND_STORAGE_TANKS,
		"LEAKY_UNDERGROUND_STORAGE_TANKS":                                                                node.LEAKY_UNDERGROUND_STORAGE_TANKS,
		"GTET_90_PERCENTILE_ASTHMA_AND_LOW_INCOME":                                                       node.GTET_90_PERCENTILE_ASTHMA_AND_LOW_INCOME,
		"PERCENTILE_ASTHMA_GTET_18_YOA":                                                                  node.PERCENTILE_ASTHMA_GTET_18_YOA,
		"PERCENT_ASTHMA_GTET_18_YOA":                                                                     node.PERCENT_ASTHMA_GTET_18_YOA,
		"GTET_90_PERCENTILE_FOR_DIABETES_AND_LOW_INCOME":                                                 node.GTET_90_PERCENTILE_FOR_DIABETES_AND_LOW_INCOME,
		"PERCENTILE_DIAGNOSED_DIABETES_GTET_18_YOA":                                                      node.PERCENTILE_DIAGNOSED_DIABETES_GTET_18_YOA,
		"DIAGNOSED_DIABETES_GTET_18_YOA":                                                                 node.DIAGNOSED_DIABETES_GTET_18_YOA,
		"GTET_90_PERCENTILE_HEART_DISEASE_AND_LOW_INCOME":                                                node.GTET_90_PERCENTILE_HEART_DISEASE_AND_LOW_INCOME,
		"PERCENTILE_CORONARY_HEART_DISEASE_GTET_18_YOA":                                                  node.PERCENTILE_CORONARY_HEART_DISEASE_GTET_18_YOA,
		"PERCENT_CORONARY_HEART_DISEASE_GTET_18_YOA":                                                     node.PERCENT_CORONARY_HEART_DISEASE_GTET_18_YOA,
		"GTET_90_PERCENTILE_LOW_LIFE_EXPECTANCY_AND_LOW_INCOME":                                          node.GTET_90_PERCENTILE_LOW_LIFE_EXPECTANCY_AND_LOW_INCOME,
		"PERCENTILE_LOW_LIFE_EXPECTANCY":                                                                 node.PERCENTILE_LOW_LIFE_EXPECTANCY,
		"LIFE_EXPECTANCY_YEARS":                                                                          node.LIFE_EXPECTANCY_YEARS,
		"GTET_THE_90_PERCENTILE_FOR_LOW_MEDIAN_INCOME_AS_PERCENT_OF_AREA_MEDIAN_AND_LOW_HS_ATTAINMENT":   node.GTET_THE_90_PERCENTILE_FOR_LOW_MEDIAN_INCOME_AS_PERCENT_OF_AREA_MEDIAN_AND_LOW_HS_ATTAINMENT,
		"PERCENTILE_LOW_MEDIAN_INCOME_AS_PERCENT_AREA_MEDIAN_INCOME":                                     node.PERCENTILE_LOW_MEDIAN_INCOME_AS_PERCENT_AREA_MEDIAN_INCOME,
		"MEDIAN_INCOME_AS_PERCENT_OF_AREA_MEDIAN":                                                        node.MEDIAN_INCOME_AS_PERCENT_OF_AREA_MEDIAN,
		"GTET_THE_90_PERCENTILE_HOUSEHOLDS_IN_LINGUISTIC_ISOLATION_AND_LOW_HS_ATTAINMENT":                node.GTET_THE_90_PERCENTILE_HOUSEHOLDS_IN_LINGUISTIC_ISOLATION_AND_LOW_HS_ATTAINMENT,
		"PERCENTILE_LINGUISTIC_ISOLATION":                                                                node.PERCENTILE_LINGUISTIC_ISOLATION,
		"PERCENT_LINGUISTIC_ISOLATION":                                                                   node.PERCENT_LINGUISTIC_ISOLATION,
		"GTET_90_PERCENTILE_UNEMPLOYMENT_AND_LOW_HS":                                                     node.GTET_90_PERCENTILE_UNEMPLOYMENT_AND_LOW_HS,
		"PERCENTILE_UNEMPLOYMENT":                                                                        node.PERCENTILE_UNEMPLOYMENT,
		"PERCENT_UNEMPLOYMENT":                                                                           node.PERCENT_UNEMPLOYMENT,
		"GTET_90_PERCENTILE_HOUSEHOLDS_BELOW_100_PERCENT_FEDERAL_POVERTY_LEVEL_AND_LOW_HS_ATTAINMENT":    node.GTET_90_PERCENTILE_HOUSEHOLDS_BELOW_100_PERCENT_FEDERAL_POVERTY_LEVEL_AND_LOW_HS_ATTAINMENT,
		"PERCENTILE_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE_TWO":                                          node.PERCENTILE_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE_TWO,
		"PERCENT_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE_TWO":                                             node.PERCENT_BELOW_200_PERCENT_FEDERAL_POVERTY_LINE_TWO,
		"PERCENTILE_BELOW_100_PERCENT_FEDERAL_POVERTY_LINE":                                              node.PERCENTILE_BELOW_100_PERCENT_FEDERAL_POVERTY_LINE,
		"PERCENT_BELOW_100_PERCENT_FEDERAL_POVERTY_LINE":                                                 node.PERCENT_BELOW_100_PERCENT_FEDERAL_POVERTY_LINE,
		"PERCENTILE_AGE_25_PLUS_LESS_THAN_HIGH_SCHOOL":                                                   node.PERCENTILE_AGE_25_PLUS_LESS_THAN_HIGH_SCHOOL,
		"PERCENT_AGE_25_PLUS_LESS_THAN_HIGH_SCHOOL":                                                      node.PERCENT_AGE_25_PLUS_LESS_THAN_HIGH_SCHOOL,
		"PERCENT_RESIDENTS_NOT_CURRENTLY_ENROLLED_IN_HIGHER_ED":                                          node.PERCENT_RESIDENTS_NOT_CURRENTLY_ENROLLED_IN_HIGHER_ED,
		"UNEMPLOYMENT_PERCENT_ISLAND_AREAS_AND_STATES_AND_PR":                                            node.UNEMPLOYMENT_PERCENT_ISLAND_AREAS_AND_STATES_AND_PR,
		"PERCENTAGE_HOUSEHOLDS_BELOW_100_PERCENT_OF_FEDERAL_POVERTY_LINE_ISLAND_AREAS_AND_STATES_AND_PR": node.PERCENTAGE_HOUSEHOLDS_BELOW_100_PERCENT_OF_FEDERAL_POVERTY_LINE_ISLAND_AREAS_AND_STATES_AND_PR,
		"GTET_90_PERCENTILE_FOR_UNEMPLOYMENT_AND_LOW_HS_EDUCATION_ISLAND_AREAS":                          node.GTET_90_PERCENTILE_FOR_UNEMPLOYMENT_AND_LOW_HS_EDUCATION_ISLAND_AREAS,
		"GTET_90_PERCENTILE_BELOW_100_PERCENT_FEDERAL_POVERTY_LEVEL_AND_LOW_HS_EDUCATION_ISLAND_AREAS":   node.GTET_90_PERCENTILE_BELOW_100_PERCENT_FEDERAL_POVERTY_LEVEL_AND_LOW_HS_EDUCATION_ISLAND_AREAS,
		"GTET_90_PERCENTILE_LOW_AREA_MEDIAN_HOUSEHOLD_INCOME_AND_LOW_HS_EDUCATION_ISLAND_AREAS":          node.GTET_90_PERCENTILE_LOW_AREA_MEDIAN_HOUSEHOLD_INCOME_AND_LOW_HS_EDUCATION_ISLAND_AREAS,
		"NUMBER_OF_TRIBAL_AREAS_ALASKA":                                                                  node.NUMBER_OF_TRIBAL_AREAS_ALASKA,
		"NAMES_TRIBAL_AREAS":                                                                             node.NAMES_TRIBAL_AREAS,
		"PERCENT_WITHIN_TRIBAL_AREAS":                                                                    node.PERCENT_WITHIN_TRIBAL_AREAS,
	}
	return testData
}

func CreateDataNodes(filepath string, session neo4j.SessionWithContext, ctx context.Context) {
	testData, _ := loadData(filepath)
	for _, d := range testData {
		//dat := d
		if err := graph.CreateNode(allData(&d), "TestData", session, ctx); err != nil {
			log.Printf("err = %+v\n", err)
		}
	}
}
