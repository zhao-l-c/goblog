<!DOCTYPE html>

<html>
<head>
  <title>模板使用测试</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <style>
    .title {
      color : #428bca;
    }
  </style>
<body>
  <span class="title">直接获取模板绑定值:</span>
  <p>{{.sayHello}}</p>

  <span class="title">条件判断：</span>
  <p>
    {{if .TrueCond}}
      true condition
    {{end}}
    <br>
    {{if .FalseCond}} 

    {{else}}
      false condition
    {{end}}
  </p>

  <span class="title">获取对象属性：</span>
  <p>
  name: {{.user.Name}}, age: {{.user.Age}}
  </p>

  <span class="title">使用with结构获取对象属性：</span>
  <p>
  {{with .user}} 
    name is {{.Name}}, age is {{.Age}}
  {{end}}
  </p>

  <span class="title">打印数组或者切片：</span>
  <p>
    {{.nums}}
  </p>
  <span class="title">循环打印数组/切片元素：</span>
  <p>
  {{range .nums}} 
    {{.}}<br>
  {{end}}
  </p>

  <span class="title">模板变量：</span>
  <p>
  {{$str := .tplVar}}
  {{$str}}
  </p>

  <span class="title">str2html：</span>
  <p>
  {{str2html .html}} 
  </p>

  <span class="title">htmlquote：</span>
  <p>
  {{htmlquote .html}}
  </p>

  <span class="title">管道：</span>
  <p>{{.html | str2html}}</p>

  <span class="title">定义模板:</span>
  <p>{{template "test"}}</p>

  <br>
  <br>
  <br>
  <br>


</body>
</html>


<!-- 定义一个模板 -->
{{define "test"}} 
  <div>
  <h3>this is a template</h3>
  <div>
{{end}}

