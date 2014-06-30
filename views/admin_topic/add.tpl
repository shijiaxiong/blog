<div class="col-xs-12 col-sm-9" style="width: 80%">
    <div class="panel panel-default">
        <div class="panel-heading">{{.PageTitle}}</div>
        <div class="panel-body">
            
            <form action="/admintopic" method="post">
                <table class="table table-striped">
                    <tbody>
                    <tr>
                        <td align="right" style="vertical-align:middle;">Category:</td>
                        <td>
                            <input type="hidden" name="op" value="add">
                            <select class="form-control input-lg" name="category">
                                <!-- <option value="0">Select Category</option> -->
                                {{range .Categories}}
                                    <option value="{{.Id}}">{{.Title}}</option>
                                {{end}}
                            </select>
                            </td>
                    </tr>
                        <tr>
                            <td width="10%" align="right" style="vertical-align:middle;">Title:</td>
                            <td>
                                <div class="input-group-lg">
                                    <input type="text" name="title" class="form-control" id="Title" placeholder="Title" required autofocus></div>
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
                                    <button type="submit" class="btn btn-lg btn-primary">Add</button>
                                </div>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </form>
        </div>
    </div>
</div>
</div>