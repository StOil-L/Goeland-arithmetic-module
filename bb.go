package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isInteger(54.4))	

}

func isInteger(nombre float64) bool{
	var estEntier bool 
	estEntier=false
	var monFlottant float64
	monFlottant=math.Mod(nombre,1)
	if(nombre - (monFlottant)==nombre){
	estEntier=true
	}

return estEntier

}

func branch_bound(solution []float64){

	var vous_vouliez_de_l_iteratif bool
	vous_vouliez_de_l_iteratif=false

	for ! vous_vouliez_de_l_iteratif{
		var tamponBooleen bool
		tamponBooleen=true
		for index,element := range solution	{
			if(!isInteger(element)){
				tamponBooleen=false
			}	
		}
		if(tamponBooleen){
			vous_vouliez_de_l_iteratif=true
			break
		}
		
		//branch and bound ! 
		for index,element := range solution	{
		
			//pour chaque solution non entiere appeler le simplexe
			// avec contrainte floor et ceil.	
			}	
		}
		
		
	}

}