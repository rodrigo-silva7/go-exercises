package main

import (
	"fmt"
	"io"
	"strings"

	//"reflect"
	"net/http"
	"os"
	"time"

	"io/ioutil"
	"bufio"

   "strconv"
)

const (
   VERSAO = 1.1
   NOME = "Rodrigo"
   DELAY = 2
)

func main() {
   intro()
   for {
      switch awaitInstructions() {
      case 1:
         monitor()
      case 2:
         showLogs()
      case 0:
         fmt.Println("Saindo do programa...")
         os.Exit(0)
      default:
         fmt.Println("Não conheço esse comando!")
         os.Exit(-1)
      }
   }
}


func intro() {
	fmt.Println("Versao:", VERSAO)
	fmt.Println("Olá,", NOME)
	//fmt.Println("type is", reflect.TypeOf(versao))
}

func awaitInstructions() int  {
   var comando int

   fmt.Println("1- Iniciar o monitoramento")
   fmt.Println("2- Exibir logs")
   fmt.Println("0- Sair do programa")

   fmt.Scan(&comando)
   fmt.Println("O comando escolhido foi:", comando)
   return comando
}

func monitor() {
   fmt.Println("Iniciando monitoramento...")

   sites := readSitesfromFile()
   //fmt.Println(reflect.TypeOf(sites))
   //fmt.Println("O tamanho do array é:", len(sites))
   for index, site := range sites {
      testSite(site,index)
      time.Sleep(DELAY * time.Second)
   }
   fmt.Println("")
}

func testSite(site string, index int) {
      fmt.Println(index + 1, "- Monitorando", site)
      resp, err:= http.Get(site)

      if err != nil {
         fmt.Println("Ocorreu um erro:", err)
      }

      if resp.StatusCode == 200 {
         fmt.Println("Site:", site, "foi carregado com sucesso!")
         saveLogs(site,true)
      } else {
         fmt.Println("Site:", site, "está com problemas! StatusCode:", resp.StatusCode)
         saveLogs(site,false)
      }
}

func readSitesfromFile() []string {
   var sites []string
   arquivo, err := os.Open("sites.txt")
   //arquivo, err := ioutil.ReadFile("sites.txt")
   //fmt.Println(string(arquivo))
   if err != nil {
      fmt.Println("Ocorreu um erro:", err)
   }
   reader := bufio.NewReader(arquivo)
   for {
      line, err := reader.ReadString('\n')
      line = strings.TrimSpace(line)
      if err == io.EOF {
         break
      }
      if err != nil {
         fmt.Println("Ocorreu um erro:", err)
      }
      //fmt.Println(line)
      sites = append(sites,line)
   }
   //fmt.Println(sites)
   arquivo.Close()
   return sites;
}

func saveLogs(site string, status bool) {
   file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
   if err != nil {
      fmt.Println("Ocorreu um erro:", err)
   }
   file.WriteString(site + "- online: " + strconv.FormatBool(status) +" - time: "+ time.Now().Format("02/01/2006 15:04:05") + "\n")
   //fmt.Println(file)
   file.Close()
}

func showLogs() {
   fmt.Println("Exibindo logs...")
   file, err := ioutil.ReadFile("log.txt")
   if err != nil {
      fmt.Println("Arquivo não encontrado! erro:", err)
   }
   fmt.Println(string(file))
}

//func showNames() {
   //nomes := []string{"Douglas", "Rafael", "Rodrigo"}
   //fmt.Println(nomes)
   //fmt.Println(reflect.TypeOf(nomes))
   //fmt.Println("O tamanho do slice é:", len(nomes),"e sua capacidade é:", cap(nomes))
//
   //fmt.Println("aumentando slice")
   //nomes = append(nomes, "Aparecida")
   //fmt.Println(nomes)
   //fmt.Println(reflect.TypeOf(nomes))
   //fmt.Println("O tamanho do slice é:", len(nomes),"e sua capacidade é:", cap(nomes))
//}
