package channelDetailService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils"
	"adtec/backend/src/utils/database"
)

func GetAll(filter FindManyInput) ([]db.ChannelDetailsModel, error) {

	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	var limit int32 = 20

	if filter.Filter != nil && filter.Filter.Limit != nil {

		limit = utils.LimitSetter(int32(*filter.Filter.Limit))
	}

	fetchChannelDetail := client.ChannelDetails.FindMany().Take(int(limit))
	if filter.Filter != nil {
		if filter.Filter.After != nil {
			fetchChannelDetail = client.ChannelDetails.FindMany().Skip(1).Take(int(limit)).Cursor(
				db.ChannelDetails.ChannelID.Cursor(string(*filter.Filter.After)),
			)
		}
		if filter.Filter.Before != nil {
			fetchChannelDetail = client.ChannelDetails.FindMany().Skip(1).Take(int(limit)).Cursor(
				db.ChannelDetails.ChannelID.Cursor(string(*filter.Filter.Before)),
			)
		}
	}
	chennelDetail, err := fetchChannelDetail.Exec(ctx)
	if err != nil {
		return nil, err
	}
	return chennelDetail, nil

}
