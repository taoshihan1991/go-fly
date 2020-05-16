package tmpl

import (
	"html/template"
	"net/http"
)

func RenderLogin(w http.ResponseWriter, render interface{}) {
	const html = `
<!doctype html>
<html lang="cn">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="Mark Otto, Jacob Thornton, and Bootstrap contributors">
    <meta name="generator" content="Jekyll v3.8.6">
    <title>邮箱IMAP-登录页</title>
    <!-- Bootstrap core CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">

    <style>
        .bd-placeholder-img {
            font-size: 1.125rem;
            text-anchor: middle;
            -webkit-user-select: none;
            -moz-user-select: none;
            -ms-user-select: none;
            user-select: none;
        }

        @media (min-width: 768px) {
            .bd-placeholder-img-lg {
                font-size: 3.5rem;
            }
        }
        html,
        body {
            height: 100%;
        }

        body {
            display: -ms-flexbox;
            display: flex;
            -ms-flex-align: center;
            align-items: center;
            padding-top: 40px;
            padding-bottom: 40px;
            background-color: #f5f5f5;
        }

        .form-signin {
            width: 100%;
            max-width: 330px;
            padding: 15px;
            margin: auto;
        }
        .form-signin .checkbox {
            font-weight: 400;
        }
        .form-signin .form-control {
            position: relative;
            box-sizing: border-box;
            height: auto;
            padding: 10px;
            font-size: 16px;
        }
        .form-signin .form-control:focus {
            z-index: 2;
        }
        .form-signin input[type="email"] {
            margin-bottom: -1px;
            border-bottom-right-radius: 0;
            border-bottom-left-radius: 0;
        }
        .form-signin input[type="password"] {
            margin-bottom: 10px;
            border-top-left-radius: 0;
            border-top-right-radius: 0;
        }
    </style>

</head>
<body class="text-center">
<form class="form-signin" action="/login" method="post">
    <h1 class="h3 mb-3 font-weight-normal">邮箱IMAP工具</h1>
    <label for="inputServer" class="sr-only">IMAP服务器:</label>
    <input type="text" name="server" id="inputServer" class="form-control" placeholder="IMAP服务器" required autofocus>
    <label for="inputEmail" class="sr-only">邮箱地址:</label>
    <input type="email" id="inputEmail" name="email" class="form-control" placeholder="邮箱地址" required autofocus>
    <label for="inputPassword" class="sr-only">密码:</label>
    <input type="password" name="password" id="inputPassword" class="form-control" placeholder="密码" required>
    {{if .}}<div class="alert alert-danger" role="alert">{{.}}</div>{{end}}
    <button class="btn btn-lg btn-primary btn-block" type="submit">登陆</button>
    <p class="mt-5 mb-3 text-muted">&copy; 2020</p>
</form>
</body>
</html>
`
	t, _ := template.New("list").Parse(html)
	t.Execute(w, render)
}
