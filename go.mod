module github.com/coroot/coroot

go 1.23

require (
	github.com/ClickHouse/ch-go v0.62.0
	github.com/ClickHouse/clickhouse-go/v2 v2.8.3
	github.com/DataDog/golz4 v1.3.0
	github.com/PagerDuty/go-pagerduty v1.6.0
	github.com/atc0005/go-teams-notify/v2 v2.7.0
	github.com/buger/jsonparser v1.1.1
	github.com/coroot/logparser v1.1.4
	github.com/dustin/go-humanize v1.0.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/snappy v0.0.4
	github.com/google/pprof v0.0.0-20241210010833-40e02aabc2ad
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/grafana/pyroscope-go/godeltaprof v0.1.8
	github.com/hako/durafmt v0.0.0-20210608085754-5c1018a4e16b
	github.com/jpillora/backoff v1.0.0
	github.com/klauspost/compress v1.17.11
	github.com/lib/pq v1.10.7
	github.com/matoous/go-nanoid v1.5.0
	github.com/mattn/go-sqlite3 v1.14.15
	github.com/opsgenie/opsgenie-go-sdk-v2 v1.2.14
	github.com/prometheus/client_golang v1.20.5
	github.com/prometheus/common v0.61.0
	github.com/prometheus/prometheus v0.300.0-beta.1
	github.com/sirupsen/logrus v1.9.3
	github.com/slack-go/slack v0.11.3
	github.com/stretchr/testify v1.10.0
	github.com/uptrace/opentelemetry-go-extra/otelsql v0.3.2
	github.com/xhit/go-str2duration/v2 v2.1.0
	go.elastic.co/apm/module/apmotel/v2 v2.6.3
	go.opentelemetry.io/collector/semconv v0.116.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.55.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.55.0
	go.opentelemetry.io/otel v1.30.0
	go.opentelemetry.io/proto/otlp v1.4.0
	golang.org/x/crypto v0.31.0
	golang.org/x/exp v0.0.0-20240119083558-1b970713d09a
	golang.org/x/net v0.32.0
	golang.org/x/term v0.27.0
	gonum.org/v1/gonum v0.12.0
	google.golang.org/protobuf v1.36.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v3 v3.0.1
	inet.af/netaddr v0.0.0-20230525184311-b8eac61e914a
	k8s.io/klog v1.0.0
)

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20240927000941-0f3dac36c52b // indirect
	github.com/andybalholm/brotli v1.0.5 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dennwc/varint v1.0.0 // indirect
	github.com/dmarkham/enumer v1.5.9 // indirect
	github.com/elastic/go-sysinfo v1.7.1 // indirect
	github.com/elastic/go-windows v1.0.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.7.1 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grafana/regexp v0.0.0-20240518133315-a468a5bfb3bc // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.24.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/joeshaw/multierror v0.0.0-20140124173710-69b34d4ec901 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pascaldekloe/name v1.0.1 // indirect
	github.com/paulmach/orb v0.9.0 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	go.elastic.co/apm/module/apmhttp/v2 v2.6.3 // indirect
	go.elastic.co/apm/v2 v2.6.3 // indirect
	go.elastic.co/fastjson v1.1.0 // indirect
	go.opentelemetry.io/otel/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/sdk v1.30.0 // indirect
	go.opentelemetry.io/otel/sdk/metric v1.30.0 // indirect
	go.opentelemetry.io/otel/trace v1.30.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	go4.org/intern v0.0.0-20211027215823-ae77deb06f29 // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20230525183740-e7c30c78aeb2 // indirect
	golang.org/x/mod v0.22.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.28.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/tools v0.28.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241216192217-9240e9c98484 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241209162323-e6fa225c2576 // indirect
	google.golang.org/grpc v1.69.0-dev // indirect
	howett.net/plist v1.0.0 // indirect
)
