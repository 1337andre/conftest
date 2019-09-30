module github.com/instrumenta/conftest

go 1.12

require (
	cuelang.org/go v0.0.4
	github.com/BurntSushi/toml v0.3.1
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/agext/levenshtein v1.2.2 // indirect
	github.com/bugsnag/bugsnag-go v1.5.1 // indirect
	github.com/containerd/containerd v1.3.0-0.20190426060238-3a3f0aac8819
	github.com/deislabs/oras v0.5.1-0.20190510174428-2836a4314d4a
	github.com/docker/cli v0.0.0-20190511004558-53fc257292ad // indirect
	github.com/docker/docker-credential-helpers v0.6.2 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/ghodss/yaml v1.0.0
	github.com/go-ini/ini v1.44.0
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190430165422-3e4dfb77656c // indirect
	github.com/gorilla/mux v1.7.1 // indirect
	github.com/hashicorp/hcl v1.0.0
	github.com/hashicorp/hcl2 v0.0.0-20190618163856-0b64543c968c
	github.com/jtolds/gls v4.2.1+incompatible // indirect
	github.com/kardianos/osext v0.0.0-20190222173326-2bc1f35cddc0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.2 // indirect
	github.com/logrusorgru/aurora v0.0.0-20190417130405-e50442bb4cb5
	github.com/magiconair/properties v1.8.1 // indirect
	github.com/mitchellh/go-wordwrap v1.0.0 // indirect
	github.com/moby/buildkit v0.5.1
	github.com/open-policy-agent/opa v0.12.0
	github.com/opencontainers/image-spec v1.0.1
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2 // indirect
	github.com/prometheus/common v0.4.1 // indirect
	github.com/prometheus/procfs v0.0.0-20190523193104-a7aeb8df3389 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20181016184325-3113b8401b8a // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/smartystreets/assertions v0.0.0-20180927180507-b2de0cb4f26d // indirect
	github.com/smartystreets/goconvey v0.0.0-20180222194500-ef6db91d284a // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.4
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.3.0
	github.com/xenolf/lego v2.5.0+incompatible // indirect
	github.com/yashtewari/glob-intersection v0.0.0-20180916065949-5c77d914dd0b // indirect
	github.com/yvasiyarov/newrelic_platform_go v0.0.0-20160601141957-9c099fbc30e9 // indirect
	github.com/zclconf/go-cty v1.0.0
	golang.org/x/crypto v0.0.0-20190513172903-22d7a77e9e5f // indirect
	golang.org/x/net v0.0.0-20190628185345-da137c7871d7 // indirect
	golang.org/x/sys v0.0.0-20190626221950-04f50cda93cb // indirect
	google.golang.org/appengine v1.4.0 // indirect
	google.golang.org/genproto v0.0.0-20190620144150-6af8c5fc6601 // indirect
	gopkg.in/ini.v1 v1.42.0 // indirect
)

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999
	github.com/containerd/containerd => github.com/containerd/containerd v1.2.6
	github.com/docker/docker => github.com/docker/docker v0.0.0-20190131205458-8a43b7bb99cd
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422
)
