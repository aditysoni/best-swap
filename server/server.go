package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aditysoni/flashloan-bot/bestswap"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
)

type DexResponse struct {
	DEX    string `json:"dex"`
	Output string `json:"output"`
}

func RunServer() {
	// Set Gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Create engine with default middleware
	router := gin.Default()

	// Define routes
	router.GET("/best-dex", handleBestDex)

	// Get port from env or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address := ":" + port

	// Print startup message
	fmt.Printf("✅ Server is running on http://localhost%s\n", address)

	// Start server
	if err := router.Run(address); err != nil {
		panic("❌ Failed to start server: " + err.Error())
	}
}

func handleBestDex(c *gin.Context) {
	rpc := os.Getenv("RPC")
	if rpc == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "RPC not provided"})
		return
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to Ethereum node"})
		return
	}
    
	tokenInitial := c.Query("tokenInitial")
	tokenFinal := c.Query("tokenFinal")

	best, err := bestswap.GetBestDexOutput(client , tokenInitial , tokenFinal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting best DEX"})
		return
	}

	c.JSON(http.StatusOK, DexResponse{
		DEX:    best.DEX,
		Output: best.Output.String(),
	})
}
