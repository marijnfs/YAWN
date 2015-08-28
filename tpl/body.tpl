<html>
  <head>
    <title>YAWN/{{.Sub}}</title>

    <style type="text/css">
      body {
      background-color:#A898BC;
      color:#331D4F;
      }
      a:link {color:#2B1B17;}
      a:visited {color:#2B1B17;}
      a:hover {color:#2B1B17;}
      a:active {color:#2B1B17;}
    </style>
    <script type="text/x-mathjax-config">
      MathJax.Hub.Config({tex2jax: {inlineMath: [['$','$']]}});
    </script>
    <script type="text/javascript"
            src="https://cdn.mathjax.org/mathjax/latest/MathJax.js?config=TeX-AMS_HTML">
    </script>
  </head>

  <body>
    <h3 style="text-align:center;margin: 0px">YAWN/{{.Sub}}</h3>
    <a href="/"><u>Root</u></a> | <a href="/edit/{{.Sub}}"><u>Edit</u></a> | <a href="/logout"><u>Logout</u></a>
    <hr/>
    {{.Body}}
    <hr/>

  </body>
</html>
