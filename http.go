package main

import (
	"fmt"
	"net/http"
)
type increby1 int
type increby100 int
type fizzbuzz int
type counter interface {
	Up() int
}
func (i *fizzbuzz)Up()int{
	*i+=1
	if *i%5==0{
		return 0
	}
	return int(*i)
}
func (i *increby1)Up()int{
	*i+=1
	return int(*i)
}
func (i *increby100)Up()int{
	*i+=100
	return int(*i)
}
type pingpong struct {
	countervar counter
}
func (p pingpong) Pong(w http.ResponseWriter,r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintln(w, p.countervar.Up())
}
func main(){
	//var genral increby1
	var genraltypevar increby100
	//var genral fizzbuzz
	pingpongobj:=pingpong{&genraltypevar}
	http.Handle("/ping", http.HandlerFunc(pingpongobj.Pong))
	http.ListenAndServe("0.0.0.0:1024", nil)
}