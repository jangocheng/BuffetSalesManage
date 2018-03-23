package businesses_account_logic

import (
	"gopkg.in/mgo.v2"
	"BuffetSalesManage/BuffetSalesManage/config"
	"BuffetSalesManage/BuffetSalesManage/model/businesses.account.model"
	"gopkg.in/mgo.v2/bson"
)

func IsExists(session *mgo.Session, accountName string) bool {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	count, _ := coll.Find(bson.M{businesses_account_model.AccountName.String(): accountName}).Count()
	if count != 0 {
		return false
	}
	return true
}

func RegisterBusinesses(session *mgo.Session, accountName, password string) error {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	selector := bson.M{businesses_account_model.AccountName.String(): accountName}
	update := bson.M{businesses_account_model.Password.String(): password}
	_, err := coll.Upsert(selector, update)
	return err
}

func IsLogin(session *mgo.Session, accountName, password string) bool {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)

	selector := bson.M{businesses_account_model.AccountName.String(): accountName}
	var businessesInfo businesses_account_model.BusinessesAccount
	coll.Find(selector).One(&businessesInfo)
	if businessesInfo.Password != password {
		return false
	}
	return true
}

func ComplementInfo(session *mgo.Session, businessesInfo businesses_account_model.BusinessesAccount) error {
	coll := session.DB(config.MongoDBName).C(businesses_account_model.COLL_BUSINESSES_ACCOUNT)
	selector := bson.M{businesses_account_model.AccountName.String(): businessesInfo.AccountName}
	update := bson.M{
		"$set": bson.M{
			businesses_account_model.NameOfShop.String(): businessesInfo.NameOfShop,
			businesses_account_model.Location.String():   businessesInfo.Location,
			businesses_account_model.Street.String():     businessesInfo.Street,
		},
	}

	err := coll.Update(selector, update)
	return err
}
