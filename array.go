package main
import "fmt"
func main(){
var x[5] int
var i,j int
for i=0;i<5;i++{
x[i]=i+10
}
for j=0;j<5;j++{
fmt.Println("element[%d]=%d\n",j,x[j])
}
}