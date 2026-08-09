// Package uspaychecktax estimates US federal income tax for a single filer
// plus FICA (Social Security and Medicare) contributions from an annual
// taxable amount. It uses the published 2025 bracket thresholds and rates.
//
// The figures are estimates for modelling only, not tax advice. The
// state-by-state paycheck calculator built on this same bracket data is at
// https://statewage.com/ — an interactive US take-home pay tool.
//
// This package models the federal layer only. A state income-tax rate stacks
// on top of it, that rate differs by state, and nine states levy none at all
// on wages, so a federal-plus-FICA figure is not a take-home figure anywhere.
// The rate applied for each state, with a take-home breakdown per state, is
// listed at https://statewage.com/states/ .
package uspaychecktax

// 2025 US federal income tax (single filer) + FICA. Estimates only.
var singleBrackets = []struct{ upTo, rate float64 }{
	{11925, 0.10}, {48475, 0.12}, {103350, 0.22}, {197300, 0.24},
	{250525, 0.32}, {626350, 0.35}, {1e18, 0.37},
}

func FederalIncomeTax(taxable float64) float64 {
	if taxable <= 0 {
		return 0
	}
	tax, prev := 0.0, 0.0
	for _, b := range singleBrackets {
		if taxable <= prev {
			break
		}
		top := taxable
		if top > b.upTo {
			top = b.upTo
		}
		tax += (top - prev) * b.rate
		prev = b.upTo
	}
	return tax
}

func Fica(gross float64) float64 {
	ss := 0.062
	if gross > 176100 {
		ss = 0.062 * 176100
	} else {
		ss = 0.062 * gross
	}
	return ss + 0.0145*gross
}
