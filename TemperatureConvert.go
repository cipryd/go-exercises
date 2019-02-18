package main

type celsius float64
type kelvin float64
type fahrenheit float64

// celsius converts °K to °C
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

// celsius converts °F to °C
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

// kelvin converts °C to °K
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

// kelvin converts °F to °K
func (f fahrenheit) kelvin() kelvin {
	return kelvin((f + 459.67) * 5.0 / 9.0)
}

// fahrenheit converts °C to °F
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

// fahrenheit converts °K to °F
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(k*9.0/5.0 - 459.67)
}
