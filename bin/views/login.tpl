{{template "common/header.tpl" .}}


  <div class="card ms-auto me-auto my-3" style="width: 24rem;">
    <div class="card-body">
      <h5 class="card-title text-center">请输入密码</h5>
      <div>
        <form method="post" >
          <div class="mb-3">
          <input class="form-control" type="password" name="password" placeholder="请输入密码" >
          </div>
          <div class="mb-3 text-end">
          <button type="submit" class="btn btn-primary">确认</button>
          </div>
        </form>
      </div>
    </div>
  </div>

{{template "common/footer.tpl" .}}