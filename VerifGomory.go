// déclaration de la structure Gomory
type Gomory struct {
	variable string
	borne *big.Rat
}

//Verifier si on a les conditions pour faire une Gomory cut
//NombreVariables = len(TabVar)
func VerifGomory(tabVar []string ,posVarTableau []string, alphaTab map[string]*big.Rat )([]Gomory){

  var tab_gomory = make([]Gomory,0)
  var GomoryStruct Gomory 

  //Test tous les variables hors base
  for i := len(posVarTableau)-1; i >= len(posVarTableau)- len(tabVar);i--{
    
    //Vérifie si ce sont que des variables d'écarts
    if posVarTableau[i][0]!='e' || alphaTab[posVarTableau[i]]{
      return tab_gomory
    }
  }
  



  //Si c'est vrai, deuxième vérification, il faut que l'affectation d'une des variables initiales soit non entières
  var i = 0;
  for i <= len(tabVar){
    //Regarde si ce n'est pas un entier
    if alphaTab[tabVar[i]].IsInt() {
      //Si c'est vrai, renvoyer la struct avec le booléen à true / la variable qui a son affection non entière / La nouvelle contrainte qui est crée
      GomoryStruct.variable = tabVar[i];
      var contrainte float64
      contrainte,_ = alphaTab[tabVar[i]].Float64()
      borne := math.Floor(contrainte) + 1
      GomoryStruct.borne = new(Rat).SetFloat64(borne)
      tab_gomory=append(tab_gomory,GomoryStruct)
    }
    i++
  }
  
  
  return tab_gomory
} 