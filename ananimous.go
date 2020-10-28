package main
import(
 "fmt"
)
func main(){
number:=10
squarenum:=func()(int){
number*=number
return number
}
fmt.Println(squarenum())
fmt.Println(squarenum())
}

