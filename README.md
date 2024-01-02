 # go-ocr-img

O Projeto é visa apenas entender o quanto é possivel retirar textos de imagens em diversas ocasiões com Tesseract.
Utilizamos apenas a linguagem golang para realização onde tratamos a imagem convertendo para uma escala de cinza
assim a renderização e avaliação por pixel imagem fica mais assertiva em relação a cores. 

O modelo consegue fazer retirada de grande parte de imagens na rede.

O projeto foi realizado através da plataforma samsung dex utilizando os seguintes passos:

* Termux
* Tesseract
* Leptonica
* Golang 1.21.5
* github.com/disintegration/imaging
* github.com/otiai10/gosseract/v2
* github.com/xuri/excelize/v2
* Tesseract Languages

Para utlizar o projeto bastar fazer um git clone deste repositório e mudar as imagens presentes nas pastas imgs .

No final o Arquivo Book1.xls vai ser gerado podendo ser utilizado para o que for necessário.


## O Arquivo

O arquivo gerado vai conter duas colunas, sendo o primeiro o nome da imagem analisada, o segundo texto que foi extraido. 
Algumas imagens podem não ser entendidas perfeitamente, devido ao fundo, tipo de fonte, qualidade da imagem entre outros.
É possível que você adapte as configurações de acordo com o que desejar, assim como incluir mais casos de colunas para arquivo final.

## Tesseract Languages 

Para instalar mais linguagens em relação a OCR treinado para reconhecimento podemos fazer um git do repositório contendo todas as informações necessárias.

1 - Primeiro vamos clonar o repositório com todos os idiomas :
   git clone https://github.com/tesseract-ocr/tessdata

2 - Caso esteja no linux entre no diretório com comando cd e utilize o pwd para buscar todo o caminho da pasta
   EX: baixamos o arquivo na pasta home do usuário
      $ cd tesseract 
      $ pwd 
    
Copiar todo o endereço que aparece.

 3 - Editar o .bashrc para indicar os idiomas.
   
  Para isso vamos utilizar um editor de texto para editar um arquivo oculto na home do usuário criando uma váriavel de ambiente com diretório
  das linguagens OCR.

  no terminal linux 
  $ nano .bashrc 
   
  export TESSDATA_PREFIX=cole aqui o caminho utilizado pelo comando pwd /tessdata
  exemplo : export TESSDATA_PREFIX=/Users/adrianrosebrock/Desktop/tessdata 

 OBS : Dentro do arquivo main ou outro do seu projeto existe uma configuração SetLanguage("por") onde podemos adicionar mais caso precise.
       É importante baixar os arquivos de linguagens, pois, está configurado para portugues diferente do tradicional eng. 

## Rodar o programa

Depois que instalou todas as depêndencias e fez um git do projeto o comando go run main.go vai gerar os arquivos necessários.
