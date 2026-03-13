package tracker

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Tracker(t *testing.T) {
	t.Parallel()

	t.Run("When_check_link_leak", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "",
			Name: "First Item",
		}
		itemRsl, err := tracker.AddItem(item)

		sliceItems := tracker.GetItems()

		assert.Equal(t, err, nil)
		assert.Equal(t, []Item{itemRsl}, sliceItems)
	})

	t.Run("When_add_item`s_Then_success", func(t *testing.T) {
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

		sliceItems := tracker.GetItems()

		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)
		assert.Equal(t, 2, len(sliceItems))
		assert.Equal(t, sliceItems[1].Name, "Second Item")
		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)
	})

	t.Run("When_try_add_exist_Item_Then_should_return_errIllegalArgument", func(t *testing.T) {
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

		assert.Equal(t, errTwo, nil)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, 2, len(tracker.GetItems()))
		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)
		assert.Equal(t, errThree, ErrIllegalArgument)
	})

	t.Run("When_add_Item_and_use_toString_Then_equals", func(t *testing.T) {
		t.Parallel()

		tracker := NewTracker()
		item := Item{
			ID:   "",
			Name: "First Item",
		}
		itemRsl, _ := tracker.AddItem(item)

		sliceItems := tracker.GetItems()
		rsl := sliceItems[0].ToString()
		itemID := itemRsl.ID
		exp := "id: " + itemID + ", name: First Item"

		assert.Equal(t, exp, rsl)
	})

	t.Run("When_name_prefix_matches_Then_return_true", func(t *testing.T) {
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
		_, ok := tracker.FindByPrefixName("Boris Johnson")

		assert.Equal(t, ok, true)
		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)
	})

	t.Run("When_name_prefix_misses_Then_returns_false", func(t *testing.T) {
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
		_, ok := tracker.FindByPrefixName("Bo Johnson")

		assert.Equal(t, ok, false)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)

		assert.Equal(t, itemOne.Name, itemOneRsl.Name)
		assert.Equal(t, itemTwo.Name, itemTwoRsl.Name)

	})

	t.Run("When_try_to_update_non-existent_item_Then_get_false", func(t *testing.T) {
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
		boolResult := tracker.UpdateItem(itemUpdate)
		sliceItems := tracker.GetItems()

		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)
		assert.Equal(t, boolResult, false)
		assert.Equal(t, sliceItems[1].Name, "John Johnson")
	})

	t.Run("When_try_to_update_existent_item_Then_get_true", func(t *testing.T) {
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
		itemOneRsl, errAddOne := tracker.AddItem(item)
		itemTwoRsl, errAddTwo := tracker.AddItem(itemTwo)

		itemUpdate := Item{
			ID:   itemTwoRsl.ID,
			Name: "Donald Trump",
		}

		boolResult := tracker.UpdateItem(itemUpdate)
		sliceItems := tracker.GetItems()

		assert.Equal(t, nil, errAddOne)
		assert.Equal(t, nil, errAddTwo)
		assert.Equal(t, boolResult, true)
		assert.Equal(t, item.Name, itemOneRsl.Name)
		assert.Equal(t, sliceItems[1].Name, "Donald Trump")
	})

	t.Run("When_delete_by_id_Then_success", func(t *testing.T) {
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

		sliceItems := tracker.GetItems()

		assert.Equal(t, err, nil)
		assert.Equal(t, errOne, nil)
		assert.Equal(t, errTwo, nil)
		assert.Equal(t, 1, len(sliceItems))
		assert.Equal(t, "Alice Marcuse", sliceItems[0].Name)
		assert.Equal(t, itemOneRsl.ID, sliceItems[0].ID)
	})
}
