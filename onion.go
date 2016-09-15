//go:generate ginger $GOFILE
package main

//@ginger
type Grs struct {
	Ginger_Created int32 `json:"ginger_created"`
	Ginger_Id      int32 `json:"ginger_id" gorm:"primary_key"`

	Code    string `json:"code"`
	Name string `json:"name"`
}
