package rest

type Handler interface {
	Prepare()
	Finish()
	Get()
	Post()
	Patch()
	Put()
	Delete()
	setChild(hd Handler)
}
