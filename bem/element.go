package bem

type Element struct {
	Tag       string
	Name      string
	Elements  []Element
	Modifiers []Modifier
}
