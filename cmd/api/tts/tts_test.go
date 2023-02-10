package tts

import (
	b64 "encoding/base64"
	"github.com/golang/protobuf/proto"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"log"
	"testing"
	"wcrwg-iot-ingress/pkg/types"
	"wcrwg-iot-ingress/pkg/utils"
)

func TestDecodeV3(t *testing.T) {
	postbodies := []string{
		//`{"end_device_ids":{"device_id":"cricket-002","application_ids":{"application_id":"jpm-crickets"},"dev_addr":"260BEAC2"},"correlation_ids":["as:up:01F55DCJ0AAJPSX90MWGKX0A41","gs:conn:01F4S0KGVPX2QJQFH6K8KE47FR","gs:up:host:01F4S0KGZTYC78NWW5Q86GK1S7","gs:uplink:01F55DCHS8EA3P0HBMKPV9QP0V","ns:uplink:01F55DCHSKX3D9CYNKRV6RSG3W","rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01F55DCHSJRM33M74T3WG5HPYF","rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01F55DCJ0A688GR1D7QZWH545S"],"received_at":"2021-05-08T07:17:07.723112244Z","uplink_message":{"f_cnt":23171,"frm_payload":"","decoded_payload":{},"rx_metadata":[{"gateway_ids":{"gateway_id":"eui-000080029c09dd87","eui":"000080029C09DD87"},"timestamp":1408363515,"rssi":-27,"channel_rssi":-27,"snr":10.5,"location":{"latitude":-33.93667538260562,"longitude":18.871081173419956,"source":"SOURCE_REGISTRY"},"uplink_token":"CiIKIAoUZXVpLTAwMDA4MDAyOWMwOWRkODcSCAAAgAKcCd2HEPvXx58FGgwI8/XYhAYQx8nF7AEg+JiHyP6FXw==","channel_index":3},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01F55DCHTCMTM0SSJ1RSN3BBJ9","forwarder_net_id":"000013","forwarder_tenant_id":"ttn","forwarder_cluster_id":"ttn-v2-eu-3","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1","hops":[{"received_at":"2021-05-08T07:17:07.532727295Z","sender_address":"40.113.68.198","receiver_name":"router-dataplane-57d9d9bddd-dsrjj","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.533400574Z","sender_name":"router-dataplane-57d9d9bddd-dsrjj","sender_address":"forwarder_uplink","receiver_name":"router-5b5dc54cf7-psxlt","receiver_agent":"pbrouter/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.534011884Z","sender_name":"router-5b5dc54cf7-psxlt","sender_address":"deliver.000013_ttn_ttn-eu1.uplink","receiver_name":"router-dataplane-57d9d9bddd-f7h6k","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"}]},"rssi":-25,"channel_rssi":-25,"snr":9.5,"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU2sxaVJXeEVaRmMxV0dSRWFFZE9Sa0p4VlVaa1NVbHBkMmxrUjBadVNXcHZhVlV3ZHpSU2JsSkpWRzA1V2sweVNUTldiR2hDVWxVeFVsVlZTbEphZVVvNUxtc3RlRTh0WkhwNlJtRkdUbTV0VlZCWVZXOWtNMUV1YTBSaVZHMUJXbXhVZVROWFdtbEpVaTVoYldwVVpVZHdVVTFKWVZWT1RsSnRTR3BKWW5seFJrcFpZMUI2WDB4dVdsOUlUalJpWVcxR1psTmxRbTV3TTAxYU56a3hXblk1ZUdFMVV6QlFVbEJXYzBKbmExWk5UV1psVmxGRFRWSnRSM0JvWkdFMloxZEROMkZtWlZSbk9FVkdkbEUzWVRSelZrbzROMXBEWDJKeGIwbFJjbTFZTm1WU2JGOHdaaTFrYUZwU2RVbzBlRVZTZG1kRVUwbE1PRGxmT0dGQlNUa3lhVGg1YzJSeFpGOXdMWEUwUkU5aVh6TTRUM3BITG5ZeVpXcHFSM0Z5U2kwMUxWaHJlVTVTWkZwblIyYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19"},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01F55DCHTK8M44A02PSWXSTNF4","forwarder_net_id":"000013","forwarder_tenant_id":"ttn","forwarder_cluster_id":"ttn-v2-eu-3","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1","hops":[{"received_at":"2021-05-08T07:17:07.539820178Z","sender_address":"40.113.68.198","receiver_name":"router-dataplane-57d9d9bddd-xjszp","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.540236935Z","sender_name":"router-dataplane-57d9d9bddd-xjszp","sender_address":"forwarder_uplink","receiver_name":"router-5b5dc54cf7-mwf8m","receiver_agent":"pbrouter/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.540851394Z","sender_name":"router-5b5dc54cf7-mwf8m","sender_address":"deliver.000013_ttn_ttn-eu1.uplink","receiver_name":"router-dataplane-57d9d9bddd-f7h6k","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"}]},"rssi":-28,"channel_rssi":-28,"snr":8.5,"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU2tWa2JWbzFZVVJSTTFSdGN6Tk5NMG95V2pJMWMwbHBkMmxrUjBadVNXcHZhVTlJVms5a01rcEhWMVpXWVdGclZubGpWemx5VGxWb1dWTXpSbFJSVTBvNUxtaEVZelJGUlhaNGMyUXpaVEZuYjBvM2VYRXRjRUV1YXpkYWJ6Um5UamRRVjNsdFREWkxkUzVRVTBzMk9XdGpaMlpyY0habGRuUjZRa3hJWTBkcVJIZEhWSFY1YW5NdGRWWmlRbFozYlRFMGJuSkVkR2xZTVdOYVNHUkpaa05WVUUxdmNEVkxNSFZ5T0RsdlZHUTRkMWhMY1VsWFIzaFNNemxvWlhaVGJXbFJNbWhPUWs0d1F6VlhNRWwyYnkwNU5WbzBXbGcxYjJOSWVrNVVNVmRVTTNwUFNFMWxlREF4WnprMk5GUnJUMmN4YkdaQlZVNVJTWEZ0YVRsVmNFbzBiemhFZDBSM2QwODBhM3BOTTNGUmN6WlVaVWhvTGsxNmExRnJaMlUwYXpoYVRtWllUWHBIVTBkQ1puYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19"},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01F55DCHVF2ES39GECE99FPC8M","forwarder_net_id":"000013","forwarder_tenant_id":"ttn","forwarder_cluster_id":"ttn-v2-eu-1","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1","hops":[{"received_at":"2021-05-08T07:17:07.567259072Z","sender_address":"52.169.73.251","receiver_name":"router-dataplane-57d9d9bddd-f7h6k","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.568763883Z","sender_name":"router-dataplane-57d9d9bddd-f7h6k","sender_address":"forwarder_uplink","receiver_name":"router-5b5dc54cf7-psxlt","receiver_agent":"pbrouter/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.570193353Z","sender_name":"router-5b5dc54cf7-psxlt","sender_address":"deliver.000013_ttn_ttn-eu1.uplink","receiver_name":"router-dataplane-57d9d9bddd-dsrjj","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"}]},"time":"2021-05-08T07:17:06.010555982Z","rssi":-33,"channel_rssi":-33,"snr":9.5,"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU2pWV2JXeHJVekl4ZW1WcVJteGtWM0ExV2tWd2JrbHBkMmxrUjBadVNXcHZhV0p1YUhWaldHaHJVbXMxVWs5RVZUVmFha1pPVkd4S1MyVnVTWGhhZVVvNUxraHNablpYU1MxcVNuVjBNVzFqT0VkUlNVWmhURUV1ZURaV1FqVjVPVGN0Ykdob1pXOHhXUzVyV0doSFpYazBXbTR4Um5aaVZXbzVaMDV4Tmw5dFNFVTJSMVp2T0ZsVVZWTkhRVVYwWm1aWVFYSllZbEJKZFVGeFYyaEhha3BVTTBGdVJtSm5SblptVEZoSExVSnhhMkpOYUZGbFZFVTBlbFpYWkRCSVZHSnpWSE5oU3pKYWR6bFBOelV3VTJoWFluVTVTa3B0VjAxUmNGQk1XR052VFhGVVJsRmhlV3A0TURGSk1uTlRRMUZHY2tGS2NsZE9lblpmTjI1aldYbzJaRTlwTm00MWJrTTNPRk56ZVZwRk5UbFJSakZETGpRNWNVUTFVRmRPU1d3NWFHRmFlbU5oYjNWS1ptYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0xIn19"},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01F55DCHVTK41WZ054G87YDX87","forwarder_net_id":"000013","forwarder_tenant_id":"ttn","forwarder_cluster_id":"ttn-v2-eu-4","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1","hops":[{"received_at":"2021-05-08T07:17:07.578226151Z","sender_address":"52.169.150.138","receiver_name":"router-dataplane-57d9d9bddd-xjszp","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.582886038Z","sender_name":"router-dataplane-57d9d9bddd-xjszp","sender_address":"forwarder_uplink","receiver_name":"router-5b5dc54cf7-xh822","receiver_agent":"pbrouter/1.5.2 go/1.16.2 linux/amd64"},{"received_at":"2021-05-08T07:17:07.585417009Z","sender_name":"router-5b5dc54cf7-xh822","sender_address":"deliver.000013_ttn_ttn-eu1.uplink","receiver_name":"router-dataplane-57d9d9bddd-f7h6k","receiver_agent":"pbdataplane/1.5.2 go/1.16.2 linux/amd64"}]},"rssi":-67,"channel_rssi":-67,"snr":10,"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU1RWalJHeDJaVlUxUjJFeVNtOWhSVTVYVGxWYVFrbHBkMmxrUjBadVNXcHZhV0V6UVRCU2JHUkpaV3BDU1dGWFpFVlNNMDB6V2pOc05XVkZUbTlWVTBvNUxraG9iRFpaYmpobWRqQm9jakpuVEc5T09HSTVaVkV1VjFCVFJuUktUa1JCUW1Gc05UWlRXaTV4TTJSeE9VOUlhRlU0WTI4MVJtRnJUazEwYkdKMGQyRlpRVWhKY0cxVGMySmFUVFpRYkRSb2JtVXpiVmxrUkhCNE9YTkRTaTFCWkMxTlZXSmpaell6V0VoWVFuUnRWM0JuYm1KZmFXNURRbUpYUnpGNFNYVnJiM0JxUWtkWVQyUndMUzFuVm5CNVoyWkZNbmhIY1dWS1dIRXdaMnBSTldNd2MxWnVUbGd0WjJsRWIyVnRSRjlDYkcxaU1XUjNNR0o2Y1ZsWE4ybEZRMVoyVUhBNWNESnFjVVprYm5SV2QyUmZNVFZWTG5aWVdISjNjbEU0ZEdKeVQwTnllbmw1TTA5SlZrRT0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS00In19"}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":7}},"data_rate_index":5,"coding_rate":"4/5","frequency":"867100000","timestamp":1408363515},"received_at":"2021-05-08T07:17:07.507008642Z","consumed_airtime":"0.041216s","locations":{"user":{"latitude":-33.93623477040523,"longitude":18.871655166149143,"source":"SOURCE_REGISTRY"}}}}`,
		//`{"end_device_ids":{"device_id":"cricket-001","application_ids":{"application_id":"jpm-crickets"},"dev_addr":"26011CE4"},"correlation_ids":["as:up:01E175D2K6EHZH7GGH9TWRCVBN","gs:conn:01E16YPNYG4HEXHYJ7VFYKH2EW","gs:uplink:01E175D2AYR39QT12BY0ESMPP7","ns:uplink:01E175D2AZPJF4RDZH7A5EP2BS","rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01E175D2AYJYFSCZ6NMXKJ2QWQ"],"received_at":"2020-02-16T14:10:59.302096081Z","uplink_message":{"f_port":1,"f_cnt":527,"frm_payload":"AIj60lkC4SQAMY8=","decoded_payload":{"gps_0":{"altitude":126.87000274658203,"latitude":-33.93669891357422,"longitude":18.870800018310547}},"rx_metadata":[{"gateway_ids":{"gateway_id":"pisupply-shield","eui":"B827EBFFFED88375"},"timestamp":2732493451,"rssi":-72,"channel_rssi":-72,"snr":9.8,"uplink_token":"Ch0KGwoPcGlzdXBwbHktc2hpZWxkEgi4J+v//tiDdRCLlfqWCg=="}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":7}},"data_rate_index":5,"coding_rate":"4/5","frequency":"868100000","timestamp":2732493451},"received_at":"2020-02-16T14:10:59.039048589Z"}}`,
		//`{"end_device_ids":{"device_id":"cricket-002","application_ids":{"application_id":"jpm-crickets"},"dev_addr":"260BEAC2"},"correlation_ids":["as:up:01FD2A6GEN2F1VY4GEEVCJTM62","gs:conn:01FD1ZZQ1T6CEAMJ29PVCAKKF8","gs:up:host:01FD1ZZQ226GN31YFKSCN3SN15","gs:uplink:01FD2A6G7ZJHR911MNB5G44EYF","ns:uplink:01FD2A6G81E8TH80T8GJ9WXASM","rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01FD2A6G80R72JW4DY3QHN4N4K","rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01FD2A6GEM3Q5TWEJ8VPKQY14F"],"received_at":"2021-08-14T12:29:15.095206575Z","uplink_message":{"f_cnt":33540,"frm_payload":"","decoded_payload":{},"rx_metadata":[{"gateway_ids":{"gateway_id":"eui-000080029c09dd87","eui":"000080029C09DD87"},"timestamp":4266137228,"rssi":-33,"channel_rssi":-33,"snr":8.5,"location":{"latitude":-33.93667538260562,"longitude":18.871081173419956,"source":"SOURCE_REGISTRY"},"uplink_token":"CiIKIAoUZXVpLTAwMDA4MDAyOWMwOWRkODcSCAAAgAKcCd2HEIytoPIPGgwImu7eiAYQw9akowMg4KXgzJT2Ag==","channel_index":7},{"gateway_ids":{"gateway_id":"eui-58a0cbfffe80049a","eui":"58A0CBFFFE80049A"},"time":"2021-08-14T12:29:14.725533962Z","timestamp":2108938836,"rssi":-41,"channel_rssi":-41,"snr":8.75,"uplink_token":"CiIKIAoUZXVpLTU4YTBjYmZmZmU4MDA0OWESCFigy//+gASaENS0z+0HGgwImu7eiAYQq8H1qQMgoLCztLC3Ag=="},{"gateway_ids":{"gateway_id":"eui-60c5a8fffe71a964","eui":"60C5A8FFFE71A964"},"timestamp":3259758379,"rssi":-74,"channel_rssi":-74,"snr":8.8,"uplink_token":"CiIKIAoUZXVpLTYwYzVhOGZmZmU3MWE5NjQSCGDFqP/+calkEKvur5IMGgwImu7eiAYQ+va7sAMg+P/1xe+/NA==","channel_index":4},{"gateway_ids":{"gateway_id":"packetbroker"},"packet_broker":{"message_id":"01FD2A6GA8GDGP8GR539JHB45H","forwarder_net_id":"000013","forwarder_tenant_id":"ttnv2","forwarder_cluster_id":"ttn-v2-eu-3","forwarder_gateway_eui":"647FDAFFFE007A1F","forwarder_gateway_id":"eui-647fdafffe007a1f","home_network_net_id":"000013","home_network_tenant_id":"ttn","home_network_cluster_id":"ttn-eu1"},"rssi":-37,"channel_rssi":-37,"snr":8.2,"location":{"latitude":-33.93626794,"longitude":18.87168703},"uplink_token":"eyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU1RWTlZrcG1WVmR3U2xReFJuSmtWVlpoVTBkR1ZVbHBkMmxrUjBadVNXcHZhV050YkhWbFJXeHhUMWhzWm1KSFRubFJibFpSVkVka2JFOUZSalpSVTBvNUxrSkhha3hVTURkaE1GQnVUVEpCVTBWUlluTjRWVUV1WW5NdFZXbFRTMjV2WjBaa01VbEdaQzVSUXpGNVdtdFpTV3BrU1hKbk1TMXFUalpUWkhsRFRsTTRNMjVSYWpFNVMxaHBRemhHTFVOVmJGaE9SbGM0TUdWWU1tWklZVVZoUlZCUFRURk5PSEJTZUVjdGJWQkRNMkpUTFhaUU9EaE9ZamwxVFhkUllsVlVkM0IwUTFaNmFrMXdSR05TVDBaNFZGVjROazVoTmt3MlVISktObUZ1YVdGUFZraGtWVXBrUW5Oc2RHOTZjV05qU21rNVprVlBZemMxWDBaUlVHeExRV3hYUVZwTFZrUnNNMWxSVmpkME5VOTZaeTVJY21wdVFrVTVkRlJTYzJKTGFEUTRZWFJaY2taUiIsImEiOnsiZm5pZCI6IjAwMDAxMyIsImZ0aWQiOiJ0dG52MiIsImZjaWQiOiJ0dG4tdjItZXUtMyJ9fQ=="}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":7}},"data_rate_index":5,"coding_rate":"4/5","frequency":"867900000","timestamp":4266137228},"received_at":"2021-08-14T12:29:14.881054809Z","consumed_airtime":"0.041216s","network_ids":{"net_id":"000013","tenant_id":"ttn","cluster_id":"ttn-eu1"}}}`,
		//`{"@type":"type.googleapis.com/ttn.lorawan.v3.ApplicationUp","end_device_ids":{"device_id":"eui-3533363557307f06","application_ids":{"application_id":"selati-awt-tags"},"dev_eui":"3533363557307F06","join_eui":"49728720E2D64FB3","dev_addr":"260BE66F"},"correlation_ids":["as:up:01GB033KXR4FT6XFR3PEHN64VQ","gs:conn:01GAXCHNH2RNY6AVY0MFW3N4KD","gs:up:host:01GAXCHNH73PGWC7C6Q7KX0KQW","gs:uplink:01GB033KQ5Y0TQH0F70ZXVHX9V","ns:uplink:01GB033KQ6BKD9YS84M5KFXW3N","rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01GB033KQ6JMEK3VCRH9V4VYJ6","rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01GB033KXRF1ZGZDBBJ0HV61RC"],"received_at":"2022-08-21T11:37:46.168847136Z","uplink_message":{"session_key_id":"AYHwRyaij1T3YkMux85mbg==","f_port":2,"f_cnt":1293,"frm_payload":"k9HRILMBlRdVAgAAZgIBcWMCGQGJbnxTgdDD2QAA6EkOzDgaMUo=","decoded_payload":{"alarm":0,"battery":3.69,"global_ack":0,"gps_fix":true,"hdop":1,"latitude":-24.030005,"longitude":30.764735,"movement":0,"speed":0,"temperature":49,"timestamp":1661081857},"rx_metadata":[{"gateway_ids":{"gateway_id":"selati-04","eui":"AC1F09FFFE065276"},"time":"2022-08-21T11:37:45.811Z","timestamp":903776268,"rssi":-124,"channel_rssi":-124,"snr":-7.5,"location":{"latitude":-23.989572,"longitude":30.771441,"source":"SOURCE_REGISTRY"},"uplink_token":"ChcKFQoJc2VsYXRpLTA0EgisHwn//gZSdhCMkPquAxoMCImyiJgGEIHWwMgDIODdjeqm2BUqDAiJsoiYBhDAwduCAw==","channel_index":5,"gps_time":"2022-08-21T11:37:45.811Z","received_at":"2022-08-21T11:37:45.831825114Z"}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":12}},"coding_rate":"4/5","frequency":"868100000","timestamp":903776268,"time":"2022-08-21T11:37:45.811Z"},"received_at":"2022-08-21T11:37:45.958459135Z","confirmed":true,"consumed_airtime":"2.465792s","locations":{"frm-payload":{"latitude":-24.030005,"longitude":30.764735,"source":"SOURCE_GPS"}},"version_ids":{"band_id":"EU_863_870"},"network_ids":{"net_id":"000013","tenant_id":"ttn","cluster_id":"eu1","cluster_address":"eu1.cloud.thethings.network"}}}`,
		`{
    "@type": "type.googleapis.com/ttn.lorawan.v3.ApplicationUp",
    "end_device_ids": {
      "device_id": "1332000002",
      "application_ids": {
        "application_id": "izinto-misol-lora"
      },
      "dev_eui": "003CBCC074A82381",
      "join_eui": "70B3D57ED00194A5",
      "dev_addr": "260BA3B8"
    },
    "correlation_ids": [
      "as:up:01GF3MJ0S4NTR8QWA2Y2QXXWB7",
      "gs:conn:01GF1HDBEG1MVFV8V8HZ2KG24R",
      "gs:up:host:01GF1HDBEPN3N7QCEVYW36ECVD",
      "gs:uplink:01GF3MJ0JNHGJE114N8FMJYJN2",
      "ns:uplink:01GF3MJ0JP8J5KMC002PJSD9XM",
      "rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01GF3MJ0JP3X02Y1CQ834GHC1V",
      "rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01GF3MJ0S3SAXPDF6SST2B7G2C"
    ],
    "received_at": "2022-10-11T13:43:14.468158091Z",
    "uplink_message": {
      "session_key_id": "AX0TvN0hsyqbXlLnBpkOlA==",
      "f_port": 1,
      "f_cnt": 45906,
      "frm_payload": "AQqLCkYefBr3/T39JgAAAAAAAP//CgI=",
      "decoded_payload": {
        "barometer": 1016.4817499999999,
        "barometerMax": 1016.4829,
        "barometerMin": 1016.4806,
        "battery": 2.562,
        "firmwareVersion": 1,
        "humidity": 73.535,
        "humidityMax": 78.04,
        "humidityMin": 69.03,
        "rain": 0,
        "temperature": 26.645,
        "temperatureMax": 26.99,
        "temperatureMin": 26.3,
        "wind": 0,
        "windGust": 0
      },
      "rx_metadata": [
        {
          "gateway_ids": {
            "gateway_id": "eui-60c5a8fffe761551",
            "eui": "60C5A8FFFE761551"
          },
          "timestamp": 1701195036,
          "rssi": -83,
          "channel_rssi": -83,
          "snr": 9,
          "location": {
            "latitude": -33.88073487185602,
            "longitude": 19.050543995042034,
            "source": "SOURCE_REGISTRY"
          },
          "uplink_token": "CiIKIAoUZXVpLTYwYzVhOGZmZmU3NjE1NTESCGDFqP/+dhVREJzamKsGGgsI8uWVmgYQpYfgfCDg6oC5wYEQ",
          "channel_index": 5,
          "received_at": "2022-10-11T13:43:14.261620645Z"
        },
        {
          "gateway_ids": {
            "gateway_id": "lnx-solutions-izinto-fhk",
            "eui": "323531321B004F00"
          },
          "time": "2022-10-11T13:43:14.155983Z",
          "timestamp": 122956940,
          "rssi": -98,
          "channel_rssi": -98,
          "snr": 9.25,
          "location": {
            "latitude": -33.90775925,
            "longitude": 19.07312776,
            "altitude": 679,
            "source": "SOURCE_REGISTRY"
          },
          "uplink_token": "CiYKJAoYbG54LXNvbHV0aW9ucy1pemludG8tZmhrEggyNTEyGwBPABCM2dA6GgwI8uWVmgYQuOvfgwEg4IW4hsqLpAE=",
          "received_at": "2022-10-11T13:43:14.276297144Z"
        }
      ],
      "settings": {
        "data_rate": {
          "lora": {
            "bandwidth": 125000,
            "spreading_factor": 7,
            "coding_rate": "4/5"
          }
        },
        "frequency": "868100000",
        "timestamp": 1701195036
      },
      "received_at": "2022-10-11T13:43:14.262742556Z",
      "consumed_airtime": "0.077056s",
      "network_ids": {
        "net_id": "000013",
        "tenant_id": "ttn",
        "cluster_id": "eu1",
        "cluster_address": "eu1.cloud.thethings.network"
      }
    }
  }`,
		`
{
    "@type": "type.googleapis.com/ttn.lorawan.v3.ApplicationUp",
    "end_device_ids": {
      "device_id": "eui-3030303130373332",
      "application_ids": {
        "application_id": "wcrwg-otto-io"
      },
      "dev_eui": "3030303130373332",
      "join_eui": "70B3D57ED003E052",
      "dev_addr": "260B6564"
    },
    "correlation_ids": [
      "as:up:01GR35H53P5YZYNNHREZXCR60N",
      "gs:conn:01GR1GXJ3SNMPQ6JTNCP3X9C5A",
      "gs:up:host:01GR1GXJ9CE76JBYGJ30M14YRX",
      "gs:uplink:01GR35H4X8KGMHKTS9NAD7FTQE",
      "ns:uplink:01GR35H4X9Z2WNE05C5D29WF6P",
      "rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01GR35H4X8V45AX9Q2ANT59AAM",
      "rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01GR35H53PTV672BTFY4CW916V"
    ],
    "received_at": "2023-01-31T05:41:53.910284534Z",
    "uplink_message": {
      "session_key_id": "AYYF5fJaPYG8XyFn8znCLw==",
      "f_port": 1,
      "f_cnt": 15,
      "frm_payload": "CUIDIgMARbDYY3IS",
      "decoded_payload": {
        "analog_2": 834,
        "analog_4": 802,
        "battery": 4.722,
        "digital_out_1": 0,
        "message_type": "periodic",
        "timestamp": 1675145285
      },
      "rx_metadata": [
        {
          "gateway_ids": {
            "gateway_id": "lnx-solutions-knkop-01",
            "eui": "3436323826004400"
          },
          "time": "2022-05-17T12:11:37.926841Z",
          "timestamp": 3236781044,
          "rssi": -115,
          "channel_rssi": -115,
          "snr": -4,
          "location": {
            "latitude": -33.82608772160143,
            "longitude": 18.606270281597975,
            "altitude": 130,
            "source": "SOURCE_REGISTRY"
          },
          "uplink_token": "CiQKIgoWbG54LXNvbHV0aW9ucy1rbmtvcC0wMRIINDYyOCYARAAQ9Le1hwwaDAih1OKeBhDOg+XPAiCg4rz5mbcN",
          "channel_index": 4,
          "received_at": "2023-01-31T05:41:53.618111981Z"
        }
      ],
      "settings": {
        "data_rate": {
          "lora": {
            "bandwidth": 125000,
            "spreading_factor": 10,
            "coding_rate": "4/5"
          }
        },
        "frequency": "867300000",
        "timestamp": 3236781044,
        "time": "2022-05-17T12:11:37.926841Z"
      },
      "received_at": "2023-01-31T05:41:53.705098970Z",
      "confirmed": true,
      "consumed_airtime": "0.411648s",
      "network_ids": {
        "net_id": "000013",
        "tenant_id": "ttn",
        "cluster_id": "eu1",
        "cluster_address": "eu1.cloud.thethings.network"
      }
    }
  }`,
	}

	for _, postbody := range postbodies {
		marshaler := jsonpb.TTN()

		var packetIn ttnpb.ApplicationUp
		if err := marshaler.Unmarshal([]byte(postbody), &packetIn); err != nil {
			t.Error(err.Error())
		}

		//log.Printf("%+v", packetIn)

		packetInUplinkData := packetIn.GetUplinkMessage()
		if packetInUplinkData == nil {
			t.Error("Uplink not parsed")
		}
		//log.Printf("Uplink Message: %+v", packetIn.GetUplinkMessage())

		if packetIn.GetUplinkMessage().GetNetworkIds() == nil {
			t.Error("Network IDs not set")
			continue
		}

		sensorMessage, err := UplinkToSensorMessage(packetIn)
		if err != nil {
			t.Fatalf(err.Error())
		}

		log.Println(utils.PrettyPrint(sensorMessage))

		// 2. Gateway data
		gatewayMessages, err := UplinkToGatewayMessage(packetIn)
		if err != nil {
			t.Fatalf(err.Error())
		}
		log.Println(utils.PrettyPrint(gatewayMessages))

	}
}

