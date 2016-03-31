package main

import (
	"github.com/matwachich/iup"
	"github.com/matwachich/iup/iupcontrols"
)

var img_bits1 = []byte{
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 0, 2, 0, 2, 0, 2, 2, 0, 2, 2, 2, 0, 0, 0, 2, 2, 2, 0, 0, 2, 0, 2, 2, 0, 0, 0, 2, 2, 2,
	2, 2, 2, 0, 2, 0, 0, 2, 0, 0, 2, 0, 2, 0, 2, 2, 2, 0, 2, 0, 2, 2, 0, 0, 2, 0, 2, 2, 2, 0, 2, 2,
	2, 2, 2, 0, 2, 0, 2, 2, 0, 2, 2, 0, 2, 2, 2, 2, 2, 0, 2, 0, 2, 2, 2, 0, 2, 0, 2, 2, 2, 0, 2, 2,
	2, 2, 2, 0, 2, 0, 2, 2, 0, 2, 2, 0, 2, 2, 0, 0, 0, 0, 2, 0, 2, 2, 2, 0, 2, 0, 0, 0, 0, 0, 2, 2,
	2, 2, 2, 0, 2, 0, 2, 2, 0, 2, 2, 0, 2, 0, 2, 2, 2, 0, 2, 0, 2, 2, 2, 0, 2, 0, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 0, 2, 0, 2, 2, 0, 2, 2, 0, 2, 0, 2, 2, 2, 0, 2, 0, 2, 2, 0, 0, 2, 0, 2, 2, 2, 0, 2, 2,
	2, 2, 2, 0, 2, 0, 2, 2, 0, 2, 2, 0, 2, 2, 0, 0, 0, 0, 2, 2, 0, 0, 2, 0, 2, 2, 0, 0, 0, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 2, 2, 2, 0, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 0, 0, 0, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
}

var img_bits2 = []byte{
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 0, 3, 0, 3, 0, 3, 3, 0, 3, 3, 3, 1, 1, 0, 3, 3, 3, 0, 0, 3, 0, 3, 3, 0, 0, 0, 3, 3, 3,
	3, 3, 3, 0, 3, 0, 0, 3, 0, 0, 3, 0, 3, 0, 1, 1, 3, 0, 3, 0, 3, 3, 0, 0, 3, 0, 3, 3, 3, 0, 3, 3,
	3, 3, 3, 0, 3, 0, 3, 3, 0, 3, 3, 0, 3, 3, 1, 1, 3, 0, 3, 0, 3, 3, 3, 0, 3, 0, 3, 3, 3, 0, 3, 3,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	3, 3, 3, 0, 3, 0, 3, 3, 0, 3, 3, 0, 3, 0, 1, 1, 3, 0, 3, 0, 3, 3, 0, 0, 3, 0, 3, 3, 3, 0, 3, 3,
	3, 3, 3, 0, 3, 0, 3, 3, 0, 3, 3, 0, 3, 3, 1, 1, 0, 0, 3, 3, 0, 0, 3, 0, 3, 3, 0, 0, 0, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 0, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 0, 3, 3, 3, 0, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 1, 1, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
}

func createMatix() (mat iup.Ihandle) {
	mat = iupcontrols.Matrix().
		SetAttributes(map[string]interface{}{
		"NUMCOL":         "1",
		"NUMLIN":         "3",
		"NUMCOL_VISIBLE": "1",
		"NUMLIN_VISIBLE": "3",
		"EXPAND":         "NO",
		"SCROLLBAR":      "NOs",
	}).
		SetAttributes(map[string]string{
		"0:0": "Inflation",
		"1:0": "Medicine",
		"2:0": "Food",
		"3:0": "Energy",
		"0:1": "January 2000",
		"1:1": "5.6",
		"2:1": "2.2",
		"3:1": "7.2",
	}).
		SetAttributes(map[string]interface{}{
		"BGCOLOR":    "255 255 255",
		"BGCOLOR1:0": "255 128 0",
		"BGCOLOR2:1": [3]byte{255, 128, 0},
		"FGCOLOR2:0": "255 0 128",
		"FGCOLOR1:1": [3]byte{255, 0, 128},
	}).
		SetAttributes(`CX=600,CY=250`) /*.
	GetIhandle(&mat)*/
	return
}

func createTree() iup.Ihandle {
	return iup.Tree().SetAttributes(`FONT=COURIER_NORMAL_10,
                                     NAME=Figures,
                                     ADDBRANCH=3D,
                                     ADDBRANCH=2D,
                                     ADDLEAF1=trapeze,
                                     ADDBRANCH1=parallelogram,
                                     ADDLEAF2=diamond,
                                     ADDLEAF2=square,
                                     ADDBRANCH4=triangle,
                                     ADDLEAF5=scalenus,
                                     ADDLEAF5=isosceles,
                                     ADDLEAF5=equilateral,
                                     RASTERSIZE=180x180,
                                     VALUE=6,
                                     CTRL=ON,
                                     SHIFT=ON,
                                     CX=600,
                                     CY=10,
                                     ADDEXPANDED=NO`)
}

