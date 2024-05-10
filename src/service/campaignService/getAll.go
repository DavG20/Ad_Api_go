package campaignService

import (
	"adtec/backend/src/prisma/campaign/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
	"adtec/backend/src/utils/enumConverter"
)

func GetAll(filter CampaignFilter) ([]db.CampaignsModel, error) {

	client, ctx, err := database.GetCampaignClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	parameter := []db.CampaignsWhereParam{}

	if filter.Name != nil {
		parameter = append(parameter, db.Campaigns.Or(
			db.Campaigns.Name.Contains(*filter.Name),
			db.Campaigns.Name.Mode(db.QueryModeInsensitive),
		))
	}
	if filter.Objective != nil {
		parameter = append(parameter, db.Campaigns.Or(
			db.Campaigns.Objective.In(enumConverter.ObjectiveTypeConverter(filter.Objective))),
		)
	}

	if filter.Languages != nil {
		parameter = append(parameter, db.Campaigns.Audiences.Where(
			db.Audiences.Language.In(filter.Languages),
		))
	}
	if filter.Categories != nil {
		parameter = append(parameter, db.Campaigns.Audiences.Where(
			db.Audiences.Category.In(filter.Categories),
		))
	}

	if filter.InitialBudget != nil {
		if filter.InitialBudget.Max != nil {
			parameter = append(parameter, db.Campaigns.Budgets.Where(db.Budgets.InitialAmount.Lte(*filter.InitialBudget.Max)))
		}
		parameter = append(parameter, db.Campaigns.Budgets.Where(db.Budgets.InitialAmount.Gte(filter.InitialBudget.Min)))

	}
	if filter.UsedAmount != nil {
		if filter.UsedAmount.Max != nil {
			parameter = append(parameter, db.Campaigns.Budgets.Where(db.Budgets.UsedAmount.Lte(*filter.UsedAmount.Max)))
		}
		parameter = append(parameter, db.Campaigns.Budgets.Where(db.Budgets.UsedAmount.Gte(filter.UsedAmount.Min)))

	}

	if filter.ContentType != nil {
		parameter = append(parameter, db.Campaigns.Or(
			db.Campaigns.Contents.Where(db.Contents.ContentType.In(enumConverter.ContentTypeConverter(filter.ContentType)))),
		)
	}
	
	if filter.AdvertisementStatus != nil {
		parameter = append(parameter, db.Campaigns.Advertisements.Some(
			db.Advertisements.Status.In(enumConverter.AdvertisementStatusConverter(filter.AdvertisementStatus))),
		)

	}

	if filter.AdvertisementRate != nil {
		if filter.AdvertisementRate.Max != nil {
			parameter = append(parameter, db.Campaigns.Advertisements.Some(db.Advertisements.AdvertisementCpms.Where(db.AdvertisementCpms.Rate.Lte(*filter.AdvertisementRate.Max))))
		}
		parameter = append(parameter, db.Campaigns.Advertisements.Some(db.Advertisements.AdvertisementCpms.Where(db.AdvertisementCpms.Rate.Gte(filter.AdvertisementRate.Min))))

	}
	if filter.AdvertisementRequiredViews != nil {
		if filter.AdvertisementRequiredViews.Max != nil {
			parameter = append(parameter, db.Campaigns.Advertisements.Some(db.Advertisements.AdvertisementCpms.Where(db.AdvertisementCpms.RequiredViews.Lte(int(*filter.AdvertisementRequiredViews.Max)))))
		}
		parameter = append(parameter, db.Campaigns.Advertisements.Some(db.Advertisements.AdvertisementCpms.Where(db.AdvertisementCpms.RequiredViews.Gte(int(filter.AdvertisementRequiredViews.Min)))))

	}

	if filter.AdvertisementViews != nil {
		if filter.AdvertisementViews.Max != nil {
			parameter = append(parameter, db.Campaigns.Advertisements.Some(db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.TotalViews.Lte(int(*filter.AdvertisementViews.Max)))))
		}
		parameter = append(parameter, db.Campaigns.Advertisements.Some(db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.TotalViews.Gte(int(filter.AdvertisementViews.Min)))))

	}

	if filter.AdvertisementHours != nil {
		if filter.AdvertisementHours.Max != nil {
			parameter = append(parameter, db.Campaigns.Advertisements.Some(db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.TotalHour.Lte(int(*filter.AdvertisementHours.Max)))))
		}
		parameter = append(parameter, db.Campaigns.Advertisements.Some(db.Advertisements.AdvertisementResults.Where(db.AdvertisementResults.TotalHour.Gte(int(filter.AdvertisementHours.Min)))))

	}

	fetchCampaigns := client.Campaigns.FindMany(
		parameter[:]...,
	).Take(int(limit)).OrderBy(
		db.Campaigns.CreatedAt.Order(db.DESC),
	)
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchCampaigns = client.Campaigns.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Campaigns.ID.Cursor(string(*filter.Filter.After)),
			).OrderBy(
				db.Campaigns.CreatedAt.Order(db.DESC),
			)
		}
		if filter.Filter.Before != nil {
			fetchCampaigns = client.Campaigns.FindMany(
				parameter[:]...,
			).Skip(1).Take(int(limit)).Cursor(
				db.Campaigns.ID.Cursor(string(*filter.Filter.Before)),
			).OrderBy(
				db.Campaigns.CreatedAt.Order(db.DESC),
			)
		}
	}
	campaign, err := fetchCampaigns.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return campaign, nil

}
