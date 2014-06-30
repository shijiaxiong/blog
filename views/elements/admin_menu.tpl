<div class="col-xs-6 col-sm-3 sidebar-offcanvas" id="sidebar" role="navigation" style="padding-left: 15px; width: 20%">
    <div class="list-group">
        <a href="/adminhome" class="list-group-item {{if .IsHome}}active{{end}}">控制台</a>
        <a href="/admincategory" class="list-group-item {{if .IsCategory}}active{{end}}">分类管理</a>
        <a href="/admintopic" class="list-group-item {{if .IsTopic}}active{{end}}">文章管理</a>
    </div>
</div>