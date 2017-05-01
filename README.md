# dtemplate
Javascript templating framework, super simple.

## Super Simple

dtemplate takes a directory with a set of html files, and creates a Javascript function that will create instances of the DOM in each of those HTML files, named by the filename. It does this by pre-creating the DOM nodes, then deep-copying them. (I think this is fast.)

## data-set

Besides returning the element, it also returns an object with every DOM element referenced with a `data-set` attribute. If the key for the object in the data-set attribute starts with `$`, it will be `jQuery` wrapped.

The only other catch is that if the DOM element has a data-set attribute of `this`, then that DOM element is returned, not the root-element of the document. The reason for this is that if you want to have a template that is a table-row `tr` for instance, you cannot actually instantiate a tr ourside a table, so you need a template that looks like this:

    <table>
      <tr data-set="this">
      	<td>Name</td>
      	<td data-set="name"> - name will go here - </td>
      </tr>
    </table>

