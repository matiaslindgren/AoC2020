package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/matiaslindgren/AoC2020/util"
)

type Food struct {
	ingredients []string
	allergens []string
}

func intersectIngredients(foods []*Food) []string {
	igs := map[string]bool{}
	for _, ig := range foods[0].ingredients {
		igs[ig] = true
	}
	for _, f := range foods[1:] {
		tmp := map[string]bool{}
		for _, ig := range f.ingredients {
			if igs[ig] {
				tmp[ig] = true
			}
		}
		igs = tmp
	}
	res := []string{}
	for k := range igs {
		res = append(res, k)
	}
	return res
}

func parseFood(line string) *Food {
	parts := strings.Split(line, " (contains ")
	ig := strings.Split(parts[0], " ")
	ag := strings.Split(strings.TrimRight(parts[1], ")"), ", ")
	return &Food{ig, ag}
}

type A2I map[string][]string

func (a A2I) copy() A2I {
	b := A2I{}
	for k, v := range a {
		b[k] = v
	}
	return b
}

func (a A2I) valid() bool {
	for _, vs := range a {
		if len(vs) != 1 {
			return false
		}
	}
	seen := map[string]bool{}
	for _, vs := range a {
		ag := vs[0]
		if seen[ag] {
			return false
		}
		seen[ag] = true
	}
	return true
}

func searchAllergenMapping(ag2igs A2I, tmpMapping A2I, ags []string) A2I {
	if len(ags) == 0 {
		return tmpMapping
	}
	for _, ig := range ag2igs[ags[0]] {
		nextMapping := tmpMapping.copy()
		nextMapping[ags[0]] = []string{ig}
		if x := searchAllergenMapping(ag2igs, nextMapping, ags[1:]); x != nil && x.valid() {
			return x
		}
	}
	return nil
}

func search(lines []string) (int, string) {
	allFoods := make([]*Food, len(lines))
	for i, line := range lines {
		allFoods[i] = parseFood(line)
	}

	allergen2foods := map[string][]*Food{}
	for _, food := range allFoods {
		for _, ag := range food.allergens {
			allergen2foods[ag] = append(allergen2foods[ag], food)
		}
	}

	allergen2ingredients := A2I{}
	hasAllergen := map[string]bool{}
	for ag, foods := range allergen2foods {
		for _, ig := range intersectIngredients(foods) {
			hasAllergen[ig] = true
			allergen2ingredients[ag] = append(allergen2ingredients[ag], ig)
		}
	}

	a := 0
	for _, f := range allFoods {
		for _, ig := range f.ingredients {
			if !hasAllergen[ig] {
				a++
			}
		}
	}

	allAgs := []string{}
	for ag := range allergen2foods {
		allAgs = append(allAgs, ag)
	}
	ig2ag := map[string]string{}
	igs := []string{}
	for ag, ig := range searchAllergenMapping(allergen2ingredients, A2I{}, allAgs) {
		ig2ag[ig[0]] = ag
		igs = append(igs, ig[0])
	}

	sort.Slice(igs, func(i, j int) bool { return ig2ag[igs[i]] < ig2ag[igs[j]] })
	b := strings.Join(igs, ",")

	return a, b
}

func main() {
	input := util.SlurpStdinLines()
	a, b := search(input)
	fmt.Println(a, b)
}
