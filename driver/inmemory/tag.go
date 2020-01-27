package inmemory

import (
	"fmt"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
)

var tagStorage = map[string]bool{
	asTagID("1", "gPlake"):   true,
	asTagID("2", "sSchmidt"): true,
	asTagID("2", "mHattrup"): true,
	asTagID("2", "gPlake"):   true,
}

func existsTag(photoID model.PhotoID, userID model.UserID) bool {
	return tagStorage[asTagID(photoID, userID)]
}

func asTagID(photoID model.PhotoID, userID model.UserID) string {
	return fmt.Sprintf("%v_%v", photoID, userID)
}
