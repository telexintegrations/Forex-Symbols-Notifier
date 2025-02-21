package entity

// Date structure
type Date struct {
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// Descriptions structure
type Descriptions struct {
  AppDescription string `json:"app_description"`
	AppLogo        string `json:"app_logo"`
	AppName        string `json:"app_name"`
	AppURL         string `json:"app_url"`
	BackgroundColor string `json:"background_color"`
}

// MonitoringUser structure
type MonitoringUser struct {
	AlwaysOnline bool   `json:"always_online"`
	DisplayName  string `json:"display_name"`
}

// Permissions structure
type Permissions struct {
	MonitoringUser MonitoringUser `json:"monitoring_user"`
}

// Setting structure
type Setting struct {
	Label    string `json:"label"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Default  string `json:"default"`
}

type Data struct {
	Date                Date         `json:"date"`
	Descriptions        Descriptions `json:"descriptions"`
  IntegrationCategory string       `json:"integration_category"`
	IntegrationType     string       `json:"integration_type"`
  IsActive            bool         `json:"is_active"`
	KeyFeatures         []string     `json:"key_features"`
	Author              string       `json:"author"`
	Settings            []Setting    `json:"settings"`
  TickURL             string       `json:"tick_url"`
	TargetURL           string       `json:"target_url"`
}

// Integrationson main structure
type Integration struct {
	Data Data `json:"data"`
}

var Integrationson = Integration{
  Data: Data{
    Date: Date{
      CreatedAt: "2025-02-20",
      UpdatedAt: "2025-02-20",
    },
    Descriptions: Descriptions{
      AppDescription: "This is a notification app to give recurring updates on the price of popular Forex symbols",
      AppLogo:        "https://my-portfolio-343207.web.app/MyLogo4.png",
      AppName:        "ForexPI",
      AppURL:         "https://fun-numbers.onrender.com/integration.json",
      BackgroundColor: "#fff",
    },
    IntegrationCategory: "Monitoring & Logging",
    IntegrationType:     "interval",
    IsActive:            true,
    KeyFeatures:         []string{"Forex", "Updates"},
    Author:              "Samuel Ikoli",
    Settings: []Setting{
      {
        Label:    "interval",
        Type:     "text",
        Required: true,
        Default:  "* * * * *",
      },
    },
    TickURL:   "https://fun-numbers.onrender.com/tick",
    TargetURL: "https://ping.telex.im/v1/webhooks/01950b90-b1bf-75b7-b9e6-e831fdd18b5f",
  },
}

  type MonitorPayload struct {
    ChannelID string        `json:"channel_id,omitempty"`
    ReturnURL string        `json:"return_url,omitempty"`
    Settings  []interface{} `json:"settings,omitempty"`
  }
  