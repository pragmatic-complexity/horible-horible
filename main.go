package main

func main() {

	secrets := secretService.NewProvider()

	awesomethingy.TestableServe(secrets)
}
