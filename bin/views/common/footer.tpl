

</div>

{{if .pageMessage.success}}
<div class="alert alert-success w-100 text-center  alert-dismissible fade show position-fixed fixed-top"  role="alert" >
  {{.pageMessage.success}}
  <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
</div>
{{end}}

{{if .pageMessage.notice}}
<div class="alert alert-info w-100 text-center  alert-dismissible fade show position-fixed fixed-top" role="alert" style="width: 24rem;">
  {{.onceMessage.notice}}
  <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
</div>
{{end}}
{{if .pageMessage.warning}}
<div class="alert alert-warning w-100 text-center  alert-dismissible fade show position-fixed fixed-top" role="alert" style="width: 24rem;">
  {{.pageMessage.warning}}
  <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
</div>
{{end}}

{{if .pageMessage.error}}
<div class="alert alert-danger w-100 text-center  alert-dismissible fade show position-fixed fixed-top" role="alert" style="width: 24rem;">
  {{.pageMessage.error}}
  <button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button>
</div>
{{end}}



<script src="/static/js/jquery.min.js"></script>
  <script src="/static/js/bootstrap.bundle.min.js"></script>

</body>
</html>
