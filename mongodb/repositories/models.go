package mongodb

type Movie struct {
	ID         string   `bson:"_id"`
	Title      string   `bson:"title"`
	Genres     []string `bson:"genres"`
	Characters []string `bson:"characters"`
	Rating     float64  `bson:"rating"`
}
