// Copyright 2026 xgfone
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package currency

import (
	"fmt"
	"strings"
)

var _currencies = make(map[string]Currency, 32)

func normalizeCode(code string) string {
	return strings.ToUpper(strings.TrimSpace(code))
}

// Register registers a currency into the package-level registry.
//
// It panics if the given currency definition is invalid.
// It also panics if the same currency code is registered more than once.
func Register(code string, minorUnit int8, number int16, symbol, name string) {
	code = normalizeCode(code)

	if _, exists := _currencies[code]; exists {
		panic(fmt.Errorf("currency %s already registered", code))
	}

	currency := Currency{
		Name:      name,
		Code:      code,
		Symbol:    symbol,
		Number:    number,
		MinorUnit: minorUnit,
	}

	if err := currency.Validate(); err != nil {
		panic(err)
	}

	_currencies[code] = currency
}

// Unregister unregisters a registered currency.
func Unregister(code string) {
	code = normalizeCode(code)
	delete(_currencies, code)
}

// Get returns the registered currency by code.
//
// It returns nil if the currency is not found.
func Get(code string) (currency Currency, ok bool) {
	currency, ok = _currencies[normalizeCode(code)]
	return
}

// GetAll returns all the registered currencies.
func GetAll() []Currency {
	currencies := make([]Currency, 0, len(_currencies))
	for _, currency := range _currencies {
		currencies = append(currencies, currency)
	}
	return currencies
}

// GetAllCodes returns all the registered currency codes.
func GetAllCodes() []string {
	codes := make([]string, 0, len(_currencies))
	for code := range _currencies {
		codes = append(codes, code)
	}
	return codes
}

// IsSupported reports whether the given currency code is supported.
func IsSupported(code string) bool {
	_, ok := _currencies[normalizeCode(code)]
	return ok
}

// FormatMinorToMajor looks up a currency by code and
// formats a minor amount int64 into a major amount string.
//
// It returns an error if the currency is not found.
func FormatMinorToMajor(minorAmount int64, currencyCode string) (string, error) {
	currency, ok := Get(currencyCode)
	if !ok {
		return "", fmt.Errorf("%w: not found currency %s", ErrUnsupportedCurrency, currencyCode)
	}
	return currency.FormatMinorToMajor(minorAmount)
}

// ParseMajorToMinor looks up a currency by code and
// parses a major amount string into a minor amount int64.
func ParseMajorToMinor(majorAmount string, currencyCode string) (int64, error) {
	currency, ok := Get(currencyCode)
	if !ok {
		return 0, fmt.Errorf("%w: not found currency %s", ErrUnsupportedCurrency, currencyCode)
	}
	return currency.ParseMajorToMinor(majorAmount)
}
