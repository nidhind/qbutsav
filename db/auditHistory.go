package db

const AuctionHistoryColl = "auditHistory"

type AuctionHistory struct {
	UserID     string `bson:"user_id"`
	TeamID     string `bson:"team_id"`
	TeamPoints int `bson:"team_points"`
	UserPoints int `bson:"user_points"`
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
