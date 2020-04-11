package Trans

import "fmt"

type TransModel struct {
	From string `json:"from"`
	To string  `json:"to"`
	Money int `json:"m"`
}
func NewTransModel() *TransModel{
	return &TransModel{}
}
func(this *TransModel) String() string {
	return fmt.Sprintf("%s转账给%s,金额是:%d\n",this.From,this.To,this.Money)
}
