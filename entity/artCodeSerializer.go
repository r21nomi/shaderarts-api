package entity

import "github.com/r21nomi/shaderarts-api/datastore"

type ArtCodeSerializer struct {
	datastore.Art
}

type ArtCodeEntity struct {
	ID    string       `json:"id"`
	Codes []CodeEntity `json:"codes"`
}

func (self *ArtCodeSerializer) Entity() ArtCodeEntity  {
	codesSerializer := CodesSerializer{self.Codes}

	return ArtCodeEntity{
		ID:    self.ID,
		Codes: codesSerializer.Entities(),
	}

}