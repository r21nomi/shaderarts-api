package entity

import "time"
import "github.com/r21nomi/arto-api/datastore"

type ArtSerializer struct {
	datastore.Art
}

type ArtsSerializer struct {
	Arts []datastore.Art
}

type ArtEntity struct {
	ID          string       `json:"id"`
	Title       string       `json:"title"`
	Type        int          `json:"type"`
	Thumb       string       `json:"thumb"`
	Description string       `json:"description"`
	Star        int          `json:"star"`
	IsStarred   bool         `json:"isStarred"`
	CreatedAt   time.Time    `json:"createdAt"`
	UpdatedAt   time.Time    `json:"updatedAt"`
	User        UserEntity   `json:"user"`
	Codes       []CodeEntity `json:"codes"`
	Tags        []TagEntity  `json:"tags"`
}

func (self *ArtSerializer) Entity() ArtEntity {
	userSerializer := UserSerializer{self.User}
	codesSerializer := CodesSerializer{self.Codes}
	tagsSerializer := TagsSerializer{self.Tags}

	return ArtEntity{
		ID:          self.ID,
		Title:       self.Title,
		Type:        self.Type,
		Thumb:       self.Thumb,
		Description: self.Description,
		Star:        self.Star,
		IsStarred:   false, // FIXME
		CreatedAt:   self.CreatedAt,
		UpdatedAt:   self.UpdatedAt,
		User:        userSerializer.Entity(),
		Codes:       codesSerializer.Entities(),
		Tags:        tagsSerializer.Entities(),
	}
}

func (self *ArtsSerializer) Entities() []ArtEntity {
	entities := []ArtEntity{}
	for _, art := range self.Arts {
		serializer := ArtSerializer{art}
		entities = append(entities, serializer.Entity())
	}
	return entities
}
