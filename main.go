package main

import (
	"fmt"
	"html/template"

	"github.com/go-macaron/macaron"
)

func main() {
	m := macaron.Classic()
	m.Use(macaron.Static("gallery"))
	m.Use(macaron.Renderer(macaron.RenderOptions{
		Funcs: []template.FuncMap{
			{},
		},
	}))
	m.Group("/", func() {
		m.Get("", func(ctx *macaron.Context) {
			fmt.Println("This is toukii,cat")
			ctx.SetTemplatePath("", "gallery")
			ctx.Data["gallery_cat"] = gallery_cat
			ctx.HTML(200, "cat")
		})
	})

	m.Run(8080)
}

type CatGll struct {
	Photo string
	Desc  string
}

var (
	gallery_cat = []CatGll{
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat28.pic.jpg", `我的名字叫番茄，
因为被捡到的那天，主银吃了这道国菜：番茄炒蛋。
我是一只流浪猫，
脖子上戴着一环粉红色的项圈；
据说我那时有1个多月大，看上去不太干净。Orz...`},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat27.pic_hd.jpg", "有点丑吗？"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat26.pic.jpg", "“我可能是帅的不太明显，而且那只是从前吖。”"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat25.pic.jpg", "“我觉得作为一只喵，长得怎么样是次要的，关键还是会卖萌，会摆POS（也会讲英语哦😯）。”😊"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat24.pic.jpg", "“每个季末的下午，我稀饭坐着思考喵生，比如‘我的妈妈在哪里？’，‘我还有没有哥哥或者妹妹，TA有没有和我一样有吃有喝？’”"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat23.pic.jpg", "“主银，请上传我的漂亮照片，麻烦把我不喜欢的照片删掉。阿里嘎肚！”"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat22.pic.jpg", "“嗯，我最喜欢这个故事是：<strong>丑小鸭</strong>，我的内心一定住着一只漂亮的白天鹅。”"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat21.pic.jpg", "随着时间的流逝，番茄渐渐长出了羽毛，哦不，是毛发渐渐变顺变亮了。"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat20.pic.jpg", "主银给她换了一个大红色的项圈，虽然她对粉色留恋，但有新项圈戴不要太开心啊。<br>“番茄，你笑起来比哭都难看。😄” “😯”"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat19.pic.jpg", "“我笑起来比哭都难看...”<br>“我笑起来比哭都难看...”<br>“我笑起来比哭都难看...”"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat18.pic.jpg", "“我明白啦，那么那我哭起来是不是比笑好看？😊😊😊😊😊😊😊😊😊😊😊😊😊😊😊，😭😭😭😭😭😭😭😭😭😭😭😭😭”<br>“你哭哭／开心就好。”"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat17.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat16.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat15.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat14.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat13.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat12.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat11.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat10.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat9.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat7.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat6.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat5.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat4.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat3.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat2.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat1.pic.jpg", ""},
	}
)
