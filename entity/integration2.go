package entity

var Integrationson2 = Integration{
  Data: Data{
    Date: Date{
      CreatedAt: "2025-02-20",
      UpdatedAt: "2025-02-20",
    },
    Descriptions: Descriptions{
      AppDescription: "This is a notification app to give recurring updates on the price of popular Forex symbols",
      AppLogo:        "https://my-portfolio-343207.web.app/MyLogo3.png",
      AppName:        "SAMForexPI2",
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
        Default:  "*/5 * * * *",
      },
    },
    TickURL:   "https://fun-numbers.onrender.com/tick",
    TargetURL: "https://ping.telex.im/v1/webhooks/01950b90-b1bf-75b7-b9e6-e831fdd18b5f",
  },
}
