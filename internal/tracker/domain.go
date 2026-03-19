package tracker

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

type Item struct {
	ID   string
	Name string
}

func (i Item) ToString() string {
	return fmt.Sprintf("id: %s, name: %s", i.ID, i.Name)
}

type Tracker struct {
	items []Item
}

func NewTracker() *Tracker {
	return &Tracker{}
}

func (t *Tracker) UpdateItem(item Item) error {
	index, ok := t.indexOf(item.ID)
	if !ok {
		return ErrNotFound
	}

	t.items[index] = item
	return nil
}

func (t *Tracker) indexOf(id string) (int, bool) {
	for i, item := range t.items {
		if item.ID == id {
			return i, true
		}
	}
	return -1, false
}

func (t *Tracker) AddItem(item Item) (Item, error) {
	var res Item

	_, ok := t.indexOf(item.ID)
	if ok {
		return Item{}, ErrAlreadyExists
	}

	item.ID = uuid.New().String()
	t.items = append(t.items, item)
	res = item

	return res, nil
}

// GetItems - important! return COPY(not origin) []Items using func copy(dest, resource)
func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)

	return res
}

// FindByPrefixName - поиск заявки по частичному совпадению имени.
func (t *Tracker) FindByPrefixName(name string) (Item, bool) {
	ok := false
	itm := Item{}

	for _, item := range t.GetItems() {
		if item.Name == name {
			ok = true
			itm = item
			break
		} else {
			ok = strings.HasPrefix(item.Name, name)
			if ok {
				itm = item
				break
			}
		}
	}

	return itm, ok
}

func (t *Tracker) DeleteItem(id string) error {
	i, ok := t.indexOf(id)
	if !ok {
		return ErrItemNotFound
	}

	t.items[i] = Item{}
	t.items = append(t.items[:i], t.items[i+1:]...) // Удалить, но сохранить порядок

	return nil
}
