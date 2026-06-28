package main
import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// Mathmatical operations

func add(x int,y int) int{
	return x + y
}

func sub(x int,y int) int{
	return x - y
}

func mul(x int,y int) int{
	return x * y
}

func div(x int,y int)(int ,error){
	if y==0{
		return 0,errors.New("division by zero")
	}
	return x / y,nil
}

// HTTP Handlers

func adds(w http.ResponseWriter,r *http.Request){
	xstr:=r.URL.Query().Get("x")
	ystr:=r.URL.Query().Get("y")
	x,err1:=strconv.Atoi(xstr)
	if err1!=nil{
		http.Error(w,"Invalid value for x",http.StatusBadRequest)
		return
	}
	y,err2:=strconv.Atoi(ystr)
	if err2!=nil{
		http.Error(w,"Invalid value for y",http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w,"Addition of",x,"and",y,"is",add(x,y));
}

func subs(w  http.ResponseWriter,r *http.Request){
	xstr:=r.URL.Query().Get("x")
	ystr:=r.URL.Query().Get("y")

	x,err1:=strconv.Atoi(xstr)
	if err1!=nil{
		http.Error(w,"Invalid value for x",http.StatusBadRequest)
		return
	}
	y,err2:=strconv.Atoi(ystr)
	if err2!=nil{
		http.Error(w,"Invalid value for y",http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w,"Subtraction of",x,"and",y,"is",sub(x,y));
}

func divs(w  http.ResponseWriter,r *http.Request){
	xstr:=r.URL.Query().Get("x")
	ystr:=r.URL.Query().Get("y")
	x,err1:=strconv.Atoi(xstr)
	if err1!=nil{
		http.Error(w,"Invalid value for x",http.StatusBadRequest)
		return
	}
	y,err2:=strconv.Atoi(ystr)
	if err2!=nil{
		http.Error(w,"Invalid value for y",http.StatusBadRequest)
		return
	}
	result,err:=div(x,y)
	if err!=nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w,"Division of",x,"and",y,"is",result);	
}

func muls(w  http.ResponseWriter,r *http.Request){
	xstr:=r.URL.Query().Get("x")
	ystr:=r.URL.Query().Get("y")		

	x,err1:=strconv.Atoi(xstr)
	if err1!=nil{
		http.Error(w,"Invalid value for x",http.StatusBadRequest)
		return
	}
	y,err2:=strconv.Atoi(ystr)
	if err2!=nil{
		http.Error(w,"Invalid value for y",http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w,"Multiplication of",x,"and",y,"is",mul(x,y));
}

// home handler

func home(w http.ResponseWriter,r *http.Request){
	fmt.Fprintln(w,"Welcome to the calculator API!")
}

// main function to start the server

func main(){

	http.HandleFunc("/",home)
	http.HandleFunc("/add",adds)
	http.HandleFunc("/sub",subs)
	http.HandleFunc("/mul",muls)
	http.HandleFunc("/div",divs)

	err:=http.ListenAndServe(":8080",nil)
	if err==nil{
		fmt.Println("Server started on port 8080")
	}
	if err!=nil{
		fmt.Println("Error starting server:",err)
	}
}