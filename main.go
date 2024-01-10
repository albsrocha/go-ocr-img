package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/gift"
	"github.com/disintegration/imaging"
	"github.com/otiai10/gosseract/v2"
	"github.com/xuri/excelize/v2"
)

type chatData struct{
  name string
  texto string
}

func main(){


   fs := imgGray() 
  
  var data = []chatData{}
  exec(fs, data)
  
}

func exec(fs []string, data []chatData){
  for x:= 1;x <= len(fs); x++{
  nweData := chatData{fs[x-1], imgReader(fs[x-1])}
  data = append(data, nweData)
  }
  excel(data)
}

func errorHandle(err *error){
   if *err != nil{
  log.Println("O error gerado:", err)
  }else{}
  }

func imgReader(name string) string{
  client := gosseract.NewClient()
  client.SetLanguage("por")
  client.SetPageSegMode(1)
  defer client.Close()
  client.SetImage(fmt.Sprint(name))
	
   text, _ := client.Text() 
   return fmt.Sprint(text)
}


func excel(data []chatData){

f := excelize.NewFile()
  defer func() {
     err := f.Close()
   errorHandle(&err)
  }()
 //  Create a new sheet.
   //Set value of a cell.
  

  for x:=1; x <=len(data); x++{
    
  f.SetCellValue("Sheet1",fmt.Sprint("A", x), data[x-1].name)
  f.SetCellValue("Sheet1", fmt.Sprint("B", x), data[x-1].texto)
  }
   
    //Set active sheet of the workbook.

  // Save spreadsheet by the given path.
    if err := f.SaveAs("Book1.xlsx"); err != nil {
      fmt.Println(err)
    }
}


func imgGray() []string{
 fs, err := os.ReadDir("imgs")
 errorHandle(&err)
 
 var data []string
 
 for x:= 1; x <= len(fs); x++{

  src := loadImage(fs[x-1].Name())
  g := gift.New(
    gift.Sigmoid(0.5,7),
    gift.Grayscale(),
    gift.Contrast(10),
   )

   dst := image.NewRGBA(g.Bounds(src.Bounds()))
   g.DrawAt(dst, src, dst.Bounds().Min, gift.CopyOperator)
   saveImage(fmt.Sprint("imgs/","gray",fs[x-1]), dst)
   data = append(data, fmt.Sprint("imgs/","gray",fs[x-1]))
 }
 return data
}



func loadImage(filename string) image.Image {
	f, err := os.Open(fmt.Sprint("imgs/",filename))

	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	img,  err := imaging.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("os.Create failed: %v", err)
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		log.Fatalf("png.Encode failed: %v", err)
	}
}