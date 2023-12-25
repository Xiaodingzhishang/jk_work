package main

func YourName(name string, aliases ...string) {

}
func CallYourName() {
	YourName("大名")
	YourName("大名", "小康")
	YourName("大名", "小康", "小米")
	aliases := []string{"大名", "小康"}
	YourName("大名", aliases...)
}
