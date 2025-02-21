# Telex-Go-integration


## 🚀 Overview

This repository provides a seamless integration between Telex and your application, enabling real-time monitoring, alerting, and automated notifications for critical events. The integration leverages Telex's HTTP webhooks and API endpoints to enhance observability and incident management. In this particular project, telex is used to create an interval integration which send notifications on Forex trading pair values to the channel hourly. The time can be modified via the JSON integration setting with cron syntax.
 
## 🎯 Features

📡 Real-time event tracking and notifications

⚡ AI-powered insights into critical events

🔗 Easy integration with Telex webhooks and APIs

🛡️ Secure authentication using API tokens

⚙️ Configurable notification channels

## 🛠️ Tech Stack


Node.js (>= 14.x)

npm or yarn

A valid Telex API key (Get it here)

### 📝 Requirements

1. This project requires go 1.23.3, Gin backend framework for go
2.  OPTIONALLY Air module for live reload. 
3. An API_KEY environment variable from [twelvedemo's website](https://twelvedata.com/)
4. A TELEX_WEBHOOK environment variable gotten for [Telex](https://telex.im/)

### 💻 Running Locally

1. Clone this repository by running:
   ```bash
   git clone https://github.com/samuelIkoli/Telex-Go-integration
   cd Telex-Go-integration
   ```
2. Install the dependencies:
   ```bash
   go mod tidy
   ```
3. Start the server in dev mode (if you have air installed):
   ```bash
   air
   ```
   or (if you do not want to use air live reload)
   ```bash
   go run main.go
   ```

### 💻 Testing

Online API testing tools such as **Postman** and **Thunderclient** can be used to test the endpoints or just easily make a get request from your browser to the endpoints in the documentation and expect responses as given in the documentation. Or better still, use your terminal and Curl 😈 . Ultimately, your TELEX_WEBHOOK environment variable is gotten after creating an account, then an organisation and then a channel on [Telex](https://telex.im/)

## 📖 Documentation

## Endpoints


- **URL**: `/integration.json` where you replace {number} with an actual integer
- **Method**: `GET`
- **Body**: No body but a QUERY in form of an integer should be passed
- **Response**:
  ```json
    {
        "data": {
            "date": {
            "created_at": "2025-02-20",
            "updated_at": "2025-02-20"
            },
        "descriptions": {
        "app_description": "This is a notification app to give recurring updates on the price of popular Forex symbols",
          "app_logo": "https://my-portfolio-343207.web.app/MyLogo4.png",
          "app_name": "ForexPI",
          "app_url": "https://fun-numbers.onrender.com/integration.json",
          "background_color": "#fff"
        },
        "integration_category": "Monitoring & Logging",
        "integration_type": "interval",
        "is_active": true,
        "key_features": [
          "Forex",
          "Updates"
        ],
        "author": "Samuel Ikoli",
        "settings": [
          {
            "label": "interval",
            "type": "text",
            "required": true,
            "default": "* * * * *"
          }
        ],
        "tick_url": "https://fun-numbers.onrender.com/tick",
        "target_url": "https://ping.telex.im/v1/webhooks/01950b90-b1bf-75b7-b9e6-e831fdd18b5f"
        }
    }
  ```


- **URL**: `/tick` where you replace {number} with an actual integer
- **Method**: `POST`
- **Body**: No body but a QUERY in form of an integer should be passed
- **Response1**:
  ```json
    {
        "event_name": "Forex Update",
        "message": "📈 **Forex Exchange Rates**\n--------------------------\nEUR/USD → 💹 Rate: 1.04701\nUSD/CAD → 💹 Rate: 1.41940\nEUR/GBP → 💹 Rate: 0.82770002\nEUR/JPY → 💹 Rate: 157.58000\nGBP/USD → 💹 Rate: 1.26510\nUSD/CHF → 💹 Rate: 0.89910001\nAUD/CAD → 💹 Rate: 0.90649998\nEUR/CHF → 💹 Rate: 0.94129997\n",
        "status": "success",
        "username": "Samex Forex Update"
    }
  ```
  **Response2**: On telex'x UI, the response should look like:
  ![alt text](<Screenshot 2025-02-21 at 11.24.10 AM.png>)



## 🔗 Link(s)

- [Telex](https://telex.im/)
- [twelvedemo's website](https://twelvedata.com/)
- [Backlink] (https://hng.tech/hire/golang-developers)
Built by SAMUEL IKOLI
