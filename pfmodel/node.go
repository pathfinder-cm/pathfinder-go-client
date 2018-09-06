package pfmodel

type Node struct {
	Hostname   string
	Ipaddress  string
	CreatedAt  string
	MemFreeMb  uint64
	MemUsedMb  uint64
	MemTotalMb uint64
}
