package main

import (
	"fmt"

	art "github.com/kelvinCB/FirstStepsWithGo/007-sub-folder" //Local package
	"github.com/kelvinCB/puppy"                               // Remote package
)

func main() {
	fmt.Println("hello,", Name)          //Local variable into same local package
	fmt.Println("Hello,", art.Article)   //Local varible into another local package
	fmt.Println("Hello,", puppy.Bark())  //Remote metodo from another repo on Github
	fmt.Println("Hello,", puppy.Barks()) //Remote metodo from another repo on Github
}

//Notes:

/*
To import Remote package, import package that you want
and then go to your terminal,
go to your project and write command:
go get -a

With that command your go.mod file will include remote package
and you can use its methods and variables

info learned in:
Local: https://www.youtube.com/watch?v=Nv8J_Ruc280&t=132s
Remote: https://www.youtube.com/watch?v=rHjlaHYd1F4

Udemy Section 8 until Modular code, #1
link: https://www.udemy.com/course/learn-how-to-code/learn/lecture/37652382#announcements
*/
