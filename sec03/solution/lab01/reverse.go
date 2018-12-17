package main


func reverse(s string) string {
	// revisiting 'string' iteration
	sar := []rune(s)
	l := len(sar)
	for i:=0;i<l/2;i++{
		sar[i],sar[l-(i+1)]=sar[l-(i+1)], sar[i]
	}
	return string(sar)
}