package components

type HomeComponent struct{}

func (h HomeComponent) Render() string {
	return "<h2>Home Component</h2>"
}
