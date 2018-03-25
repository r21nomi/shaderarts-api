package entity

import "github.com/r21nomi/arto-api/datastore"

type TagSerializer struct {
	datastore.Tag
}

type TagsSerializer struct {
	tags []datastore.Tag
}

type TagEntity struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

func (self *TagSerializer) Entity() TagEntity {
	return TagEntity{
		ID:   self.ID,
		Text: self.Text,
	}
}

func (self *TagsSerializer) Entities() []TagEntity {
	entities := []TagEntity{}
	for _, tag := range self.tags {
		serializer := TagSerializer{tag}
		entities = append(entities, serializer.Entity())
	}
	return entities
}
