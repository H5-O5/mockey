module code.byted.org/dc/meego_ai

go 1.24

replace (
	code.byted.org/middleware/dynamicgo => code.byted.org/middleware/dynamicgo v0.3.0
	github.com/apache/thrift => github.com/apache/thrift v0.13.0
	github.com/mattn/go-sqlite3 => github.com/mattn/go-sqlite3 v1.14.16
	golang.org/x/sync => golang.org/x/sync v0.11.0
)

replace github.com/volcengine/ve-tos-golang-sdk/v2 => github.com/volcengine/ve-tos-golang-sdk/v2 v2.6.4

require (
	code.byted.org/bits/utils v1.7.5
	code.byted.org/bits/workitem-sdk-go v0.0.49
	code.byted.org/bytehouse/clickhouse-go v1.16.1-0.20230406040433-6f6123a86ea8
	code.byted.org/data/mario_collector v1.3.1
	code.byted.org/data/xllm-client-go v0.0.0-20260408040750-261a2294fec2
	code.byted.org/dc/lark-project-cli v1.0.1-0.20260513134731-1d0f5815cfaf
	code.byted.org/dc/meego-analytic-language v0.0.0-20260704171115-44240a36f56e
	code.byted.org/dc/meego_ai_sdk v0.1.3-0.20260805085901-6a4a80e05536
	code.byted.org/ee/utils/environ v0.0.0-20220916112247-31dbc7ce64f8
	code.byted.org/flow/eino-byted-ext/callbacks/fornax v0.2.4
	code.byted.org/flow/eino-byted-ext/components/model/bytedgemini v0.1.18
	code.byted.org/flow/eino-byted-ext/components/model/bytedgpt v0.1.14
	code.byted.org/flow/eino-byted-ext/components/prompt/prompthub v0.1.7
	code.byted.org/flow/flow-telemetry-common/go v0.0.0-20260113113300-0887ab5e838a
	code.byted.org/flow/mcp_client v1.4.6
	code.byted.org/flowdevops/fornax_sdk v1.2.61
	code.byted.org/gopkg/consul v1.2.10
	code.byted.org/gopkg/context v0.0.1
	code.byted.org/gopkg/ctxvalues v0.7.0
	code.byted.org/gopkg/env v1.7.21
	code.byted.org/gopkg/facility v1.0.14
	code.byted.org/gopkg/gomonkey v0.1.4
	code.byted.org/gopkg/jsonx v0.3.3
	code.byted.org/gopkg/lang v0.21.8
	code.byted.org/gopkg/lang/v2 v2.1.4
	code.byted.org/gopkg/logid v0.0.0-20241008043456-230d03adb830
	code.byted.org/gopkg/logs v1.2.27
	code.byted.org/gopkg/metainfo v0.1.5
	code.byted.org/gopkg/metrics/v3 v3.2.0
	code.byted.org/gopkg/mysql-driver v1.2.7
	code.byted.org/gopkg/thrift v1.14.2
	code.byted.org/inf/bytedai-go v1.1.52
	code.byted.org/inf/sarama v1.5.3
	code.byted.org/kite/kitex v1.21.1
	code.byted.org/kite/kitutil v3.8.8+incompatible
	code.byted.org/kv/goredis v5.7.3+incompatible
	code.byted.org/kv/redis-v6 v1.1.5
	code.byted.org/lagrange/viking_go_client v0.0.60
	code.byted.org/lang/gg v0.22.1
	code.byted.org/lark/errorx v1.2.1
	code.byted.org/lark/errorx_gen/bytedance/bits/project/errgen v0.0.0-20250604025546-2037050622ac
	code.byted.org/lark/errorx_gen/bytedance/bits/search/errgen v0.0.0-20260326142947-61db7b652b31
	code.byted.org/lark/errorx_gen/bytedance/bits/workitem/errgen v0.0.0-20250320093320-e02620c39261
	code.byted.org/lark/errorx_gen/lark/meego/ai/errgen v0.0.0-20260807032223-26433964d3f5
	code.byted.org/lark_scs_pkg/guardian_pep_sdk v1.1.17
	code.byted.org/lark_search/enterprise_qa v0.0.0-20250701034623-71dda722e44e
	code.byted.org/lidar/agent v0.2.44
	code.byted.org/meegopkg/bytedtrace v0.2.2
	code.byted.org/meegopkg/config v0.1.53
	code.byted.org/meegopkg/ctxvalue v0.1.28
	code.byted.org/meegopkg/env v0.1.13
	code.byted.org/meegopkg/fg v0.1.29
	code.byted.org/meegopkg/file v1.0.20
	code.byted.org/meegopkg/kitexmw/larkerrorxmw v0.1.11
	code.byted.org/meegopkg/kitexmw/meegotracemw v0.0.11
	code.byted.org/meegopkg/kits/gflow v0.0.4-0.20251029074958-dc91c3c0cd0f
	code.byted.org/meegopkg/logger v0.2.1
	code.byted.org/meegopkg/mgf v0.1.196
	code.byted.org/meegopkg/permission v1.2.4
	code.byted.org/meegopkg/rocketmq-proxy v0.1.27-gama-4
	code.byted.org/meegopkg/rpccli v0.1.22
	code.byted.org/meegopkg/rpcmodel v0.0.0-20260819110505-def8f139f86f
	code.byted.org/meegopkg/starling v1.0.11
	code.byted.org/overpass/bear_doc_agg_api v0.0.0-20260528111624-b5e032dab958
	code.byted.org/overpass/bear_server_suite_integration v0.0.0-20260528113644-bb1b34084f2e
	code.byted.org/overpass/bytedance_bits_collect v0.0.0-20260707071805-ceda9dbb077b
	code.byted.org/overpass/bytedance_bits_facade v0.0.0-20260506043907-93e63ee7e99a
	code.byted.org/overpass/bytedance_bits_measure v0.0.0-20260512034703-3571ec190daf
	code.byted.org/overpass/bytedance_bits_meego_integration v0.0.0-20260625141312-5a51c77d411f
	code.byted.org/overpass/bytedance_bits_project v0.0.0-20260806090946-06c6f0d22288
	code.byted.org/overpass/bytedance_bits_search v0.0.0-20260813093316-364a260aaceb
	code.byted.org/overpass/bytedance_bits_search_infra v0.0.0-20260325064955-e1a336d4de5e
	code.byted.org/overpass/bytedance_bits_user v0.0.0-20260318063449-91390b6d676f
	code.byted.org/overpass/bytedance_bits_workitem v0.0.0-20260810033613-2e5035ce0c21
	code.byted.org/overpass/creation_doc_common v0.0.0-20260528105917-e19b42162159
	code.byted.org/overpass/lark_meego_ai v0.0.0-20260819110608-be33ef9240b0
	code.byted.org/overpass/lark_meego_permission v0.0.0-20260611073955-3b6df25d3147
	code.byted.org/security/volczti-helper v1.5.7
	code.byted.org/security/zti-jwt-helper-golang v1.0.18
	code.byted.org/toutiao/easygo v0.2.19
	code.byted.org/toutiao/elastic/v7 v7.0.47
	code.byted.org/ttarch/lego_client/host v1.17.8
	code.byted.org/ttarch/lego_client/wasm v0.0.12
	git.byted.org/ee/apacana v0.0.0-20260528102241-1ad52dfdd3a9
	git.byted.org/ee/apacana/kitc/clients/bear/server/explorer v0.0.0-20240814025450-279ee8bc628a
	git.byted.org/ee/apacana/kitc/thrift_gen/bear/server/explorer v0.0.0-20240809073037-cfc35384b88f
	git.byted.org/ee/go/kitex_gen v1.5.3-0.20260611120423-8f18df1d8f2f
	git.byted.org/ee/go/kitex_gen/bear/doc/agg_api v0.0.0-20260202092826-07a7fd775889
	git.byted.org/ee/go/kitex_gen/creation/doc/common v0.0.0-20260526081526-14d8b5d57605
	git.byted.org/ee/go/kitex_gen/creation/docx/engine v0.0.0-20260203065130-49e585e15b3a
	git.byted.org/ee/go/kitex_gen/lark/oapi/app_platform_runtime v0.0.0-20231009075842-8dd3c619a6f7
	git.byted.org/ee/go/kitex_gen/lark/scs/guardian v0.0.0-20260414061907-c76b7a1bfe8e
	git.byted.org/ee/go/kitex_gen/suite/passport/application v0.0.0-20260607025325-fa9c1783fbcd
	git.byted.org/ee/go/kitex_gen/todo/svc/core v0.0.0-20260804132206-6d4b08b60e9a
	git.byted.org/ee/gopkg v1.7.2
	git.byted.org/ee/lark/common v1.0.844
	github.com/ClickHouse/clickhouse-go/v2 v2.29.0
	github.com/agent-infra/sandbox-sdk-go v0.0.5
	github.com/agiledragon/gomonkey/v2 v2.13.0
	github.com/alicebob/miniredis/v2 v2.35.0
	github.com/avast/retry-go/v4 v4.6.0
	github.com/brianvoe/gofakeit/v6 v6.26.0
	github.com/bytedance/gopkg v0.1.3
	github.com/bytedance/mockey v1.3.0
	github.com/bytedance/sonic v1.15.0
	github.com/cloudwego/eino v0.8.4
	github.com/cloudwego/eino-ext/components/model/ark v0.1.64
	github.com/cloudwego/eino-ext/components/model/gemini v0.1.27
	github.com/cloudwego/eino-ext/components/model/openai v0.1.13
	github.com/cloudwego/fastpb v0.0.5
	github.com/cloudwego/kitex v0.16.0
	github.com/dgraph-io/ristretto v0.1.0
	github.com/eapache/queue v1.1.1-0.20180227141424-093482f3f8ce
	github.com/eino-contrib/jsonschema v1.0.3
	github.com/expr-lang/expr v1.17.6
	github.com/getkin/kin-openapi v0.118.0
	github.com/glycerine/goconvey v0.0.0-20190410193231-58a59202ab31
	github.com/go-git/go-billy/v5 v5.6.2
	github.com/go-git/go-git/v5 v5.13.2
	github.com/go-sql-driver/mysql v1.7.2-0.20231213112541-0004702b931d
	github.com/golang-module/carbon v1.7.3
	github.com/golang/mock v1.6.0
	github.com/google/uuid v1.6.0
	github.com/hashicorp/go-uuid v1.0.3
	github.com/invopop/jsonschema v0.13.0
	github.com/jedib0t/go-pretty/v6 v6.5.9
	github.com/jinzhu/copier v0.4.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.12
	github.com/kaptinlin/jsonrepair v0.2.1
	github.com/klauspost/compress v1.18.6
	github.com/larksuite/cli v1.0.72
	github.com/larksuite/oapi-sdk-go/v3 v3.9.5
	github.com/mattn/go-runewidth v0.0.19
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/mitchellh/mapstructure v1.5.0
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826
	github.com/nikolalohinski/gonja/v2 v2.3.1
	github.com/parquet-go/parquet-go v0.25.1
	github.com/philippgille/chromem-go v0.7.0
	github.com/pkg/errors v0.9.2-0.20201214064552-5dd12d0cfe7f
	github.com/pkoukk/tiktoken-go v0.1.7
	github.com/samber/lo v1.53.0
	github.com/santhosh-tekuri/jsonschema/v5 v5.3.1
	github.com/sashabaranov/go-openai v1.41.2
	github.com/smartystreets/goconvey v1.8.1
	github.com/spf13/cobra v1.10.2
	github.com/spf13/pflag v1.0.9
	github.com/stretchr/testify v1.11.1
	github.com/vmihailenco/msgpack v4.0.4+incompatible
	github.com/vmihailenco/msgpack/v5 v5.4.1
	github.com/volcengine/volc-sdk-golang v1.0.172
	github.com/volcengine/volcengine-go-sdk v1.2.12
	github.com/wk8/go-ordered-map/v2 v2.1.8
	github.com/yanyiwu/gojieba v1.4.6
	go.uber.org/dig v1.18.0
	golang.org/x/exp v0.0.0-20250128182459-e0ece0dbea4c
	golang.org/x/sync v0.15.0
	gonum.org/v1/gonum v0.15.1
	google.golang.org/genai v1.36.0
	google.golang.org/protobuf v1.36.8
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/driver/sqlite v1.5.4
	gorm.io/gorm v1.30.1
	gorm.io/hints v1.1.2
	gorm.io/plugin/dbresolver v1.6.2
)

