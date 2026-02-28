package tracker

import "fmt"

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

func (t *Tracker) AddItem(item Item) {
	t.items = append(t.items, item)
}

// GetItems - important! return COPY []Items
func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items)
	return res
}
