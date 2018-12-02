package rest

var (
	defaultOffset = 0
	defaultLimit  = 10
)

func SetOptions(opts ...Option) {
	for _, opt := range opts {
		opt()
	}
}

type Option func()

func DefaultOffset(offset int) Option {
	return func() {
		defaultOffset = offset
	}
}

func DefaultLimit(limit int) Option {
	return func() {
		defaultLimit = limit
	}
}
