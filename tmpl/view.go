package tmpl

import (
	"github.com/taoshihan1991/imaptool/tools"
	"html/template"
	"net/http"
)

func RenderView(w http.ResponseWriter, render interface{}) {
	const html = `
<!doctype html>
<html lang="cn">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <meta name="description" content="">
    <meta name="author" content="Mark Otto, Jacob Thornton, and Bootstrap contributors">
    <meta name="generator" content="Jekyll v3.8.6">
    <title>邮箱IMAP-详情</title>
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
		<table class="table table-hover">
		  <tbody>
			<tr>
			  <th scope="row" width="100">日期:</th>
			  <td>{{.Date}}</td>
			</tr>
			<tr>
			  <th scope="row">发件人:</th>
			  <td>{{.From}}</td>
			</tr>
			<tr>
			  <th scope="row">收件人:</th>
			  <td>{{.To}}</td>
			</tr>
			<tr>
			  <th scope="row">主题:</th>
			  <td>{{.Subject}}</td>
			</tr>
			<tr>
			  <th scope="row">内容:</th>
			  <td>
							{{.HtmlBody}}
              </td>
			</tr>
		  </tbody>
		</table>    
    </div>
</div>
</body>
</html>
`
	html1 := tools.FileGetContent("html/view.html")
	t, _ := template.New("view").Parse(html1)
	t.Execute(w, render)
}
