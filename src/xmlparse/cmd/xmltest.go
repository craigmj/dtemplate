package main

import (
	"fmt"
	"log"

	"xmlparse"
)

func main() {
	doc, err := xmlparse.Parse([]byte(`<!doctype html >
<html lang="en">
	<head>
	</head>
	<script><![CDATA[Here's some script]]>
	console.log('this should work');
	</script>
	<!-- And a comment-->
	<!-- A second comment-->
	<body>
		<h1>This is a header</h1>
		<p>And here is a paragraph</p>
		<div id="first">Here is a first div with an id &amp; and an entity and a badly formed & without the entity, 
		and another &just by itself.</div>
		<p>And here is some text and a <br> auto-closing tag.</p>
	</body>
</html>`))
	if nil != err {
		log.Fatal(err)
	}
	fmt.Println(doc.String())
}
