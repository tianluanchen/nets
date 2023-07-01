package tools

import (
	"nets/pkg"
	"sync"
)

type ScanGroup struct {
	IP   string
	Port int
}
type ScanResult struct {
	ScanGroup
	Open bool
}
type Scanner struct {
	Groups []ScanGroup
}

func (s *Scanner) Scan() {
	for _, group := range s.Groups {
		pkg.StandardScanSocket(group.IP, group.Port)
	}
}

func (s *Scanner) Add(ip string, port int) {
	s.AddGroup(ScanGroup{
		IP:   ip,
		Port: port,
	})
}
func (s *Scanner) AddGroup(group ScanGroup) {
	s.Groups = append(s.Groups, group)
}

func (s *Scanner) Len() int {
	return len(s.Groups)
}

func (s *Scanner) Run() []ScanResult {
	wg := sync.WaitGroup{}
	var results []ScanResult
	for _, group := range s.Groups {
		wg.Add(1)
		go func(group ScanGroup) {
			defer wg.Done()
			results = append(results, ScanResult{
				ScanGroup: group,
				Open:      pkg.StandardScanSocket(group.IP, group.Port),
			})
		}(group)
	}
	wg.Wait()
	return results
}