func TestTtsV3Pb(t *testing.T) {
	protoMessages := [][]byte{}
	data, _ := b64.StdEncoding.DecodeString("CiMKC2NyaWNrZXQtMDAyEg4KDGpwbS1jcmlja2V0czIEJgvqwhIgYXM6dXA6MDFGNTZGQ0RKMTU0R0tCUFlKWjNUWkNDMkgSImdzOmNvbm46MDFGNFMwS0dWUFgyUUpRRkg2SzhLRTQ3RlISJWdzOnVwOmhvc3Q6MDFGNFMwS0daVFlDNzhOV1c1UTg2R0sxUzcSJGdzOnVwbGluazowMUY1NkZDREFZWDFFTjRNN1gwRlFXODM4UBIkbnM6dXBsaW5rOjAxRjU2RkNEQVpNVDRYU1A3WDJGTU4wVzg0EkBycGM6L3R0bi5sb3Jhd2FuLnYzLkdzTnMvSGFuZGxlVXBsaW5rOjAxRjU2RkNEQVpaSERaUDJBRzNZMDBEU00zEkBycGM6L3R0bi5sb3Jhd2FuLnYzLk5zQXMvSGFuZGxlVXBsaW5rOjAxRjU2RkNESjFDU1NGU1E5REo1N0RZREs5YgwIsozbhAYQoqjX5wIazCIYgLcBKgAykgEKIAoUZXVpLTAwMDA4MDAyOWMwOWRkODcSCAAAgAKcCd2HIJz/sIUKRQAAuMFNAAC4wV3NzCRBahQJNaGb+uT3QMARAQAALf/eMkAoA3pACiIKIAoUZXVpLTAwMDA4MDAyOWMwOWRkODcSCAAAgAKcCd2HEJz/sIUKGgwIsozbhAYQmOv0+gEg4PLhhrqTZ4gBBjKSCAoOCgxwYWNrZXRicm9rZXJFAACAwk0AAIDCXc3MDEF6nAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTbEJpYlRWMVducG9lbFZGYkVkaVYwNVdZbFk1TkVscGQybGtSMFp1U1dwdmFXUkhSVE5qTURWelpEQTFhR1JHYnpGYU0xSnNWSHBrY0Zac2FFNVJVMG81TGpOU1ZsOUVPRFo1ZGtkaFRpMVRNVFpxWlZsYWNVRXVTRGRhUlVWbVR6SmlOMlJSVjFGM2RTNU1WblJRT0dSbWFqbGhMVEJLY25GRFJGTm5NMFJwYVZWaWMxRkthRU5ZUkhwVFUzVkdSWFpXTmpReE1uZFZlakYwYm1jeWJteGlkVlZpZWxFelRYUXdSMkZPUm1aVlZqSnFlVUUxVkc1bFZqbGthRzAxUjBGZldVaHdNMGxhYm5KaFZGVkxaRGxLUVZvNU5IQkNZMUZxTm5CR1kyaHdlRVF6YVRjeVNXRnFTekJ3Y1Mxa1ltRmZjRk41U3pkUk9HOTZiMnRHY1VKNFQyaFpSR2hxWnpkSmFWSndaVTVmV21sVmRVZENMbEpWTVdaeE9FTnZSMjVHZVZsak1HdGFOSE5oWmtFPSIsImEiOnsiZm5pZCI6IjAwMDAxMyIsImZ0aWQiOiJ0dG4iLCJmY2lkIjoidHRuLXYyLWV1LTQifX2SAdADChowMUY1NkZDREMxSEgxRFdLMUs5N0JaRjlCQhIDAAATGgN0dG4iC3R0bi12Mi1ldS00KgMAABMyA3R0bjpqCgwIsozbhAYQrdvViwIaDjUyLjE2OS4xNTAuMTM4IiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqCAQoMCLKM24QGEP7Q+4sCEiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoaEGZvcndhcmRlcl91cGxpbmsiF3JvdXRlci01YjVkYzU0Y2Y3LW13ZjhtKiRwYnJvdXRlci8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6lgEKDAiyjNuEBhDa17mMAhIXcm91dGVyLTViNWRjNTRjZjctbXdmOG0aIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTEykQgKDgoMcGFja2V0YnJva2VyRQAADMJNAAAMwl3NzAxBepwEeyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU2xwbGJtUlJUbGh3U21Wc1ZuRmtWMFV4V1ZWMFdrbHBkMmxrUjBadVNXcHZhVkZZUWs5VlNHUnpZMnBKZVZac1kzbGtSRTV4V0RFNWFWcEhPV3RSVTBvNUxsYzRaVUYyYkU5a05WZ3plbGRTWWkwdFZtdGxaR2N1TURWV2EwbGlabDg1WVVsS2FGQnpWeTR6VWpkR2VtdG1SWFJJY25GeVdXbzVSUzAzYTIwd2IwbFZXVWQyV2xKWGFuTjJTWGh5YTNrM1ZVcFJjVlE0WjJGbU16bDRaMEZXYjFZMFRVVnZVRUZCUkhkV1NqWnRRM0Z6VnpaSVNHSjZUVjlzUkhOMll6ZFhXVVJrUWxjMldrTlZiRk56YUZGUVVscHNWMk5hVjFOb1RtWlNOa0ZvUmpsVlZuTm5UMnRPUVZwQlRDMTNiRkZ1YlhobmIxQmpiVnBOVGpaS1NuTlBZVVZVTWtoS1IydExUakExWW1sTk9FRXdhWFp5TG5sRFdFeEVZVk5JV0ZCdGRHRlJWa28xVkZWdWVYYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19kgHPAwoaMDFGNTZGQ0RDOEpCV1FYV0NDTVg4SEQxS0ISAwAAExoDdHRuIgt0dG4tdjItZXUtMyoDAAATMgN0dG46aQoMCLKM24QGEM6SqI8CGg00MC4xMTMuNjguMTk4IiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQteGpzenAqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqCAQoMCLKM24QGEPTN3Y8CEiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQteGpzenAaEGZvcndhcmRlcl91cGxpbmsiF3JvdXRlci01YjVkYzU0Y2Y3LW13ZjhtKiRwYnJvdXRlci8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6lgEKDAiyjNuEBhDyhJKQAhIXcm91dGVyLTViNWRjNTRjZjctbXdmOG0aIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTEykQgKDgoMcGFja2V0YnJva2VyRQAAkMFNAACQwV0AABhBepwEeyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU25kU2JURkxUbFZXZFZSc1JuSmlSRkp1VkROUk1VbHBkMmxrUjBadVNXcHZhV0pWVWxKaldHUk5VMGhrTmxwNU1UTlBWWGhGV1hwbk5WVXdjRTVhZVVvNUxqQTRYM015UkcwNFRqQnlWVlJVVjFaT2NWSlpTMmN1VEdSRFJXeGtUV3hOZW1NemFETmlhaTQyVFhGTVQyVXphVFJCTVZGWk56WndSbGt5WVc5VVFUUlpWSEUxYldOQk9HZGxVMFI2V2xwTk5XZzVhbTU2YzNoWGJVbEdhMlJUVXpsVmNqTTJkVkkwVmpKdE1YUjNlREpTYURKVFRsYzVOWGRGUlVOaFJEUkxXR0UyWDFZemRYcGFRelpCVldod1ltUnVhamhvTmxCd2NqWnFjQzFMYkZwT1JrOVhZM05QUkc1SGFYQmtkVUZ3YkRsUVNEWnJPWGsxWjFGSlYwcG1PVUp1VFZwNVRUbE1jSEJzZDFsTWNuVkdSMmxwTGpGcmRHNVJNbXBQYkZOWVVtbFBWMWRSTUdnMVZXYz0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19kgHPAwoaMDFGNTZGQ0RDR01QSkJBOFpHUkNLTlhZNlcSAwAAExoDdHRuIgt0dG4tdjItZXUtMyoDAAATMgN0dG46aQoMCLKM24QGEJrk95ICGg00MC4xMTMuNjguMTk4IiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqCAQoMCLKM24QGEPSyp5MCEiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoaEGZvcndhcmRlcl91cGxpbmsiF3JvdXRlci01YjVkYzU0Y2Y3LXBzeGx0KiRwYnJvdXRlci8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6lgEKDAiyjNuEBhCfkLiUAhIXcm91dGVyLTViNWRjNTRjZjctcHN4bHQaIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWY3aDZrKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTEynwgKDgoMcGFja2V0YnJva2VyGgwIsIzbhAYQ9Y3lwgNFAABAwk0AAEDCXQAAGEF6nAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTazFrUlZaR1UyNUNUV1JzVG1samJGWndaVVJTYlVscGQybGtSMFp1U1dwdmFVNVVXbk5UVm1zelVqQm9WbU13YUZWVVZYUTFWbXN4ZEdGRk1VWlJVMG81TGpKV1ZteFVkV2RMWTBaSU9HYzJVR0pxV0U5WmIzY3VORlF0YUZwTFMyZERZVWhqUjBRdFdTNWFhVkJCYTFCNU5UWjVVRlJtYkc5MmVGSkhUbXBQVkZSc1JscFZSalF3VVMxeGQzVXRNRjlTVFZGclVteEhNbGxrZFdwdWVHaEhMVzgxWVhaaGIybzBkMWx6ZGpOelpsVndWVzlTYjFrMU1uZFZNRVJpU0dkVWIyaFRXV1ozUlRJNFpVMTVTMVZKTWtndFpXUnBaRUpNYUZSVlprOXRWbmxKYzBrdFEzZEhiRlpaTVRkYWFsWmxjeTFaYlU4dGNtbHJWRlpQU2xjd2RGTTJSbEIzTjA5RWFWTXhkVWhEV2pNNFprSlpMbGhWT1ZOcmIwczNka1IyVFZGRVdEZDZOMlZUVEhjPSIsImEiOnsiZm5pZCI6IjAwMDAxMyIsImZ0aWQiOiJ0dG4iLCJmY2lkIjoidHRuLXYyLWV1LTEifX2SAc8DChowMUY1NkZDRENSR0RZVjVXOVdFMFJFTVJYUxIDAAATGgN0dG4iC3R0bi12Mi1ldS0xKgMAABMyA3R0bjppCgwIsozbhAYQkpDylgIaDTUyLjE2OS43My4yNTEiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoIBCgwIsozbhAYQiq6JmQISIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqahoQZm9yd2FyZGVyX3VwbGluayIXcm91dGVyLTViNWRjNTRjZjctbXdmOG0qJHBicm91dGVyLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqWAQoMCLKM24QGEKXrsJkCEhdyb3V0ZXItNWI1ZGM1NGNmNy1td2Y4bRohZGVsaXZlci4wMDAwMTNfdHRuX3R0bi1ldTEudXBsaW5rIiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZjdoNmsqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NEIHdHRuLWV1MTodCggKBgjI0AcQBxAFGgM0LzUgoJrgnQMwnP+whQpCDAiyjNuEBhDBwbn7AWoFEIDQ0xNyHAoEdXNlchIUCf4qsqJa+kDAEQEAAFChcDZAKAM=")
	protoMessages = append(protoMessages, data)
	data, _ = b64.StdEncoding.DecodeString("CiMKC2NyaWNrZXQtMDAyEg4KDGpwbS1jcmlja2V0czIEJgvqwhIgYXM6dXA6MDFGNTZGQ0dZM1BERlNWU1hDWTFCVzdGREsSImdzOmNvbm46MDFGNFMwS0dWUFgyUUpRRkg2SzhLRTQ3RlISJWdzOnVwOmhvc3Q6MDFGNFMwS0daVFlDNzhOV1c1UTg2R0sxUzcSJGdzOnVwbGluazowMUY1NkZDRzhNMjlHQ1I2QTkzS1ZYTjZXWRIkbnM6dXBsaW5rOjAxRjU2RkNHOE5FMU41VkY1NFM0VjRGNTI1EkBycGM6L3R0bi5sb3Jhd2FuLnYzLkdzTnMvSGFuZGxlVXBsaW5rOjAxRjU2RkNHOE0xS0FZMjg5U1RIUjJYM1c0EkBycGM6L3R0bi5sb3Jhd2FuLnYzLk5zQXMvSGFuZGxlVXBsaW5rOjAxRjU2RkNHWTM0S1RSWk5KUVg4OU5BU0pHYgsItozbhAYQ6JOtZRqjEhiBtwEqADKSAQogChRldWktMDAwMDgwMDI5YzA5ZGQ4NxIIAACAApwJ3Ycgg5nohgpFAACowU0AAKjBXQAAEEFqFAk1oZv65PdAwBEBAAAt/94yQCgDekAKIgogChRldWktMDAwMDgwMDI5YzA5ZGQ4NxIIAACAApwJ3YcQg5nohgoaDAi1jNuEBhCAjfL5ASC4v4OexZNniAEBMpEICg4KDHBhY2tldGJyb2tlckUAAATCTQAABMJdAADwQHqcBHsiZyI6IlpYbEthR0pIWTJsUGFVcENUVlJKTkZJd1RrNVRNV05wVEVOS2JHSnRUV2xQYVVwQ1RWUkpORkl3VGs1SmFYZHBZVmhaYVU5cFNUTmlNRVozWVcxa1JsWlZTWHBTZWtab1VYcGtVMGxwZDJsa1IwWnVTV3B2YVdJelVrZFpXRlp6VGtoT01FMUVVblJOZW14V1VXcEJNMDVJWkU5VlUwbzVMak5IYlVSbE4zbDJjbE5TTURKRlowWTRaV1ZuTTFFdVpIZDZNbWxaWmxOamRqQTBTa0UwUWk1NGRrRjBkMjE1WDNwVWFXZ3lTR1p3Vm05UWRGRmZWRWhzUlVSclUyVjBhemxCVDNnMVJWTnNhRUZRYUVsM1gza3dVVVZDUmtoMGR5MXpPVmRPUXkxd1pFNUZNRFZSZURWcWEweE1URlJ5VVZSeWN6aEZhRzAwTTAwNVJrbDJiMUphY1RacWRHOXRlalV3Ym00MGRXRTNMV1Z2WTI4eU9Xc3dWRmxsV1hRNVNHVTNiVlJuYzBwMldFczJkRTVmVjB4cmR6ZDVRV05IUWpkYWNrVjVla2xTUXpnNFgyUkhiRkpJT1Vod0xsVnFRalZhZWxsS1UzbEVlVGs0Y2pWSk9FMXdVMUU9IiwiYSI6eyJmbmlkIjoiMDAwMDEzIiwiZnRpZCI6InR0biIsImZjaWQiOiJ0dG4tdjItZXUtMyJ9fZIBzwMKGjAxRjU2RkNHQTkyTjY4OUJaMFhGTkRZRjA4EgMAABMaA3R0biILdHRuLXYyLWV1LTMqAwAAEzIDdHRuOmkKDAi1jNuEBhCv+76TAhoNNDAuMTEzLjY4LjE5OCIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLXhqc3pwKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6ggEKDAi1jNuEBhDL9u6UAhIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLXhqc3pwGhBmb3J3YXJkZXJfdXBsaW5rIhdyb3V0ZXItNWI1ZGM1NGNmNy14aDgyMiokcGJyb3V0ZXIvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OpYBCgwItYzbhAYQh9C2lQISF3JvdXRlci01YjVkYzU0Y2Y3LXhoODIyGiFkZWxpdmVyLjAwMDAxM190dG5fdHRuLWV1MS51cGxpbmsiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1mN2g2ayoncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0Qgd0dG4tZXUxMp8ICg4KDHBhY2tldGJyb2tlchoMCLOM24QGELvQ2L4DRQAAIMJNAAAgwl0AAAxBepwEeyJnIjoiWlhsS2FHSkhZMmxQYVVwQ1RWUkpORkl3VGs1VE1XTnBURU5LYkdKdFRXbFBhVXBDVFZSSk5GSXdUazVKYVhkcFlWaFphVTlwU2tOVlZVNU9UVVJXVldKRlRqVk9ibEpJVG0weFZrbHBkMmxrUjBadVNXcHZhVTVHUm5SYWJrSkZZek5PY1ZSdGFETlZSMnhPWkZaU1dtSnROREZhZVVvNUxuVkZTMnR2YjI1U00wSndRWEpyYkc5dllWUmxZMmN1ZDJZNVJFUkhOMFYwTUU5cmN6bDFkQzVPVTNsb1FsWlhWbWh5Wmw5NVZuTkVaRTh0ZGsxalNUZzVZa3RMUjNrMFQwcEZkVlZ5Y2xWd01HWmpSRTkyZEdaVUxWbFZWRmhXTlVKQk0zQkxNVlEyUjJsTmVqbE5aMmRoWVZNd1lqVmlaa2czYTAxeFNYSlNkV1kwZEdsT1RrUlplVFJSVVU5c1YzUnNabkJpTjBSdmMyYzRNa2h3WVZseFNFSkVVbkJKVUV4VlFXeHZhRTU1ZVdKTWIzZG9UM2RTWW5WTGVtNWpZakpWV2tKRFdsbGlVamxLTm05TVV6ZzFRMHMwTG1sdExXcFdla3B0V25scU16TTBPRXQxU3preE5GRT0iLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0xIn19kgHPAwoaMDFGNTZGQ0dBNk05NUgyR1lORkUzQkRSWDYSAwAAExoDdHRuIgt0dG4tdjItZXUtMSoDAAATMgN0dG46aQoMCLWM24QGEP6WlJICGg01Mi4xNjkuNzMuMjUxIiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZjdoNmsqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqCAQoMCLWM24QGEOaprpICEiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZjdoNmsaEGZvcndhcmRlcl91cGxpbmsiF3JvdXRlci01YjVkYzU0Y2Y3LW13ZjhtKiRwYnJvdXRlci8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6lgEKDAi1jNuEBhCzwN+TAhIXcm91dGVyLTViNWRjNTRjZjctbXdmOG0aIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWY3aDZrKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTE6HQoICgYIyNAHEAcQBRoDNC81IODphJ4DMIOZ6IYKQgwItYzbhAYQ9oG2+gFqBRCA0NMTchwKBHVzZXISFAn+KrKiWvpAwBEBAABQoXA2QCgD")
	protoMessages = append(protoMessages, data)
	data, _ = b64.StdEncoding.DecodeString("CiMKC2NyaWNrZXQtMDAyEg4KDGpwbS1jcmlja2V0czIEJgvqwhIgYXM6dXA6MDFGNTZGQ0tERlQ5UDU0WUtCTU1WM0haVjYSImdzOmNvbm46MDFGNFMwS0dWUFgyUUpRRkg2SzhLRTQ3RlISJWdzOnVwOmhvc3Q6MDFGNFMwS0daVFlDNzhOV1c1UTg2R0sxUzcSJGdzOnVwbGluazowMUY1NkZDSzZGOUJOS0tERERUQzVUTVdCThIkbnM6dXBsaW5rOjAxRjU2RkNLNkdTNjU2QktRQjNTVEo5MDJOEkBycGM6L3R0bi5sb3Jhd2FuLnYzLkdzTnMvSGFuZGxlVXBsaW5rOjAxRjU2RkNLNkdXRzRFR1JTMVNUSjhTVFQwEkBycGM6L3R0bi5sb3Jhd2FuLnYzLk5zQXMvSGFuZGxlVXBsaW5rOjAxRjU2RkNLREZKS1lIUzBXRzIwWlY3RFNOYgwIuIzbhAYQhbbV5gIazCIYgrcBKgAykgEKIAoUZXVpLTAwMDA4MDAyOWMwOWRkODcSCAAAgAKcCd2HIIyzn4gKRQAAmMFNAACYwV2amflAahQJNaGb+uT3QMARAQAALf/eMkAoA3pACiIKIAoUZXVpLTAwMDA4MDAyOWMwOWRkODcSCAAAgAKcCd2HEIyzn4gKGgwIuIzbhAYQu6Ci+wEg4JWntdCTZ4gBBTKRCAoOCgxwYWNrZXRicm9rZXJFAACIwU0AAIjBXWZm5kB6nAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTWFJUVjBwS1RXMTRkR0pFUW10T2JXaE1ZbFZLVUVscGQybGtSMFp1U1dwdmFWUkhXbkZhVmpneVZXeFZlRk16VGxoaE1taFBWREJTU0ZFeVJuUlZVMG81TGxoMFdpMXpVbkpHYzFaVE1taE9aa1ZSWDJGdGFGRXVaakp5YzI5MmIwWXlWUzFLVm1SeFdDNHhVSHBWWnpoZk9HeEJjblpJTFc1TVRtSXRPVWhKTTNWTFIwUndhMWRCUzNoclFuVkJiMjlGZWtSdVkyeGFSemx6VGs1NWVWbDBNV2M0TTI1RWEyUkxlV05DZG5RdFlVNXlja2RTWm10aGIwVkNORUV4ZEhBd1R6ZEhRbmd3VnpCSlMxRmFkV1prVjJGRmNVWllXRFpyZERCWVNUZDNORmg0YzJ4bFJ5MVZlVU5LZVdaNVFqaG9XSFJVZERWeU0wUTJZbUYzUjFFM2FHcDJha2QzYjBwUExVRm9RMnhEUkVSTFFsVjRMakp5U0RGRVVHdHlkelZwZDJaVVJtOXVOMUpXYTBFPSIsImEiOnsiZm5pZCI6IjAwMDAxMyIsImZ0aWQiOiJ0dG4iLCJmY2lkIjoidHRuLXYyLWV1LTMifX2SAc8DChowMUY1NkZDSzhBSlJZN1BBRDFQV1JKS0NEWRIDAAATGgN0dG4iC3R0bi12Mi1ldS0zKgMAABMyA3R0bjppCgwIuIzbhAYQvbbVlwIaDTQwLjExMy42OC4xOTgiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC14anN6cConcGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoIBCgwIuIzbhAYQ5uTSmAISIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC14anN6cBoQZm9yd2FyZGVyX3VwbGluayIXcm91dGVyLTViNWRjNTRjZjctcHN4bHQqJHBicm91dGVyLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqWAQoMCLiM24QGEJ71ypwCEhdyb3V0ZXItNWI1ZGM1NGNmNy1wc3hsdBohZGVsaXZlci4wMDAwMTNfdHRuX3R0bi1ldTEudXBsaW5rIiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NEIHdHRuLWV1MTKRCAoOCgxwYWNrZXRicm9rZXJFAADIwU0AAMjBXZqZ6UB6nAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTa05WTURsSFZucE9hMDFZWjNSak1tODFXbXRhZDBscGQybGtSMFp1U1dwdmFVOVhiek5rTUd0NVUwUlJlVTFHUm5aamJrSlRWbTVhYzJKR1dYZGFlVW81TGtrek5VeDZjMUpSVXpSRFVVc3pTMU01ZDAxcE4yY3VhVUpCUkRad1FuaG5iRVpxV0dWaVlpNVBZMDA1TVU5WFJHWnlRVmhvU2xsSVFXVjFTR05WUW5nMmEwWktZblJVY1daUVVGVTJWVVF0U0VrM2ExWjBRM2MyY1RkaU9YbzJiVEJtYlc5dWFFUTFZMk56UTNkV09WVXdjbVZJYWw5cFVVMTZhM0pTVkZKUVNYaFNNME5oYVZWZmJXcEljVWxETTB3elEwRjVTbWRpWkdnMVZtNTRXVVZzUmpkdFlrcDJNVzVuYmtvdFdGbGZZeTEwUmxsc1gxUXllVEIyVWpCNFNYVnlPWEJHUlZCMWRHTXRSek5JWW5ndGEzZFdMa2xwUm1kVWVYa3laWFpWV0dkUFdURkNja3R1ZWxFPSIsImEiOnsiZm5pZCI6IjAwMDAxMyIsImZ0aWQiOiJ0dG4iLCJmY2lkIjoidHRuLXYyLWV1LTMifX2SAc8DChowMUY1NkZDSzg4RVFZN0pOR1Y4SjJYQjJXQRIDAAATGgN0dG4iC3R0bi12Mi1ldS0zKgMAABMyA3R0bjppCgwIuIzbhAYQycvblgIaDTQwLjExMy42OC4xOTgiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1mN2g2ayoncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoIBCgwIuIzbhAYQ/6ivmAISIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1mN2g2axoQZm9yd2FyZGVyX3VwbGluayIXcm91dGVyLTViNWRjNTRjZjctbXdmOG0qJHBicm91dGVyLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqWAQoMCLiM24QGEMf1hpsCEhdyb3V0ZXItNWI1ZGM1NGNmNy1td2Y4bRohZGVsaXZlci4wMDAwMTNfdHRuX3R0bi1ldTEudXBsaW5rIiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NEIHdHRuLWV1MTKSCAoOCgxwYWNrZXRicm9rZXJFAACGwk0AAIbCXc3MBEF6nAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTbVpQUldRelVrZDBTazFITVRKa1JrNXhZVzVhVDBscGQybGtSMFp1U1dwdmFXRlhSa1poTWtaWlpFUmFhMWx0Um5oaWFrWktWVVJzUzJKRVZubGFlVW81TGpoWVMxcGpiVEpFV0VkV1FXVjNURTl5TFdrMGVsRXViVEpoVW5sUGJXODJVV2hZZVhacFNpNVVjbEJJVjJaVVZ6aFFTVU4xWmxvM09WVnhWV05EYkhsZlgwMW5RbXhuUWw5Tk1sQmZWalZCTVhWR1FXRlpaMjFVWkZkWFFsSXdSa1ZUUjNWRk1GQnBRVEpFTjFOWFlsTnVRVzlJTlZNdFVHaGpia3RSTFhoVk9VOUVXWFJGUTBvNVJVUmtiM1ZZVkhsaFIyaEJUME5HUXpORFJEZDZlWFpwTkV0MlkwVjFVVVZCVFVKUVJGTm5NRk5rYmxsTVVIcENka0o1WWpWVmJVVnpTek55UkhkNFNFeGFVVjgwVTA0MVFqTmhMakJ4YVc1bFpEZzRVVzFMWlhSU1ZqQlBhVWR3YjNjPSIsImEiOnsiZm5pZCI6IjAwMDAxMyIsImZ0aWQiOiJ0dG4iLCJmY2lkIjoidHRuLXYyLWV1LTQifX2SAdADChowMUY1NkZDSzhON1JFRllHRFg4Q1Y5SzI5RhIDAAATGgN0dG4iC3R0bi12Mi1ldS00KgMAABMyA3R0bjpqCgwIuIzbhAYQ18X/nAIaDjUyLjE2OS4xNTAuMTM4IiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqCAQoMCLiM24QGELO82Z0CEiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoaEGZvcndhcmRlcl91cGxpbmsiF3JvdXRlci01YjVkYzU0Y2Y3LXhoODIyKiRwYnJvdXRlci8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6lgEKDAi4jNuEBhDmrpanAhIXcm91dGVyLTViNWRjNTRjZjcteGg4MjIaIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWY3aDZrKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTEynwgKDgoMcGFja2V0YnJva2VyGgwItozbhAYQ5OqjvgNFAAAcwk0AABzCXQAAyEB6nAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTa0phUlc5NlpWUkNlbE5yTkhkWWVrSjJaRVprV2tscGQybGtSMFp1U1dwdmFVOUhTa05OTVZaTlVraEZkMkpWZUZGTVZtUnJWVVp2ZUZSNlJUUlJVMG81TGtOa1l6bExTalZ3WjNWNk5sWTFaM2xIVlhZNVpIY3VaVmhYV0d0b1JVZHhhRTR6VTBseFF5NVphbVpJUXkxck9WVk9jVkJIUVdOWk9WSnJYMGxNV0ZoSFZXa3RSMmN5U3pkeVVHUkJjR3BYV0V3dFJreGpaVzh5UWpReWIxbFZjRFJmWjJSM2RFOW9OMnhuWVVwMFNra3ROeTEyV0ZNemVDMVNSR1JWWVhSZlNUWklSVWhSVmpWS2NHOWFTeTA1ZG01TWJURnFTRWh6UW1KSGRuaFhWMVZyWkc1Zk1qSXRVakpuVVdSMFZrNXRNVlZPTkRGaWRFOXRNVVl3ZDBWb1JqVjVRVVJmTURWbVoyb3pZVXBmUWt4T2MwaDBMbUpGU1ZGTE0zVlpiakJEVjFjNE0wUkVNVFJTVkdjPSIsImEiOnsiZm5pZCI6IjAwMDAxMyIsImZ0aWQiOiJ0dG4iLCJmY2lkIjoidHRuLXYyLWV1LTEifX2SAc8DChowMUY1NkZDSzg4SjQ5OU5INVlEREc5QVYwSxIDAAATGgN0dG4iC3R0bi12Mi1ldS0xKgMAABMyA3R0bjppCgwIuIzbhAYQ29XYlgIaDTUyLjE2OS43My4yNTEiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoIBCgwIuIzbhAYQw6zQmAISIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqahoQZm9yd2FyZGVyX3VwbGluayIXcm91dGVyLTViNWRjNTRjZjctcHN4bHQqJHBicm91dGVyLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqWAQoMCLiM24QGEJvrt5wCEhdyb3V0ZXItNWI1ZGM1NGNmNy1wc3hsdBohZGVsaXZlci4wMDAwMTNfdHRuX3R0bi1ldTEudXBsaW5rIiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZjdoNmsqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NEIHdHRuLWV1MTodCggKBgjI0AcQBxAFGgM0LzUg4P/TnQMwjLOfiApCDAi4jNuEBhDK//H7AWoFEIDQ0xNyHAoEdXNlchIUCf4qsqJa+kDAEQEAAFChcDZAKAM=")
	protoMessages = append(protoMessages, data)
	data, _ = b64.StdEncoding.DecodeString("CiMKC2NyaWNrZXQtMDAyEg4KDGpwbS1jcmlja2V0czIEJgvqwhIgYXM6dXA6MDFGNTZGRThKMkdOTjI5Vlc1V0gyVlNNNFISImdzOmNvbm46MDFGNFMwS0dWUFgyUUpRRkg2SzhLRTQ3RlISJWdzOnVwOmhvc3Q6MDFGNFMwS0daVFlDNzhOV1c1UTg2R0sxUzcSJGdzOnVwbGluazowMUY1NkZFOEJBNEpUSDBQQVFHSE5LUDZGUBIkbnM6dXBsaW5rOjAxRjU2RkU4QkJWWU5TRzJUVlEyRFhSUkpOEkBycGM6L3R0bi5sb3Jhd2FuLnYzLkdzTnMvSGFuZGxlVXBsaW5rOjAxRjU2RkU4QkJSOFQzMzFSNDU3RkJaRVQyEkBycGM6L3R0bi5sb3Jhd2FuLnYzLk5zQXMvSGFuZGxlVXBsaW5rOjAxRjU2RkU4SjE5VEJFV0tYSlNEQllNUDFLYgsI74zbhAYQwdG4URqwIhiAuQEqADKSAQogChRldWktMDAwMDgwMDI5YzA5ZGQ4NxIIAACAApwJ3Ycgu7KZogpFAACowU0AAKjBXZqZ2UBqFAk1oZv65PdAwBEBAAAt/94yQCgDekAKIgogChRldWktMDAwMDgwMDI5YzA5ZGQ4NxIIAACAApwJ3YcQu7KZogoaDAjujNuEBhCwwZPHAyD4nLKWm5VniAECMo4ICg4KDHBhY2tldGJyb2tlckUAANjBTQAA2MFdzcwcQXqcBHsiZyI6IlpYbEthR0pIWTJsUGFVcENUVlJKTkZJd1RrNVRNV05wVEVOS2JHSnRUV2xQYVVwQ1RWUkpORkl3VGs1SmFYZHBZVmhaYVU5cFNuaFVhM0JZVkZSYVRrMHdNVXROYkZwelRrUk5NVWxwZDJsa1IwWnVTV3B2YVU5RVVtNU5SV3hQVEZZNVlXUkhUbTVQUjBwelRWUldVVkpWZUhwUlUwbzVMbVJKY0hCbVJrRllUVVJVUjFaMWIwVlBNSGh2YlhjdVpGaDFXR3BWVWpWc1YxRjVlRXRKYmk1bk9IUmpORGxwYjI5b1NtbFZVakZCTkZWMllrZFZjekUxVjI1WmVWaEdSRXBKVTNCSVFrODVlWHBpWjNwTlFscEpSbmREUVhsRlFuTnRVbVJWWjFsNVZFSnVjWEJDYTNCYVJIRjZXRlZrUTNVMGRYcE1NM05yV1ZOYVpGVTJYMWRVUkRaQmFWbHVORzFoZDFoaGQwNW1ZVlpDUmpOdlMzSnNSMVZ3UTBJM1FrOVRhRGRGZFdob2QxbHVORXd6ZFhwRGRYaFZSVkJtY0VKek9HZE1ORE5OYWxGRVlVazBVbmRrVWtoVUxtaElPRUZSY0U1dFNXc3lTemwxWVVWRVFqWTJWM2M9IiwiYSI6eyJmbmlkIjoiMDAwMDEzIiwiZnRpZCI6InR0biIsImZjaWQiOiJ0dG4tdjItZXUtMyJ9fZIBzAMKGjAxRjU2RkU4Q1cxS0JDN1IxUVpFUUFWMEZEEgMAABMaA3R0biILdHRuLXYyLWV1LTMqAwAAEzIDdHRuOmgKCwjvjNuEBhDAnPUBGg00MC4xMTMuNjguMTk4IiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQteGpzenAqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqBAQoLCO+M24QGEOmFlAISIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC14anN6cBoQZm9yd2FyZGVyX3VwbGluayIXcm91dGVyLTViNWRjNTRjZjctcHN4bHQqJHBicm91dGVyLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqVAQoLCO+M24QGEIL51AISF3JvdXRlci01YjVkYzU0Y2Y3LXBzeGx0GiFkZWxpdmVyLjAwMDAxM190dG5fdHRuLWV1MS51cGxpbmsiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0Qgd0dG4tZXUxMo8ICg4KDHBhY2tldGJyb2tlckUAAIrCTQAAisJdmpnZQHqcBHsiZyI6IlpYbEthR0pIWTJsUGFVcENUVlJKTkZJd1RrNVRNV05wVEVOS2JHSnRUV2xQYVVwQ1RWUkpORkl3VGs1SmFYZHBZVmhaYVU5cFNuUlRXR2gwVjFjMVNrMHdaSEJpV0ZFd1pGUnJORWxwZDJsa1IwWnVTV3B2YVdGWVRrcGxWMmhTWXpOQ1dtUXhVbGhoTWpGUFdtMDFVV0p0TlRWYWVVbzVMa2RTTTFoS1EyTlZTelpzTm14SmJrNWZiMWx5ZEhjdWVHWTVkVnBtTVZWblZ5MTNNMjl3Umk1SFoyNU9OR055UVRVM1RIVlpYMjVMWVRWcFVHd3hSMjVOZFdwVE5GWndTVE5uTUdsNWNqSk9TRzA0TlVKNFEyTXhPSGs0VVdjdGFrVnVPWFpsUXpsVU1GZFZRVTVTUVhwbFFsUllkbE5wYWtGbWVWWmFUR2xNU1dkVllqaHBTa3h3ZGxwd1ZGVXdSVVpKTUZONU1WZEZTMHRUTVcwNGR6VjJXRU5ZV1ZwWWVYVkphV3Q2Ykc1ak0zWkJha0UyYmxKT2VXSlpaVFpJYm5CSmNYSk1OMHBpYjBNeWFHZFJTVmMxUldObkxrWk5RWGwyY0ZOMVZtcDFVWGc0WjFacVVrWkdNRkU9IiwiYSI6eyJmbmlkIjoiMDAwMDEzIiwiZnRpZCI6InR0biIsImZjaWQiOiJ0dG4tdjItZXUtNCJ9fZIBzQMKGjAxRjU2RkU4Q1o2WUU3U1Q0SkJKV0I0TVdEEgMAABMaA3R0biILdHRuLXYyLWV1LTQqAwAAEzIDdHRuOmkKCwjvjNuEBhDxoLgDGg41Mi4xNjkuMTUwLjEzOCIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6gQEKCwjvjNuEBhCpgr4EEiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoaEGZvcndhcmRlcl91cGxpbmsiF3JvdXRlci01YjVkYzU0Y2Y3LXBzeGx0KiRwYnJvdXRlci8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6lQEKCwjvjNuEBhCxjoUFEhdyb3V0ZXItNWI1ZGM1NGNmNy1wc3hsdBohZGVsaXZlci4wMDAwMTNfdHRuX3R0bi1ldTEudXBsaW5rIiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZjdoNmsqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NEIHdHRuLWV1MTKGCAoOCgxwYWNrZXRicm9rZXJFAAAUwk0AABTCXQAAAEF6lAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTbkZOVjJSRlRraFpkMk5HV2xOaGVsVXdWMGhrY0VscGQybGtSMFp1U1dwdmFXUnBNVk5TYmtKWVkydG9WazFHWkZGTlNFb3lZVzFPZVdWR1pFbGFlVW81TG10S1dIQkxPVzlSZUVzMGMwbExUekZmZFhGRlpuY3ViRFI2TmxkNlkwWTVOSFl3UWxCNFZDNVVja1prUm5WMk9IUjFNVkIzUTFKdGIzTjBObmczVW5saFJYQjVZVkozVlZoc1lsRnFNSEU1WVRodmQxRTBiSGszTFdsRU1WUnJXamhLUTFJd1ZYUm5TM3A2WXpKVFoyMVlSbGw0VkdaUk4wMU9OMHRGY3pnMVNrdHVUbXRNU0hSQ09WcFBTV2h1VVRSWlVtUkNSVnBVU1VsVGNreDVRVWRPZVU1eWFXZEJRMmhmVUMxQmQxRkdUMjFNTVcxdVdFWnRhRXhUVlRoVlRqVnhURW81VFdGcmNXRXpYM040T0M1QmNHbHlURFo1TlhWSldHWkdaVGhHYm5oYWJIbFIiLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19kgHMAwoaMDFGNTZGRThEOTMzRFRQNUg5WDk2TTdFN0sSAwAAExoDdHRuIgt0dG4tdjItZXUtMyoDAAATMgN0dG46aAoLCO+M24QGEMu9kAgaDTQwLjExMy42OC4xOTgiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoEBCgsI74zbhAYQ/fz0DBIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqGhBmb3J3YXJkZXJfdXBsaW5rIhdyb3V0ZXItNWI1ZGM1NGNmNy14aDgyMiokcGJyb3V0ZXIvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OpUBCgsI74zbhAYQ4bHHEBIXcm91dGVyLTViNWRjNTRjZjcteGg4MjIaIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTEylAgKDgoMcGFja2V0YnJva2VyGgwI7YzbhAYQ9e+6rgFFAAD4wU0AAPjBXQAA4EB6lAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTVFJhYlU1Q1lXc3dlR0ZHY0ZGamVscFlZMGRvTTBscGQybGtSMFp1U1dwdmFWTjZXVFJqUjJ3elRucE9hRkZYTVVoaVJUbFdUMVZ6TlZGcWFHMVJVMG81TGxGRVNISlhXbDk2T1Rad1VqSk9NbkZ2WkRaak5tY3VOVmcxZDE5R1ZFaFNTMVZrV1d0VWNpNWlVa0pMZHpWRE1GSnplWE5TVURGdVJUSldiMlZNZW5KT1JXeEdVMjkyY1ZCTWJURlFTWFJNWTJJMmFtOXZVRlZJTjFWNlQyVldNRk5zVlZsaFl6RndVRTkzYm1oNFJXTmlhMDFRUTBkV1dESmZUamRtU1ZKTmFqWTVjSGhKWjA1WVdqTnpOMnhwYUY5a1pVRmpTa3RsYWxKMU1sWkNXbVJCT0hGd1RGQk5hbkY2Y0RkWFRYcENXa3RYY2xSZk4xaFlTekUwVlRKNVNqWnFkVFIwVG5OVlpIVk1PRlZwWXk1elRISm5iQzFJUlRWc1RFRlpOblF6UVhoc2NFTlIiLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0xIn19kgHMAwoaMDFGNTZGRThEOVZQNk04SldIN0JERzVCSkQSAwAAExoDdHRuIgt0dG4tdjItZXUtMSoDAAATMgN0dG46aAoLCO+M24QGELGqqggaDTUyLjE2OS43My4yNTEiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoEBCgsI74zbhAYQ1dG5DBIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqGhBmb3J3YXJkZXJfdXBsaW5rIhdyb3V0ZXItNWI1ZGM1NGNmNy1td2Y4bSokcGJyb3V0ZXIvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OpUBCgsI74zbhAYQp97DEBIXcm91dGVyLTViNWRjNTRjZjctbXdmOG0aIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTE6HQoICgYIyNAHEAcQBRoDNC81IKCEkZ4DMLuymaIKQgwI7ozbhAYQgZ3TxwNqBRCA0NMTchwKBHVzZXISFAn+KrKiWvpAwBEBAABQoXA2QCgD")
	protoMessages = append(protoMessages, data)

	for _, message := range protoMessages {
		var packetIn ttnpb.ApplicationUp
		if err := proto.Unmarshal(message, &packetIn); err != nil {
			t.Error(err.Error())
		}

		if packetIn.GetUplinkMessage().GetNetworkIds() == nil {
			t.Error("Network IDs not set")
			continue
		}

		//log.Printf("Uplink Message: %+v", packetIn)

		var packetOut types.IotMessage
		AddNetworkMetadataFields(packetIn, &packetOut)
		log.Println(utils.PrettyPrint(packetOut))
	}
}
