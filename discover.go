package ec

// DiscoverService discovery interface
type DiscoverService interface {
	GetService(key string) (string, error)
}
