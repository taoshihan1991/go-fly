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
	<link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/element-ui@2.13.1/lib/theme-chalk/index.css">
    <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/element-ui@2.13.1/lib/index.js"></script>
	<script src="https://cdn.jsdelivr.net/npm/jquery/dist/jquery.min.js"></script>
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
<div id="app" style="width:100%">
    <template>
		<div class="form-signin">
			<h1 class="h3 mb-3 font-weight-normal">邮箱IMAP工具</h1>
			<label for="inputServer" class="sr-only">IMAP服务器:</label>
			<input type="text" v-model="imapServer" name="server" id="inputServer" class="form-control" placeholder="IMAP服务器如imap.sina.net:143" required autofocus>
			<label for="inputEmail" class="sr-only">邮箱地址:</label>
			<input type="email" v-model="imapEmail"  id="inputEmail" name="email" class="form-control" placeholder="邮箱地址" required autofocus>
			<label for="inputPassword" class="sr-only">密码:</label>
			<input type="password" v-model="imapPass"  name="password" id="inputPassword" class="form-control" placeholder="密码" required>
			{{if .}}<div class="alert alert-danger" role="alert">{{.}}</div>{{end}}
			<el-button :loading="false" type="primary" v-on:click="checkEmailPass">登陆</el-button>
			<p class="mt-5 mb-3 text-muted">&copy; 2020</p>
		</div>
</template>           
</div>
</body>
<script>
	new Vue({
		el: '#app',
		data: {
			imapEmail:"",
			imapServer:"",
			imapPass:"",
			loading:false,
		},
		methods: {
			checkEmailPass: function () {
				var data={}
				data.server=this.imapServer;
				data.email=this.imapEmail;
				data.password=this.imapPass;
				let _this=this;
				this.loading=true;
				$.post("/check",data,function(data){
					if(data.code==200){
						_this.$message({
						  message: data.msg,
						  type: 'success'
						});
					}else{
						_this.$message({
						  message: data.msg,
						  type: 'error'
						});
					}
					_this.loading=false;
				});

			}
		}
	})


</script>
</html>
`
	t, _ := template.New("list").Parse(html)
	t.Execute(w, render)
}
