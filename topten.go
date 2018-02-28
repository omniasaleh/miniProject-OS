// gets top 10 most followed on twitter dataset.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"bufio"
	"sort"
)

// Implementing sort interface.
type SortUser struct {
	Followers map[int]int
	Keys      []int
}

// Gets length of map.
func (su *SortUser) Len() int {
	// TODO: Implement Len function.
	return len((*su).Keys)
}

// Condition for sorting to compare between values of keys.
func (su * SortUser) Less(i, j int) bool {

	return (*su).Followers[(*su).Keys[i]] > (*su).Followers[(*su).Keys[j]]
	// TODO: Implement Less function.
}

// Swaps two keys in keys array.
func (su * SortUser) Swap(i, j int) {
	(*su).Keys[i],(*su).Keys[j]=(*su).Keys[j],(*su).Keys[i]
	// TODO: Implement Swap function.
}

// Sorts Keys based on number of followers in descending order.
func sortKeys(m map[int]int) []int {
	// TODO: Implement sortKeys function.
	keys:= make([]int,0,len(m))
	for k := range m {
		keys = append(keys,k)
	}
  su:=SortUser{m,keys}
  sort.Sort(&su)
	return su.Keys
}

// Calculates top 10 most followed for input file
// and returns array of user id (int) for top 10.
func topTen(dataInput string) []int {
	// TODO: Implement topTen function.
	fd, err := os.Open(dataInput)
	 if err != nil {
			 panic(fmt.Sprintf("open %s: %v", dataInput, err))
	 }
	var  followed,user int
	//scanner:= bufio.NewScanner(fd)
	var m  map[int]int
	m=make(map[int]int)
	now:=0
	scanner := bufio.NewScanner(fd)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        x,err := strconv.Atoi(scanner.Text())
        if err != nil {
            fmt.Println(err)
        }
				if now%2==0 {
					user=x
					var tmp=m[user]
					m[user]=tmp
				} else{
					followed=x
				m[followed]++
			}
				now++
    }
		fd.Close()
		 var ids [] int=sortKeys(m)
		 size:=10
		 if len(ids)<10 {
			 size=len(ids)
		 }
		 return ids[0:size]
}

// Connects to remote service through internet to convert user id
// to username.
func getUsername(userId string) string {
	response, err := http.PostForm("https://tweeterid.com/ajax.php",
		url.Values{"input": {userId}})
	if err != nil {
		fmt.Println("Error getting username("+userId+"): ", err)
		return ""
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in response("+userId+"): ", err)
		return ""
	}
	if string(body) == "error" {
		return ""
	}
	return string(body)
}

func main() {
	fmt.Println("Calculating top 10 most followed...")
	topId:=topTen("/share/dataset.txt")
 	fmt.Println("topTen length: ", len(topId))

	fmt.Println("Getting and printing screen name for top 10...")
	for i := 0; i < 10 && i < len(topId); i++ {
		fmt.Printf("%-15d%s\n", topId[i], getUsername(strconv.Itoa(topId[i])))
	}
}
