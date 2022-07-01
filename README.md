# dtemplate - HTML templating without magic, for vanilla JS
Javascript templating framework, super simple. The template framework should do nothing you cannot do by hand. There's no magic, just a bit of processing to create a simple Javascript function you could create yourself.

## Simple Example

`dtemplate -dir src/es6/admin -lang js -name "AdminTemplates" -out src/es6/admin/AdminTemplates.js -separator '/'`

Will create a function `AdminTemplates` in `src/es6/admin/AdminTemplates.js`. You retrieve a DOM Element for any HTML file in `src/esc/admin` by the filename (without .html), for example `AdminTemplates('MainPage')` will return the contents of `src/es6/admin/MainPage.html` as a DOM Element. It also returns an associative array of all Node in that HTML file that are marked with `data-set` attributes.

## Super Simple

dtemplate takes a directory with a set of html files, and creates a single Javascript function that will create instances of the DOM in each of those HTML files, named by the filename. It does this by pre-creating the DOM nodes, then deep-copying them.

## data-set or id

Besides returning the element, it also returns an object with every DOM element referenced with an `id` or `data-set` attribute: `data-set` attributes override `id` attributes, if both are present. If the key for the object in the data-set attribute starts with `$`, it will be `jQuery` wrapped.

### table and 'this'

The one 'catch' is that if the DOM element has a data-set attribute of `this`, then that DOM element is returned, not the root-element of the document. The reason for this is that if you want to have a template that is a table-row `tr`, you cannot actually instantiate a tr ourside a table. Then you need a template that looks like this:

    <table>
      <tr data-set="this">
        <td>Name</td>
        <td data-set="name"> - name will go here - </td>
      </tr>
    </table>

## dtemplate-include and dtemplate-process

You can include another file into the file you are processing with the `dtemplate-include` directive, and you can process that file before inclusion. Simply dtemplate finds the file, and replaces the innerHTML of your tag with the c

````
<template>
  <style dtemplate-include="my-css.scss" dtemplate-process="scss"> </style>
</template>
````

For processing, dtemplate uses a `dtemplate.yml` file in the directory where it runs. It resolves the dtemplate-process attribute in this file, or if the process isn't defined, assumes the process is an executable name itself.

My `dtemplate.yml` looks like this:

````
process:
  scss: 
    exec: sassc -I %.%/src/scss/public/ --style compressed
    prefix: |
      @import "settings";

````
I define the `scss` process, which in my case actually runs *sassc* rather than *scss*, and pass some parameters. Most importantly is the `%.%` macro which expands to the directory where `dtemplate.yml` is found. This allows me to
define an absolute path - dtemplate itself resolves filenames relative to the template being processed.

Note another small feature - the `prefix` allows me to prefix some code before any processing happens. In this case, I 
import some sass from my `_settings.scss` file which I'm defined in `src/scss/public` - so that my site-wide scss matches the scss using in my javascript processed css.

## dtemplate-child

dtemplate also includes child-templates. There's no magic here, a child template is marked with a `dtemplate-child` attribute, and gets the name of it's parent template, a `.` and the value of it's `dtemplate-child` attribute. It is then removed from the parent template. It's just a convenience to avoid having numerous small template files.

Consider `MyTable.html`:

```

<div id="table-container">
  <div class="header">
    <div>Name</div><div>Age</div>
  </div>
  <div dtemplate-child="TableRow">
    <div data-set="name"> </div><div data-set="age"> </div>
  </div>
</div>
```

This is exactly the same as having 2 template files:

MyTable.html:
```

<div id="table-container">
  <div class="header">
    <div>Name</div><div>Age</div>
  </div>
</div>
```
and MyTable.TableRow.html:
````
<div dtemplate-child="TableRow">
  <div data-set="name"> </div><div data-set="age"> </div>
</div>
````
