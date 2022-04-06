package simulator

type Simulator interface {
	Name() string
	Run()
}
