package list

type Pusher interface {
	Push(key string, value string)
}
