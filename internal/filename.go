package internal

import "fmt"


func makeFilename(dbname string, num uint64, suffix string) string {
	return  fmt.Sprintf("%s-%06d.%s", dbname, num, suffix)
}


func TableFilename(dbname string, num uint64)string  {
	return makeFilename(dbname, num, "ldb")
}


func DescriptorFilename( num uint64) string {
	return  fmt.Sprintf("MANIFEST-%06d", num)
}


func CurrentFilename()string  {
    return  fmt.Sprintf("CURRENT")
}


func TempfileName(dbname string, num uint64)string  {
	 return makeFilename(dbname, num, "tmp")
}


