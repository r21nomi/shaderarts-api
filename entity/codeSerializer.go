package entity

import "github.com/r21nomi/shaderarts-api/datastore"

type CodeSerializer struct {
	datastore.Code
}

type CodesSerializer struct {
	Codes []datastore.Code
}

type CodeEntity struct {
	ID    string `json:"id"`
	Type  int    `json:"type"`
	Text  string `json:"text"`
	ArtID string `json:"artId"`
}

func (self *CodeSerializer) Entity() CodeEntity {
	return CodeEntity{
		ID:    self.ID,
		Type:  self.Type,
		Text:  self.Text,
		ArtID: self.ArtID,
	}
}

func (self *CodesSerializer) Entities() []CodeEntity {
	entities := []CodeEntity{}
	for _, code := range self.Codes {
		serializer := CodeSerializer{code}
		entities = append(entities, serializer.Entity())
	}
	return entities
}
