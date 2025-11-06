package domain

type Drawer interface {
	Render(errors int) string
	MaxStages() int
}
