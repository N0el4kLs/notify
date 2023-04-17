package notifications

type Notifier interface {
	Notice() error
}
