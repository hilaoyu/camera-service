{{template "common/header.tpl" .}}


<div class="d-flex">
    <div class="card my-3" style="width: 8rem">
        {{template "common/left-menu.tpl" .}}
    </div>
    <div class="flex-grow-1 ">
        <div class="card ms-auto me-auto my-3">
            <div class="card-body">
                <h3 class="card-title">录像列表</h3>
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
                        <table class="table">

                            <thead>
                            <tr>

                                <th scope="col">名称</th>
                                <th scope="col">时间</th>
                                <th scope="col">大小</th>
                                <th scope="col">操作</th>
                            </tr>
                            </thead>


                            <tbody>
                            {{range .files}}
                            <tr>

                                <td>{{.SubPath}}</td>
                                <td>{{.ModTimeFormat}}</td>
                                <td>{{.SizeFormat}}</td>
                                <td>
                                    <button type="button" class="btn btn-primary record-play-btn"
                                            data-video-name="{{.SubPath}}"
                                            data-video-url="/console/camera/records/{{.SubPath}}">
                                        播放
                                    </button>

                                    <a href="/console/camera/records/{{.SubPath}}"  class="link-info btn">
                                        下载
                                    </a>

                                </td>

                            </tr>
                            {{end}}
                            </tbody>

                        </table>
                    </div>
                </div>
            </div>
        </div>

    </div>
    <!-- Modal -->
    <div class="modal fade " id="playModal" tabindex="-1" aria-labelledby="playModalLabel">
        <div class="modal-dialog modal-xl">
            <div class="modal-content">
                <div class="modal-header">
                    <h1 class="modal-title fs-5" id="playModalLabel">record file</h1>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                </div>
                <div class="modal-body">
                    <video
                            width="100%"
                            id="camera-record-player"
                            src=""
                            controls autoplay>
                    </video>
                </div>
            </div>
        </div>
    </div>
</div>
{{template "common/footer.tpl" .}}
<script type="text/javascript">
    $(function (){
        var playModal = $("#playModal")
        var playModalTitle = $("#playModalLabel")
        var videoPlayer = $("#camera-record-player");
        $(".record-play-btn").click(function (){
            let videoUrl = $(this).data("video-url")
            let videoName = $(this).data("video-name")


            if ("" == videoUrl) {
                return
            }
            if("" == videoName ){
                videoName = videoUrl.lastIndexOf("/")
            }
            playModalTitle.text(videoName)

            videoPlayer.attr("src",videoUrl)

            playModal.modal('show')
        })

        playModal.on("hide.bs.modal",function (){
            videoPlayer.attr("src","")
        })

    })

</script>