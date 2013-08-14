package fatman

import "time"

type Account struct {
	Id        int
	Name      string
	LastLogin time.Time
}