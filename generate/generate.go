// build +generate
package main

import (
	"fmt"
	"os"

	"github.com/unknown321/hashing"
)

var extensions = []string{
	"1.ftexs",
	"1.nav2",
	"2.ftexs",
	"3.ftexs",
	"4.ftexs",
	"5.ftexs",
	"6.ftexs",
	"ag.evf",
	"cc.evf",
	"eng.lng",
	"fre.lng",
	"fx.evf",
	"ger.lng",
	"ita.lng",
	"jpn.lng",
	"por.lng",
	"rus.lng",
	"sd.evf",
	"spa.lng",
	"vo.evf",
	"aia",
	"aib",
	"aibc",
	"aig",
	"aigc",
	"aim",
	"aip",
	"ait",
	"atsh",
	"bnd",
	"bnk",
	"clo",
	"csnav",
	"dat",
	"des",
	"dnav",
	"dnav2",
	"ese",
	"evb",
	"evf",
	"fag",
	"fage",
	"fago",
	"fagp",
	"fagx",
	"fclo",
	"fcnp",
	"fcnpx",
	"fdes",
	"fdmg",
	"ffnt",
	"fmdl",
	"fmdlb",
	"fmtt",
	"fnt",
	"fova",
	"fox",
	"fox2",
	"fpk",
	"fpkd",
	"fpkl",
	"frdv",
	"frig",
	"frt",
	"fsd",
	"fsm",
	"fsml",
	"fsop",
	"fstb",
	"ftex",
	"fv2",
	"fxp",
	"gani",
	"geom",
	"gpfp",
	"grxla",
	"grxoc",
	"gskl",
	"htre",
	"info",
	"json",
	"lad",
	"ladb",
	"lani",
	"las",
	"lba",
	"lng",
	"lpsh",
	"lua",
	"mas",
	"mbl",
	"mog",
	"mtar",
	"mtl",
	"nav2",
	"nta",
	"obr",
	"obrb",
	"param",
	"parts",
	"path",
	"pftxs",
	"ph",
	"phep",
	"phsd",
	"qar",
	"rbs",
	"rdb",
	"rdf",
	"rnav",
	"sad",
	"sand",
	"sani",
	"sbp",
	"sdf",
	"sim",
	"simep",
	"snav",
	"spch",
	"sub",
	"subp",
	"tgt",
	"tre2",
	"txt",
	"uia",
	"uif",
	"uig",
	"uigb",
	"uil",
	"uilb",
	"utxl",
	"veh",
	"vfx",
	"vfxbin",
	"vfxdb",
	"vnav",
	"vpc",
	"wem",
	"wmv",
	"xml",
}

func main() {
	var res = make([]uint64, len(extensions))
	for i, v := range extensions {
		res[i] = hashing.HashFileName(v, false) & 0x1fff
	}

	var out *os.File
	var err error
	if out, err = os.OpenFile("extension.go", os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0644); err != nil {
		panic(err)
	}
	defer out.Close()

	if _, err = out.WriteString(`package hashing
var Extensions = map[string]uint64{
`); err != nil {
		panic(err)
	}

	for i := range res {
		if _, err = out.WriteString(fmt.Sprintf("\"%s\":0x%x,\n", extensions[i], res[i])); err != nil {
			panic(err)
		}
	}

	if _, err = out.WriteString("}\n"); err != nil {
		panic(err)
	}

	out.WriteString("var ExtensionsByHash = map[uint64]string{\n")
	for i := range res {
		if _, err = out.WriteString(fmt.Sprintf("0x%x: \"%s\",\n", res[i], extensions[i])); err != nil {
			panic(err)
		}
	}
	if _, err = out.WriteString("}\n"); err != nil {
		panic(err)
	}
}
