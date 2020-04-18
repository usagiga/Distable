package entity

type ServerType int

const (
	Master ServerType = iota // 0 (Default value)
	Slave                    // 1
)
