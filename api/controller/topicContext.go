package controller

import "RPLL/api/model"

type SearchContext struct {
	strategy SearchStrategy
}

func (sc *SearchContext) SetStrategy(strategy SearchStrategy) {
	sc.strategy = strategy
}

func (sc *SearchContext) PerformSearch(query string) []model.Topic {
	return sc.strategy.Search(query)
}
