package api

type storage interface {
	Write(key string, value interface{}) error
	Read(key string, value interface{}) error
	List(kind interface{}) []string
	Delete(key string, kind interface{}) error
}
