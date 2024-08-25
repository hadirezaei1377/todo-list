package memorystore

import (
	"todo-list/entity"
)

type Category struct {
	categories []entity.Category
}

func (c *Category) DoesThisUserHaveThisCategoryID(userID, categoryID int) bool {
	isFound := false
	for _, c := range c.categories {
		if c.ID == categoryID && c.UserID == userID {
			isFound = true

			break
		}
	}

	return isFound
}
