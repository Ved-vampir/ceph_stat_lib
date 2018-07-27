package ceph_collect

import (
	"fmt"
	"ceph_data"
	"perf_counters"
)

//temp
func readFilters(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

func collect_perfcounters() {
    counters, err = perfDump()

    filters, err := readLines("filters.txt")
    fmt.Println("Len: ", len(filters))

    fmap, err := filterMap(str, filters)
    if err != nil {
        fmt.Print(err)
    }
    fmt.Println("Len: ", len(fmap))

    for key, value := range fmap {
    	fmt.Println("Key:", key)
    	for k, v := range value {
    		fmt.Println("k:", k, "Value:", v)
		}
}
}