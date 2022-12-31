package tts

import (
	"bytes"
	b64 "encoding/base64"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wcrwg-iot-ingress/pkg/types"
	"wcrwg-iot-ingress/pkg/utils"
)

func TestHandlerJson(t *testing.T) {
	var publishChannel = make(chan types.IotMessage, 1)

	go PrintChannelContent(publishChannel)

	context := &Context{PublishChannel: publishChannel}

	postbodies := []string{
		`{"@type":"type.googleapis.com/ttn.lorawan.v3.ApplicationUp","end_device_ids":{"device_id":"eui-3030303130373135","application_ids":{"application_id":"wcrwg-otto-io"},"dev_eui":"3030303130373135","join_eui":"70B3D57ED003E052","dev_addr":"260B4ED9"},"correlation_ids":["as:up:01GNKHRNGZN3NHKCQQHB3PSMJB","gs:conn:01GM50FMQSH6K9JY5XNQYBK06F","gs:up:host:01GM50FMR17K2MK3VTBTGNG8GW","gs:uplink:01GNKHRNADMYBCV2ZHH8BZ22VJ","ns:uplink:01GNKHRNAEVVTTZ9VDNT30X5HB","rpc:/ttn.lorawan.v3.GsNs/HandleUplink:01GNKHRNAEDV4H1XCJ4APGZ9HZ","rpc:/ttn.lorawan.v3.NsAs/HandleUplink:01GNKHRNGYGCH4CHMPM6JKCK3B"],"received_at":"2022-12-31T07:36:28.446971766Z","uplink_message":{"session_key_id":"AYVj87hTxU/cxOe2HalTpQ==","f_port":1,"f_cnt":116,"frm_payload":"CT8DIAMA6OavY6sS","decoded_payload":{"analog_2":831,"analog_4":800,"battery":4.779,"digital_out_1":0,"message_type":"periodic","timestamp":1672472296},"rx_metadata":[{"gateway_ids":{"gateway_id":"lnx-solutions-klnmnd-01","eui":"3135323533002400"},"time":"2022-12-31T07:36:28.121842Z","timestamp":192353316,"rssi":-115,"channel_rssi":-115,"snr":2.5,"location":{"latitude":-34.339333,"longitude":19.034111,"altitude":20,"source":"SOURCE_REGISTRY"},"uplink_token":"CiUKIwoXbG54LXNvbHV0aW9ucy1rbG5tbmQtMDESCDE1MjUzACQAEKSo3FsaCwj8zL+dBhDoqKZxIKDZmsnMmvAC","channel_index":7,"received_at":"2022-12-31T07:36:28.135334954Z"},{"gateway_ids":{"gateway_id":"lnx-solutions-betbay-01","eui":"343632385C003700"},"time":"2022-12-31T07:36:28.123778Z","timestamp":2755043420,"rssi":-115,"channel_rssi":-115,"snr":-5,"location":{"latitude":-34.35025,"longitude":18.961417,"altitude":465,"source":"SOURCE_REGISTRY"},"uplink_token":"CiUKIwoXbG54LXNvbHV0aW9ucy1iZXRiYXktMDESCDQ2MjhcADcAENzA2qEKGgsI/My/nQYQmLWzcyDgzomrl4nlAg==","channel_index":7,"received_at":"2022-12-31T07:36:28.242014872Z"},{"gateway_ids":{"gateway_id":"lnx-solutions-jskop-01","eui":"3133303746002A00"},"time":"2022-12-31T07:36:28.127535Z","timestamp":3887686692,"rssi":-115,"channel_rssi":-115,"snr":-4.75,"location":{"latitude":-33.97178,"longitude":19.50633,"altitude":1700,"source":"SOURCE_REGISTRY"},"uplink_token":"CiQKIgoWbG54LXNvbHV0aW9ucy1qc2tvcC0wMRIIMTMwN0YAKgAQpMjlvQ4aCwj8zL+dBhDX0Nl3IKDZvOGSjiE=","channel_index":7,"received_at":"2022-12-31T07:36:28.157374949Z"},{"gateway_ids":{"gateway_id":"lnx-solutions-pbp-01","eui":"3436323825004700"},"time":"1971-11-08T16:59:28.965485Z","timestamp":409823220,"rssi":-114,"channel_rssi":-114,"snr":-3.25,"location":{"latitude":-34.35208487,"longitude":18.82273625,"altitude":130,"source":"SOURCE_REGISTRY"},"uplink_token":"CiIKIAoUbG54LXNvbHV0aW9ucy1wYnAtMDESCDQ2MjglAEcAEPTPtcMBGgsI/My/nQYQr9n5fyCgovja9sJB","channel_index":7,"received_at":"2022-12-31T07:36:28.173520830Z"}],"settings":{"data_rate":{"lora":{"bandwidth":125000,"spreading_factor":7,"coding_rate":"4/5"}},"frequency":"867900000","timestamp":192353316,"time":"2022-12-31T07:36:28.121842Z"},"received_at":"2022-12-31T07:36:28.238415358Z","confirmed":true,"consumed_airtime":"0.061696s","network_ids":{"net_id":"000013","tenant_id":"ttn","cluster_id":"eu1","cluster_address":"eu1.cloud.thethings.network"}}}`,
	}

	for _, postbody := range postbodies {
		// Create a request to pass to our handler.
		req, err := http.NewRequest("POST", "", strings.NewReader(postbody))
		if err != nil {
			t.Fatal(err)
		}

		// Set request headers
		req.Header.Set("TTNMAPPERORG-USER", "test@ttnmapper.org")
		req.Header.Set("TTNMAPPERORG-EXPERIMENT", "test-experiment")
		req.Header.Set("COntent-Type", "application/json")
		req.Header.Set("X-Tts-Domain", "test.cloud.thethings.network")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(context.PostV3Uplink)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		log.Println(rr.Body.String())

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusAccepted {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusAccepted)
		}

		// Check the response body is what we expect.
		expected := `{"message":"New packet accepted into queue","success":true}`
		if strings.TrimSpace(rr.Body.String()) != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}

	}
}

