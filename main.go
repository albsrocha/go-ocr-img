 package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
   "path/filepath"
	"os"
	"strings"
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
  return strings.ReplaceAll(text, "\n", "")
  
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
    imgPath := fmt.Sprint("imgs/",fs[x-1].Name())
    fs, err := os.Open(imgPath) 
    errorHandle(&err) 
    defer fs.Close() 
    img,_, err:= image.Decode(fs)
    errorHandle(&err)
    size := img.Bounds().Size()  
    rect := image.Rect(0, 0, size.X, size.Y)  
    wImg := image.NewRGBA(rect)
    // loop though all the x  
    for x := 0; x < size.X; x++ { 
        // and now loop thorough all of this x's y 
        for y := 0; y < size.Y; y++ { 
            pixel := img.At(x, y) 
            originalColor := color.RGBAModel.Convert(pixel).
                (color.RGBA) 
            // Offset colors a little, adjust it to your taste 
            r := float64(originalColor.R) * 0.92126 
            g := float64(originalColor.G) * 0.97152 
            b := float64(originalColor.B) * 0.90722 
            // average
            grey := uint8((r + g + b) / 3) 
            c := color.RGBA{ 
                R: grey, G: grey, B: grey, A: originalColor.A, 
            } 
            wImg.Set(x, y, c)
            
        } 
    }
    ext := filepath.Ext(imgPath)  
name := strings.TrimSuffix(filepath.Base(imgPath), ext)  
newImagePath := fmt.Sprintf("%s/%s_gray%s", filepath.Dir(imgPath), name, ext)  
fg, err := os.Create(newImagePath)  
defer fg.Close()  
errorHandle(&err)  
  err = jpeg.Encode(fg,imaging.Invert(wImg), nil)  
errorHandle(&err)

data = append(data, newImagePath)
  }
 return data
}

