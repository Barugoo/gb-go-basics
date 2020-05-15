package main

func main() {
	func mirroredQuery() string {
		responses := make(chan string, 3)
		go func() {
			responses <- request("asia.site.io")
		}()
		go func() {
			responses <- request("europe.site.io")
		}()
		go func() {
			responses <- request("americas.site.io")
		}()
		return <-responses // возврат самого быстрого ответа
	}
	
	func request(hostname string) string { /* */ }
}
