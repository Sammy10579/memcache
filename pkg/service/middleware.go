package service

type Storage interface {
	Set(key, value string) (err error)
	Get(key string) (value string, err error)
	Delete(key string) (err error)
}
