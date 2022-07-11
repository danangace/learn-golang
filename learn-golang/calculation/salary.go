package calculation

func CalcSalary(currSalary int, tax float32) int {
	annualSalary := calculateAnnualSalary(currSalary)
	taxableSalary := calculateTaxableSalary(annualSalary)
	return taxableSalary
}

func calculateAnnualSalary(salary int) int {
	return salary * 12
}

func calculateTaxableSalary(annualSalary int) int {
	if annualSalary <= 54000000 {
		return 0
	} else {
		return annualSalary - 54000000
	}
}
