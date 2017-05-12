package configuration

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strconv"

	"github.com/AlexsJones/gok8s/util"
	"github.com/olekukonko/tablewriter"
)

//MapConfiguration ...
type MapConfiguration struct {
	maps     []*item
	tableMap []string
}

//NewMapConfiguration ...
func NewMapConfiguration() *MapConfiguration {
	m := MapConfiguration{}
	m.tableMap = []string{"Step", "Resource locator", "Validated", "Executed", "Successful"}
	return &m
}

//Clear ...
func (m *MapConfiguration) Clear() {
	m.maps = nil
}

//Push ...
func (m *MapConfiguration) Push(uri string) {
	i := item{uri: uri}
	retb, _ := util.Exists(uri)
	_, err := url.ParseRequestURI(uri)
	if (err == nil) || (retb == true) {
		i.Validated(false)
	} else {
		i.Validated(true)
	}
	i.Executed(false)
	i.success = "?"
	m.maps = append(m.maps, &i)
}

//List ...
func (m *MapConfiguration) List() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(m.tableMap)

	var data [][]string

	var inc = 1
	for _, current := range m.maps {
		if current.uri != "" {
			data = append(data, []string{strconv.Itoa(inc), current.uri,
				fmt.Sprint(current.validated), fmt.Sprintf(current.executed), current.success})
			inc++
		}
	}
	for _, current := range data {
		table.Append(current)
	}

	if len(data) >= 1 {
		table.Render()
	} else {
		fmt.Println("Nothing scheduled...")
	}
}

func (m *MapConfiguration) run(i *item) {
	c := exec.Command(i.uri) //temppppp
	c.Stdout = nil
	c.Stderr = os.Stdout
	err := c.Run()
	if err != nil {
		i.Success(false)
		return
	}
	i.Success(true)
}

//Retry ...
func (m *MapConfiguration) Retry(i int) {

	if len(m.maps) < i || i < 0 {
		fmt.Println("Index out of bounds")
		return
	}

	m.run(m.maps[i])
}

//Run ...
func (m *MapConfiguration) Run() {

	var data [][]string
	var inc = 1

	for _, current := range m.maps {
		if current.uri != "" {

			m.run(current)
			current.Executed(true)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			data = append(data, []string{strconv.Itoa(inc), current.uri,
				fmt.Sprint(current.validated), fmt.Sprintf(current.executed), current.success})
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader(m.tableMap)
			for _, v := range data {
				table.Append(v)
			}
			table.Render()
		}
	}
}
