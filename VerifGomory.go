package main


import (
    "fmt"
    "regexp"
  "math/big"
  "math"
  "strconv"
)

// déclaration de la structure Gomory
 type Gomory struct {
	valide bool
	variable string
	borne float64
}

//Verifier si on a les conditions pour faire une Gomory cut
//NombreVariables = len(TabVar)
func VerifGomory(NombreVariables int ,posVarTableau []string, alphaTab map[string]*big.Rat,tabConst []*big.Rat )(Gomory){

  GomoryStruct := Gomory{false,"0",0}
  
  var valideVariableEcart = regexp.MustCompile(`[e][0-9]`)
  var cpt int
  
  
  //Test tous les variables dans la base
  for i := len(posVarTableau)-1; i >= len(posVarTableau)- NombreVariables;i--{
    var contraint string
    contraint = posVarTableau[i][1:]
    Contrainte,_ := strconv.Atoi(contraint)
    //Vérifie si ce sont que des variables d'écarts
    if valideVariableEcart.MatchString(posVarTableau[i]) && (alphaTab[posVarTableau[i]] == tabConst[Contrainte]) {
      cpt++
    }
  }
  var TestEcart bool
  TestEcart = (cpt == NombreVariables) 
  
  //Si c'est vrai, deuxième vérification, il faut que l'affectation d'une des variables initiales soit non entières
  if(TestEcart){
      var i = 0;
    for i <= NombreVariables{
  	xi := fmt.Sprint("x", i)
      //Regarde si ce n'est pas un entier
      if(alphaTab[xi].IsInt()){
        //Si c'est vrai, renvoyer la struct avec le booléen à true / la variable qui a son affection non entière / La nouvelle contrainte qui est crée
        GomoryStruct.valide = true;
        GomoryStruct.variable = xi;
        var contrainte float64
        contrainte,_ = alphaTab[xi].Float64()
        GomoryStruct.borne = math.Floor(contrainte) + 1
      }
      i++
    }
  }
  
  return GomoryStruct
} 

func main (){
  //VerifGomory(3,[x0,x1,e0,e1],[x0:2/1,x1:1/1,e0:1/1,e1:0:1])
}