package Trans

import "fmt"

type TransModel struct {
	Tid int `db:"tid"`
	From string `json:"from" db:"from"`
	To string  `json:"to" db:"to"`
	Money int `json:"m" db:"money"`
}
func NewTransModel() *TransModel{
	return &TransModel{}
}
func(this *TransModel) String() string {
	return fmt.Sprintf("%s转账给%s,金额是:%d\n",this.From,this.To,this.Money)
}
