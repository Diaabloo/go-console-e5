package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Diaabloo/go-console-e5/weatherApi"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("👉 Entrez le nom d'une ville : ")
	city, _ := reader.ReadString('\n')
	city = strings.TrimSpace(city) // <-- supprime \r, \n et espaces

	weather, err := weatherApi.GetWeather(city)
	if err != nil {
		log.Fatal(err)
	}

	if len(weather.Weather) == 0 {
		log.Fatalf("❌ Pas de données météo trouvées pour %s", city)
	}

	fmt.Println("\n🌍 Rapport météo")
	fmt.Println("──────────────────────────────")
	fmt.Printf("📍 Ville       : %s\n", weather.Name)
	fmt.Printf("🌡️  Température : %.1f °C\n", weather.Main.Temp)
	fmt.Printf("💧 Humidité    : %d %%\n", weather.Main.Humidity)
	fmt.Printf("🔽 Pression    : %d hPa\n", weather.Main.Pressure)
	fmt.Printf("🌬️ Vent        : %.1f m/s\n", weather.Wind.Speed)
	fmt.Printf("☁️  Ciel        : %s\n", weather.Weather[0].Description)
	fmt.Println("──────────────────────────────")
}
