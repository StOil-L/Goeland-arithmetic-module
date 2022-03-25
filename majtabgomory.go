// Met à jour les coeff, const et alpha en cas de Gomory Cut
func majTabGomory(cut Gomory, tableau [][]*big.Rat, tabConst []*big.Rat, varInit []string, solution map[string]*big.Rat)([][]*big.Rat, []*big.Rat, map[string]*big.Rat){
    if cut.valide {
        var coeffLine := make([]*big.Rat, len(tableau[0]))
        for i := 0; i < len(coeffLine); i++{
            if varInit[i] == cut.variable {
                coeffLine[i] = big.NewRat(1, 1)
            }
            else {
            coeffLine[i] = big.NewRat(0, 1)
            }
        }
        append(tableau, coeffLine)
        
        var ratborne := big.NewRat(cut.borne, 1)
        append(tabConst, ratborne)
        
        solution[cut.variable] = big.NewRat(0, 1)
    }
    return tableau, tabConst, solution
}