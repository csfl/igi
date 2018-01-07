package node

import "strings"

type Conf struct {
	HttpHost  string
	UdpHost   string
	TcpHost   string
	Debug     bool
	Testnet   bool
	Neighbors MultiString
}

type MultiString []string

func (m MultiString) String() string {
	return strings.Join([]string(m), ", ")
}

func (m *MultiString) Set(value string) error {
	a := []string(*m)
	*m = MultiString(append(a, value))
	return nil
}
