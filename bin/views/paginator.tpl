{{if gt .paginator.PageNums 1}}
<nav aria-label="Page navigation example">

    <ul class="pagination">
        <li class="page-item disabled"><a class="page-link">总:{{.paginator.Nums}}条 / {{.paginator.PageNums}}页</a></li>

        {{if .paginator.HasPrev}}
        <li class="page-item"><a class="page-link" href="{{.paginator.PageLinkFirst}}">首页</a></li>
        <li class="page-item"><a class="page-link" href="{{.paginator.PageLinkPrev}}">&laquo;</a></li>
        {{else}}
        <li class="page-item disabled"><a class="page-link">首页</a></li>
        <li class="page-item disabled"><a class="page-link">&laquo;</a></li>
        {{end}}
        {{range $index, $page := .paginator.Pages}}
        <li class="page-item {{if $.paginator.IsActive .}} active{{end}}">
            <a class="page-link" href="{{$.paginator.PageLink $page}}">{{$page}}</a>
        </li>
        {{end}}
        {{if .paginator.HasNext}}
        <li class="page-item"><a class="page-link" href="{{.paginator.PageLinkNext}}">&raquo;</a></li>
        <li class="page-item"><a class="page-link" href="{{.paginator.PageLinkLast}}">尾页</a></li>
        {{else}}
        <li class="page-item disabled"><a>&raquo;</a></li>
        <li class="page-item disabled"><a>尾页</a></li>
        {{end}}
    </ul>
</nav>
{{end}}