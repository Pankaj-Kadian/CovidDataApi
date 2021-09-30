package Config

var state_code map[string]string = map[string]string{"Total": "TT",
	"Andaman & Nicobar Islands": "AN",
	"Andhra Pradesh":            "AP",
	"Arunachal Pradesh":         "AR",
	"Assam":                     "AS",
	"Bihar":                     "BR",
	"Chandigarh":                "CH",
	"Chhattisgarh":              "CT",
	"Delhi":                     "DL",
	"Goa":                       "GA",
	"Gujarat":                   "GJ",
	"Haryana":                   "HR",
	"Himachal Pradesh":          "HP",
	"Jammu and Kashmir":         "JK",
	"Jharkhand":                 "JH",
	"Karnataka":                 "KA",
	"Kerala":                    "KL",
	"Ladakh":                    "LA",
	"Lakshadweep":               "LD",
	"Madhya Pradesh":            "MP",
	"Maharashtra":               "MH",
	"Manipur":                   "MN",
	"Meghalaya":                 "ML",
	"Mizoram":                   "MZ",
	"Nagaland":                  "NL",
	"Odisha":                    "OR",
	"Puducherry":                "PY",
	"Punjab":                    "PB",
	"Rajasthan":                 "RJ",
	"Sikkim":                    "SK",
	"Tamil Nadu":                "TN",
	"Telangana":                 "TG",
	"Tripura":                   "TR",
	"Uttar Pradesh":             "UP",
	"Uttarakhand":               "UT",
	"West Bengal":               "WB"}

func GetStateCodes() map[string]string {
	return state_code
}

func stateCodeFromStateName() map[string]string {
	state_name := make(map[string]string)
	for k, v := range state_code {
		state_name[v] = k
	}
	return state_name
}

// func StateCode() map[string]string{

// 	return stateCode
// }
