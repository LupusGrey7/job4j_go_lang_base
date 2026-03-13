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

func (t *Tracker) UpdateItem(item Item) bool {
	idx := t.indexOf(item.ID)
	if idx == -1 {
		return false
	}

	t.items[idx] = item
	return true
}

func (t *Tracker) indexOf(id string) int {
	for i, item := range t.items {
		if item.ID == id {
			return i
		}
	}
	return -1
}

func (t *Tracker) AddItem(item Item) (Item, error) {
	var itemResult Item

	idx := t.indexOf(item.ID)
	if idx != -1 {
		return Item{}, ErrIllegalArgument
	}

	item.ID = uuid.New().String()
	t.items = append(t.items, item)
	itemResult = item

	return itemResult, nil
}

// GetItems - important! return COPY []Items using func copy(dest, resource)
func (t *Tracker) GetItems() []Item {
	res := make([]Item, len(t.items))
	copy(res, t.items) //The copy built-in function copies elements from a source slice into a destination slice (in this case is `res`).

	return res
}

// FindByPrefixName - поиск заявки по частичному совпадению имени.
func (t *Tracker) FindByPrefixName(name string) (Item, bool) {
	flag := false
	itemRsl := Item{}

	for _, item := range t.GetItems() {

		if item.Name == name {
			flag = true
			itemRsl = item
			break
		} else {
			itemRsl, flag = t.compareByPrefix(name, item)

			if flag {
				return itemRsl, flag
			}
		}
	}

	return itemRsl, flag
}

func (t *Tracker) DeleteItem(id string) error {

	idx := t.indexOf(id)
	if idx == -1 {
		return ErrIllegalArgument
	}

	t.items[idx] = Item{}                               // Удаляем элемент с индексом i
	t.items = append(t.items[:idx], t.items[idx+1:]...) // Удалить, но сохранить порядок

	return nil
}

/*
compareByPrefix - private func, сравнение имени по префику(первые символы в имени)
*/
func (t *Tracker) compareByPrefix(name string, item Item) (Item, bool) {
	flag := false
	index := 0
	coutPrefix := 0
	matchCounter := 0

	// Конвертируем строку в слайс рун (символов)
	runesItemName := []rune(item.Name)

	for _, r := range name { // ПРЯМО по string!
		fmt.Println(string(r)) // Выводит каждый символ отдельно

		if len(runesItemName) == 0 || unicode.IsSpace(r) {
			return Item{}, false
		}

		if string(r) == string(runesItemName[index]) { // что если длина имени и Item имени не совпала
			coutPrefix++
			index++
			matchCounter++
		} else {
			coutPrefix++
			index++
		}
		if coutPrefix == 3 && matchCounter == 3 {
			flag = true
			break
		}
	}
	return item, flag
}
