package utils


func Map_in(i string,m map[string]interface{}) bool{
	for k,_ := range m {
		if i == k {
			return true
		}
	}
	return false
}
