package util

func Check(err error) {
	if err != nil {
		// log.Print(err.Error())
		panic(err)
	}
}
