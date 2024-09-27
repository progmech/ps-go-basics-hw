package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList []Bin

func newBin(id string, private bool, name string) *Bin {
	res := &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}
	return res
}
