package main
import (
    "fmt"
  //  "strings"
    "regexp"
//    "strconv"
)

func main() {





    var tableauInequation =make([]string,0)
    var tableauInequaBienEcrit =make([]string,0)
    
    tableauInequation=append(tableauInequation,"x+3y<=2\x00")
    tableauInequation=append(tableauInequation,"x+y>=2")
    tableauInequation=append(tableauInequation,"x+y>=2")
    tableauInequation=append(tableauInequation,"x+y=2")

    fmt.Println("ici",tableauInequation)
    var position =make([]string,0)
    for _,element := range tableauInequation{
        for j,element2 := range element{
          if element2 =='>'{
            position=append(position,"pas_bien")
          } else if element2=='=' && element[j-1]!='<' && element[j-1]!='>'{
            position=append(position,"=")
          } else if element2=='<' {
            position=append(position,"bien")

          }
        }

    }
    fmt.Println("position :" ,position)

    for i, element := range position {
    
        var stringTampon =make([]byte,0)
         if element=="bien"{
            stringTampon=append(stringTampon,tableauInequation[i][0])
            reV1 := regexp.MustCompile(`[a-zA-Z]`)
            vars1 := reV1.FindAll([]byte(tableauInequation[i]), -1)
            if stringTampon[0]==vars1[0][0]{
                stringTampon=append(stringTampon,' ')
                stringTampon=append(stringTampon,tableauInequation[i][1])
                reV2 := regexp.MustCompile(`[+\-]`)
                vars2 := reV2.FindAll([]byte(tableauInequation[i]), -1)
                if stringTampon[2]==vars2[0][0] {
                    reV3 := regexp.MustCompile(`[0-9]`)
                    vars3 := reV3.FindAll([]byte(tableauInequation[i]), -1)    
                    if tableauInequation[i][2]==vars3[0][0]{
                        stringTampon=append(stringTampon,vars3[0][0])
                        stringTampon=append(stringTampon,' ')
                        stringTampon=append(stringTampon,tableauInequation[i][3])    
                        stringTampon=append(stringTampon,' ') 
                        if tableauInequation[i][4] =='<'{
                            stringTampon=append(stringTampon,tableauInequation[i][4])
                            stringTampon=append(stringTampon,tableauInequation[i][5])
                            var cpt int
                            cpt=6
                            for tableauInequation[i][cpt]!='\x00'{
                                stringTampon=append(stringTampon,' ') 
                                stringTampon=append(stringTampon,tableauInequation[i][cpt])
                                cpt+=1
                            }
                            tableauInequaBienEcrit=append(tableauInequaBienEcrit,string(stringTampon))

                        } 
                    } else {
                        stringTampon=append(stringTampon,' ')
                        stringTampon=append(stringTampon,tableauInequation[i][2])
                        stringTampon=append(stringTampon,' ')
                        if tableauInequation[i][3] =='<'{
                            stringTampon=append(stringTampon,tableauInequation[i][3])
                            stringTampon=append(stringTampon,tableauInequation[i][4])
                        }
                    }
                    
            
                }
                
            }
                
            fmt.Println("stringTampon",string(stringTampon))
            fmt.Println("tableauInequationBienEcrit",tableauInequaBienEcrit)



        } else if element=="pas_bien" {
            for _,element2 := range tableauInequation[i]{
                if element2 != '>'{
            //  stringTampon= "-"+tableauInequa[element2:]
                    continue
                } else if element2 == '>'{
                     break
                }
            }

        }

    }
}