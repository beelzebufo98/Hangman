package domain

type FallbackDrawer struct{}

func (f *FallbackDrawer) Render(errors int) string {
	return ""
}

func (f *FallbackDrawer) MaxStages() int {
	return 5
}
