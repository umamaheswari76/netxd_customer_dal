package interfaces


type ITransaction interface {
	Transfer(int, int, int) (string, error)
}
