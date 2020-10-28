package main
import "fmt"
func main(){
fmt.Print("enter number:")
var input int
fmt.Scanln(&input)
fmt.Print(input)
if(input%2==0){
fmt.Printf("is even number");
}else{
fmt.Print("is add");
}

}