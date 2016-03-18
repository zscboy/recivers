package http

import (
	"github.com/open-falcon/recivers/g"
	"github.com/toolkits/file"
	"net/http"
	"strings"
)

type Event struct {
	Endpoint string `json:"endpoint"`
	Metric   string `json:"metric"`
	Status   string `json:"status"`
	Step     string `json:"step"`
	Priority string `json:"priority"`
	Time     string `json:"time"`
	TplId    string `json:"tpl_id"`
	ExpId    string `json:"exp_id"`
	StraId   string `json:"stra_id"`
	tags     string `json:"tags"`
}

func configReciversRoutes() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(g.VERSION))
	})

	http.HandleFunc("/recivers", func(w http.ResponseWriter, r *http.Request) {
		event := Event{}

		event.Endpoint = r.URL.Query().Get("endpoint")
		event.Metric = r.URL.Query().Get("metric")
		event.Status = r.URL.Query().Get("status")
		event.Step = r.URL.Query().Get("step")
		event.Priority = r.URL.Query().Get("priority")
		event.Time = r.URL.Query().Get("time")
		event.TplId = r.URL.Query().Get("tpl_id")
		event.ExpId = r.URL.Query().Get("exp_id")
		event.StraId = r.URL.Query().Get("stra_id")
		event.tags = r.URL.Query().Get("tags")

		send2Agent(evnet)
	})
}

func send2Agent(event *Event) {
	url := "http://" + event.Endpoint + ":8988"

	buf, err := json.Marshal(j)
	if err != nil {
		log.Println("encode json err: ", err)
		return
	}

	client := &http.Client{Timeout: time.Second * 60}
	req, err := http.NewRequest("GET", url+"/restore", bytes.NewBuffer(buf))

	token := genTK("@##weeee%^*@")
	req.Header.Add("tk", token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("err: ", err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println("resp.StatusCode != 200")
		return
	}

	body := make([]byte, 1024)
	read, err := resp.Body.Read(body)
	if err != nil && read < 1 {
		log.Println("read message body err: ", err)
		return
	}

}
