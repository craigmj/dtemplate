package xmlparse

import (
	`os`
	`testing`
	`strings`
	`fmt`
	`io/ioutil`
)

func TestOne(t *testing.T) {
	raw, err := ioutil.ReadFile(`../../tests/templates/test.html`)
	if nil!=err {
		t.Fatal(err.Error())
	}
	d, err := Parse(raw)
	if nil!=err {
		t.Fatal(err.Error())
	}
	tag := d.Root.FirstChild().TagName()
	if tag!=`template` {
		t.Fatalf(`Expected root tag name of %s but got %s`, `template`, tag)
	}
	if !strings.HasPrefix(d.Root.RawString(), `<template>`) {
		fmt.Fprintln(os.Stderr, d.Root.RawString())
		t.Fatalf(`Expected template as root of document, but got something else`)
	}
	fmt.Fprintln(os.Stderr, d.Root.RawString())
	t.Fatalf(`all good here`)
		
}

var badForm = `<template>
<style dtemplate-process="scss">
:host {
	position: relative;
	overflow: hidden;
}
a {
	display: flex;
	flex-direction: column;
	justify-content: flex-start;
	align-items: stretch;
}
img {
	width: 100%;
	height: 12em;
	border-radius: 8px;
}
.details {
	position: absolute;
	top: 1em;
	left: 0;
	border-radius: 0 3px 3px 0;
	background-color: var(--color-backpackers, #f57324);
	color: white;
	padding: 0.3em 1em 0.3em 0.3em;
}
.brand {
	font-weight: bold;
}
.ends {
	font-size: 0.8em;	
	display: flex;
	flex-direction: row;
	justify-content: center;
	align-items: center;
	
	color: var(--color-text);
	text-decoration: none;
	
	padding: 0.5em 10em;
}
.ends svg {
	width: 1.4em;
}
</style>
<a href="#" data-set="link"><img 
data-set="image" /><div class="details"><div data-set="brand" class="brand"> </div>
<div data-set="deal" class="deal"> </div></div><div class="ends">
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 64.521 76.02" height="287.32" width="243.86"><path d="M28.77 74.705C15.719 73.086 5.275 63.928 2.036 51.265c-3.69-14.43 3.707-29.738 17.38-35.969 2.375-1.082 5.693-2.084 7.766-2.343 2.553-.32 2.381-.08 2.381-3.337V6.723h-1.596c-2.035 0-3.274-.516-3.834-1.599-.595-1.152-.326-2.418.697-3.28l.805-.677H39.13l.773.773c1.077 1.077 1.08 2.58.008 3.78-.715.8-.901.865-2.779.96l-2.011.101v2.876c0 2.772.021 2.881.595 2.996.327.066 1.37.258 2.315.428 4.177.75 8.511 2.59 12.265 5.208l2.08 1.45 3.63-3.597c3.518-3.486 3.666-3.598 4.755-3.598 1.584 0 2.67 1.086 2.67 2.67 0 1.09-.112 1.237-3.598 4.755l-3.597 3.63 1.45 2.08c2.784 3.991 4.526 8.247 5.384 13.155.604 3.454.369 8.92-.535 12.422-3.026 11.731-12.439 20.679-24.24 23.043-2.614.524-7.05.713-9.524.406zm9.967-6.206c11.175-2.834 19.125-13.058 19.131-24.602.002-3.403-.338-5.47-1.427-8.668-1.269-3.725-2.997-6.466-5.976-9.478-5.048-5.106-11.11-7.656-18.179-7.647-8.111.01-15.645 3.716-20.412 10.043C5.926 36.04 4.971 46.45 9.386 55.274c3.666 7.328 11.238 12.751 19.384 13.885 2.69.374 7.018.087 9.967-.66zm-8.4-22.862l-.774-.773V25.468l.774-.773c1.321-1.322 3.593-.953 4.383.712.363.765.409 2.136.343 10.172l-.076 9.288-.864.772c-1.207 1.078-2.708 1.077-3.786-.002z" fill="#333"/></svg>
<span data-set="ends" > </span>
</div></a>
</template>`