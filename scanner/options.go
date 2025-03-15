package scanner

type Option[T ScannerOption] func(*T)

type ScannerOption struct {
	path   string
	filter Filter
}

func WithPath[T ScannerOption](path string) Option[T] {
	return func(opt *T) {
		switch x := any(opt).(type) {
		case *ScannerOption:
			x.path = path
		}
	}
}

func WithFilter[T ScannerOption](filter Filter) Option[T] {
	return func(opt *T) {
		switch x := any(opt).(type) {
		case *ScannerOption:
			x.filter = filter
		}
	}
}
