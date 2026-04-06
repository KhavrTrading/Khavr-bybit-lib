package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/KhavrTrading/Khavr-bybit-lib/bybit"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit/bybit_models"
	"github.com/KhavrTrading/Khavr-bybit-lib/bybit_ws"

	log "github.com/sirupsen/logrus"
)

func loadEnv(path string) {
	f, err := os.Open(path)
	if err != nil {
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			os.Setenv(parts[0], parts[1])
		}
	}
}

func main() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	loadEnv(".env")

	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	if apiKey == "" || apiSecret == "" {
		log.Fatal("API_KEY and API_SECRET must be set")
	}

	// ── REST: fetch initial wallet balance ──
	rest := bybit.NewBybitClient(apiKey, apiSecret)
	walletResp, err := rest.GetWalletBalance(&bybit_models.WalletBalanceRequest{
		AccountType: "UNIFIED",
	})
	if err != nil {
		log.Fatalf("GetWalletBalance failed: %v", err)
	}

	fmt.Println("═══ INITIAL WALLET BALANCE ═══")
	for _, acct := range walletResp.Result.List {
		fmt.Printf("Account: %s | Equity: %s | Balance: %s | Available: %s | UPL: %s\n",
			acct.AccountType, acct.TotalEquity, acct.TotalWalletBalance,
			acct.TotalAvailableBalance, acct.TotalPerpUPL)
		for _, c := range acct.Coin {
			if c.WalletBalance != "0" && c.WalletBalance != "" {
				fmt.Printf("  %-6s  balance=%-14s  equity=%-14s  uPnL=%-14s  free=%s\n",
					c.Coin, c.WalletBalance, c.Equity, c.UnrealisedPnl, c.Free)
			}
		}
	}
	fmt.Println("══════════════════════════════")

	// ── WS: live updates ──
	ws := bybit_ws.NewBybitWsClient(apiKey, apiSecret)

	positionCh := bybit_ws.NewPositionChannel(ws)
	orderCh := bybit_ws.NewOrderChannel(ws)
	executionCh := bybit_ws.NewExecutionChannel(ws)
	walletCh := bybit_ws.NewWalletChannel(ws)

	posSub := positionCh.Subscribe()
	orderSub := orderCh.Subscribe()
	execSub := executionCh.Subscribe()
	walletSub := walletCh.Subscribe()

	if err := ws.Connect(); err != nil {
		log.Fatalf("WS Connect failed: %v", err)
	}
	fmt.Println("\nWebSocket connected. Subscribing...")

	positionCh.SubscribeToPositions("position.linear")
	orderCh.SubscribeToOrders("order.linear")
	executionCh.SubscribeToExecutions("execution.linear")
	walletCh.SubscribeToWallet()

	fmt.Println("Listening for live updates... (Ctrl+C to stop)")

	go func() {
		for ev := range posSub {
			for _, p := range ev.Data {
				fmt.Printf("[POSITION] %s %s size=%s entry=%s uPnL=%s liq=%s\n",
					p.Symbol, p.Side, p.Size, p.EntryPrice, p.UnrealisedPnl, p.LiqPrice)
			}
		}
	}()

	go func() {
		for ev := range orderSub {
			for _, o := range ev.Data {
				fmt.Printf("[ORDER] %s %s %s qty=%s price=%s status=%s\n",
					o.Symbol, o.Side, o.OrderType, o.Qty, o.Price, o.OrderStatus)
			}
		}
	}()

	go func() {
		for ev := range execSub {
			for _, e := range ev.Data {
				fmt.Printf("[EXEC] %s %s qty=%s price=%s pnl=%s fee=%s\n",
					e.Symbol, e.Side, e.ExecQty, e.ExecPrice, e.ExecPnl, e.ExecFee)
			}
		}
	}()

	go func() {
		for ev := range walletSub {
			fmt.Println("[WALLET UPDATE]")
			for _, w := range ev.Data {
				fmt.Printf("  Equity=%s Balance=%s Available=%s UPL=%s\n",
					w.TotalEquity, w.TotalWalletBalance, w.TotalAvailableBalance, w.TotalPerpUPL)
				for _, c := range w.Coin {
					if c.WalletBalance != "0" && c.WalletBalance != "" {
						fmt.Printf("    %-6s balance=%s equity=%s\n", c.Coin, c.WalletBalance, c.Equity)
					}
				}
			}
		}
	}()

	go ws.ListenLoop()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nShutting down...")
	ws.Stop()
}
