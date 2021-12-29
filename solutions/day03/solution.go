package day3

type Solution struct{}

func (sln Solution) Part1(input string) (int, error) {
	report, err := NewReport(input)
	if err != nil {
		return -1, err
	}

	g, err := report.Gamma()
	if err != nil {
		return -1, err
	}
	e, err := report.Epsilon()
	if err != nil {
		return -1, err
	}

	return int(g * e), nil
}

func (sln Solution) Part2(input string) (int, error) {
	report, err := NewReport(input)
	if err != nil {
		return -1, err
	}

	o2, err := report.Oxygen()
	if err != nil {
		return -1, err
	}
	co2, err := report.C02()
	if err != nil {
		return -1, err
	}

	return int(o2 * co2), nil
}
