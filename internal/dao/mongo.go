package dao

import (
	"DiscordRolesBot/internal/model"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// unmarshal result directly
func MongoAggregateUnmarshal(model mgm.Model, pipeLine interface{}, result interface{}) error {
	c, err := mgm.Coll(model).Aggregate(mgm.Ctx(), pipeLine)
	if err != nil {
		return err
	}

	err = c.All(mgm.Ctx(), result)
	if err != nil {
		return err
	}
	return nil
}

func GetAllNFTNum(accounts []string) (data map[int]int64, err error) {
	data = make(map[int]int64, 2)
	for _, v := range accounts {
		pipeline := mongo.Pipeline{
			bson.D{{"$match", bson.D{{"owner_address", v}}}},
			bson.D{{"$group", bson.D{{"_id", "$nft_type"}, {"count", bson.D{{"$sum", 1}}}}}},
		}
		type result struct {
			ID    int64 `json:"id" bson:"_id,omitempty"`
			Count int64 `json:"count"`
		}
		var quantityResult []result
		err = MongoAggregateUnmarshal(&model.NFT{}, pipeline, &quantityResult)
		if err != nil {
			return data, err
		}

		for _, v := range quantityResult {
			switch v.ID {
			case 1:
				data[1] += v.Count
			case 2:
				data[2] += v.Count
			}
		}
	}

	return data, err
}
