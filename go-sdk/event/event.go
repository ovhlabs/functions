package event

type Event struct {
	Data    string
	Method  string
	Params  map[string]string
	Secrets map[string]string
}
