package tracker

type Usecase interface {
	Done(in Input, out Output, tracker *Tracker)
}

type AddUsecase struct{}

func (u AddUsecase) Done(in Input, out Output, tracker *Tracker) {
	out.Out("enter name:")
	name := in.Get()

	item, err := tracker.AddItem(Item{Name: name})
	if err != nil {
		out.Out(err.Error())
	}
	out.Out(item.ToString())
}

type GetUsecase struct{}

func (u GetUsecase) Done(_ Input, out Output, tracker *Tracker) {
	for _, item := range tracker.GetItems() {
		out.Out(item.ToString())
	}
}
