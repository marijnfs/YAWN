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
</head>

<body>
<h3 style="text-align:center;margin: 0px">YAWN/{{.Sub}}</h3>
<a href="/"><u>Root</u></a> | <a href="/{{.Sub}}"><u>View</u></a> | <a href="/logout"><u>Logout</u></a>
<hr/>
<form name="edit-form" action="/edit/{{.Sub}}" method="post">
<input name="button1" type="submit" value="Save" /><br/>
<textarea name="Body" class="markitup" cols="80" rows="50">{{.Body}}</textarea><br/>
<input name="button2" type="submit" value="Save" />
</form>

</body>
</html>
