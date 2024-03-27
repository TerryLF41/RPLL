package controller

import "RPLL/api/model"

type SearchStrategy interface {
	Search(query string) []model.Topic
}
