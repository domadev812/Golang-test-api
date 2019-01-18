package main

import (
    "encoding/json"
		"net/http"
		"fmt"
		"math/rand"
		"github.com/go-chi/chi"
		"github.com/go-chi/cors"
)

type LoginResultType struct {
	Success  bool
}

type LoginInfoType struct {
	Email string
	Password string
}

type StatsCardType struct {
	CardType			string		`json:"cardType"`
	Icon					string		`json:"icon"`
	Title					string		`json:"title"`
	Value					string		`json:"value"`
	FooterText		string		`json:"footerText"`
	FooterIcon		string		`json:"footerIcon"`
}

type PreferencesChartDataType struct {
	Labels 				[]string	`json:"labels"`
	Series 				[]int			`json:"series"`
}

type DashboardData struct {
	StatsCardData								[]StatsCardType							`json:"statsCardData"`
	PreferencesChartData				PreferencesChartDataType			`json:"preferencesChartData"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Login Request...\n")
	var userInfo LoginInfoType
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	if userInfo.Email == "john@smith.com" {
		fmt.Printf("Login Success...\n")
		var result = LoginResultType{Success: true}
		json.NewEncoder(w).Encode(result)
	} else {
		fmt.Printf("Login Failure...\n")
		var result = LoginResultType{Success: false}
		json.NewEncoder(w).Encode(result)
	}
}

// Testing function to get random dashboard data
func GetDashboardData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Loading Dashboard Data Request...\n")
	var statsCards []StatsCardType
	statsCards = append(statsCards, StatsCardType{CardType:"warning", Icon: "ti-server", Title: "Capacity",	Value: fmt.Sprintf("%dGB", rand.Intn(200)),	FooterText: "Updated now", FooterIcon: "ti-reload"})
	statsCards = append(statsCards, StatsCardType{CardType:"success", Icon: "ti-wallet", Title: "Revenue",	Value: fmt.Sprintf("$%d", rand.Intn(2000)),	FooterText: "Last day", FooterIcon: "ti-calendar"})
	statsCards = append(statsCards, StatsCardType{CardType:"danger", Icon: "ti-pulse", Title: "Errors",	Value: fmt.Sprintf("%d", rand.Intn(100)),	FooterText: "In the last hour", FooterIcon: "ti-timer"})
	statsCards = append(statsCards, StatsCardType{CardType:"info", Icon: "ti-twitter-alt", Title: "Followers",	Value: fmt.Sprintf("+%d", rand.Intn(60)),	FooterText: "Updated now", FooterIcon: "ti-reload"})

	var num1 = rand.Intn(100)
	var num2 = rand.Intn(100 - num1)
	var num3 = 100 - num1 - num2

	var preferencesChartData = PreferencesChartDataType {
		Labels: []string{fmt.Sprintf("%d%%", num1), fmt.Sprintf("%d%%", num2), fmt.Sprintf("%d%%", num3)},
		Series: []int{num1, num2, num3},
	}

	var dashboardData = DashboardData {
		StatsCardData: statsCards,
		PreferencesChartData: preferencesChartData,
	}
	json.NewEncoder(w).Encode(dashboardData)
}

// our main function
func main() {
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(cors.Handler)
	router.Post("/login", Login)
	router.Get("/dashboard-data", GetDashboardData)

	fmt.Printf("Listing Request :5500...\n")
	http.ListenAndServe(":5500", router)
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	/***
		* We need to update this part to make white list.
		if origin == "http://development.com" {
			return true
		} else {
			return false
		}
	*/
	return true
}