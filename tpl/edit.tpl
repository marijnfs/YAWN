<html>
<head>
<title>{{.Sub}}</title>
<style type="text/css">
body{
background-color:#B8A8CC;
color:#331D4F;
}
a:link {color:#2B1B17;}
a:visited {color:#2B1B17;}
a:hover {color:#2B1B17;}
a:active {color:#2B1B17;}

</style>
<!-- wysihtml5 parser rules -->
<script src="/static/xing-wysihtml5-fb0cfe4/parser_rules/advanced.js"></script>
<!-- Library -->
<script src="/static/xing-wysihtml5-fb0cfe4/dist/wysihtml5-0.3.0.min.js"></script>
</head>

<body>
<h3 style="text-align:center;margin: 0px">YAWN/{{.Sub}}</h3>
<a href="/"><u>Root</u></a> | <a href="/{{.Sub}}"><u>View</u></a> | <a href="/logout"><u>Logout</u></a>
<hr/>

<form name="edit-form" action="/edit/{{.Sub}}" method="post">
<div id="wysihtml5-toolbar" style="display: none;">
  <a data-wysihtml5-command="bold">bold</a>
  <a data-wysihtml5-command="italic">italic</a>

  <!-- Some wysihtml5 commands require extra parameters -->
  <a data-wysihtml5-command="foreColor" data-wysihtml5-command-value="red">red</a>
  <a data-wysihtml5-command="foreColor" data-wysihtml5-command-value="green">green</a>
  <a data-wysihtml5-command="foreColor" data-wysihtml5-command-value="blue">blue</a>

  <!-- Some wysihtml5 commands like 'createLink' require extra paramaters specified by the user (eg. href) -->
<a data-wysihtml5-command="formatBlock" data-wysihtml5-command-value="h1">h1</a> |
    <a data-wysihtml5-command="formatBlock" data-wysihtml5-command-value="h2">h2</a> |
    <a data-wysihtml5-command="insertUnorderedList">insertUnorderedList</a> |
    <a data-wysihtml5-command="insertOrderedList">insertOrderedList</a> |
  <a data-wysihtml5-command="createLink">insert link</a>
  <div data-wysihtml5-dialog="createLink" style="display: none;">
    <label>
      Link:
      <input data-wysihtml5-dialog-field="href" value="http://" class="text">
    </label>
    <a data-wysihtml5-dialog-action="save">OK</a> <a data-wysihtml5-dialog-action="cancel">Cancel</a>
  </div>
</div>
<input name="button1" type="submit" value="Save" /><br/>
<!--<textarea name="Body" class="markitup" cols="80" rows="50">{{.Body}}</textarea><br/>-->
<input name="button2" type="submit" value="Save" />
<textarea name="Body" class="wysiwyg-color-fuchsia" id="wysihtml5-textarea" autofocus="autofocus" cols="80" rows="50">{{.Body}}</textarea><br/>

</form>





<script>
var editor = new wysihtml5.Editor("wysihtml5-textarea", { // id of textarea element
  toolbar:      "wysihtml5-toolbar", // id of toolbar element
  parserRules:  wysihtml5ParserRules // defined in parser rules set 
});
</script>
</body>
</html>
