package main


import ("fmt")


func main(){
	fmt.Println("Hello, World!")
	firstRelease := map[string]int{
		"Go": 2009,
		"Rust": 2010,
		"Python": 1991,
		"C++": 1985,
		"C#": 2000,
	}
	for k, v :=range(firstRelease){
		fmt.Printf("%s was first released in %d\n", k, v)
	}
}
