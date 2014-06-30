<!DOCTYPE html>
<html lang="zh-cn">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="description" content="">
    <meta name="author" content="">
    <link rel="shortcut icon" href="/static/img/favicon.png">

    <title>{{.PageTitle}}</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="/static/css/offcanvas.css" rel="stylesheet">
    <link href="/static/css/form.css" rel="stylesheet">

    <!-- Just for debugging purposes. Don't actually copy this line! -->
    <!--[if lt IE 9]>
    <script src="../../docs-assets/js/ie8-responsive-file-warning.js"></script>
    <![endif]-->

    <!-- HTML5 shim and Respond.js IE8 support of HTML5 elements and media queries -->
    <!--[if lt IE 9]>
    <script src="http://cdn.bootcss.com/html5shiv/3.7.0/html5shiv.min.js"></script>
    <script src="http://cdn.bootcss.com/respond.js/1.3.0/respond.min.js"></script>
    <![endif]-->
</head>

<body>
    <div class="navbar navbar-fixed-top navbar-inverse" role="navigation">
        <div class="container" style="width: 100%">
            <div class="navbar-header">
                <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-collapse">
                    <span class="sr-only">Toggle navigation</span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                    <span class="icon-bar"></span>
                </button>
                <a class="navbar-brand" href="/adminhome">Gray Blog</a>
            </div>
            <div class="collapse navbar-collapse">
                <ul class="nav navbar-nav navbar-right">
                    <li>
                        <a href="/adminhome">Welcome, {{.Admin}}</a>
                    </li>
                    <li>
                        <a href="/adminlogin?exit=yes">Sign Out</a>
                    </li>
                </ul>
            </div>
            <!-- /.nav-collapse --> </div>
        <!-- /.container --> </div>
    <!-- /.navbar -->

    <div class="container" style="width: 100%; padding: 0;">
        {{template "elements/admin_menu.tpl" .}}
        {{.LayoutContent}}
    </div>
    <!--/.container-->
</body>
</html>