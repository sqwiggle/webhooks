package main

type Attributer interface {
	TableName() string
	Attributes() ([]string, []interface{})
}