require (
	code.byted.org/bytetim/larktim-go v1.0.15 // indirect
	code.byted.org/eventbus/pkg/tcc v0.0.0-20250902032048-f2c5438d44d6 // indirect
	code.byted.org/eventbus/pkg/timeoutv2 v0.0.0-20260721081657-d5db2261774d // indirect
	code.byted.org/eventbus/pkg/v2 v2.0.0-20260721081657-d5db2261774d // indirect
	code.byted.org/eventbus/proto v1.3.69 // indirect
	code.byted.org/gopkg/bytekv v1.1.31 // indirect
	code.byted.org/gopkg/commons v0.1.5 // indirect
	code.byted.org/gopkg/metrics/generic v1.0.6 // indirect
	code.byted.org/kv/goredis/v5 v5.6.6 // indirect
	code.byted.org/lark/ares/config/v2 v2.4.4 // indirect
	code.byted.org/lark/brandctx v1.0.4 // indirect
	code.byted.org/lark/cipher_sdk v1.8.8 // indirect
	code.byted.org/lark/enigma_common v1.0.0 // indirect
	code.byted.org/lark/gorm_cipher v1.8.3 // indirect
	code.byted.org/lark_admin/crossconf v0.0.0-20230807114326-84d55b79329a // indirect
	code.byted.org/lark_admin/gopkg/alarm v0.2.6 // indirect
	code.byted.org/lark_admin/gopkg/cacheproxy/common v1.1.1-0.20240312062432-bb1ea8919fe7 // indirect
	code.byted.org/lark_admin/gopkg/cacheproxy/v2 v2.8.4 // indirect
	code.byted.org/lark_admin/gopkg/errex v1.3.9-0.20230808064425-c25a573ca035 // indirect
	code.byted.org/lark_admin/gopkg/i18n v1.2.5 // indirect
	code.byted.org/lark_admin/gopkg/redis v1.3.1 // indirect
	code.byted.org/lark_admin/gopkg/retry v0.1.1 // indirect
	code.byted.org/lark_admin/gopkg/rpc_helper v0.1.1 // indirect
	code.byted.org/lark_admin/gopkg/skip v0.0.0-20220808050720-168325d7ec1e // indirect
	code.byted.org/lark_admin/gopkg/unit_config v1.0.12 // indirect
	code.byted.org/lark_admin/shield v1.1.12 // indirect
	code.byted.org/lark_admin/strategy_constant v0.0.0-20220224085027-116773311e60 // indirect
	code.byted.org/lark_scs/crypto_sdk/v2 v2.8.2 // indirect
	code.byted.org/lark_scs_pkg/gopkg/misc/metrics4 v1.0.2 // indirect
	code.byted.org/lark_scs_pkg/gopkg/misc/simplespooling v1.1.2 // indirect
	code.byted.org/lark_scs_pkg/gopkg/misc/syncx v1.0.0 // indirect
	code.byted.org/larkapigw/domain_sdk v1.2.14 // indirect
	code.byted.org/larkapigw/domain_sdk/core v1.2.14 // indirect
	code.byted.org/larkapigw/domain_sdk/pure v1.2.14 // indirect
	code.byted.org/larkarch/lark_dac v0.16.25 // indirect
	code.byted.org/larkarch/lark_dac_flow_control v0.0.0-20230713084029-ebd27c65612e // indirect
	code.byted.org/larkarch/meta_info/model v0.0.0-20260615115239-7ab46f5567ed // indirect
	code.byted.org/larkarch/metadatahttp v1.0.533 // indirect
	code.byted.org/larkim/common/tcc2 v0.0.0-20230406065614-1c233a6d490c // indirect
	code.byted.org/larkorg/kitex v0.2.62 // indirect
	code.byted.org/middleware/dynamicgo v0.0.1 // indirect
	code.byted.org/overpass/ttarch_spd_plugin_for_lark v0.0.0-20230629041102-13d81557a333 // indirect
	code.byted.org/security/ecm-sdk-golang v1.0.10 // indirect
	code.byted.org/security/fastOpeGo v1.0.2 // indirect
	code.byted.org/seed/common_idls/model_api v1.0.15 // indirect
	code.byted.org/ttarch/spd_action_for_lark v1.0.11-0.20230918092027-5c77657e7a57 // indirect
	code.byted.org/ttarch/spd_action_sdk v1.1.7 // indirect
	code.byted.org/ttarch/spd_plugin_for_lark/pkg/client v0.0.0-20240130065123-a821aa49f895 // indirect
	code.byted.org/ttarch/spd_plugin_for_lark/pkg/codec v0.0.0-20240130065123-a821aa49f895 // indirect
	git.byted.org/ee/apacana/kitc/thrift_gen/bear/server/permission v0.0.0-20240506044209-52d4a54fe1ca // indirect
	git.byted.org/ee/apacana/kitc/thrift_gen/bear/server/share_space v0.0.0-20221122025116-d8f4bc42b4fe // indirect
	git.byted.org/ee/apacana/kitc/thrift_gen/box/constants v0.0.0-20260525100903-750c85d0bb94 // indirect
	git.byted.org/ee/apacana/kitc/thrift_gen/box/external/route_call v0.0.0-20220120064111-60dd71ebcc7d // indirect
	git.byted.org/ee/ares/crossunit/v2 v2.1.37-0.20230202093946-0a7383ddb78e // indirect
	git.byted.org/ee/ares/genid/misc v1.0.2 // indirect
	git.byted.org/ee/go/clients v0.0.0-20260519082259-7abcf3ee7050 // indirect
	git.byted.org/ee/go/clients/mail/rule/dlp v0.0.0-20211214080835-e7ef09f05e7d // indirect
	git.byted.org/ee/go/clients/xunit v0.0.6 // indirect
	git.byted.org/ee/go/config v0.0.0-20190920095112-9d97055a3977 // indirect
	git.byted.org/ee/go/kitex_gen/bear/rce/syncroom v0.0.0-20260531073057-9edaa14025b1 // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/comment v0.0.0-20250326115355-977052150a5b // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/explorer v0.0.0-20250414122056-d629c391d20f // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/import_export_cli v0.0.0-20240617032826-fcc39c7617fe // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/import_export_facade v0.0.0-20240617032826-fcc39c7617fe // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/meta v0.0.0-20240411061148-d0c06496ea9a // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/obj_stats v0.0.0-20250210151450-bb459080b45d // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/permission v0.0.0-20250428073357-b5e17dba2738 // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/review v0.0.0-20250401101637-4c9fe267cca0 // indirect
	git.byted.org/ee/go/kitex_gen/bear/server/user v0.0.0-20250528082854-26682106eec2 // indirect
	git.byted.org/ee/go/kitex_gen/box/auth_common v0.0.0-20240411061148-d0c06496ea9a // indirect
	git.byted.org/ee/go/kitex_gen/box/constants v0.0.4 // indirect
	git.byted.org/ee/go/kitex_gen/box/external/file v0.0.0-20241225120509-34e0c6c53f20 // indirect
	git.byted.org/ee/go/kitex_gen/box/external/plan_cli v0.0.0-20240617032826-fcc39c7617fe // indirect
	git.byted.org/ee/go/kitex_gen/box/external/route_call v0.0.0-20240617032826-fcc39c7617fe // indirect
	git.byted.org/ee/go/kitex_gen/box/svc/resource_integration v0.0.0-20250221102444-d891ba78091c // indirect
	git.byted.org/ee/go/kitex_gen/commonentity v0.0.0-20250610070852-36fc702aaf2c // indirect
	git.byted.org/ee/go/kitex_gen/creation/docx/virtualblock v0.0.0-20260608084112-9182c263c98a // indirect
	git.byted.org/ee/go/kitex_gen/creation/docx/virtualblock_v2 v0.0.0-20260608084112-9182c263c98a // indirect
	git.byted.org/ee/go/kitex_gen/creation/explorer/display v0.0.0-20260530105329-a753364a29b7 // indirect
	git.byted.org/ee/go/kitex_gen/creation/explorer/space v0.0.0-20240617032826-fcc39c7617fe // indirect
	git.byted.org/ee/go/kitex_gen/creation/kam/migrate v0.0.0-20250428064326-b7ff74bd5701 // indirect
	git.byted.org/ee/go/kitex_gen/creation/platform/bypass_check v0.0.0-20240411061148-d0c06496ea9a // indirect
	git.byted.org/ee/go/kitex_gen/creation/spacex/shield v0.0.0-20240910140207-a960f1a8431a // indirect
	git.byted.org/ee/go/kitex_gen/creation/spacex/taskx v0.0.0-20240411061148-d0c06496ea9a // indirect
	git.byted.org/ee/go/kitex_gen/creation/wiki/core v0.0.0-20260601025609-781ebd7c84dc // indirect
	git.byted.org/ee/go/kitex_gen/creation/wiki/schedule v0.0.0-20240411061148-d0c06496ea9a // indirect
	git.byted.org/ee/go/kitex_gen/ee/infra/cron v0.0.0-20240617032826-fcc39c7617fe // indirect
	git.byted.org/ee/go/kitex_gen/lark/arch/domain v0.0.0-20231211023631-5b505c20cf9f // indirect
	git.byted.org/ee/go/kitex_gen/lark/arch/dts_api v0.0.0-20231121030415-dd1f654377b5 // indirect
	git.byted.org/ee/go/kitex_gen/lark/oapi/app_platform_admin v0.0.0-20260623140035-8792601eaf6e // indirect
	git.byted.org/ee/go/kitex_gen/lark/permission/ccm_common v0.0.0-20250428073357-b5e17dba2738 // indirect
	git.byted.org/ee/go/kitex_gen/lark/permission/ccm_mgr v0.0.0-20250428074216-47924f667d14 // indirect
	git.byted.org/ee/go/kitex_gen/lark/scs/common v0.0.0-20250424082309-49b143b72bac // indirect
	git.byted.org/ee/go/kitex_gen/lark/scs/encipherment v0.0.0-20251110024636-aea9cd2107f6 // indirect
	git.byted.org/ee/go/kitex_gen/lark/scs/octo v0.0.0-20251108105544-30811cf35e1c // indirect
	git.byted.org/ee/go/kitex_gen/lark/tns/task v0.0.0-20260330072621-3deedf4da05d // indirect
	git.byted.org/ee/go/kitex_gen/suite/feature_gating/svc v0.0.0-20250529081131-2752c7e56c66 // indirect
	git.byted.org/ee/go/kitex_gen/suite/security/engine v0.0.0-20251012135816-651f598ec9e6 // indirect
	git.byted.org/ee/go/kitex_gen/suite/security/unified_kms v0.0.0-20251110024636-aea9cd2107f6 // indirect
	git.byted.org/ee/go/kitex_gen/xunit v0.1.13-0.20220714075408-d9d97bd6de83 // indirect
	git.byted.org/ee/go/loader v0.0.0-20190925074519-f96ac9fcad39 // indirect
	git.byted.org/ee/go/thrift_gen v1.0.1 // indirect
	git.byted.org/ee/go/thrift_gen/mail/rule/dlp v0.0.0-20211214080834-86fd6eacbbb9 // indirect
	git.byted.org/ee/lark/kiteclient/mail v0.0.0-20210225065640-552300e133eb // indirect
	git.byted.org/ee/suite/security/sign_sdk v0.3.6 // indirect
	git.byted.org/lark/gopkg/fg v1.2.8 // indirect
	github.com/allegro/bigcache v1.2.1 // indirect
	github.com/allegro/bigcache/v3 v3.1.0 // indirect
	github.com/atotto/clipboard v0.1.4 // indirect
	github.com/avast/retry-go v3.0.0+incompatible // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/bmatcuk/doublestar/v4 v4.10.0 // indirect
	github.com/catppuccin/go v0.3.0 // indirect
	github.com/cbroglie/mustache v1.4.0 // indirect
	github.com/charmbracelet/bubbles v0.21.1-0.20250623103423-23b8fd6302d7 // indirect
	github.com/charmbracelet/bubbletea v1.3.6 // indirect
	github.com/charmbracelet/colorprofile v0.2.3-0.20250311203215-f60798e515dc // indirect
	github.com/charmbracelet/huh v1.0.0 // indirect
	github.com/charmbracelet/lipgloss v1.1.0 // indirect
	github.com/charmbracelet/x/ansi v0.9.3 // indirect
	github.com/charmbracelet/x/cellbuf v0.0.13 // indirect
	github.com/charmbracelet/x/exp/strings v0.0.0-20240722160745-212f7b056ed0 // indirect
	github.com/charmbracelet/x/term v0.2.1 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20230717121745-296ad89f973d // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/danieljoos/wincred v1.2.3 // indirect
	github.com/erikgeiser/coninput v0.0.0-20211004153227-1c3628e74d0f // indirect
	github.com/godbus/dbus/v5 v5.2.2 // indirect
	github.com/gofrs/flock v0.8.1 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/itchyny/gojq v0.12.17 // indirect
	github.com/itchyny/timefmt-go v0.1.6 // indirect
	github.com/julienschmidt/httprouter v1.3.0 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/luci/go-render v0.0.0-20160219211803-9a04cc21af0f // indirect
	github.com/mark3labs/mcp-go v0.43.0 // indirect
	github.com/mattn/go-localereader v0.0.1 // indirect
	github.com/mitchellh/hashstructure/v2 v2.0.2 // indirect
	github.com/muesli/ansi v0.0.0-20230316100256-276c6243b2f6 // indirect
	github.com/muesli/cancelreader v0.2.2 // indirect
	github.com/muesli/termenv v0.16.0 // indirect
	github.com/openai/openai-go/v3 v3.31.0 // indirect
	github.com/pingcap/log v0.0.0-20210906054005-afc726e70354 // indirect
	github.com/pingcap/tidb/parser v0.0.0-20221110094352-5ccc10beb8b4 // indirect
	github.com/prometheus-community/pro-bing v0.1.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/sanity-io/litter v1.5.5 // indirect
	github.com/shirou/gopsutil/v3 v3.24.5 // indirect
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	github.com/timandy/routine v1.1.4 // indirect
	github.com/xo/terminfo v0.0.0-20220910002029-abceb7e1c41e // indirect
	github.com/yosida95/uritemplate/v3 v3.0.2 // indirect
	github.com/zalando/go-keyring v0.2.8 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.27.0 // indirect
	golang.org/x/term v0.32.0 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
)

