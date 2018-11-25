package main

import "fmt"

func main(){
  //usuario: definiendo el tipo de dato
  var nombre, apellido string
  nombre ="Alexys"
  apellido ="Lozada"
  fmt.Println(nombre, apellido)
  //Go: definiendo el tipo de dato
  nombre2 :="Carlos"
  apellido2 :="Pineda"
  fmt.Println(nombre2, apellido2)
  //definiendo valores en una sola línea
  nombre3,apellido3 :="Poncho","Zuleta"
  fmt.Println(nombre3, apellido3)
  //cambiando valor de una variable [sin := ]
  nombre2 ="George"
  fmt.Println(nombre2, apellido2)
  /*
    cambiando valor de una variable [con := ]
    siempre que se cre una nueva variable
  */
  nombre3,edad3 :="Manuel",25
  fmt.Println("Nombre: ",nombre3," edad: ",edad3)
  
}
