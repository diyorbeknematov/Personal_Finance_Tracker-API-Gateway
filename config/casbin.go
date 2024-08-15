package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

func CasbinEnforcer() (*casbin.Enforcer, error) {
	cfg := Load()

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Error connecting to DB: ", err)
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS casbin;")
	if err != nil {
		log.Println("Error dropping DB: ", err)
		return nil, err
	}

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)

	// Adapter'ni inicializatsiya qilish
	adapter, err := xormadapter.NewAdapter("postgres", conn)
	log.Println(conn, err)
	if err != nil {
		log.Println("Error initializing Casbin adapter: ", err)
		return nil, err
	}

	// Casbin Enforcer'ni yaratish
	enforcer, err := casbin.NewEnforcer("config/model.conf", adapter)
	if err != nil {
		log.Println("Error initializing Casbin enforcer: ", err)
		return nil, err
	}

	// Policy'larni yuklash
	err = enforcer.LoadPolicy()
	if err != nil {
		log.Println("Error loading policy from database: ", err)
		return nil, err
	}

	policies := [][]string{
		{"admin", "/api/v1/users/profile", "GET"},
		{"admin", "/api/v1/users/profile", "PUT"},
		{"admin", "/api/v1/users/password", "PUT"},
		{"admin", "/api/v1/users/", "GET"},

		{"admin", "/api/v1/accounts/:id", "GET"},
		{"admin", "/api/v1/accounts/", "POST"},
		{"admin", "/api/v1/accounts/:id", "PUT"},
		{"admin", "/api/v1/accounts/:id", "DELETE"},
		{"admin", "/api/v1/accounts/", "GET"},

		{"admin", "/api/v1/budgets/", "POST"},
		{"admin", "/api/v1/budgets/:id", "GET"},
		{"admin", "/api/v1/budgets/:id", "PUT"},
		{"admin", "/api/v1/budgets/:id", "DELETE"},
		{"admin", "/api/v1/budgets/", "GET"},

		{"admin", "/api/v1/categories/", "POST"},
		{"admin", "/api/v1/categories/:id", "PUT"},
		{"admin", "/api/v1/categories/:id", "DELETE"},
		{"admin", "/api/v1/categories/", "GET"},
		{"admin", "/api/v1/categories/:id", "GET"},

		{"admin", "/api/v1/transactions/", "POST"},
		{"admin", "/api/v1/transactions/:id", "GET"},
		{"admin", "/api/v1/transactions/:id", "PUT"},
		{"admin", "/api/v1/transactions/:id", "DELETE"},
		{"admin", "/api/v1/transactions/", "GET"},

		{"admin", "/api/v1/goals/", "POST"},
		{"admin", "/api/v1/goals/:id", "GET"},
		{"admin", "/api/v1/goals/:id", "PUT"},
		{"admin", "/api/v1/goals/:id", "DELETE"},
		{"admin", "/api/v1/goals/", "GET"},

		{"admin", "/api/v1/reporting/spending", "GET"},
		{"admin", "/api/v1/reporting/income", "GET"},
		{"admin", "/api/v1/reporting/budget-performance", "GET"},
		{"admin", "/api/v1/reporting/goal-progress", "GET"},

		{"admin", "/api/v1/notification/send", "POST"},
		{"admin", "/api/v1/notification/list", "GET"},
		{"admin", "/api/v1/notification/:id", "GET"},
		{"admin", "/api/v1/notification/:id/read", "PUT"},
		{"admin", "/api/v1/notification/:id", "DELETE"},

		{"user", "/api/v1/users/profile", "GET"},
		{"user", "/api/v1/users/profile", "PUT"},
		{"user", "/api/v1/users/password", "PUT"},

		{"user", "/api/v1/accounts/", "POST"},
		{"user", "/api/v1/accounts/:id", "GET"},
		{"user", "/api/v1/accounts/:id", "PUT"},
		{"user", "/api/v1/accounts/:id", "DELETE"},
		{"user", "/api/v1/accounts/", "GET"},

		{"user", "/api/v1/budgets/", "POST"},
		{"user", "/api/v1/budgets/:id", "GET"},
		{"user", "/api/v1/budgets/:id", "PUT"},
		{"user", "/api/v1/budgets/:id", "DELETE"},

		{"user", "/api/v1/categories/:id", "GET"},
		{"user", "/api/v1/categories/", "GET"},

		{"user", "/api/v1/transactions/", "POST"},
		{"user", "/api/v1/transactions/:id", "GET"},
		{"user", "/api/v1/transactions/:id", "PUT"},
		{"user", "/api/v1/transactions/:id", "DELETE"},
		{"user", "/api/v1/transactions/", "GET"},

		{"user", "/api/v1/goals/", "POST"},
		{"user", "/api/v1/goals/:id", "GET"},
		{"user", "/api/v1/goals/:id", "PUT"},
		{"user", "/api/v1/goals/:id", "DELETE"},
		{"user", "/api/v1/goals/", "GET"},

		{"user", "/api/v1/reporting/spending", "GET"},
		{"user", "/api/v1/reporting/income", "GET"},
		{"user", "/api/v1/reporting/budget-performance", "GET"},
		{"user", "/api/v1/reporting/goal-progress", "GET"},

		{"user", "/api/v1/notification/", "GET"},
		{"user", "/api/v1/notification/:id", "GET"},
		{"user", "/api/v1/notification/:id/", "PUT"},
	}

	// Policy'larni qo'shish
	ok, err := enforcer.AddPolicies(policies)
	if err != nil {
		log.Println("Error adding policies to Casbin: ", err)
		return nil, err
	}
	if !ok {
		log.Println("Error adding policies to Casbin")
	}

	// Policy'larni saqlash
	err = enforcer.SavePolicy()
	if err != nil {
		log.Println("Error saving policy to database: ", err)
		return nil, err
	}

	return enforcer, nil
}
