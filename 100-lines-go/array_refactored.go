// NOTE: Use `range` in for loop
package main

import "fmt"

func main() {
	/* Define an array and let the compiler count its size
	** previous code **
	DeploymentOptions := [4]string{"R-pi", "AWS", "GCP", "Azure"} */
	DeploymentOptions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
	for index, option := range DeploymentOptions {
		fmt.Println(index, option)
	}

}
