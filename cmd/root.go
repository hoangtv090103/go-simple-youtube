package cmd

import (
	"fmt"
	"gosimpleyoutube/config"
	"gosimpleyoutube/internal/repository/postgres"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, world!")
	},
}

func Execute() {
	cfg := config.GetConfig()
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	fmt.Printf("config: %v", cfg)

	db, err := postgres.NewPostgresRepository(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	videoRepo := postgres.NewPostgresVideoRepository(db)

	log.Printf("videoRepo: %v", videoRepo)

	log.Printf("cfg: %v", cfg)

	rootCmd.Execute()
}
