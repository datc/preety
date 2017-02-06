package main

import (
	"html/template"
	"fmt"

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
			ctx.SetTemplatePath("","gallery")
			ctx.Data["gattery_cat"] = gattery_cat
			ctx.HTML(200,"cat")
		})
	})

	m.Run(8080)
}


var(
	gattery_cat = []string{
		"http://7xku3c.com1.z0.glb.clouddn.com/cat28.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat27.pic_hd.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat26.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat25.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat24.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat23.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat22.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat21.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat20.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat19.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat18.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat17.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat16.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat15.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat14.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat13.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat12.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat11.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat10.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat9.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat8.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat7.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat6.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat5.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat4.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat3.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat2.pic.jpg",
		"http://7xku3c.com1.z0.glb.clouddn.com/cat1.pic.jpg",
	}
)
