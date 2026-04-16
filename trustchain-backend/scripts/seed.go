package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type seedCampaign struct {
	Name          string
	Description   string
	GoalUSD       string
	CreatorWallet string
	Milestones    []seedMilestone
}

type seedMilestone struct {
	Name             string
	Description      string
	AmountUSD        string
	DaysFromNow      int
	RequiredEvidence string
}

func main() {
	reset := flag.Bool("reset", false, "clear campaigns/milestones/admin/activity/donations before seeding")
	flag.Parse()

	_ = godotenv.Load(".env")
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("db connect failed: %v", err)
	}
	defer pool.Close()

	tx, err := pool.Begin(ctx)
	if err != nil {
		log.Fatalf("begin tx failed: %v", err)
	}
	defer tx.Rollback(ctx)

	if *reset {
		if _, err := tx.Exec(ctx, `
			TRUNCATE TABLE campaign_activity, admin_actions, access_grants, donations, milestones, campaigns RESTART IDENTITY CASCADE
		`); err != nil {
			log.Fatalf("reset failed: %v", err)
		}
	}

	campaigns := []seedCampaign{
		{
			Name:          "Clean Water Project - Nairobi",
			Description:   "Providing clean water access to 3 rural villages in Kenya through borehole drilling and purification systems.",
			GoalUSD:       "16000",
			CreatorWallet: "0x0000000000000000000000000000000000000001",
			Milestones: []seedMilestone{
				{Name: "Survey and permitting", Description: "Hydrogeological survey, permits, and local council approvals.", AmountUSD: "5300", DaysFromNow: 21, RequiredEvidence: "Permit documents and survey report"},
				{Name: "Drilling and pump installation", Description: "Drill boreholes and install solar-powered pumps.", AmountUSD: "5400", DaysFromNow: 45, RequiredEvidence: "Contractor invoice and site photos"},
				{Name: "Filtration and community training", Description: "Install purification systems and train local operators.", AmountUSD: "5300", DaysFromNow: 75, RequiredEvidence: "Equipment receipts and training attendance log"},
			},
		},
		{
			Name:          "School Rebuild - Rural Kenya",
			Description:   "Rebuilding a primary school damaged by flooding. 4 classrooms for 200 students.",
			GoalUSD:       "25600",
			CreatorWallet: "0x0000000000000000000000000000000000000002",
			Milestones: []seedMilestone{
				{Name: "Site clearing and foundation", Description: "Debris removal, grading, and reinforced foundations.", AmountUSD: "6400", DaysFromNow: 20, RequiredEvidence: "Civil works invoice and geotagged photos"},
				{Name: "Classroom structure", Description: "Walls, roofing, and weatherproofing for 4 classrooms.", AmountUSD: "6400", DaysFromNow: 50, RequiredEvidence: "Material receipts and structural inspection note"},
				{Name: "Interior fit-out", Description: "Flooring, doors, windows, blackboards, and basic furnishing.", AmountUSD: "6400", DaysFromNow: 80, RequiredEvidence: "Supplier invoices and completion photos"},
				{Name: "Utilities and safety", Description: "Water, sanitation, solar lighting, and safety certification.", AmountUSD: "6400", DaysFromNow: 110, RequiredEvidence: "Compliance certificate and utility install proof"},
			},
		},
		{
			Name:          "Medical Aid - Gaza",
			Description:   "Emergency medical supplies for a field hospital serving displaced families.",
			GoalUSD:       "32000",
			CreatorWallet: "0x0000000000000000000000000000000000000003",
			Milestones: []seedMilestone{
				{Name: "Critical supplies procurement", Description: "Procure trauma kits, antibiotics, and emergency consumables.", AmountUSD: "10600", DaysFromNow: 14, RequiredEvidence: "Procurement receipts and inventory list"},
				{Name: "Cold-chain and logistics", Description: "Secure storage and transport for temperature-sensitive medicine.", AmountUSD: "10700", DaysFromNow: 35, RequiredEvidence: "Logistics manifest and cold-chain records"},
				{Name: "Deployment and replenishment", Description: "Distribute to field hospital and reserve stock for 6 weeks.", AmountUSD: "10700", DaysFromNow: 60, RequiredEvidence: "Distribution logs and signed receipt forms"},
			},
		},
	}

	for _, sc := range campaigns {
		var campaignID string
		err := tx.QueryRow(ctx, `
			INSERT INTO campaigns (
				creator_wallet, category, name, description, goal_amount_usd,
				goal_amount_wei, status, creator_name, creator_email,
				creator_country, creator_bio, document_file_names
			)
			VALUES ($1, 'Humanitarian Aid', $2, $3, $4, 0, 'live', 'Seed Creator', 'seed@trustchain.dev', 'Kenya', 'Seed data for local development.', ARRAY['proposal.pdf'])
			RETURNING id::text
		`, sc.CreatorWallet, sc.Name, sc.Description, sc.GoalUSD).Scan(&campaignID)
		if err != nil {
			log.Fatalf("insert campaign failed: %v", err)
		}

		for i, sm := range sc.Milestones {
			deadline := time.Now().UTC().AddDate(0, 0, sm.DaysFromNow)
			if _, err := tx.Exec(ctx, `
				INSERT INTO milestones (
					campaign_id, sequence_index, name, description, amount_usd, amount_wei, deadline, required_evidence, status
				)
				VALUES ($1, $2, $3, $4, $5, 0, $6, $7, 'pending')
			`, campaignID, i, sm.Name, sm.Description, sm.AmountUSD, deadline, sm.RequiredEvidence); err != nil {
				log.Fatalf("insert milestone failed: %v", err)
			}
		}

		if _, err := tx.Exec(ctx, `
			INSERT INTO campaign_activity (type, campaign_id, campaign_name, wallet)
			VALUES ('campaign_live', $1::uuid, $2, $3)
		`, campaignID, sc.Name, sc.CreatorWallet); err != nil {
			log.Fatalf("insert activity failed: %v", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		log.Fatalf("commit failed: %v", err)
	}

	fmt.Println("seed completed successfully")
}
