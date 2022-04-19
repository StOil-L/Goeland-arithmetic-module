/**************************/
/* normalization_rules.go */
/**************************/
/**
* This file provides rules to manage normalization rules
**/

package ari

import (
	"ARI/types"
	"strconv"
)

// TODO : Compléter avec les règle int-LT et Int-GT si ce sont des entiers !
func normalizationRule(p types.Pred) []types.FormList {
	res := []types.FormList{}
	arg1 := p.GetArgs()[0]
	arg2 := p.GetArgs()[1]
	res_form_list := types.MakeEmptyFormList()

	switch p.GetID().GetName() {
	case types.Id_eq.GetName():
		res_form_list = append(res_form_list, types.MakePred(types.MakerId("lesseq"), []types.Term{arg1.Copy(), arg2.Copy()}))
		res_form_list = append(res_form_list, types.MakePred(types.MakerId("lesseq"), []types.Term{arg2.Copy(), arg1.Copy()}))
		return append(res, res_form_list)

	case types.Id_neq.GetName():
		res_form_list = append(res_form_list, types.MakePred(types.MakerId("less"), []types.Term{arg1.Copy(), arg2.Copy()}))
		res_form_list = append(res_form_list, types.MakePred(types.MakerId("greater"), []types.Term{arg2.Copy(), arg1.Copy()}))
		return append(res, res_form_list)
//int-lt
	case "less":
		y, _ := strconv.Atoi(arg2.GetName())
		Sy := strconv.Itoa(y - 1)
		nouveau_arg2 := types.MakerConst(types.MakerId(Sy), tInt)

		res_form_list = append(res_form_list, types.MakePred(types.MakerId("lesseq"), []types.Term{arg1.Copy(), nouveau_arg2.Copy()}))
		return append(res, res_form_list)
//int-gt
	case "greater":
		y, _ := strconv.Atoi(arg2.GetName())
		Sy := strconv.Itoa(y + 1)
		nouveau_arg2 := types.MakerConst(types.MakerId(Sy), tInt)

		res_form_list = append(res_form_list, types.MakePred(types.MakerId("greateq"), []types.Term{arg1.Copy(), nouveau_arg2.Copy()}))
		return append(res, res_form_list)

	default:
		return append(res, types.MakeSingleElementList(p))
	}
}

func normalizationNegRule(p types.Pred) []types.FormList {
	res := []types.FormList{}
	arg1 := p.GetArgs()[0]
	arg2 := p.GetArgs()[1]
	switch p.GetID().GetName() {
	case "less":
		res_form_list := types.MakeEmptyFormList()
		res_form_list = append(res_form_list, types.MakePred(types.MakerId("greateq"), []types.Term{arg1.Copy(), arg2.Copy()}))
		return append(res, res_form_list)

	case "greater":
		res_form_list := types.MakeEmptyFormList()
		res_form_list = append(res_form_list, types.MakePred(types.MakerId("lesseq"), []types.Term{arg1.Copy(), arg2.Copy()}))
		return append(res, res_form_list)

	case "lesseq":
		res_form_list := types.MakeEmptyFormList()
		res_form_list = append(res_form_list, types.MakePred(types.MakerId("greater"), []types.Term{arg1.Copy(), arg2.Copy()}))
		return append(res, res_form_list)

	case "greatereq":
		res_form_list := types.MakeEmptyFormList()
		res_form_list = append(res_form_list, types.MakePred(types.MakerId("less"), []types.Term{arg1.Copy(), arg2.Copy()}))
		return append(res, res_form_list)
	}
	return res
}
