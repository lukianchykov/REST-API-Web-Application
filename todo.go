package todo

import "errors"

/**
 * Todo Model
 */

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	UpdateDone  bool   `json:"update_done" db:"update_done"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	UpdateDone  bool   `json:"update_done" db:"update_done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	UpdateDone  *bool   `json:"update_done"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	UpdateDone  *bool   `json:"update_done"`
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.UpdateDone == nil {
		return errors.New("update structure has no values to be updated")
	}

	return nil
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.UpdateDone == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
