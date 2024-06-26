package controller

import (
	"RPLL/api/model"
	"log"
)

type LatestTopicSearchStrategy struct{}

func (rs *LatestTopicSearchStrategy) Search(query string) []model.Topic {
	db := connect()
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}

	var topic model.Topic
	var topicList []model.Topic
	for rows.Next() {
		if err := rows.Scan(
			&topic.TopicNo, &topic.TopicTitle, &topic.TopicDesc, &topic.CreateDate, &topic.BanStatus, &topic.TopicPicture, &topic.ThreadCount); err != nil {
			log.Println(err)
			return nil
		} else {
			topicList = append(topicList, topic)
		}
	}
	return topicList
}