func PrintChannelContent(publishChannel chan types.IotMessage) {
	for {
		// Check if a packet was written to the queue
		select {
		case packetOut, ok := <-publishChannel:
			if ok {
				log.Println(utils.PrettyPrint(packetOut))
			} else {
				log.Print("Channel closed!")
			}
			//default:
			//	log.Println("No value read, moving on.")
		}
	}
}

func TestHandlerProtobuf(t *testing.T) {
	var publishChannel = make(chan types.IotMessage, 1)
	context := &Context{PublishChannel: publishChannel}

	// Create a request to pass to our handler.
	data, _ := b64.StdEncoding.DecodeString("CiMKC2NyaWNrZXQtMDAyEg4KDGpwbS1jcmlja2V0czIEJgvqwhIgYXM6dXA6MDFGNTZGRThKMkdOTjI5Vlc1V0gyVlNNNFISImdzOmNvbm46MDFGNFMwS0dWUFgyUUpRRkg2SzhLRTQ3RlISJWdzOnVwOmhvc3Q6MDFGNFMwS0daVFlDNzhOV1c1UTg2R0sxUzcSJGdzOnVwbGluazowMUY1NkZFOEJBNEpUSDBQQVFHSE5LUDZGUBIkbnM6dXBsaW5rOjAxRjU2RkU4QkJWWU5TRzJUVlEyRFhSUkpOEkBycGM6L3R0bi5sb3Jhd2FuLnYzLkdzTnMvSGFuZGxlVXBsaW5rOjAxRjU2RkU4QkJSOFQzMzFSNDU3RkJaRVQyEkBycGM6L3R0bi5sb3Jhd2FuLnYzLk5zQXMvSGFuZGxlVXBsaW5rOjAxRjU2RkU4SjE5VEJFV0tYSlNEQllNUDFLYgsI74zbhAYQwdG4URqwIhiAuQEqADKSAQogChRldWktMDAwMDgwMDI5YzA5ZGQ4NxIIAACAApwJ3Ycgu7KZogpFAACowU0AAKjBXZqZ2UBqFAk1oZv65PdAwBEBAAAt/94yQCgDekAKIgogChRldWktMDAwMDgwMDI5YzA5ZGQ4NxIIAACAApwJ3YcQu7KZogoaDAjujNuEBhCwwZPHAyD4nLKWm5VniAECMo4ICg4KDHBhY2tldGJyb2tlckUAANjBTQAA2MFdzcwcQXqcBHsiZyI6IlpYbEthR0pIWTJsUGFVcENUVlJKTkZJd1RrNVRNV05wVEVOS2JHSnRUV2xQYVVwQ1RWUkpORkl3VGs1SmFYZHBZVmhaYVU5cFNuaFVhM0JZVkZSYVRrMHdNVXROYkZwelRrUk5NVWxwZDJsa1IwWnVTV3B2YVU5RVVtNU5SV3hQVEZZNVlXUkhUbTVQUjBwelRWUldVVkpWZUhwUlUwbzVMbVJKY0hCbVJrRllUVVJVUjFaMWIwVlBNSGh2YlhjdVpGaDFXR3BWVWpWc1YxRjVlRXRKYmk1bk9IUmpORGxwYjI5b1NtbFZVakZCTkZWMllrZFZjekUxVjI1WmVWaEdSRXBKVTNCSVFrODVlWHBpWjNwTlFscEpSbmREUVhsRlFuTnRVbVJWWjFsNVZFSnVjWEJDYTNCYVJIRjZXRlZrUTNVMGRYcE1NM05yV1ZOYVpGVTJYMWRVUkRaQmFWbHVORzFoZDFoaGQwNW1ZVlpDUmpOdlMzSnNSMVZ3UTBJM1FrOVRhRGRGZFdob2QxbHVORXd6ZFhwRGRYaFZSVkJtY0VKek9HZE1ORE5OYWxGRVlVazBVbmRrVWtoVUxtaElPRUZSY0U1dFNXc3lTemwxWVVWRVFqWTJWM2M9IiwiYSI6eyJmbmlkIjoiMDAwMDEzIiwiZnRpZCI6InR0biIsImZjaWQiOiJ0dG4tdjItZXUtMyJ9fZIBzAMKGjAxRjU2RkU4Q1cxS0JDN1IxUVpFUUFWMEZEEgMAABMaA3R0biILdHRuLXYyLWV1LTMqAwAAEzIDdHRuOmgKCwjvjNuEBhDAnPUBGg00MC4xMTMuNjguMTk4IiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQteGpzenAqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqBAQoLCO+M24QGEOmFlAISIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC14anN6cBoQZm9yd2FyZGVyX3VwbGluayIXcm91dGVyLTViNWRjNTRjZjctcHN4bHQqJHBicm91dGVyLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NDqVAQoLCO+M24QGEIL51AISF3JvdXRlci01YjVkYzU0Y2Y3LXBzeGx0GiFkZWxpdmVyLjAwMDAxM190dG5fdHRuLWV1MS51cGxpbmsiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0Qgd0dG4tZXUxMo8ICg4KDHBhY2tldGJyb2tlckUAAIrCTQAAisJdmpnZQHqcBHsiZyI6IlpYbEthR0pIWTJsUGFVcENUVlJKTkZJd1RrNVRNV05wVEVOS2JHSnRUV2xQYVVwQ1RWUkpORkl3VGs1SmFYZHBZVmhaYVU5cFNuUlRXR2gwVjFjMVNrMHdaSEJpV0ZFd1pGUnJORWxwZDJsa1IwWnVTV3B2YVdGWVRrcGxWMmhTWXpOQ1dtUXhVbGhoTWpGUFdtMDFVV0p0TlRWYWVVbzVMa2RTTTFoS1EyTlZTelpzTm14SmJrNWZiMWx5ZEhjdWVHWTVkVnBtTVZWblZ5MTNNMjl3Umk1SFoyNU9OR055UVRVM1RIVlpYMjVMWVRWcFVHd3hSMjVOZFdwVE5GWndTVE5uTUdsNWNqSk9TRzA0TlVKNFEyTXhPSGs0VVdjdGFrVnVPWFpsUXpsVU1GZFZRVTVTUVhwbFFsUllkbE5wYWtGbWVWWmFUR2xNU1dkVllqaHBTa3h3ZGxwd1ZGVXdSVVpKTUZONU1WZEZTMHRUTVcwNGR6VjJXRU5ZV1ZwWWVYVkphV3Q2Ykc1ak0zWkJha0UyYmxKT2VXSlpaVFpJYm5CSmNYSk1OMHBpYjBNeWFHZFJTVmMxUldObkxrWk5RWGwyY0ZOMVZtcDFVWGc0WjFacVVrWkdNRkU9IiwiYSI6eyJmbmlkIjoiMDAwMDEzIiwiZnRpZCI6InR0biIsImZjaWQiOiJ0dG4tdjItZXUtNCJ9fZIBzQMKGjAxRjU2RkU4Q1o2WUU3U1Q0SkJKV0I0TVdEEgMAABMaA3R0biILdHRuLXYyLWV1LTQqAwAAEzIDdHRuOmkKCwjvjNuEBhDxoLgDGg41Mi4xNjkuMTUwLjEzOCIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6gQEKCwjvjNuEBhCpgr4EEiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZHNyamoaEGZvcndhcmRlcl91cGxpbmsiF3JvdXRlci01YjVkYzU0Y2Y3LXBzeGx0KiRwYnJvdXRlci8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjQ6lQEKCwjvjNuEBhCxjoUFEhdyb3V0ZXItNWI1ZGM1NGNmNy1wc3hsdBohZGVsaXZlci4wMDAwMTNfdHRuX3R0bi1ldTEudXBsaW5rIiFyb3V0ZXItZGF0YXBsYW5lLTU3ZDlkOWJkZGQtZjdoNmsqJ3BiZGF0YXBsYW5lLzEuNS4yIGdvLzEuMTYuMiBsaW51eC9hbWQ2NEIHdHRuLWV1MTKGCAoOCgxwYWNrZXRicm9rZXJFAAAUwk0AABTCXQAAAEF6lAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTbkZOVjJSRlRraFpkMk5HV2xOaGVsVXdWMGhrY0VscGQybGtSMFp1U1dwdmFXUnBNVk5TYmtKWVkydG9WazFHWkZGTlNFb3lZVzFPZVdWR1pFbGFlVW81TG10S1dIQkxPVzlSZUVzMGMwbExUekZmZFhGRlpuY3ViRFI2TmxkNlkwWTVOSFl3UWxCNFZDNVVja1prUm5WMk9IUjFNVkIzUTFKdGIzTjBObmczVW5saFJYQjVZVkozVlZoc1lsRnFNSEU1WVRodmQxRTBiSGszTFdsRU1WUnJXamhLUTFJd1ZYUm5TM3A2WXpKVFoyMVlSbGw0VkdaUk4wMU9OMHRGY3pnMVNrdHVUbXRNU0hSQ09WcFBTV2h1VVRSWlVtUkNSVnBVU1VsVGNreDVRVWRPZVU1eWFXZEJRMmhmVUMxQmQxRkdUMjFNTVcxdVdFWnRhRXhUVlRoVlRqVnhURW81VFdGcmNXRXpYM040T0M1QmNHbHlURFo1TlhWSldHWkdaVGhHYm5oYWJIbFIiLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0zIn19kgHMAwoaMDFGNTZGRThEOTMzRFRQNUg5WDk2TTdFN0sSAwAAExoDdHRuIgt0dG4tdjItZXUtMyoDAAATMgN0dG46aAoLCO+M24QGEMu9kAgaDTQwLjExMy42OC4xOTgiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoEBCgsI74zbhAYQ/fz0DBIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqGhBmb3J3YXJkZXJfdXBsaW5rIhdyb3V0ZXItNWI1ZGM1NGNmNy14aDgyMiokcGJyb3V0ZXIvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OpUBCgsI74zbhAYQ4bHHEBIXcm91dGVyLTViNWRjNTRjZjcteGg4MjIaIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTEylAgKDgoMcGFja2V0YnJva2VyGgwI7YzbhAYQ9e+6rgFFAAD4wU0AAPjBXQAA4EB6lAR7ImciOiJaWGxLYUdKSFkybFBhVXBDVFZSSk5GSXdUazVUTVdOcFRFTktiR0p0VFdsUGFVcENUVlJKTkZJd1RrNUphWGRwWVZoWmFVOXBTVFJhYlU1Q1lXc3dlR0ZHY0ZGamVscFlZMGRvTTBscGQybGtSMFp1U1dwdmFWTjZXVFJqUjJ3elRucE9hRkZYTVVoaVJUbFdUMVZ6TlZGcWFHMVJVMG81TGxGRVNISlhXbDk2T1Rad1VqSk9NbkZ2WkRaak5tY3VOVmcxZDE5R1ZFaFNTMVZrV1d0VWNpNWlVa0pMZHpWRE1GSnplWE5TVURGdVJUSldiMlZNZW5KT1JXeEdVMjkyY1ZCTWJURlFTWFJNWTJJMmFtOXZVRlZJTjFWNlQyVldNRk5zVlZsaFl6RndVRTkzYm1oNFJXTmlhMDFRUTBkV1dESmZUamRtU1ZKTmFqWTVjSGhKWjA1WVdqTnpOMnhwYUY5a1pVRmpTa3RsYWxKMU1sWkNXbVJCT0hGd1RGQk5hbkY2Y0RkWFRYcENXa3RYY2xSZk4xaFlTekUwVlRKNVNqWnFkVFIwVG5OVlpIVk1PRlZwWXk1elRISm5iQzFJUlRWc1RFRlpOblF6UVhoc2NFTlIiLCJhIjp7ImZuaWQiOiIwMDAwMTMiLCJmdGlkIjoidHRuIiwiZmNpZCI6InR0bi12Mi1ldS0xIn19kgHMAwoaMDFGNTZGRThEOVZQNk04SldIN0JERzVCSkQSAwAAExoDdHRuIgt0dG4tdjItZXUtMSoDAAATMgN0dG46aAoLCO+M24QGELGqqggaDTUyLjE2OS43My4yNTEiIXJvdXRlci1kYXRhcGxhbmUtNTdkOWQ5YmRkZC1kc3JqaioncGJkYXRhcGxhbmUvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OoEBCgsI74zbhAYQ1dG5DBIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqGhBmb3J3YXJkZXJfdXBsaW5rIhdyb3V0ZXItNWI1ZGM1NGNmNy1td2Y4bSokcGJyb3V0ZXIvMS41LjIgZ28vMS4xNi4yIGxpbnV4L2FtZDY0OpUBCgsI74zbhAYQp97DEBIXcm91dGVyLTViNWRjNTRjZjctbXdmOG0aIWRlbGl2ZXIuMDAwMDEzX3R0bl90dG4tZXUxLnVwbGluayIhcm91dGVyLWRhdGFwbGFuZS01N2Q5ZDliZGRkLWRzcmpqKidwYmRhdGFwbGFuZS8xLjUuMiBnby8xLjE2LjIgbGludXgvYW1kNjRCB3R0bi1ldTE6HQoICgYIyNAHEAcQBRoDNC81IKCEkZ4DMLuymaIKQgwI7ozbhAYQgZ3TxwNqBRCA0NMTchwKBHVzZXISFAn+KrKiWvpAwBEBAABQoXA2QCgD")
	req, err := http.NewRequest("POST", "", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}

	// Set request headers
	req.Header.Set("TTNMAPPERORG-USER", "test@ttnmapper.org")
	req.Header.Set("Content-Type", "application/octet-stream") // TTS uses octet-stream for protobufs
	req.Header.Set("X-Tts-Domain", "test.cloud.thethings.network")

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(context.PostV3Uplink)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	log.Println(rr.Body.String())

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

	// Check the response body is what we expect.
	expected := `{"message":"New packet accepted into queue","success":true}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}

	// Check if a packet was written to the queue
	select {
	case packetOut, ok := <-publishChannel:
		if ok {
			log.Println(utils.PrettyPrint(packetOut))
		} else {
			t.Error("Channel closed!")
		}
	default:
		t.Error("No value ready, moving on.")
	}
}