require (
	cloud.google.com/go v0.116.0 // indirect
	cloud.google.com/go/auth v0.9.3 // indirect
	cloud.google.com/go/compute/metadata v0.5.2 // indirect
	code.byted.org/ad/delta_generation v0.0.0-20240613062043-0a68745699de // indirect
	code.byted.org/aiops/apm_vendor_byted v0.0.43 // indirect
	code.byted.org/aiops/metrics_codec v0.0.33 // indirect
	code.byted.org/aiops/monitoring-common-go v0.0.7 // indirect
	code.byted.org/aweme-go/hstruct v0.1.1 // indirect
	code.byted.org/bcc/bcc-go-client v0.1.51 // indirect
	code.byted.org/bcc/conf_engine v0.0.0-20230510030051-32fb55f74cf1 // indirect
	code.byted.org/bcc/pull_json_model v1.0.22 // indirect
	code.byted.org/bcc/tools v0.0.21 // indirect
	code.byted.org/bits/library v0.6.20 // indirect
	code.byted.org/bits/workitem-cache v1.7.5 // indirect
	code.byted.org/bytedtrace-contrib/kitex-go v1.1.52 // indirect
	code.byted.org/bytedtrace/bytedtrace-client-go v1.3.2 // indirect
	code.byted.org/bytedtrace/bytedtrace-common/go v0.0.15 // indirect
	code.byted.org/bytedtrace/bytedtrace-compatible-client-go v0.0.16 // indirect
	code.byted.org/bytedtrace/bytedtrace-compatible-lightweight-go v1.0.2 // indirect
	code.byted.org/bytedtrace/bytedtrace-conf-provider-client-go v0.0.27 // indirect
	code.byted.org/bytedtrace/bytedtrace-gls-switch v1.3.0 // indirect
	code.byted.org/bytedtrace/interface-go v1.0.20 // indirect
	code.byted.org/bytedtrace/serializer-go v1.0.1-pre // indirect
	code.byted.org/bytees/olivere_elastic/v7 v7.0.36 // indirect
	code.byted.org/bytehouse/parser v1.1.1 // indirect
	code.byted.org/ccm-platform/log_guard v1.0.26 // indirect
	code.byted.org/data-arch/gotbase v1.0.8-0.20220905113555-b9d46a7dc975 // indirect
	code.byted.org/data/databus_client v1.3.10 // indirect
	code.byted.org/dp/mario_common v1.0.7 // indirect
	code.byted.org/ee/idgenerator_client v1.0.20 // indirect
	code.byted.org/ee/logrus_sentry v0.0.0-20190902082707-ef0e387902de // indirect
	code.byted.org/ep/common_lib v1.2.9-0.20230407073713-559c2928ddee // indirect
	code.byted.org/flow/eino-byted-ext/byted v0.3.13 // indirect
	code.byted.org/flow/eino-byted-ext/callbacks/metrics v0.1.2 // indirect
	code.byted.org/flow/eino-byted-ext/components/model/llmgateway v0.1.12 // indirect
	code.byted.org/flowdevops/errorx v0.0.8 // indirect
	code.byted.org/flowdevops/errorx/code/gen/flow/devops/agent_server v0.0.0-20241012084451-47d6baaffb45 // indirect
	code.byted.org/gin-gonic/gin v1.0.5 // indirect
	code.byted.org/gin/ginex v1.8.0 // indirect
	code.byted.org/golf/buffer_pool v0.1.0 // indirect
	code.byted.org/golf/consul v2.1.15+incompatible // indirect
	code.byted.org/golf/metrics v0.1.0 // indirect
	code.byted.org/gopkg/apm_vendor_interface v0.0.4-beta // indirect
	code.byted.org/gopkg/asyncache v0.0.0-20210129072708-1df5611dba17 // indirect
	code.byted.org/gopkg/asynccache v0.0.0-20210422090342-26f94f7676b8 // indirect
	code.byted.org/gopkg/bytedmysql v1.1.15 // indirect
	code.byted.org/gopkg/circuitbreaker v3.8.1+incompatible // indirect
	code.byted.org/gopkg/debug v0.10.1 // indirect
	code.byted.org/gopkg/etcd_util v2.3.3+incompatible // indirect
	code.byted.org/gopkg/etcdproxy v0.1.1 // indirect
	code.byted.org/gopkg/gopool v0.13.1 // indirect
	code.byted.org/gopkg/gorm v2.0.1+incompatible // indirect
	code.byted.org/gopkg/localcache v0.10.3 // indirect
	code.byted.org/gopkg/localcache/base v0.10.0 // indirect
	code.byted.org/gopkg/localcache/contributes/freecache v0.7.5 // indirect
	code.byted.org/gopkg/localcache/contributes/gcache v0.8.2 // indirect
	code.byted.org/gopkg/localcache/contributes/vfastcache v0.2.0 // indirect
	code.byted.org/gopkg/logs/v2 v2.2.3
	code.byted.org/gopkg/metrics v1.4.25 // indirect
	code.byted.org/gopkg/metrics/v4 v4.1.10 // indirect
	code.byted.org/gopkg/metrics_core v0.0.58 // indirect
	code.byted.org/gopkg/metricx v0.5.4 // indirect
	code.byted.org/gopkg/net2 v1.5.1 // indirect
	code.byted.org/gopkg/pkg v0.1.0 // indirect
	code.byted.org/gopkg/rand v0.0.0-20241104022029-57b15c5bb0fa // indirect
	code.byted.org/gopkg/retry v0.0.0-20230209024914-cf290f094aa7 // indirect
	code.byted.org/gopkg/singleflight v0.0.0-20200508020530-47758763028a // indirect
	code.byted.org/gopkg/stats v1.2.13 // indirect
	code.byted.org/gopkg/tccclient v1.6.8 // indirect
	code.byted.org/gopkg/tccclient/v3 v3.0.4 // indirect
	code.byted.org/gopkg/tos v1.6.4 // indirect
	code.byted.org/gorm/bytedgorm v0.9.11 // indirect
	code.byted.org/hystrix/hystrix-go v0.0.0-20190214095017-a2a890c81cd5 // indirect
	code.byted.org/ies/i18n-sdk-go v0.2.1 // indirect
	code.byted.org/ies/starling-i18n-go v0.2.3 // indirect
	code.byted.org/ies/starling_goclient v0.5.9 // indirect
	code.byted.org/ies/starling_sdk_api_http v0.0.9 // indirect
	code.byted.org/iesarch/simple_rate_limit/v2 v2.0.40 // indirect
	code.byted.org/iespkg/bytedkits-go/goext v0.4.2 // indirect
	code.byted.org/iespkg/retry-go v0.1.4 // indirect
	code.byted.org/inf/authcenter v1.5.2 // indirect
	code.byted.org/inf/infsecc v1.0.3 // indirect
	code.byted.org/kite/dpstoken v0.0.5 // indirect
	code.byted.org/kite/edgeproxy v0.1.10 // indirect
	code.byted.org/kite/endpoint v3.7.5+incompatible // indirect
	code.byted.org/kite/kitc v3.10.26+incompatible // indirect
	code.byted.org/kite/kite v3.9.39+incompatible // indirect
	code.byted.org/kite/kitex-overpass-suite v0.0.36 // indirect
	code.byted.org/kite/kitex/pkg/protocol/bthrift v0.0.0-20260508052613-95b8d0cb5cf1 // indirect
	code.byted.org/kite/rpal v0.2.6 // indirect
	code.byted.org/kitex/apache_monitor v0.1.1 // indirect
	code.byted.org/kitex/dpstoken v0.1.1 // indirect
	code.byted.org/kv/backoff v0.0.0-20191031070508-5d868504e646 // indirect
	code.byted.org/kv/circuitbreaker v0.0.0-20200212034351-d3f51a5b9165 // indirect
	code.byted.org/lang/tangofeatures v0.1.9 // indirect
	code.byted.org/lang/trace v0.0.3 // indirect
	code.byted.org/lark/ares/config v1.6.1 // indirect
	code.byted.org/lark/ares/crossunit/v2 v2.4.13-0.20240709061947-946281312131 // indirect
	code.byted.org/lark/ares/crypto v1.6.0 // indirect
	code.byted.org/lark/ares/genid/misc v1.6.0 // indirect
	code.byted.org/lark/ares/gentoken/misc v1.8.0 // indirect
	code.byted.org/lark/ares/internal/iconfig v1.6.0 // indirect
	code.byted.org/lark/ares/internal/ijson v1.6.1 // indirect
	code.byted.org/lark/ares/internal/iyaml v1.6.1 // indirect
	code.byted.org/lark/ares/internal/mqmarshal v1.6.0 // indirect
	code.byted.org/lark/ares/json v1.6.0 // indirect
	code.byted.org/lark/ares/rocketmq/v2 v2.4.7 // indirect
	code.byted.org/lark/ares/utils v1.6.1 // indirect
	code.byted.org/lark/ares/vars v1.6.5 // indirect
	code.byted.org/lark/chaincache v0.0.0-20230227062013-9b5ba872408c // indirect
	code.byted.org/lark/chaincache/metacache v0.0.3 // indirect
	code.byted.org/lark/epsdk v0.1.5 // indirect
	code.byted.org/lark/epsdk/edgeproxy/kite v0.0.3 // indirect
	code.byted.org/lark/epsdk/edgeproxy/kitex v0.0.6 // indirect
	code.byted.org/lark/errorx/contrib/kitex v0.3.0 // indirect
	code.byted.org/lark/errorx_gen/lark/meego/file/errgen v0.0.0-20250311132020-0ac9cb94b4aa // indirect
	code.byted.org/lark/globalmetacore v1.2.82 // indirect
	code.byted.org/lark/globalmetasdk v1.0.27 // indirect
	code.byted.org/lark/globalmetautil/bloom v0.0.2-0.20230828064544-989585b2c7f2 // indirect
	code.byted.org/lark/globalmetautil/genunitparser v0.0.59 // indirect
	code.byted.org/lark/globalsdk v1.0.99 // indirect
	code.byted.org/lark/globalsdk/entitymetactx v0.0.17 // indirect
	code.byted.org/lark/globalsdk/globalctx v0.0.6 // indirect
	code.byted.org/lark/globalsdk/routermeta v0.0.23 // indirect
	code.byted.org/lark/kitex_gen/xunit v0.0.12 // indirect
	code.byted.org/lark/logger v1.1.8 // indirect
	code.byted.org/lark_as/common v1.0.16 // indirect
	code.byted.org/lark_as/common/v3 v3.3.48 // indirect
	code.byted.org/lark_search/common v0.0.0-20250630114709-a9f13e49035b // indirect
	code.byted.org/larkarch/metaclient v0.7.5 // indirect
	code.byted.org/larkbackend/sre_common/env v0.0.0-20220118064305-8e996f6b4ee0 // indirect
	code.byted.org/larkbackend/sre_common/sql_monitor_sdk v0.2.6-0.20260601153423-21e6c671192d // indirect
	code.byted.org/larkbackend/sre_common/tbind v0.0.0-20220831051613-51dc9fb4db05 // indirect
	code.byted.org/larkim/audit_log v1.0.4-0.20240228072955-a36301b300db // indirect
	code.byted.org/larkim/uuidcodec v1.0.0 // indirect
	code.byted.org/larkvc/byteview-gopkg v1.3.6 // indirect
	code.byted.org/larkvc/byteview-gopkg/global/mm_token v0.0.0-20231121083734-1e4c206f5dc3 // indirect
	code.byted.org/lidar/profiler v0.4.7 // indirect
	code.byted.org/lidar/profiler/hertz v0.4.7 // indirect
	code.byted.org/lidar/profiler/kitex v0.4.7 // indirect
	code.byted.org/log_market/gosdk v0.0.0-20230524072203-e069d8367314 // indirect
	code.byted.org/log_market/loghelper v0.1.12 // indirect
	code.byted.org/log_market/tracelog v0.1.9 // indirect
	code.byted.org/log_market/ttlogagent_gosdk v0.0.8 // indirect
	code.byted.org/log_market/ttlogagent_gosdk/v4 v4.0.56 // indirect
	code.byted.org/meegopkg/cache v0.0.21 // indirect
	code.byted.org/meegopkg/clickhouse-go/v2 v2.29.1-0.20260402085305-4e04eab04178 // indirect
	code.byted.org/meegopkg/config/pkg/configinstance v0.0.6 // indirect
	code.byted.org/meegopkg/i18n v0.15.2-0.20251120075545-0d5ee99e4ce9 // indirect
	code.byted.org/meegopkg/kits/domain v0.0.6 // indirect
	code.byted.org/meegopkg/kits/goext v0.1.7 // indirect
	code.byted.org/meegopkg/kits/panics v0.0.2-0.20251124085239-0cdc0a2a18d2 // indirect
	code.byted.org/meegopkg/logger/plugin/databus v0.1.4 // indirect
	code.byted.org/meegopkg/logger/plugin/zen v0.1.4 // indirect
	code.byted.org/meegopkg/meegotrace v1.1.3 // indirect
	code.byted.org/meegopkg/security_protect_sdk v1.0.4 // indirect
	code.byted.org/middleware/fic_client v0.2.8 // indirect
	code.byted.org/middleware/gocaller v0.0.7 // indirect
	code.byted.org/middleware/hertz v1.14.2
	code.byted.org/obric/flow_telemetry_go v1.1.7 // indirect
	code.byted.org/overpass/bytedance_bits_app_center v0.0.0-20250409063759-25a166abd508 // indirect
	code.byted.org/overpass/bytedance_bits_auto v0.0.0-20260709113345-da84eaf68b89
	code.byted.org/overpass/bytedance_bits_bql v0.0.0-20250409064000-9cb2af8f6ef9 // indirect
	code.byted.org/overpass/bytedance_bits_migration v0.0.0-20250318080903-3ffa654c0d5d // indirect
	code.byted.org/overpass/bytedance_bits_relation v0.0.0-20250328071122-d8f300775146 // indirect
	code.byted.org/overpass/bytedance_bits_river v0.0.0-20241111160157-5f12d9ede21d // indirect
	code.byted.org/overpass/common v0.0.0-20241127033622-79f193603286
	code.byted.org/overpass/data_aml_llmflow_engine v0.0.0-20241107145550-f2da45272e96 // indirect
	code.byted.org/overpass/flow_im_model v0.0.0-20260305090515-fd2553c10c9d // indirect
	code.byted.org/overpass/flow_mcp_core v0.0.0-20260305090521-3516c88567bf // indirect
	code.byted.org/overpass/flow_mcp_proxy v0.0.0-20260312080321-3c75d4482e60 // indirect
	code.byted.org/overpass/lark_idl_common v0.0.0-20260528085459-a31ae97c5433 // indirect
	code.byted.org/overpass/lark_meego_builder v0.0.0-20251105082105-f9bfb07a573f // indirect
	code.byted.org/overpass/lark_meego_commercial_public v0.0.0-20260617073722-73abd8433290 // indirect
	code.byted.org/overpass/lark_meego_dispatcher v0.0.0-20241111120544-4ce5ad8772bc // indirect
	code.byted.org/overpass/lark_meego_growth v0.0.0-20250306083623-2bc658b52fc1 // indirect
	code.byted.org/overpass/stone_llm_gateway v0.0.0-20250702072857-8b8e0655afa8 // indirect
	code.byted.org/overpass/suite_feature_gating_svc v0.0.0-20241224081300-bd75b35d687b // indirect
	code.byted.org/rocketmq/rocketmq-go-proxy v1.6.0 // indirect
	code.byted.org/rocketmq/rocketmq-go-proxy-mqmesh-interceptor v1.0.23 // indirect
	code.byted.org/security/certinfo v1.0.2 // indirect
	code.byted.org/security/cryptoutils v1.1.3
	code.byted.org/security/dlp/v2 v2.0.22 // indirect
	code.byted.org/security/go-spiffe-v2 v1.0.9 // indirect
	code.byted.org/security/kms-v2-sdk-golang v1.2.98 // indirect
	code.byted.org/security/memfd v0.0.2 // indirect
	code.byted.org/security/ope v0.0.0-20210820082458-fd993bd0d39a // indirect
	code.byted.org/security/sensitive_finder_engine v0.3.18 // indirect
	code.byted.org/security/spiffe_spire v0.0.0-20201116193931-c566c1c41bdf // indirect
	code.byted.org/security/zen v0.2.8 // indirect
	code.byted.org/security/zen-ext/logv2 v0.1.1 // indirect
	code.byted.org/security/zen-ext/profile v0.1.5 // indirect
	code.byted.org/security/zen-extensions v0.0.14 // indirect
	code.byted.org/security/zero-trust-identity-helper v1.0.14 // indirect
	code.byted.org/seed/common_idls/llmserver v1.0.9 // indirect
	code.byted.org/service_mesh/mesh_transport v1.0.1 // indirect
	code.byted.org/service_mesh/shmipc v0.2.25 // indirect
	code.byted.org/shield/balancer v0.0.0-20210128081154-01fb65e45023 // indirect
	code.byted.org/shield/common v1.0.11 // indirect
	code.byted.org/starling/makeplural v0.0.1 // indirect
	code.byted.org/starling/messageformat v0.0.1 // indirect
	code.byted.org/tiktok/buildinfo v0.0.2 // indirect
	code.byted.org/trace/trace-client-go v1.3.7 // indirect
	code.byted.org/ttarch/byteconf-cel-go v0.0.3 // indirect
	code.byted.org/ttarch/cds_sdk v0.3.3 // indirect
	code.byted.org/ttarch/lego_client/generic v0.0.3 // indirect
	code.byted.org/ttarch/lego_client/global v0.0.1 // indirect
	code.byted.org/ttarch/lego_client/mock v1.0.4 // indirect
	code.byted.org/ttarch/spd_kitex_section v1.0.1 // indirect
	code.byted.org/videoarch/vfastcache v1.0.10 // indirect
	code.byted.org/webcast/libs_anycache v1.6.7 // indirect
	code.byted.org/webcast/libs_anycache/plugin/cache/base v0.1.1-0.20221212082232-7c36e6844ac9 // indirect
	code.byted.org/webcast/libs_anycache/plugin/cache/objectcache v0.0.1 // indirect
	code.byted.org/webcast/libs_anycache/plugin/codec/base v0.1.0 // indirect
	code.byted.org/webcast/libs_anycache/plugin/refresh v0.1.3 // indirect
	code.byted.org/webcast/libs_sync v0.1.2 // indirect
	dario.cat/mergo v1.0.0 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	git.byted.org/ee/ares/config v1.1.2 // indirect
	git.byted.org/ee/ares/crypto v1.0.2 // indirect
	git.byted.org/ee/ares/gentoken/misc v1.0.4 // indirect
	git.byted.org/ee/ares/internal/iconfig v1.0.2 // indirect
	git.byted.org/ee/ares/internal/ijson v1.0.3 // indirect
	git.byted.org/ee/ares/internal/iyaml v1.0.3 // indirect
	git.byted.org/ee/ares/internal/mqmarshal v1.0.2 // indirect
	git.byted.org/ee/ares/json v1.0.2 // indirect
	git.byted.org/ee/ares/rocketmq/v2 v2.3.11-0.20220808103418-2ceb61732c99 // indirect
	git.byted.org/ee/ares/utils v1.0.3 // indirect
	git.byted.org/ee/ares/vars v1.1.27 // indirect
	git.byted.org/ee/go/kitex_gen/ee/idgenerator/meta v0.0.0-20221123035746-e93dbac99faf // indirect
	git.byted.org/ee/go/kitex_gen/kitexclient v0.0.0-20250825120924-f6db9c21b6a2 // indirect
	git.byted.org/ee/go/kitex_gen/lark/ai/llm v0.0.0-20241216112104-f5dbe1f6cea1 // indirect
	git.byted.org/ee/go/kitex_gen/lark/ai/llm_callback v0.0.0-20241202065224-114f2475b5bc // indirect
	git.byted.org/ee/go/kitex_gen/lark/arch/dts v0.0.0-20251125032515-9e503ceff96e // indirect
	git.byted.org/ee/go/kitex_gen/lark/arch/meta_data v0.0.0-20251110190017-bef542b83a18 // indirect
	git.byted.org/ee/go/kitex_gen/lark/facade/chat v0.0.0-20241022141749-8e78f1fa6886 // indirect
	git.byted.org/ee/go/kitex_gen/lark/facade/chatter v0.0.0-20240905131015-76fdcb6171b5 // indirect
	git.byted.org/ee/go/kitex_gen/lark/global/dyconf v0.0.0-20240802080619-056692420157 // indirect
	git.byted.org/ee/go/kitex_gen/lark/global/egw v0.0.0-20230414090614-c0026497b49e // indirect
	git.byted.org/ee/go/kitex_gen/lark/global/meta v0.0.3 // indirect
	git.byted.org/ee/go/kitex_gen/lark/global/meta_db v0.0.0-20230523122230-4566703f1578 // indirect
	git.byted.org/ee/go/kitex_gen/lark/http/gateway v0.0.0-20230710091248-31aff160c317 // indirect
	git.byted.org/ee/go/kitex_gen/lark/im/message v0.0.0-20241128093943-f3410ac27d1c // indirect
	git.byted.org/ee/go/kitex_gen/lark/oapi/api_runtime v0.0.0-20221219113339-c11577fa0beb
	git.byted.org/ee/go/kitex_gen/lark/persconn/push v0.0.0-20250402082103-14cacf1ce7bf // indirect
	git.byted.org/ee/go/kitex_gen/lark/svc/file v0.0.0-20250514071657-885da2e6d8f7 // indirect
	git.byted.org/ee/lark/im-protobuf v1.1.925 // indirect
	git.byted.org/ee/lark/kiteclient v0.0.0-20250627085239-6dc3bb358186 // indirect
	git.byted.org/ee/suite/security/crypto_sdk v1.12.10 // indirect
	git.byted.org/ee/suite/security/sm v0.2.0 // indirect
	git.byted.org/ee/suite/security/utils v0.0.0-20220922072431-d8f787c7fd5c // indirect
	git.byted.org/lark/egwsdk v0.0.0-20221108113243-8126db7eb66f // indirect
	git.byted.org/lark/egwsdk/clients/suite/security/unified_kms v0.0.0-20220527102457-885778343f22 // indirect
	git.byted.org/lark/egwsdk/generic v1.0.2 // indirect
	git.byted.org/lark/gopkg/crossunit v1.0.42 // indirect
	git.byted.org/lark/gopkg/dimension v0.0.11 // indirect
	git.byted.org/lark/gopkg/kitext v1.0.5 // indirect
	git.byted.org/lark/gopkg/lodash v1.0.1 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/azcore v1.17.0 // indirect
	github.com/Azure/azure-sdk-for-go/sdk/internal v1.10.0 // indirect
	github.com/BurntSushi/toml v1.3.2 // indirect
	github.com/ClickHouse/ch-go v0.61.5 // indirect
	github.com/Knetic/govaluate v3.0.1-0.20171022003610-9aa49832a739+incompatible // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/OrlovEvgeny/go-mcache v0.0.0-20200121124330-1a8195b34f3a // indirect
	github.com/ProtonMail/go-crypto v1.1.6 // indirect
	github.com/RoaringBitmap/roaring v1.2.1 // indirect
	github.com/alibaba/sentinel-golang v1.0.4 // indirect
	github.com/aliyun/aliyun-oss-go-sdk v3.0.2+incompatible // indirect
	github.com/andeya/ameda v1.5.3 // indirect
	github.com/andeya/goutil v1.0.1 // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/antonmedv/expr v1.15.5 // indirect
	github.com/apache/rocketmq-client-go/v2 v2.1.2 // indirect
	github.com/apache/thrift v0.21.0
	github.com/apaxa-go/helper v0.0.0-20180607175117-61d31b1c31c3 // indirect
	github.com/bahlo/generic-list-go v0.2.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.24.3 // indirect
	github.com/bits-and-blooms/bloom/v3 v3.7.1 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/bluele/gcache v0.0.2 // indirect
	github.com/bufbuild/protocompile v0.14.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/bytedance/go-tagexpr/v2 v2.9.6 // indirect
	github.com/bytedance/sonic/loader v0.5.0 // indirect
	github.com/caarlos0/env/v6 v6.10.1 // indirect
	github.com/cenk/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/certifi/gocertifi v0.0.0-20210507211836-431795d63e8d // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/choleraehyq/pid v0.0.22 // indirect
	github.com/choleraehyq/rwlock v0.0.16 // indirect
	github.com/clipperhouse/uax29/v2 v2.6.0 // indirect
	github.com/cloudflare/circl v1.6.1 // indirect
	github.com/cloudflare/golz4 v0.0.0-20150217214814-ef862a3cdc58 // indirect
	github.com/cloudwego/base64x v0.1.6 // indirect
	github.com/cloudwego/configmanager v0.2.3 // indirect
	github.com/cloudwego/dynamicgo v0.8.0 // indirect
	github.com/cloudwego/eino-ext/libs/acl/openai v0.1.17 // indirect
	github.com/cloudwego/frugal v0.3.1 // indirect
	github.com/cloudwego/gopkg v0.1.8 // indirect
	github.com/cloudwego/hertz v0.10.3 // indirect
	github.com/cloudwego/kitex/pkg/protocol/bthrift v0.0.0-20260525122007-574e04ad3c04 // indirect
	github.com/cloudwego/localsession v0.2.1 // indirect
	github.com/cloudwego/netpoll v0.7.2 // indirect
	github.com/cloudwego/runtimex v0.1.1 // indirect
	github.com/cloudwego/thriftgo v0.4.3 // indirect
	github.com/coocood/freecache v1.2.4 // indirect
	github.com/coze-dev/cozeloop-go v0.1.21 // indirect
	github.com/coze-dev/cozeloop-go/spec v0.1.8 // indirect
	github.com/cyphar/filepath-securejoin v0.4.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/dgrijalva/jwt-go v3.2.1-0.20180921172315-3af4c746e1c2+incompatible // indirect
	github.com/dlclark/regexp2 v1.11.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/emirpasic/gods v1.18.1 // indirect
	github.com/evalphobia/logrus_sentry v0.8.2 // indirect
	github.com/evanphx/json-patch v0.5.2 // indirect
	github.com/facebookgo/clock v0.0.0-20150410010913-600d898af40a // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fatih/structtag v1.2.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/getsentry/raven-go v0.2.0 // indirect
	github.com/getsentry/sentry-go v0.35.0 // indirect
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.9.1 // indirect
	github.com/go-faster/city v1.0.1 // indirect
	github.com/go-faster/errors v0.7.1 // indirect
	github.com/go-git/gcfg v1.5.1-0.20230307220236-3a3c6141e376 // indirect
	github.com/go-jose/go-jose/v3 v3.0.4 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.1 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-playground/validator/v10 v10.16.0 // indirect
	github.com/go-task/slim-sprig v0.0.0-20230315185526-52ccab3ef572 // indirect
	github.com/gobuffalo/envy v1.7.0 // indirect
	github.com/gobuffalo/packd v1.0.0 // indirect
	github.com/gobuffalo/packr v1.30.1 // indirect
	github.com/goccy/go-json v0.10.5 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.5.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/golang/glog v1.2.2 // indirect
	github.com/golang/groupcache v0.0.0-20241129210726-2c02b8208cf8 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v1.0.0 // indirect
	github.com/gomarkdown/markdown v0.0.0-20250311123330-531bef5e742b // indirect
	github.com/google/flatbuffers v25.1.24+incompatible // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/google/pprof v0.0.0-20240827171923-fa2c70bbbfe5 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/goph/emperror v0.17.2 // indirect
	github.com/gopherjs/gopherjs v1.17.2 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/gotnospirit/makeplural v0.0.0-20180622080156-a5f48d94d976 // indirect
	github.com/gotnospirit/messageformat v0.0.0-20221001023931-dfe49f1eb092 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-hclog v1.2.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.1-vault-3 // indirect
	github.com/hbollon/go-edlib v1.7.0 // indirect
	github.com/henrylee2cn/ameda v1.5.1 // indirect
	github.com/henrylee2cn/flagx v1.5.4 // indirect
	github.com/henrylee2cn/goutil v0.0.0-20221115092640-94288b660c35 // indirect
	github.com/hertz-contrib/http2 v0.1.1 // indirect
	github.com/hertz-contrib/localsession v0.1.0 // indirect
	github.com/huandu/skiplist v1.2.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/influxdata/influxdb1-client v0.0.0-20220302092344-a9ab5670611c // indirect
	github.com/invopop/yaml v0.1.0 // indirect
	github.com/jbenet/go-context v0.0.0-20150711004518-d14ea06fba99 // indirect
	github.com/jhump/protoreflect v1.17.0 // indirect
	github.com/jinzhu/configor v1.2.1 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/jtolds/gls v4.20.0+incompatible // indirect
	github.com/kevinburke/ssh_config v1.2.0 // indirect
	github.com/klauspost/cpuid/v2 v2.3.0 // indirect
	github.com/klauspost/crc32 v1.2.0 // indirect
	github.com/klauspost/pgzip v1.2.6 // indirect
	github.com/kuangchanglang/graceful v1.0.2 // indirect
	github.com/larksuite/botframework-go v0.0.0-20210409135442-ff14b99e324b // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/logrusorgru/aurora v2.0.3+incompatible // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/meguminnnnnnnnn/go-openai v0.1.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mozillazg/go-httpheader v0.2.1 // indirect
	github.com/mschoch/smat v0.2.0 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.1.2 // indirect
	github.com/nikolalohinski/gonja v1.5.3 // indirect
	github.com/nyaruka/phonenumbers v1.3.2 // indirect
	github.com/olivere/elastic/v7 v7.0.32 // indirect
	github.com/onsi/ginkgo/v2 v2.20.1 // indirect
	github.com/onsi/gomega v1.35.1 // indirect
	github.com/openai/openai-go/v2 v2.7.1 // indirect
	github.com/opentracing/opentracing-go v1.2.1-0.20210726034734-bdbb7cc3a1c0 // indirect
	github.com/orcaman/concurrent-map v1.0.0 // indirect
	github.com/panjf2000/ants/v2 v2.10.0 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/paulmach/orb v0.11.1 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pelletier/go-toml/v2 v2.0.9 // indirect
	github.com/perimeterx/marshmallow v1.1.4 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pierrec/lz4/v4 v4.1.22 // indirect
	github.com/pingcap/errors v0.11.5-0.20240311024730-e056997136bb // indirect
	github.com/pjbgf/sha1cd v0.3.2 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20240221224432-82ca36839d55 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/reactivex/rxgo/v2 v2.5.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/rubyist/circuitbreaker v2.2.1+incompatible // indirect
	github.com/segmentio/asm v1.2.0 // indirect
	github.com/sergi/go-diff v1.4.0 // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/shopspring/decimal v1.4.0 // indirect
	github.com/sirupsen/logrus v1.9.4 // indirect
	github.com/skeema/knownhosts v1.3.1 // indirect
	github.com/slongfield/pyfmt v0.0.0-20220222012616-ea85ff4c361f // indirect
	github.com/smarty/assertions v1.15.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.16.0 // indirect
	github.com/spiffe/spire-api-sdk v1.9.6 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/teivah/onecontext v0.0.0-20200513185103-40f981bfd775 // indirect
	github.com/tencentyun/cos-go-sdk-v5 v0.7.30 // indirect
	github.com/thoas/go-funk v0.9.3 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.2.0 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	github.com/tinylib/msgp v1.1.9 // indirect
	github.com/tklauser/go-sysconf v0.3.13 // indirect
	github.com/tklauser/numcpus v0.7.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.12 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fastrand v1.1.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/vmihailenco/msgpack/v4 v4.3.12 // indirect
	github.com/vmihailenco/tagparser v0.1.2 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/volcengine/ve-tos-golang-sdk/v2 v2.7.20 // indirect
	github.com/xanzy/ssh-agent v0.3.3 // indirect
	github.com/yargevad/filepathx v1.0.0 // indirect
	github.com/yuin/gopher-lua v1.1.1 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	github.com/zeebo/errs v1.4.0 // indirect
	github.com/ztrue/tracerr v0.3.0 // indirect
	go.mozilla.org/pkcs7 v0.0.0-20210826202110-33d05740a352 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel v1.31.0 // indirect
	go.opentelemetry.io/otel/trace v1.31.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	go4.org/unsafe/assume-no-moving-gc v0.0.0-20231121144256-b99613f794b6 // indirect
	golang.org/x/arch v0.16.0 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	golang.org/x/time v0.12.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241104194629-dd2ea8efbc28 // indirect
	google.golang.org/grpc v1.69.2 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1 // indirect
	gopkg.in/vmihailenco/msgpack.v2 v2.9.2 // indirect
	gopkg.in/warnings.v0 v0.1.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/driver/mysql v1.5.7 // indirect
	stathat.com/c/consistent v1.0.0 // indirect
)

replace github.com/bytedance/mockey => /mockey
