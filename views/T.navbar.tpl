{{define "navbar"}}
<div class="navbar navbar-inverse navbar-fixed-top" >
    <div class="container">
        <a class="navbar-brand pull-left" href="/">PasswordKeeper</a>
        <ul class="nav navbar-nav pull-right">
            {{if .IsLogin}}
            <li class="active"><a href="/{{.UserUrl}}">{{.UserName}}</a></li>
            {{end}}
            <li {{if .IsAbout}}class="active"{{end}}><a href="/">About</a></li>
            {{if .IsLogin}}
            <li {{if .IsInOut}} class="active" {{end}}><a href="/login?exit=true">Sign out</a></li>
            {{else}}
            <li {{if .IsInOut}} class="active" {{end}}><a href="/login">Sign in</a></li>
            {{end}}
            <li {{if .IsUp}} class="active" {{end}}><a href="/register">Sign up</a></li>
        </ul>
    </div>
</div>
{{end}}