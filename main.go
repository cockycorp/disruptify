package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	Text   string `json:"quote"`
	Author string `json:"author"`
}

var techbroQuotes = []Quote{
	{Text: "We're not just disrupting the market, we're literally reinventing the entire paradigm of synergistic blockchain-enabled AI solutions.", Author: "Chad Thunderbro, CEO of SynergyChain"},
	{Text: "Our growth hacking strategies leverage quantum machine learning to achieve 10x ROI on our viral coefficient.", Author: "Brad Hustle, Chief Disruption Officer"},
	{Text: "We're pivoting from B2B to B2C2B2G with our revolutionary AI-powered, blockchain-secured, cloud-native microservices architecture.", Author: "Derek Visionary, Founder & Thought Leader"},
	{Text: "Sleep is for people who don't have a burning passion to disrupt the napkin industry with blockchain technology.", Author: "Kyle Grindset, Serial Entrepreneur"},
	{Text: "We're not selling software, we're selling a lifestyle revolution powered by deep learning algorithms and growth mindset synergies.", Author: "Tyler Moonshot, Innovation Evangelist"},
	{Text: "Our unicorn-level valuation is driven by our ability to leverage agile methodologies in a hyperscale DevOps environment.", Author: "Connor Rockstar, CTO & Code Ninja"},
	{Text: "I don't do meetings, I do strategic alignment sessions to optimize our go-to-market velocity.", Author: "Austin Hustle, Growth Hacker Extraordinaire"},
	{Text: "We're building the Uber of AI-powered blockchain solutions for the enterprise metaverse.", Author: "Hunter Disrupt, Startup Visionary"},
	{Text: "Failure is just success in beta mode. We're iterating on our MVP in the market of life.", Author: "Mason Pivot, Lean Startup Guru"},
	{Text: "Our revolutionary SaaS platform uses machine learning to disrupt the disruptors who are disrupting disruption.", Author: "Logan Synergy, Chief Evangelist"},
	{Text: "I rise at 4am to cold-plunge while consuming bulletproof coffee and contemplating our Series C fundraising strategy.", Author: "Blake Optimizer, Biohacker CEO"},
	{Text: "We're not just first movers, we're first movers in a blue ocean of untapped vertical integration opportunities.", Author: "Jayden Rockstar, Market Dominator"},
	{Text: "Our culture isn't just innovative, it's a hyper-collaborative ecosystem of rockstar ninjas crushing KPIs.", Author: "Brayden Culture, Chief People Officer"},
	{Text: "I'm not a CEO, I'm a vision architect orchestrating the symphony of exponential growth through mindful disruption.", Author: "Aiden Mindful, Conscious Capitalism Pioneer"},
	{Text: "We're leveraging AI and blockchain to create a paradigm shift in how humanity interfaces with cloud-native solutions.", Author: "Caden Futures, Futurist & Thought Leader"},
	{Text: "Hustle isn't a lifestyle, it's a data-driven methodology for achieving product-market fit in your personal brand.", Author: "Grayson Grind, Personal Brand Architect"},
	{Text: "Our exit strategy? We don't exit, we transcend into a higher plane of market capitalization.", Author: "Carson Moon, Unicorn Whisperer"},
	{Text: "I don't network, I build synergistic relationships with fellow disruptors in the innovation ecosystem.", Author: "Easton Connect, Relationship Hacker"},
	{Text: "We're not just scaling, we're achieving hypergrowth through our proprietary algorithm for virality optimization.", Author: "Landon Scale, Hypergrowth Specialist"},
	{Text: "The only KPI that matters is how many industries you've disrupted before breakfast.", Author: "Parker Metrics, Dashboard Ninja"},
}

func getRandomQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	randomIndex := rand.Intn(len(techbroQuotes))
	quote := techbroQuotes[randomIndex]

	json.NewEncoder(w).Encode(quote)
}

func getQuoteHTML(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	randomIndex := rand.Intn(len(techbroQuotes))
	quote := techbroQuotes[randomIndex]

	html := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Disruptify - Disrupting Quote Generation</title>
    <style>
        body {
            font-family: 'Helvetica Neue', Arial, sans-serif;
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            display: flex;
            justify-content: center;
            align-items: center;
            min-height: 100vh;
            margin: 0;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 20px;
            padding: 50px;
            max-width: 800px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            text-align: center;
        }
        h1 {
            color: #667eea;
            font-size: 3em;
            margin-bottom: 10px;
            font-weight: 800;
        }
        .tagline {
            color: #666;
            font-size: 1.2em;
            margin-bottom: 40px;
            font-style: italic;
        }
        .quote {
            font-size: 1.8em;
            line-height: 1.6;
            color: #333;
            margin: 30px 0;
            font-weight: 300;
            font-style: italic;
        }
        .author {
            font-size: 1.2em;
            color: #764ba2;
            font-weight: 600;
            margin-top: 20px;
        }
        .refresh-btn {
            background: linear-gradient(135deg, #667eea 0%%, #764ba2 100%%);
            color: white;
            border: none;
            padding: 15px 40px;
            font-size: 1.1em;
            border-radius: 50px;
            cursor: pointer;
            margin-top: 30px;
            font-weight: 600;
            transition: transform 0.2s;
        }
        .refresh-btn:hover {
            transform: scale(1.05);
        }
        .api-hint {
            margin-top: 40px;
            padding-top: 30px;
            border-top: 2px solid #eee;
            color: #999;
            font-size: 0.9em;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>🚀 Disruptify</h1>
        <div class="tagline">Disrupting the Quote with AI-Powered Synergistic Solutions</div>
        <div class="quote">"%s"</div>
        <div class="author">— %s</div>
        <button class="refresh-btn" onclick="location.reload()">Get Another Disruption</button>
        <div class="api-hint">
            Pro tip: Hit <strong>/api/quote</strong> for JSON responses to integrate with your blockchain-enabled microservices
        </div>
    </div>
</body>
</html>
`, quote.Text, quote.Author)

	fmt.Fprint(w, html)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "disrupting successfully",
		"message": "All systems are synergistically aligned",
	})
}

func main() {
	rand.Seed(time.Now().UnixNano())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", getQuoteHTML)
	http.HandleFunc("/api/quote", getRandomQuote)
	http.HandleFunc("/health", healthCheck)

	log.Printf("🚀 Disruptify is disrupting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
