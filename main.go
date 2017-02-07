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
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat28.pic.jpg", "番茄小时候"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat27.pic_hd.jpg", "有点丑吗？"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat26.pic.jpg", "是帅吗？"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat25.pic.jpg", "摆个POS!!"},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat24.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat23.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat22.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat21.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat20.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat19.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat18.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat17.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat16.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat15.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat14.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat13.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat12.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat11.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat10.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat9.pic.jpg", ""},
		// CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat8.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat7.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat6.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat5.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat4.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat3.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat2.pic.jpg", ""},
		CatGll{"http://7xku3c.com1.z0.glb.clouddn.com/cat1.pic.jpg", ""},
	}
)
