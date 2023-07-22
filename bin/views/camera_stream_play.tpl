{{template "common/header.tpl" .}}


<div class="d-flex">
    <div class="card my-3" style="width: 8rem">
        {{template "common/left-menu.tpl" .}}
    </div>
    <div class="flex-grow-1 ">
        <div class="card ms-auto me-auto my-3">
            <div class="card-body">
                <h3 class="card-title">实时观看</h3>
                <div>
                    <form class="" method="get">

                        <div class="row mt-2">

                            <!--<div class="col input-group">
                                <button type="submit" class="btn btn-primary btn-sm">
                                    搜索
                                </button>
                            </div>-->

                        </div>


                    </form>
                </div>

            </div>

            <div>

                <div class="card mb-2">
                    <div class="card-body">
                        <iframe src="/console/camera/stream" width="{{.frameWidth}}" height="{{.frameHeight}}"></iframe>
                    </div>
                </div>
            </div>
        </div>

    </div>

</div>
{{template "common/footer.tpl" .}}
