package region

import (
	"context"
	"vball/internal/models"
)

func GetRegions_Service() ([]models.Region, error) {
	return GetAllRegions_Repo(context.Background())
}

func CreateRegion_Service(regionName string) (string, error) {
	return CreateRegion_Repo(context.Background(), regionName)
}

func UpdateRegion_Service(regionName string, regionId string) error {
	return UpdateRegion_Repo(context.Background(), regionId, regionName)
}

func DeleteRegion_Service(regionId string) error {
	return DeleteRegion_Repo(context.Background(), regionId)
}
