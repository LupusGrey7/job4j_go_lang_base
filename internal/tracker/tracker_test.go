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
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		res[0].Name = "Second Item"

		assert.Equal(t, []Item{item}, tracker.GetItems())
	})

	t.Run("when add Item`s Than get size []Items", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item1 := Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item1)
		item2 := Item{
			ID:   "2",
			Name: "Second Item",
		}
		tracker.AddItem(item2)

		rsl := tracker.GetItems()
		rsl[0].Name = "Second Item"

		assert.Equal(t, 2, len(tracker.GetItems()))
	})

	t.Run("when add Item Than use ToString", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "1",
			Name: "First Item",
		}
		tracker.AddItem(item)

		res := tracker.GetItems()
		rsl := res[0].ToString()

		exp := "id: 1, name: First Item"

		assert.Equal(t, exp, rsl)
	})

	t.Run("when item name is equals Than true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   uuid.NewString(),
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   uuid.NewString(),
			Name: "Boris Johnson",
		}
		tracker.AddItem(item)
		tracker.AddItem(itemTwo)
		_, flag := tracker.FindByPrefixName("First Item")

		assert.Equal(t, flag, true)
	})

	t.Run("when find Item BY prefix Than true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   uuid.NewString(),
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   uuid.NewString(),
			Name: "Boris Johnson",
		}
		tracker.AddItem(item)
		tracker.AddItem(itemTwo)
		_, flag := tracker.FindByPrefixName("Boris Johnson")

		assert.Equal(t, flag, true)
	})

	t.Run("when find Item BY prefix Than false", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   uuid.NewString(),
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   uuid.NewString(),
			Name: "Boris Johnson",
		}
		tracker.AddItem(item)
		tracker.AddItem(itemTwo)
		_, flag := tracker.FindByPrefixName("Bo Johnson")

		assert.Equal(t, flag, false)
	})

	t.Run("when update Item Than true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   uuid.NewString(),
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   uuid.NewString(),
			Name: "John Johnson",
		}

		itemUpdate := Item{
			ID:   itemTwo.ID,
			Name: "Donald Trump",
		}
		tracker.AddItem(item)
		tracker.AddItem(itemTwo)
		rsl, flag := tracker.UpdateItem(itemUpdate)

		assert.Equal(t, flag, true)
		assert.Equal(t, rsl.Name, "Donald Trump")
	})

	t.Run("when update Item Than false", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   uuid.NewString(),
			Name: "First Item",
		}
		itemTwo := Item{
			ID:   uuid.NewString(),
			Name: "John Johnson",
		}

		itemUpdate := Item{
			ID:   uuid.NewString(),
			Name: "Martin Scorsese",
		}
		tracker.AddItem(item)
		tracker.AddItem(itemTwo)
		rsl, flag := tracker.UpdateItem(itemUpdate)

		assert.Equal(t, flag, false)
		assert.Equal(t, rsl.Name, "Martin Scorsese")
	})

	t.Run("when delete Item BY ID Than true", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   uuid.NewString(),
			Name: "Alice Marcuse",
		}
		itemTwo := Item{
			ID:   uuid.NewString(),
			Name: "Boris Johnson",
		}
		tracker.AddItem(item)
		tracker.AddItem(itemTwo)
		tracker.DeleteItem(itemTwo.ID)
		rslItems := tracker.GetItems()

		assert.Equal(t, 1, len(rslItems))
		assert.Equal(t, "Alice Marcuse", rslItems[0].Name)
		assert.Equal(t, item.ID, rslItems[0].ID)

	})
}
