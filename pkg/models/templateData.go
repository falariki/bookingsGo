package models

// Create a new type that is going to hold the template data
// TemplateData holds Data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	//cross site request forgery token
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
