package client

type protoRwc struct{}

func newProtoRwc() *protoRwc {
	return &protoRwc{}
}

func (r protoRwc) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (r protoRwc) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (r protoRwc) Close() error {
	return nil
}
