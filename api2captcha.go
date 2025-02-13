		Invisible  bool
		Enterprise bool
		Version    string
		Action     string
		DataS      string
		Score      float64
	}

	Rotate struct {
		Base64          string
		File            string
		Files           []string
		Angle           int
		Lang            string
		HintText        string
		HintImageBase64 string
		HintImageFile   string
	}

	Text struct {
		Text string
		Lang string
	}

	AmazonWAF struct {
		Iv              string
		SiteKey         string
		Url             string
		Context         string
		ChallengeScript string
		CaptchaScript   string
	}

	GeeTestV4 struct {
		CaptchaId string
		Url       string
		ApiServer string
		Challenge string
	}

	Lemin struct {
		CaptchaId string
		DivId     string
		Url       string
		ApiServer string
	}

	CloudflareTurnstile struct {
		SiteKey   string
		Url       string
		Data      string
		PageData  string
		Action    string
		UserAgent string
	}

	CyberSiARA struct {
		MasterUrlId string
		Url         string
		UserAgent   string
	}

	DataDome struct {
		Url        string
		CaptchaUrl string
		Proxytype  string
		Proxy      string
		UserAgent  string