package tracker

import (
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

		assert.Equal(
			t,
			[]Item{item},
			tracker.GetItems(),
		)
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

		assert.Equal(
			t,
			2,
			len(tracker.GetItems()),
		)
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

		assert.Equal(
			t,
			exp,
			rsl,
		)
	})

}
