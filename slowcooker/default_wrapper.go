package slowcooker

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/hyperpilotio/slow_cooker_wrapper/utils"
)

var wg sync.WaitGroup

// Parameter : parameter need to post to Run method
type Parameter struct {
	Title               string
	Target              string
	Qps                 int
	Concurrency         int
	Method              string
	Interval            string
	Noreuses            bool
	Compress            bool
	NoLatencySummary    bool
	ReportLatenciesCSV  string
	TotalRequests       int
	Header              map[string]string
	Host                string
	Data                string
	RunOrder            int
	HashSampleRate      float64
	HashValue           uint64
	MetricAddr          string
	MetricServerBackend string
	InfluxUserName      string
	InfluxPassword      string
	InfluxDatabase      string
}

// CommandArray : convert parameter into slow cooker's command string
func (p Parameter) CommandArray() []string {
	result := make([]string, 0)
	if p.Qps > 0 {
		result = append(result, "-qps")
		result = append(result, strconv.Itoa(p.Qps))
	}
	if p.Concurrency > 0 {
		result = append(result, "-concurrency")
		result = append(result, strconv.Itoa(p.Concurrency))
	}
	if p.Compress {
		result = append(result, "-compress")
		result = append(result, "true")
	}
	if p.Data != "" {
		result = append(result, "-data")
		result = append(result, p.Data)
	}
	if len(p.Header) > 0 {
		for k, v := range p.Header {
			result = append(result, "-header")
			result = append(result, k+": "+v)
		}
	}
	if p.Host != "" {
		result = append(result, "-host")
		result = append(result, p.Host)
	}
	if p.Interval != "" {
		result = append(result, "-interval")
		result = append(result, p.Interval)
	}
	if p.Method != "" {
		result = append(result, "-method")
		result = append(result, strings.ToUpper(p.Method))
	}
	if p.MetricAddr != "" {
		result = append(result, "-metric-addr")
		result = append(result, p.MetricAddr)
	}
	if p.NoLatencySummary {
		result = append(result, "-noLatencySummary")
	}
	if p.Noreuses {
		result = append(result, "-noreuse")
		result = append(result, "true")
	}
	if p.ReportLatenciesCSV != "" {
		result = append(result, "-reportLatenciesCSV")
		result = append(result, p.ReportLatenciesCSV)
	}
	if p.TotalRequests > 0 {
		result = append(result, "-totalRequests")
		result = append(result, strconv.Itoa(p.TotalRequests))
	}
	if p.HashSampleRate > float64(0.0) {
		result = append(result, "-hashSampleRate")
		result = append(result, strconv.FormatFloat(p.HashSampleRate, 'f', 1, 1))
	}
	if p.HashValue > uint64(0) {
		result = append(result, "-hashValue")
		result = append(result, string(p.HashValue))
	}
	if p.MetricServerBackend != "" {
		result = append(result, "-metric-server-backend")
		result = append(result, p.MetricServerBackend)
	}
	if p.InfluxUserName != "" {
		result = append(result, "-influx-username")
		result = append(result, p.InfluxUserName)
	}
	if p.InfluxPassword != "" {
		result = append(result, "-influx-password")
		result = append(result, p.InfluxPassword)
	}
	if p.InfluxDatabase != "" {
		result = append(result, "-influx-database")
		result = append(result, p.InfluxDatabase)
	}

	result = append(result, p.Target)

	return result
}

// Parameters : slice of Parameter
type Parameters []Parameter

// Runs : implement Wrapper
func (v Parameters) Runs() {

	// cast to interface array
	var l = make([]interface{}, 0)
	for _, item := range v {
		l = append(l, item)
	}
	// get distinct run order

	runOrder := utils.Unify(utils.Map(l, func(i interface{}) interface{} {
		var rs interface{}
		rs = i.(Parameter).RunOrder
		return rs
	}))

	// sort run order
	utils.Sort(runOrder, func(arg1 interface{}, arg2 interface{}) bool {
		return arg1.(int) < arg2.(int)
	})

	// run jobs by run order
	for _, i := range runOrder {
		// time.Sleep(1 * time.Second)
		// fmt.Println("get output")
		jobs := utils.Filter(l, func(item interface{}) bool {
			return item.(Parameter).RunOrder == i
		})
		utils.Each(jobs, func(arg2 interface{}) interface{} {
			wg.Add(1)
			go arg2.(Parameter).run()
			return arg2
		})
		wg.Wait()
	}
}

func (p Parameter) run() {
	fmt.Printf("command: slow_cooker %v \n", p.CommandArray())
	// cmd := exec.Command("ping", "localhost")
	defer wg.Done()
	cmd := exec.Command("slow_cooker")
	for _, v := range p.CommandArray() {
		cmd.Args = append(cmd.Args, v)
	}

	// var out outStream
	// cmd.Stdout = out

	reader, err := cmd.StdoutPipe()
	// cmd.Stdout = os.Stdout
	if err != nil {
		log.Fatal(err)
	}

	errReader, erro := cmd.StderrPipe()
	if erro != nil {
		log.Fatal(erro)
	}
	scanner := bufio.NewScanner(io.MultiReader(reader, errReader))
	go func() {
		for scanner.Scan() {
			fmt.Println(time.Now(), " ", p.Title, " ", scanner.Text())
		}
	}()

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

}
