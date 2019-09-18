// dtemplate generated - do not edit
export namespace EL {
	export type test =	HTMLUnknownElement;
	
}
export namespace R {
	export interface test {
		name: HTMLDivElement,
		owner: HTMLSpanElement,
		content: HTMLDivElement,
		};
	
}	// end namespace R
export class test {
	public static _template : HTMLUnknownElement;
	public el : HTMLUnknownElement;
	public $ : R.test;
	constructor() {
		let t = test._template;
		if (! t ) {
			let d = document.createElement('div');
			d.innerHTML = `<div class="one">
	<div class="name"> </div>
	<div class="copy">Copyright &copy; 2018 <span>Craig</span></div>
	<div class="content"> </div>
</div>
`;
			t = d.firstElementChild as HTMLUnknownElement;
			test._template = t;
		}
		let n = t.cloneNode(true) as HTMLUnknownElement;
		this.$ = {
			name: n.childNodes[0].childNodes[1] as HTMLDivElement,
			owner: n.childNodes[0].childNodes[3].childNodes[3] as HTMLSpanElement,
			content: n.childNodes[0].childNodes[5] as HTMLDivElement,
		};
		this.el = n;
	}
}

