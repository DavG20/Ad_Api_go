package advertisementService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/enumConverter"
)

func GetAll(filter AdvertisementFilter) ([]db.AdvertisementsModel, error) {

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.AdvertisementsWhereParam{}

	if filter.Status != nil {
		parameter = append(parameter, db.Advertisements.Status.In(enumConverter.AdvertisementStatusConverter(filter.Status)))
	}

	if filter.ChannelID != nil {
		parameter = append(parameter, db.Advertisements.ChannelID.Equals(*filter.ChannelID))
	}

	if filter.CompanyID != nil {
		parameter = append(parameter, db.Advertisements.Campaign.Where(db.Campaigns.CompanyID.Equals(*filter.CompanyID)))
	}

	if filter.Rate != nil {
		if filter.Rate.Max != nil {
			parameter = append(parameter, db.Advertisements.AdvertisementCpms.Where(db.AdvertisementCpms.Rate.Lte(*filter.Rate.Max)))
		}
		parameter = append(parameter, db.Advertisements.AdvertisementCpms.Where(db.AdvertisementCpms.Rate.Gte(filter.Rate.Min)))

	}
	if filter.ChannelBudgetShare != nil {
		if filter.ChannelBudgetShare.Max != nil {
			parameter = append(parameter, db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.ChannelBudgetShare.Lte(*filter.ChannelBudgetShare.Max)))
		}
		parameter = append(parameter, db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.ChannelBudgetShare.Gte(filter.ChannelBudgetShare.Min)))

	}
	if filter.RequiredViews != nil {
		if filter.RequiredViews.Max != nil {
			parameter = append(parameter, db.Advertisements.AdvertisementCpms.Where(db.AdvertisementCpms.RequiredViews.Lte(int(*filter.RequiredViews.Max))))
		}
		parameter = append(parameter, db.Advertisements.AdvertisementCpms.Where(db.AdvertisementCpms.RequiredViews.Gte(int(filter.RequiredViews.Min))))

	}
	if filter.Hours != nil {
		if filter.Hours.Max != nil {
			parameter = append(parameter, db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.TotalHour.Lte(int(*filter.Hours.Max))))
		}
		parameter = append(parameter, db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.TotalHour.Gte(int(filter.Hours.Min))))

	}
	if filter.Views != nil {
		if filter.Views.Max != nil {
			parameter = append(parameter, db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.TotalViews.Lte(int(*filter.Views.Max))))
		}
		parameter = append(parameter, db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.TotalViews.Gte(int(filter.Views.Min))))

	}

	fetchAdvertisements := client.Advertisements.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.Advertisements.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchAdvertisements = client.Advertisements.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Advertisements.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.Advertisements.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchAdvertisements = client.Advertisements.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Advertisements.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.Advertisements.CreatedAt.Order(db.DESC),
			)
		}
	}
	advertisement, err := fetchAdvertisements.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return advertisement, nil

}
