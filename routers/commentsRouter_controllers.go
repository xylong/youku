package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["youku/controllers:BaseController"] = append(beego.GlobalControllerRouter["youku/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelRegion",
            Router: "/channel/region",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:BaseController"] = append(beego.GlobalControllerRouter["youku/controllers:BaseController"],
        beego.ControllerComments{
            Method: "ChannelType",
            Router: "/channel/type",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:CommentController"] = append(beego.GlobalControllerRouter["youku/controllers:CommentController"],
        beego.ControllerComments{
            Method: "List",
            Router: "/comment/list",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:CommentController"] = append(beego.GlobalControllerRouter["youku/controllers:CommentController"],
        beego.ControllerComments{
            Method: "Save",
            Router: "/comment/save",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["youku/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["youku/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["youku/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["youku/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:ObjectController"] = append(beego.GlobalControllerRouter["youku/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:objectId",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:TopController"] = append(beego.GlobalControllerRouter["youku/controllers:TopController"],
        beego.ControllerComments{
            Method: "ChannelTop",
            Router: "/channel/top",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:TopController"] = append(beego.GlobalControllerRouter["youku/controllers:TopController"],
        beego.ControllerComments{
            Method: "TypeTop",
            Router: "/type/top",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:UserController"] = append(beego.GlobalControllerRouter["youku/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: "/login/do",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:UserController"] = append(beego.GlobalControllerRouter["youku/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/register/save",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelAdvert",
            Router: "/channel/advert",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelHotList",
            Router: "/channel/hot",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelRecommendList",
            Router: "/channel/recommend/region",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "GetChannelRecommendTypeList",
            Router: "/channel/recommend/type",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "ChannelVideo",
            Router: "/channel/video",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoEpisodesList",
            Router: "/video/episodes/list",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["youku/controllers:VideoController"] = append(beego.GlobalControllerRouter["youku/controllers:VideoController"],
        beego.ControllerComments{
            Method: "VideoInfo",
            Router: "/video/info",
            AllowHTTPMethods: []string{"*"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
