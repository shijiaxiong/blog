

      <div class="row row-offcanvas row-offcanvas-right">

          <div class="col-xs-12 col-sm-9">
              <p class="pull-right visible-xs">
                  <button type="button" class="btn btn-primary btn-xs" data-toggle="offcanvas">Toggle nav</button>
              </p>
              <div class="jumbotron">
                  <h1>Hello, world!</h1>
                  <p>
                      This is an example to show the potential of an offcanvas layout pattern in Bootstrap. Try some responsive-range viewport sizes to see it in action.
                  </p>
              </div>
              <div class="row">
                  {{range .Topics}}
                  <div class="col-6 col-sm-6 col-lg-4">

                      <h2>{{.title}}</h2>
                      
                      <p>
                          {{.content}}.
                      </p>
                      <p>
                          <a class="btn btn-default" href="/topic/view?id={{.id}}" role="button">View details</a>
                      </p>
                      
                  </div>
                  {{end}}
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