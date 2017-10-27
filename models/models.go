package models

//Response struct
type Response struct {
	IPAddress   string `json:"ip_address"`
	CountryCode string `json:"country_code"`
	Language    string `json:"language"`
	IsoString   string `json:"iso"`
}

//Error struct contains an error string
type Error struct {
	Error string `json:"error"`
}

//Language contains a languages array, and a default language
type Language struct {
	Languages []Languages `json:"languages"`
	Default   Languages   `json:"default"`
}

//Languages contains a Counyty and Language as strings
type Languages struct {
	Language string `json:"language"`
	Country  string `json:"country"`
}

//DBLocation PRAMS
//Type: @{STRING}
//Location: @{STRING}
type DBLocation struct {
	Type     string
	Location string
	S3Config S3Config
}

//S3Config PARAMS
//Bucket: @{STRING}
//Location: @{STRING}
//Region: @{STRING}
type S3Config struct {
	Bucket string
	Key    string
	Region string
}

//LanguageConfig ...
var LanguageConfig Language

//ConfigFile is a string
var ConfigFile string

//CurrentHash is the current hash
var CurrentHash string
