package main

import (
	"fmt"
	"html"
)

func main() {
	s1 := `"Fran & Freddie's Diner" <tasty@example.com>`
	s2 := html.EscapeString(s1)
	s3 := html.UnescapeString(s2)
	fmt.Println(s2) // &#34;Fran &amp; Freddie&#39;s Diner&#34; &lt;tasty@example.com&gt;
	fmt.Println(s3) // "Fran & Freddie's Diner" <tasty@example.com>
}
