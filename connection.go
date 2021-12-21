package amt

// Connection properties for a Client
type Connection struct {
	Host  string
	Port  uint32
	Path  string
	User  string
	Pass  string
	Debug bool
}
