package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func last(name string, direct string,f bool)  (bool){
	rr:=0
	files, err := ioutil.ReadDir(direct)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Name()==name {
			rr++
		}
		if rr>0 {
			if f {
				rr++
			} else {
				if file.IsDir() {
					rr++
				}
			}
		}
	}
	rr-=1
	if rr==1{
		return  true
	}else{
		return false
	}
}

func print(name string, direct string, bumper string, f bool,size int64,PrintSize bool) {
	if last(name,direct,f) {
		bumper  += "└───"
	} else {
		bumper  += "├───"
	}
	if PrintSize{
		fmt.Println(bumper+name,"(",size,")")
	}else {
		fmt.Println(bumper + name)
	}
}

func DirTree(direct string,bumper string, f bool) {
	files, err := ioutil.ReadDir(direct)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		BackBumper:=bumper
		GoBack:=direct
		if file.IsDir()==true {
			print(file.Name(),direct,bumper,f,0,false)
			if last(file.Name(),direct,f){
				bumper+=" \t"
			}else{
				bumper+="│\t"
			}
			direct += "\\" + file.Name()
			DirTree(direct,bumper, f)
			direct=GoBack
		}
		if f && file.IsDir()==false {
			print(file.Name(), direct, bumper,f, file.Size(),true)
		}
		bumper = BackBumper
	}
}

func main() {
	var direct string
	var option string
	fmt.Print("Enter directory: ")
	fmt.Fscan(os.Stdin, &direct)
	fmt.Print("Enter option: ")
	fmt.Fscan(os.Stdin, &option)

	if option=="-f" {
		fmt.Println("go run main.go . -f")
		DirTree(direct,"\t", true)
	}else {
		fmt.Println("go run main.go . ")
		DirTree(direct,"\t", false)
	}
}