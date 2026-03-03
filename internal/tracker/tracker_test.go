package tracker

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("check link leak", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "",
			Name: "First Item",
		}
		itemRsl, err := tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t, []Item{itemRsl}, tracker.GetItems())
		assert.Equal(t, err, nil)
	})

	t.Run("when add Item`s Than get size []Items", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		itemOne := Item{
			ID:   "",
			Name: "First Item",
		}
		_, errOne := tracker.AddItem(itemOne)
		itemTwo := Item{
			ID:   "",
			Name: "Second Item",
		}
		_, errTwo := tracker.AddItem(itemTwo)

		rsl := tracker.GetItems()
		rsl[0].Name = "Second Item"

		assert.Equal(t, 2, len(tracker.GetItems()))
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)

	})

	t.Run("add Item - true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		itemOne := Item{
			ID:   "",
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   "",
			Name: "Second Item",
		}

		itemOneRsl, errOne := tracker.AddItem(itemOne)
		itemTwoRsl, errTwo := tracker.AddItem(itemTwo)

		assert.Equal(t, 2, len(tracker.GetItems()))
		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)
		assert.Equal(t, errTwo, nil)
	})

	t.Run("add exist Item - false", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		itemOne := Item{
			ID:   "",
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   "",
			Name: "Second Item",
		}

		itemOneRsl, errOne := tracker.AddItem(itemOne)
		itemTwoRsl, errTwo := tracker.AddItem(itemTwo)
		itemThree := Item{
			ID:   itemTwoRsl.ID,
			Name: "Third Item",
		}

		_, errThree := tracker.AddItem(itemThree)

		assert.Equal(t, 2, len(tracker.GetItems()))
		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)
		assert.Equal(t, errTwo, nil)
		assert.Equal(t, errThree, ErrIllegalArgument)
	})

	t.Run("when add Item Than use ToString", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "",
			Name: "First Item",
		}
		itemRsl, _ := tracker.AddItem(item)

		res := tracker.GetItems()
		rsl := res[0].ToString()
		idRsl := itemRsl.ID
		exp := "id: " + idRsl + ", name: First Item"

		assert.Equal(t, exp, rsl)
	})

	t.Run("when item name is equals Than true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		itemOne := Item{
			ID:   "",
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   "",
			Name: "Boris Johnson",
		}
		itemOneRsl, errOne := tracker.AddItem(itemOne)
		itemTwoRsl, errTwo := tracker.AddItem(itemTwo)
		_, flag := tracker.FindByPrefixName("First Item")

		assert.Equal(t, flag, true)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)
		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)
	})

	t.Run("when find Item BY prefix Than true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		itemOne := Item{
			ID:   uuid.NewString(),
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   uuid.NewString(),
			Name: "Boris Johnson",
		}
		itemOneRsl, errOne := tracker.AddItem(itemOne)
		itemTwoRsl, errTwo := tracker.AddItem(itemTwo)
		_, flag := tracker.FindByPrefixName("Boris Johnson")

		assert.Equal(t, flag, true)
		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)
	})

	t.Run("when find Item BY prefix Than false", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		itemOne := Item{
			ID:   "",
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   "",
			Name: "Boris Johnson",
		}
		itemOneRsl, errOne := tracker.AddItem(itemOne)
		itemTwoRsl, errTwo := tracker.AddItem(itemTwo)
		_, flag := tracker.FindByPrefixName("Bo Johnson")

		assert.Equal(t, flag, false)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)

		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)

	})

	t.Run("error update - not found", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}

		err := tracker.UpdateItem(item)
		assert.ErrorIs(t, err, ErrNotFound)
	})

	t.Run("update Item - true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "",
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   "",
			Name: "John Johnson",
		}
		itemOneRsl, errOne := tracker.AddItem(item)
		itemTwoRsl, errTwo := tracker.AddItem(itemTwo)

		itemUpdate := Item{
			ID:   itemTwoRsl.ID,
			Name: "Donald Trump",
		}

		errRsl := tracker.UpdateItem(itemUpdate)
		rsl := tracker.GetItems()

		assert.Equal(t, nil, errOne)
		assert.Equal(t, nil, errTwo)
		assert.Equal(t, item.Name, itemOneRsl.Name)
		assert.Equal(t, errRsl, nil)
		assert.Equal(t, rsl[1].Name, "Donald Trump")
	})

	t.Run("update Item - false", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		itemOne := Item{
			ID:   "",
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   "",
			Name: "John Johnson",
		}

		_, errOne := tracker.AddItem(itemOne)
		_, errTwo := tracker.AddItem(itemTwo)

		itemUpdate := Item{
			ID:   uuid.New().String(),
			Name: "Martin Scorsese",
		}
		errRsl := tracker.UpdateItem(itemUpdate)
		rsl := tracker.GetItems()

		assert.Equal(t, errRsl, ErrNotFound)
		assert.Equal(t, rsl[1].Name, "John Johnson")
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)

	})

	t.Run("delete Item BY ID - true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "",
			Name: "Alice Marcuse",
		}
		itemTwo := Item{
			ID:   "",
			Name: "Boris Johnson",
		}
		itemOneRsl, errOne := tracker.AddItem(item)
		itemTwoRsl, errTwo := tracker.AddItem(itemTwo)
		err := tracker.DeleteItem(itemTwoRsl.ID)
		rslItems := tracker.GetItems()

		assert.Equal(t, err, nil)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)
		assert.Equal(t, 1, len(rslItems))
		assert.Equal(t, "Alice Marcuse", rslItems[0].Name)
		assert.Equal(t, itemOneRsl.ID, rslItems[0].ID)
	})
}
