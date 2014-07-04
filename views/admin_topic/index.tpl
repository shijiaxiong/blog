<style>
    table td, table th { text-align: center;}
</style>
<div class="col-xs-12 col-sm-9" style="width: 80%">
    <div class="panel panel-default">
        <div class="panel-heading">
            {{.PageTitle}} 
            <a href="/admintopic/add" class="btn btn-primary" role="button">Add</a>
        </div>
        <div class="panel-body">
             <table class="table table-striped">
                 <thead>
                     <tr>
                         <th>#</th>
                         <th>Categroy</th>
                         <th>Title</th>
                         <th>Content</th>
                         <th>Created</th>
                         <th>Updated</th>
                         <th>Views</th>
                         <th>Author</th>
                         <th>Reply<br>Time</th>
                         <th>Reply<br>Count</th>
                         <th>ReplyLast<br>UserId</th>
                         <th>Operating</th>
                     </tr>
                 </thead>
                 <tbody>
                    {{range .Topics}}
                     <tr>
                         <td>{{.id}}</td>
                         <td>{{.CateTitle}}</td>
                         <td>{{.title}}</td>
                         <td>{{.content}}</td>
                         <td>{{.created}}</td>
                         <td>{{.updated}}</td>
                         <td>{{.views}}</td>
                         <td>{{.author}}</td>
                         <td>{{.reply_time}}</td>
                         <td>{{.reply_count}}</td>
                         <td>{{.reply_last_user_id}}</td>
                         <td><a href="/admintopic/edit?id={{.id}}">Edit</a>　<a href="">Del</a></td>
                     </tr>
                     {{end}}
                 </tbody>
             </table>
            
        </div>
    </div>
</div>