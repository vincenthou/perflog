package main

import (
  "github.com/ant0ine/go-json-rest/rest"
  flag "github.com/spf13/pflag"
  log "github.com/cihub/seelog"
  "net/http"
  "strings"
)

var (
  port *string = flag.String("port", "8080", "Binding port for the service")
)

func init() {
  logger, err := log.LoggerFromConfigAsFile("seelog.xml")
  if err == nil {
    log.ReplaceLogger(logger)
  }
}

func formatMessage(data *map[string]string) {
  messgae := ""
  for key,value := range *data {
    messgae += "\"" + key + ":" + value + "\" "
  }
  log.Info(messgae)
}

func main() {
  flag.Parse()

  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)

  router, err := rest.MakeRouter(
      &rest.Route{"GET", "/ping", func(w rest.ResponseWriter, req *rest.Request) {
          req.ParseForm()
          queryString := ""
          for key, value := range req.Form {
            queryString += key + "=" + value[0] + ","
          }
          ip := strings.Split(req.RemoteAddr, ":")[0]
          formatMessage( &map[string]string{"ip": ip, "params": queryString})
          w.WriteJson(map[string]string{"status": "ok"})
      }},
  )
  if err != nil {
      log.Error(err)
  }

  api.SetApp(router)
  http.ListenAndServe(":" + *port, api.MakeHandler())
}
