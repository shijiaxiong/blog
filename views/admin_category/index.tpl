<style>
    table th, table td { text-align: center;} 
</style>
<div class="col-xs-12 col-sm-9" style="width: 80%">
    <div class="panel panel-default">
        <div class="panel-heading">{{.PageTitle}}</div>
        <div class="panel-body">
            
            <!-- table start -->
            <table class="table table-striped">
                <thead>
                    <tr>
                        <th>#</th>
                        <th>Title</th>
                        <th>Dateline</th>
                        <th>Views</th>
                        <th>TopicCount</th>
                        <th>Operating</th>
                    </tr>
                </thead>
                <tbody>
                    {{range .Categories}}
                    <tr>
                        <td>{{.Id}}</td>
                        <td>{{.Title}}</td>
                        <td>{{.Dateline}}</td>
                        <td>{{.Views}}</td>
                        <td>{{.TopicCount}}</td>
                        <td><a href="/admincategory?op=del&id={{.Id}}">Del</a></td>
                    </tr>
                    {{end}}
                </tbody>
            </table>

            <!-- form start -->
            <div class="panel panel-default">
                <div class="panel-body">
                    
                    <form class="form-inline" role="form" action="/admincategory" method="get">
                        <div class="input-group input-group-lg">
                            <input type="text" class="form-control" name="title" placeholder="Title" required autofocus>
                            <input type="hidden" name="op" value="add">                     
                            <span class="input-group-addon" style="padding: 0; border: 0;">
                                <button type="submit" class="btn btn-primary btn-lg" style="border-bottom-left-radius: 0; border-top-left-radius: 0">Add</button>
                            </span>
                        </div>
                    </form>
                </div>
            </div>

        </div>
    </div>
</div>