package dtemplate

var jsTemplate = `// dtemplate generated - do not edit
let {{.Class}} = (function() {
{{if not .includeQuerySelect}}
	class joinIterators {
		constructor (iters) {
			this.iters = iters;
			this.i = 0;
		}
		next() {
			if (this.i == this.iters.length) {
				return { 'value':undefined, 'done':true }
			}
			let r = this.iters[this.i].next();
			if (!r.done) {
				return r;
			}
			this.i++;
			return this.next();
		}
		[Symbol.iterator]() {
			return this;
		}
	}

	class querySelectorAllIterator {
		constructor(qs) {
			this.qs = qs;
			this.i = 0;
		}
		next() {
			if (this.i == this.qs.length) {
				return { 'value':undefined, 'done':true }
			}
			return { 'value': this.qs.item(this.i++), 'done':false };
		}
		[Symbol.iterator]() {
			return this;
		}
	}

	let QuerySelectorAllIterate = function(el, query) {
		let els = [];
		if ('function'==typeof el.matches) {
			if (el.matches(query)) {
				els.push(el);
			}
		} else if ('function'==typeof el.matchesSelector) {
			if (el.matchesSelector(query)) {
				els.push(el);
			}
		}
		let qs = el.querySelectorAll(query);
		let i = qs[Symbol.iterator];
		if ('function'==typeof i) {
			return new joinIterators([els[Symbol.iterator](), qs[Symbol.iterator]()])
		}
		return new joinIterators([els[Symbol.iterator](), new querySelectorAllIterator(qs)]);
	}
{{- end}}	

	let templates =
		{{.T | JS}};

	let mk = function(k, html) {
		let el = document.createElement('div');
		el.innerHTML = html;

		let c = el.firstElementChild;
		while ((null!=c) && (Node.ELEMENT_NODE!=c.nodeType)) {
			c = c.nextSibling;
		}
		if (null==c) {
			console.error("FAILED TO FIND ANY ELEMENT CHILD OF ", k, ":", el)
			return mk('error', '<em>No child elements in template ' + k + '</em>');
		}
		el = c;
		let et = el.querySelector('[data-set="this"]');
		if (null!=et) {
			el = et;
			el.removeAttribute('data-set');
		}
		return el;
	}

	for (let i in templates) {
		templates[i] = mk(i, templates[i]);
	}

	return function(t, dest={}) {
		// Return a deep copy of the node
		let n = templates[t];
		if (n.content) {
			// console.log("template " + t + " is a TEMPLATE");
			n = n.content;
		} else {
			// console.log("templates[t] = ", n);
		}
		n = n.cloneNode(true);
		try {
			for (let el of QuerySelectorAllIterate(n, '[data-set]')) {
				let a = el.getAttribute('data-set');
				if (a.substr(0,1)=='$') {
					a = a.substr(1);
					el = jQuery(el);
				}
				dest[a] = el;
			}
		} catch (err) {
			console.error("ERROR in DTemplate(" + t + "): ", err);
			debugger;
		}
		return [n,dest];
	}
})();
`