package channelCategoryService

import (
	"adtec/backend/src/prisma/channel/db"
	"adtec/backend/src/utils/database"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"
)

// TODO: CreateMany Validation

type ChannelCategories struct {
	ChannelId string `json:"channelId" validate:"required"`
	Category  string `json:"category" validate:"required"`
}

func CreateMany(channelCategories []ChannelCategories) ([]*db.ChannelCategoriesModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	tableName := "channelCategories"

	query := ChannelCategoryCreateManyRawQuery(channelCategories, tableName)

	if query == nil {
		log.Println("error generating the need query at channel category")
		return nil, errors.New("internal server error")
	}

	var res []*db.ChannelCategoriesModel

	err = client.Prisma.QueryRaw(*query).Exec(ctx, &res)

	if err != nil {
		log.Println(err, "create many, category channel")
		return nil, errors.New("error creating data")
	}
	return res, nil

}

func ChannelCategoryCreateManyRawQuery(channelCategories []ChannelCategories, tableName string) *string {

	if len(channelCategories) > 0 {
		var columns []string
		t := reflect.TypeOf(channelCategories[0])
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			columns = append(columns, field.Tag.Get("json"))
		}

		columnStrings := make([]string, 0, len(columns))
		// Iterate over employees to create row strings
		for _, column := range columns {
			row := fmt.Sprintf("\"%s\"", column)
			columnStrings = append(columnStrings, row)
		}

		// Start the INSERT statement
		query := "INSERT INTO " + "\"" + tableName + "\"" + " (" + strings.Join(columnStrings, ", ") + ") VALUES "

		// Slice to hold value strings
		valueStrings := make([]string, 0, len(channelCategories))

		// Iterate over employees to create row strings
		for _, channelCategory := range channelCategories {
			row := fmt.Sprintf("('%s', '%s')", channelCategory.ChannelId, channelCategory.Category)
			valueStrings = append(valueStrings, row)
		}

		// Concatenate the value strings
		query += strings.Join(valueStrings, ", ")
		log.Println(query)
		return &query
	}
	return nil
}
