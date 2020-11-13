package store

// LocationHistoryStore is an interface to work with mongodb.
type LocationHistoryStore interface {
	Insert(*LocationHistory, string) (interface{}, error)
	Update(*LocationHistory, string) error
}
