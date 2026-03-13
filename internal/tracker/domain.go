package tracker

import (
	"fmt"
	"github.com/google/uuid"
	"unicode"
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
	var itemResult Item

	_, ok := t.indexOf(item.ID)
	if ok {
		return Item{}, ErrIllegalArgument
	}

	item.ID = uuid.New().String()
	t.items = append(t.items, item)
	itemResult = item

	return itemResult, nil
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
	itemRsl := Item{}

	for _, item := range t.GetItems() {

		if item.Name == name {
			ok = true
			itemRsl = item
			break
		} else {
			itemRsl, ok = t.compareByPrefix(name, item)

			if ok {
				return itemRsl, ok
			}
		}
	}

	return itemRsl, ok
}

func (t *Tracker) DeleteItem(id string) error {

	i, ok := t.indexOf(id)
	if !ok {
		return ErrIllegalArgument
	}

	t.items[i] = Item{}
	t.items = append(t.items[:i], t.items[i+1:]...) // Удалить, но сохранить порядок

	return nil
}

/*
compareByPrefix - private func, сравнение имени по префику(первые символы в имени)
*/
func (t *Tracker) compareByPrefix(name string, item Item) (Item, bool) {
	ok := false
	index := 0
	coutPrefix := 0
	matchCounter := 0

	// Конвертируем строку в слайс рун (символов)
	runesItemName := []rune(item.Name)

	for _, r := range name {
		fmt.Println(string(r)) // Выводит каждый символ отдельно

		if len(runesItemName) == 0 || unicode.IsSpace(r) {
			return Item{}, false
		}

		if string(r) == string(runesItemName[index]) {
			coutPrefix++
			index++
			matchCounter++
		} else {
			coutPrefix++
			index++
		}
		if coutPrefix == 3 && matchCounter == 3 {
			ok = true
			break
		}
	}
	return item, ok
}
