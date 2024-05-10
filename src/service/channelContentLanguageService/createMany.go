package channelContentLanguageService

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

type ChannelContentLanguages struct {
	ChannelID string `json:"channelId"`
	Language  string `json:"language"`
}

func CreateMany(channelContentLanguage []ChannelContentLanguages) ([]*db.ChannelContentLanguagesModel, error) {
	client, ctx, err := database.GetChannelClient()

	if err != nil {
		return nil, err
	}

	tableName := "channelContentLanguages"

	query := ChannelContentLanguageCreateManyRawQuery(channelContentLanguage, tableName)

	if query == nil {
		log.Println("error generating the need query at channel Language")
		return nil, errors.New("internal server error")
	}

	var res []*db.ChannelContentLanguagesModel

	err = client.Prisma.QueryRaw(*query).Exec(ctx, &res)

	if err != nil {
		log.Println(err, "create many, category channel")
		return nil, errors.New("error creating data")
	}
	return res, nil

}

func ChannelContentLanguageCreateManyRawQuery(channelContentLanguage []ChannelContentLanguages, tableName string) *string {

	if len(channelContentLanguage) > 0 {
		var columns []string
		t := reflect.TypeOf(channelContentLanguage[0])
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
		valueStrings := make([]string, 0, len(channelContentLanguage))

		// Iterate over employees to create row strings
		for _, channelLanguage := range channelContentLanguage {
			row := fmt.Sprintf("('%s', '%s')", channelLanguage.ChannelID, channelLanguage.Language)
			valueStrings = append(valueStrings, row)
		}

		// Concatenate the value strings
		query += strings.Join(valueStrings, ", ")

		return &query
	}
	return nil
}
