package nsqmate

type IMessageProcessor interface {
	Process(msg []byte) []byte
	MultiProcess(msgs [][]byte) [][]byte

	Restore(msg []byte) ([]byte, error)
}
