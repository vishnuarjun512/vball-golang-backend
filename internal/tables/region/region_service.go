package region

import (
	"context"
)

func GetRegions_Service() ([]Region, error) {
	return GetAllRegions_Repo(context.Background())
}

func CreateRegion_Service(regionName string, regionCode string) (string, error) {
	return CreateRegion_Repo(context.Background(), regionName, regionCode)
}

func UpdateRegion_Service(regionName string, regionId string) error {
	return UpdateRegion_Repo(context.Background(), regionId, regionName)
}

func DeleteRegion_Service(regionId string) error {
	return DeleteRegion_Repo(context.Background(), regionId)
}

func GetRegion_Service(regionId string) (Region, error) {
	return GetRegion_Repo(context.Background(), regionId)
}
