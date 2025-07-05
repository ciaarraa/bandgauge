package cmd
type IPerf3Result struct {
	Start struct {
		Connected []struct {
			Socket     float64    `json:"socket"`
			LocalHost  string `json:"local_host"`
			LocalPort  float64    `json:"local_port"`
			RemoteHost string `json:"remote_host"`
			RemotePort float64    `json:"remote_port"`
		} `json:"connected"`
		Version    string `json:"version"`
		SystemInfo string `json:"system_info"`
		Timestamp  struct {
			Time     string `json:"time"`
			Timesecs float64    `json:"timesecs"`
		} `json:"timestamp"`
		ConnectingTo struct {
			Host string `json:"host"`
			Port float64    `json:"port"`
		} `json:"connecting_to"`
		Cookie        string `json:"cookie"`
		TCPMssDefault float64    `json:"tcp_mss_default"`
		TargetBitrate float64    `json:"target_bitrate"`
		FqRate        float64    `json:"fq_rate"`
		SockBufsize   float64    `json:"sock_bufsize"`
		SndbufActual  float64    `json:"sndbuf_actual"`
		RcvbufActual  float64    `json:"rcvbuf_actual"`
		TestStart     struct {
			Protocol      string `json:"protocol"`
			NumStreams    float64    `json:"num_streams"`
			Blksize       float64    `json:"blksize"`
			Omit          float64    `json:"omit"`
			Duration      float64    `json:"duration"`
			Bytes         float64    `json:"bytes"`
			Blocks        float64    `json:"blocks"`
			Reverse       float64    `json:"reverse"`
			Tos           float64    `json:"tos"`
			TargetBitrate float64    `json:"target_bitrate"`
			Bidir         float64    `json:"bidir"`
			Fqrate        float64    `json:"fqrate"`
			float64erval      float64    `json:"float64erval"`
		} `json:"test_start"`
	} `json:"start"`
	float64ervals []struct {
		Streams []struct {
			Socket        float64     `json:"socket"`
			Start         float64     `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         float64     `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			Omitted       bool    `json:"omitted"`
			Sender        bool    `json:"sender"`
		} `json:"streams"`
		Sum struct {
			Start         float64     `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         float64     `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			Omitted       bool    `json:"omitted"`
			Sender        bool    `json:"sender"`
		} `json:"sum"`
	} `json:"float64ervals"`
	End struct {
		Streams []struct {
			Sender struct {
				Socket        float64     `json:"socket"`
				Start         float64     `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         float64     `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
				Sender        bool    `json:"sender"`
			} `json:"sender"`
			Receiver struct {
				Socket        float64     `json:"socket"`
				Start         float64     `json:"start"`
				End           float64 `json:"end"`
				Seconds       float64 `json:"seconds"`
				Bytes         float64     `json:"bytes"`
				BitsPerSecond float64 `json:"bits_per_second"`
				Sender        bool    `json:"sender"`
			} `json:"receiver"`
		} `json:"streams"`
		SumSent struct {
			Start         float64     `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         float64     `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			Sender        bool    `json:"sender"`
		} `json:"sum_sent"`
		SumReceived struct {
			Start         float64     `json:"start"`
			End           float64 `json:"end"`
			Seconds       float64 `json:"seconds"`
			Bytes         float64     `json:"bytes"`
			BitsPerSecond float64 `json:"bits_per_second"`
			Sender        bool    `json:"sender"`
		} `json:"sum_received"`
		CPUUtilizationPercent struct {
			HostTotal    float64 `json:"host_total"`
			HostUser     float64 `json:"host_user"`
			HostSystem   float64 `json:"host_system"`
			RemoteTotal  float64 `json:"remote_total"`
			RemoteUser   float64 `json:"remote_user"`
			RemoteSystem float64 `json:"remote_system"`
		} `json:"cpu_utilization_percent"`
		ReceiverTCPCongestion string `json:"receiver_tcp_congestion"`
	} `json:"end"`
}