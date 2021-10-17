package dtemplate

var jsTemplate = `// dtemplate generated - do not edit
let {{.Class}} = (function() {	
	let templates =
		{{.T | JS}};

	let mk = function(k, html) {
		let el = document.createElement('div');
		el.innerHTML = html;
		//console.log("mk(",k,") html = ", html);

		let c = el.firstElementChild;
		while ((null!=c) && (Node.ELEMENT_NODE!=c.nodeType)) {
			c = c.nextSibling;
		}
		if (null==c) {
			console.error("FAILED TO FIND ANY ELEMENT CHILD OF ", k, ":", el)
			return mk('error', '<em>No child elements in template ' + k + '</em>');
		}
		el = c;
		if ('function'==typeof el.querySelector) {
			let et = el.querySelector('[data-set="this"]');
			if (null!=et) {
				el = et;
				el.removeAttribute('data-set');
			}
		}
		return el;
	}

	return function(t, dest={}) {
		// Return a deep copy of the node, created on first use
		let n = templates[t];
		if ('string'==typeof(n)) {			
			n = mk(t, n);
			templates[t] = n;
		}
		if (n.content) {
			n = n.content.cloneNode(true);
		} else {
			n = n.cloneNode(true);
		}
		try {
			for (let attr of ['id', 'data-set']) {
				let nodes = Array.from(n.querySelectorAll('[' + attr + ']'));
				if (n.hasAttribute(attr)) {
					nodes = nodes.unshift(n);
				}
				for (let el of nodes) {
					let a = el.getAttribute(attr);
					if (a.substr(0,1)=='$') {
						a = a.substr(1);
						el = jQuery(el);
						el.setAttribute(attr, a);
					}
					dest[a] = el;
				}
			}
		} catch (err) {
			console.error("ERROR in DTemplate(" + t + "): ", err);
			debugger;
		}
		return [n,dest];
	}
})();
`