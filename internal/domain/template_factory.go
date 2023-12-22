package domain

type TemplateFactory struct {
	// constraints such as length, pattern ...
}

func (t *TemplateFactory) CreateTemplate(name string) *Template {

	return &Template{Name: name}
}
