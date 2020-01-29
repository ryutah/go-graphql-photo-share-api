package model

type Tag struct {
	PhotoID PhotoID
	UserID  UserID
}

type TagSlice []*Tag

func (t TagSlice) UserIDs() []UserID {
	var (
		resultsMap = make(map[UserID]bool)
		results    = make([]UserID, 0, len(t))
	)

	for _, tag := range t {
		if resultsMap[tag.UserID] {
			continue
		}
		resultsMap[tag.UserID] = true
		results = append(results, tag.UserID)
	}
	return results
}

func (t TagSlice) PhotoIDs() []PhotoID {
	var (
		resultsMap = make(map[PhotoID]bool)
		results    = make([]PhotoID, 0, len(t))
	)

	for _, tag := range t {
		if resultsMap[tag.PhotoID] {
			continue
		}
		resultsMap[tag.PhotoID] = true
		results = append(results, tag.PhotoID)
	}
	return results
}
