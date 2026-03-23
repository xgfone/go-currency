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

func init() {
	// Common 0-decimal currencies
	Register("JPY", 0, 392, "¥", "Japanese Yen")
	Register("KRW", 0, 410, "₩", "South Korean Won")
	Register("VND", 0, 704, "₫", "Vietnamese Dong")

	// Common 2-decimal currencies
	Register("CNY", 2, 156, "¥", "Chinese Yuan")
	Register("USD", 2, 840, "$", "US Dollar")
	Register("EUR", 2, 978, "€", "Euro")
	Register("GBP", 2, 826, "£", "Pound Sterling")
	Register("HKD", 2, 344, "HK$", "Hong Kong Dollar")
	Register("SGD", 2, 702, "S$", "Singapore Dollar")
	Register("AUD", 2, 36, "A$", "Australian Dollar")
	Register("CAD", 2, 124, "C$", "Canadian Dollar")
	Register("NZD", 2, 554, "NZ$", "New Zealand Dollar")
	Register("CHF", 2, 756, "CHF", "Swiss Franc")
	Register("INR", 2, 356, "₹", "Indian Rupee")
	Register("THB", 2, 764, "฿", "Thai Baht")
	Register("MYR", 2, 458, "RM", "Malaysian Ringgit")
	Register("IDR", 2, 360, "Rp", "Indonesian Rupiah")
	Register("PHP", 2, 608, "₱", "Philippine Peso")
	Register("TWD", 2, 901, "NT$", "New Taiwan Dollar")
	Register("AED", 2, 784, "AED", "UAE Dirham")
	Register("SAR", 2, 682, "SAR", "Saudi Riyal")
	Register("BRL", 2, 986, "R$", "Brazilian Real")
	Register("MXN", 2, 484, "$", "Mexican Peso")
	Register("TRY", 2, 949, "₺", "Turkish Lira")
	Register("RUB", 2, 643, "₽", "Russian Ruble")

	// Common 3-decimal currencies
	Register("BHD", 3, 48, "BD", "Bahraini Dinar")
	Register("KWD", 3, 414, "KD", "Kuwaiti Dinar")
}
