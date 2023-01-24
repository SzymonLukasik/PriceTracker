package main

import (
	"context"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"time"

	chartrender "github.com/go-echarts/go-echarts/v2/render"

	pb "pricetracker/pkg/build/pkg/proto"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// adapted from
// https://github.com/go-echarts/go-echarts/blob/master/templates/base.go
// https://github.com/go-echarts/go-echarts/blob/master/templates/header.go
var baseTpl = `
<div class="container">
    <div class="item" id="{{ .ChartID }}" style="width:{{ .Initialization.Width }};height:{{ .Initialization.Height }};"></div>
</div>
{{- range .JSAssets.Values }}
   <script src="{{ . }}"></script>
{{- end }}
<script type="text/javascript">
    "use strict";
    let goecharts_{{ .ChartID | safeJS }} = echarts.init(document.getElementById('{{ .ChartID | safeJS }}'), "{{ .Theme }}");
    let option_{{ .ChartID | safeJS }} = {{ .JSON }};
    goecharts_{{ .ChartID | safeJS }}.setOption(option_{{ .ChartID | safeJS }});
    {{- range .JSFunctions.Fns }}
    {{ . | safeJS }}
    {{- end }}
</script>
`

type snippetRenderer struct {
	c      interface{}
	before []func()
}

func newSnippetRenderer(c interface{}, before ...func()) chartrender.Renderer {
	return &snippetRenderer{c: c, before: before}
}

func (r *snippetRenderer) Render(w io.Writer) error {
	const tplName = "chart"
	for _, fn := range r.before {
		fn()
	}

	tpl := template.
		Must(template.New(tplName).
			Funcs(template.FuncMap{
				"safeJS": func(s interface{}) template.JS {
					return template.JS(fmt.Sprint(s))
				},
			}).
			Parse(baseTpl),
		)

	err := tpl.ExecuteTemplate(w, tplName, r.c)
	return err
}

func httpserver(w http.ResponseWriter, r *http.Request) {
	sdconn, err := grpc.Dial(products, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.WithError(err).Fatal("unable to connect to service dispatcher")
		// FIXME add retry policy
	}
	defer sdconn.Close()

	products := pb.NewProductsClient(sdconn)

	q := r.URL.Query()
	product := pb.Product{
		Shop: q.Get("shop"),
		Name: q.Get("name"),
		Url:  q.Get("url"),
	}

	// product := pb.Product{
	// 	Shop: "Euro",
	// 	Name: "laptop1",
	// 	Url:  "https://www.euro.com.pl/laptopy-i-netbooki/asus-laptop-fx506-i5-16gb-512ssd-3060-w11.bhtml",
	// }

	list, err := products.GetProductPrices(context.Background(), &product)

	if err == nil {
		// create a new line instance
		line := charts.NewLine()
		line.Renderer = newSnippetRenderer(line, line.Validate)

		// set some global options like Title/Legend/ToolTip or anything else
		line.SetGlobalOptions(
			charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
			charts.WithTitleOpts(opts.Title{
				Title:    fmt.Sprintf("%s - %s", product.Name, product.Shop),
				Subtitle: fmt.Sprintf("Price is given in 1/100 of PLN."),
			}),
			// charts.WithXAxisOpts(opts.XAxis{
			// 	Type: "time",
			// }),
			charts.WithTooltipOpts(opts.Tooltip{
				Trigger:     "item",
				AxisPointer: &opts.AxisPointer{Type: "line"},
			}),
		)

		// prepare data to display
		log.WithField("product", product).WithField("prices", len(list.GetPrices())).Info("fetched prices")
		dates := make([]string, len(list.GetPrices()))
		prices := make([]opts.LineData, len(list.GetPrices()))
		for i, datePrice := range list.GetPrices() {
			dates[i] = datePrice.Ts.AsTime().Format(time.RFC822)
			prices[i].Value = datePrice.Price
		}

		// Put data into instance
		line.SetXAxis(dates).
			AddSeries("Prices", prices)

		// Set cache
		cacheControl := `max-age=20, proxy-revalidate, stale-while-revalidate, stale-if-error, public, immutable`
		w.Header().Set("Cache-Control", cacheControl)

		// render charts
		line.Render(w)
	} else {
		log.WithError(err).WithField("product", product).Info("could not fetched prices for product.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/product", httpserver)
	http.ListenAndServe("localhost:8085", nil)
}

const products = "localhost:8083"
