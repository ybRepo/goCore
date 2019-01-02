package main

import (
	"strings"
	"os"
	"fmt"
	"log"
	"io"
)

func main (){
	name := "Yashar Boosharya"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		<h1> ` + name + ` </h1>
	</body>
	</html>
	`)

	//Create new file nf and check error
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating index file ", err)
	}

	//Close the new file
	defer nf.Close()

	//Copy the the str into the nf document
	io.Copy(nf, strings.NewReader(str))
}

