package function

const (
	FieldTemplateName = "template_name"
	FIeldTemplate     = "template"
)

// Context : context that can be passed into functions
//
type Context map[string]interface{}

func (c Context) TemplateName() (string, bool) {
	name, flag := c[FieldTemplateName]
	return name.(string), flag
}