func example() {
	iup.SetHandle("img1", iup.Image(32, 32, img_bits1).SetAttributes(map[string]string{
		"0": "0 0 0",
		"1": "BGCOLOR",
		"2": "255 0 0",
	}))
	iup.SetHandle("img2", iup.Image(32, 32, img_bits2).SetAttributes(map[string]string{
		"0": "0 0 0",
		"1": "0 255 0",
		"2": "BGCOLOR",
		"3": "255 0 0",
	}))

	frame1 := iup.Frame(
		iup.Vbox(
			iup.Button("Button text").SetAttributes(`CINDEX=1`),
			iup.Button("").SetAttributes(`IMAGE=img1,CINDEX=2`),
			iup.Button("").SetAttributes(`IMAGE=img1,IMPRESS=2,CINDEX=3`),
		),
	).SetAttributes(`TITLE=IupButton,CX=10,CY=180`)

	frame2 := iup.Frame(
		iup.Vbox(
			iup.Label("Label text").SetAttributes(`CINDEX=1`),
			iup.Label("").SetAttributes(`SEPARATOR=HORIZONTAL,CINDEX=2`),
			iup.Label("").SetAttributes(`IMAGE=img1,CINDEX=3`),
		),
	).SetAttributes(`TITLE=IupLabel,CX=200,CY=250`)

	frame3 := iup.Frame(
		iup.Vbox(
			iup.Toggle("Toggle text").SetAttributes(`VALUE=ON,CINDEX=1`),
			iup.Toggle("").SetAttributes(`IMAGE=img1,IMPRESS=img2,CINDEX=2`),
			iup.Frame(
				iup.Radio(
					iup.Vbox(
						iup.Toggle("Toggle text").SetAttributes(`CINDEX=3`),
						iup.Toggle("Toggle text").SetAttributes(`CINDEX=4`),
					),
				),
			).SetAttribute("TITLE", "IupRadio"),
		),
	).SetAttributes(`TITLE=IupToggle,CX=400,CY=250`)

	text1 := iup.Text().SetAttributes(map[string]string{
		"VALUE":  "IupText text",
		"SIZE":   "80x",
		"CINDEX": "1",
		"CX":     "10",
		"CY":     "100",
	})

	ml1 := iup.MultiLine().SetAttributes(map[string]string{
		"VALUE":  "IupMultiline Text\nSecond Line\nThird Line",
		"SIZE":   "80x60",
		"CINDEX": "1",
		"CX":     "200",
		"CY":     "100",
	})

	list1 := iup.List().SetAttributes(map[string]string{
		"VALUE":  "1",
		"1":      "Item 1 Text",
		"2":      "Item 2 Text",
		"3":      "Item 3 Text",
		"CINDEX": "1",
		"CX":     "10",
		"CY":     "10",
	})

	list2 := iup.List().SetAttributes(map[string]string{
		"DROPDOWN": "YES",
		"VALUE":    "2",
		"1":        "Item 1 Text",
		"2":        "Item 2 Text",
		"3":        "Item 3 Text",
		"CINDEX":   "2",
		"CX":       "200",
		"CY":       "10",
	})

	list3 := iup.List().SetAttributes(map[string]string{
		"EDITBOX": "YES",
		"VALUE":   "3",
		"1":       "Item 1 Text",
		"2":       "Item 2 Text",
		"3":       "Item 3 Text",
		"CINDEX":  "3",
		"CX":      "400",
		"CY":      "10",
	})

	cnv1 := iup.Canvas().SetAttributes(map[string]string{
		"RASTERSIZE": "100x100",
		"POSX":       "0",
		"POSY":       "0",
		"BGCOLOR":    "128 255 0",
		"CX":         "400",
		"CY":         "150",
	})

	ctrl1 := iup.Val().SetAttributes(map[string]string{
		"CX": "600",
		"CY": "200",
	})

	cbox := iup.Cbox(
		text1, ml1, list1, list2, list3, cnv1, ctrl1, createTree(), createMatix(), frame1, frame2, frame3,
	).SetAttributes(`SIZE=480x200`)

	hbox := iup.Hbox(cbox).SetAttributes(`MARGIN=10x10`)

	dlg := iup.Dialog(hbox).SetAttributes(`TITLE="Cbox Test"`)
	iup.SetHandle("dlg", dlg)
}

func main() {
	iup.Open()
	iupcontrols.Open()
	defer iup.Close()

	example()

	iup.Show(iup.GetHandle("dlg"))
	iup.MainLoop()
}
