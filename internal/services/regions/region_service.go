package region

import (
	"context"
	"vball/internal/models"
	"vball/internal/repositories/region"
)

func GetRegions_Service() ([]models.Region, error) {
	return region.GetAllRegions_Repo(context.Background())
}

func CreateRegion_Service(regionName string) (string, error) {
	return region.CreateRegion_Repo(context.Background(), regionName)
}

func UpdateRegion_Service(regionName string, regionId string) error {
	return region.UpdateRegion_Repo(context.Background(), regionId, regionName)
}

func DeleteRegion_Service(regionId string) error {
	return region.DeleteRegion_Repo(context.Background(), regionId)
}
