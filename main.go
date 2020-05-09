//Programmer Name: Vincent G.
//Date completed: 05/08/2020
//Description: Inventory (3 points)


package main

import "fmt"

//Create an inventory struct with part number, name, unit price, and quantity. 

type Inventory struct {

    Partnumber int; Name string ; Unitprice float64; Quantity int

  }

func main() {

//Use the main to manipulate the inventory when a user "purchases" an item.

var Thing [10]Inventory

for i := 0; i <= 9; i++{
Thing[i] = Inventory{Partnumber: i, Name: "Thing", Unitprice: 1.99, Quantity: 10}
}

  for i := 0; i <= 9; i++{
    fmt.Println("Part#:", Thing[i].Partnumber, "Name:", Thing[i].Name, "Unitprice: $", Thing[i].Unitprice, "Quantity:", Thing[i].Quantity)
}

fmt.Println()

var shopping string

fmt.Println("Would you like to purchase a thing? 'y' for yes 'n' for no.")

 fmt.Scanln(&shopping)

 if shopping == "y" {

 for items := 1; shopping == "y"; items ++ {

 var selection int

 fmt.Println("[Enter the Part# of the Thing you'd like to 'purchase']")

 fmt.Scanln(&selection)
 
 Thing[selection].Quantity --
 
 for i := 0; i <= 9; i++{

    fmt.Println("Part#:", Thing[i].Partnumber, "Name:", Thing[i].Name, "Unitprice: $", Thing[i].Unitprice, "Quantity:", Thing[i].Quantity)
   }

   fmt.Println("Would you like to continue shopping?")

   fmt.Println("'y' for yes 'n' for no")

   fmt.Scanln(&shopping)

  }
 } 
}
