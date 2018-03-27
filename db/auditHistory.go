package db

const AuctionHistoryColl = "auditHistory"

type AuctionHistory struct {
	UserID     string `bson:"user_id"`
	UserName string `json:"user_name"`
	TeamID     string `bson:"team_id"`
	TeamName string `json:"team_name"`
	TeamPoints int `bson:"team_points"`
	UserPoints int `bson:"user_points"`
	Action string `bson:"action"`
	At         int64 `bson:"at"`
}

func InsertAuctionHistory(ah *AuctionHistory) error {
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(AuctionHistoryColl)

	err := c.Insert(ah)
	if err != nil {
		return err
	}
	return nil
}

func FetchAuctionHistory() (*[]AuctionHistory, error){
	s := GetSession()
	defer s.Close()
	c := s.DB(DB).C(AuctionHistoryColl)

	var auctionHistory []AuctionHistory
	err := c.Find(nil).Sort("-uat").Limit(5).All(&auctionHistory)
	if err != nil {
		return &[]AuctionHistory{}, err
	}
	return &auctionHistory, nil
}