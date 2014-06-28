

<form class="form-signin" role="form" method="post" action="/adminlogin">

    <h2 class="form-signin-heading">Please sign in</h2>
    <h6 id="msg" style="color: red">{{.Msg}}</h6>
    <input type="text" class="form-control" id="name" name="admin" placeholder="Admin" required autofocus>
    <input type="password" class="form-control" id="pwd" name="password" placeholder="Admin Password" required>

    <button class="btn btn-lg btn-primary btn-block" type="submit">Sign in</button>
</form>

<script>
    
    var oMsg = document.getElementById("msg");
    var oName = document.getElementById("name");
    var oPwd = document.getElementById("pwd");

    oName.onkeyup = function ()
    {
        oMsg.innerHTML = '';
    }

    oPwd.onkeyup = function ()
    {
        oMsg.innerHTML = '';
    }
</script>