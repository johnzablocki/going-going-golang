package main

//Check for err and panic if found
func Check(err error) {
	if err != nil {
		panic(err)
	}
}
