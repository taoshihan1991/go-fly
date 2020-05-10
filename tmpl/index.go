package tmpl

import (
	"html/template"
	"net/http"
)

func RenderList(w http.ResponseWriter,render interface{}){
	const html = `
<!doctype html>
<html lang="cn">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="Mark Otto, Jacob Thornton, and Bootstrap contributors">
    <meta name="generator" content="Jekyll v3.8.6">
    <title>邮箱IMAP-首页</title>
    <!-- Bootstrap core CSS -->
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.4.1/dist/css/bootstrap.min.css" integrity="sha384-Vkoo8x4CGsO3+Hhxv8T/Q5PaXtkKtu6ug5TOeNV6gBiFeWPGFN9MuhOf23Q9Ifjh" crossorigin="anonymous">

    <style>
        body{
            padding: 10px;
        }
    </style>

</head>
<body>
<div class="row">
    <div class="col-md-3">
        <ul class="list-group">
            <li class="list-group-item active">邮件夹</li>
            {{ range $key, $value := .Folders}}
            <li class="list-group-item d-flex justify-content-between align-items-center">
                <a href="/list?fid={{$key}}">{{$key}}</a>
                {{if ne $value 0 }}<span class="badge badge-primary badge-pill">{{$value}}</span>{{end}}
            </li>
            {{end}}
        </ul>
    </div>
    <div class="col-md-9">
        <ul class="list-group">
            <li class="list-group-item active">[{{.Fid}}]邮件列表</li>
            {{ range $key, $value := .Mails}}
                <li class="list-group-item d-flex justify-content-between align-items-center">
                    {{$value}}{{if eq $value "" }}无标题{{end}}
                </li>
            {{end}}
        </ul>
        <nav aria-label="..." style="margin:20px 0;">
            <ul class="pagination">
                <li class="page-item">
                    <a class="page-link" href="{{.PrePage}}">Previous</a>
                </li>
                {{.NumPages}}
                <li class="page-item">
                    <a class="page-link" href="{{.NextPage}}">Next</a>
                </li>
            </ul>
        </nav>
    </div>
</div>
</body>
</html>
`
	t, _ := template.New("list").Parse(html)
	t.Execute(w, render)
}
