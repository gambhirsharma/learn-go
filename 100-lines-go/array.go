package main

import "fmt"

func main()  {
	var DeploymentOptions = [4]string{"R-pi", "AWS", "GCP", "Azure"}

	for i:=0; i < len(DeploymentOptions); i++ {
		options := DeploymentOptions[i]
		fmt.Println(i,options)
	}
}
