package pets

type Pets struct {
	db any
}

func New(db any) *Pets {
	return &Pets{
		db: db,
	}
}
