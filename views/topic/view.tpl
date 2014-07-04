 <div class="row row-offcanvas row-offcanvas-right">

     <div class="col-xs-12 col-sm-9">
     <div class="panel panel-primary">
        {{range .Topic}}
        <div class="panel-heading">
        
            {{.title}} <small>Created: {{.created}}</small>
        </div>
        <div class="panel-body">
        {{.content}}
        </div>
        {{end}}
    </div>

<div class="panel panel-success">
<div class="panel-heading">Replys</div>
        <div class="panel-body">
            {{range .Replys}}
             <table class="table table-striped">
                 <tbody>
                     <tr>
                         <td>ReplyTitle: {{.title}}</td>
                         <td>NickName: {{.nick_name}}</td>
                         <td>ReplyTime: {{.reply_time}}</td>
                     </tr>
                     <tr>
                         <td>ReplyContent:</td>
                         <td colspan="3">{{.content}}</td>
                     </tr>
                 </tbody>
             </table>
             {{end}}
        </div>
</div>

<div class="panel panel-warning">
<div class="panel-heading">SubmitReplys</div>
        <div class="panel-body">
             
                 <form action="/topic" method="post">
                <table class="table table-striped">
                    <tbody>
                        <tr>
                            <td width="10%" align="right" style="vertical-align:middle;">NickName:</td>
                            <td>
                                {{range .Topic}}
                                <input type="hidden" name="tid" value="{{.id}}">
                                {{end}}
                                <div class="input-group-lg">
                                    <input type="text" name="nickname" class="form-control" id="NickName" placeholder="NickName" required autofocus></div>
                            </td>
                        </tr>
                        <tr>
                            <td width="10%" align="right" style="vertical-align:middle;">ReplyTile:</td>
                            <td>
                                <div class="input-group-lg">
                                    <input type="text" name="title" class="form-control" id="ReplyTile" placeholder="ReplyTile" required autofocus></div>
                            </td>
                        </tr>
                        <tr>
                            <td align="right" style="vertical-align:middle;">Content:</td>
                            <td>
                                <div class="form-group">
                                    <textarea name="content" id="" cols="30" rows="10" class="form-control" placeholder="Content" required autofocus></textarea>
                                </div>
                            </td>
                        </tr>
                        <tr>
                            <td></td>
                            <td>
                                <div class="form-group">
                                    <button type="submit" class="btn btn-lg btn-primary">Submit</button>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </form>
             
        </div>
        </div>
     </div>


    <!-- 同分类文章 -->
    <div class="col-xs-6 col-sm-3 sidebar-offcanvas" id="sidebar" role="navigation">
        <div class="list-group">
            <a href="#" class="list-group-item active">同分类文章</a>
            {{range .CategoryTopics}}
            <a href="/topic/view?id={{.id}}" class="list-group-item">{{.title}}</a>
            {{end}}
        </div>
    </div>

 </div>