package handler

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/telexintegrations/Forex-Symbols-Notifier/entity"
	"github.com/telexintegrations/Forex-Symbols-Notifier/services"
)


func Get_symbols(ctx *gin.Context){

	var payload entity.MonitorPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		fmt.Println("⚠️ No payload supplied, proceeding without it...")
	}

	symbols := []string{
        "GBPUSD", "EURJPY",
        "EURUSD", "EURCHF",
        "USDCHF", "EURGBP",
        "USDCAD", "AUDCAD",
	}

	var results strings.Builder
	results.WriteString("📈 **Forex Exchange Rates**\n--------------------------\n")

	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, symbol := range symbols {
		wg.Add(1)
		go func(sym string) {
			defer wg.Done()
			base := symbol[:3]  // First 3 letters
			quote := symbol[3:] // Last 3 letters
			var exchangeRate = services.TwelveDemo(base, quote)
			mu.Lock()
			defer mu.Unlock()
			results.WriteString(fmt.Sprintf("%s/%s → 💹 Rate: %s\n", base, quote, exchangeRate)) 
		}(symbol)
	}
	wg.Wait()
	fmt.Println(results.String())
	telex_data := gin.H{
		"message": strings.Join(strings.Split(results.String(), "\n"), "\n"),
		"username": "Samex Forex Update",
        "event_name": "Forex Update",
        "status": "success",
	}
	telex_url := os.Getenv("TELEX_WEBHOOK")
	telresponse, err := services.PostToReturnURL(telex_url, telex_data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(telresponse)

	ctx.JSON(200, telex_data)
}

func Webhook(ctx *gin.Context){
	ctx.JSON(200, entity.Integrationson)
}

func Webhook2(ctx *gin.Context){
	ctx.JSON(200, entity.Integrationson2)
}
