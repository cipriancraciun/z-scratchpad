

package embedded


import _ "embed"




//go:embed build/version.txt
var BuildVersion string

//go:embed build/number.txt
var BuildNumber string

//go:embed build/timestamp.txt
var BuildTimestamp string

//go:embed build/sources-md5.txt
var BuildSourcesMd5 string

//go:embed build/sources.cpio.gz
var BuildSourcesCpio []byte

