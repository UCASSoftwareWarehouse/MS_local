package utils

const (
	OSmap uint32 = 0x000000ff //operating system
	PLmap uint32 = 0x0000ff00 //programing language
	NLmap uint32 = 0x00ff0000 // natural language
	Tomap uint32 = 0xff000000 //topic
)

func GetOSValue(classifer uint32) uint8 {
	return uint8(classifer & OSmap)
}

func GetPLValue(classifer uint32) uint8 {
	return uint8((classifer & PLmap) >> 8)
}

func GetNLValue(classifer uint32) uint8 {
	return uint8((classifer & NLmap) >> 16)
}

//topic value
func GetToVaule(classifer uint32) uint8 {
	return uint8((classifer & Tomap) >> 24)
}

func GetClassifier(osValue uint8, plValue uint8, nlValue uint8, toValue uint8) uint32 {
	var res uint32
	res += uint32(osValue)
	res += (uint32(plValue) << 8)
	res += (uint32(nlValue) << 16)
	res += (uint32(toValue) << 24)
	return res
}
