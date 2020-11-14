package integrity

type Strategy interface {
	Check()
	Add()
	IsSet()
}
