<link rel="stylesheet" href="/static/css/error.css">
<div class="panel panel-primary error">
    <div class="panel-heading">
        <h3 class="panel-title">提示信息</h3>
    </div>
    <div class="panel-body">
        <h3>
            {{.Message}}
            ，
            <span id="redirecttime">
                {{.Time}}</span>
            秒后返回，或
            <a href="{{.Url}}">点击这里</a>
        </h3>
    </div>
</div>


<script type="text/javascript">
var i = {{.Time}}
var timer = null;
var redirecttime = document.getElementById('redirecttime');
var timer = setInterval(function(){
    i--;
    redirecttime.innerHTML = i;
    if(i == 0)
    {
        location.href = "{{.Url}}";
        clearInterval(timer);
    }
},1000);
</script>
