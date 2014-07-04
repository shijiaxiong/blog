
<style>
  
  .list-unstyled li { line-height: 28px; border-bottom: dotted 1px #ddd}
</style>
      <div class="row row-offcanvas row-offcanvas-right">


          <div class="col-xs-12 col-sm-9">
          <div class="panel panel-primary">
            <div class="panel-heading">{{.PageTitle}}</div>
              <div class="panel-body">
                <ul class="list-unstyled">
                  {{range .Topics}}
                  <li><span class="pull-right">{{.created}}</span><a href="/topic/view?id={{.id}}">{{.title}}</a></li>
                  {{end}}
                </ul>
              </div>
          </div>
        </div>
          <div class="col-xs-6 col-sm-3 sidebar-offcanvas" id="sidebar" role="navigation">
              <div class="list-group">
                    <a href="#" class="list-group-item active">分类列表</a>
                    {{range .Categories}}
                    <a href="/category?id={{.Id}}" class="list-group-item">{{.Title}}</a>
                    
                    {{end}}
              </div>
          </div>
          <!--/span--> </div>
      <!--/row-->      