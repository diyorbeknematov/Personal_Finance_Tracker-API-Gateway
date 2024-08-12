package config

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

func CasbinEnforcer() (*casbin.Enforcer, error) {
	cfg := Load()
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)

	adapter, err := xormadapter.NewAdapter("postgres", conn)
	if err != nil {
		log.Println("Error initializing Casbin adapter: ", err)
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("config/model.conf", adapter)
	if err != nil {
		log.Println("Error initializing Casbin enforcer: ", err)
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		log.Println("Error loading policy from database: ", err)
		return nil, err
	}

	policies := [][]string{
		{"user", "/api/v1", "POST"},
	}

	ok, err := enforcer.AddPolicies(policies)
	if err != nil {
		log.Println("Error adding policies to Casbin: ", err)
		return nil, err
	}
	if !ok {
		log.Println("Error adding policies to Casbin")
	}

	err = enforcer.SavePolicy()
	if err != nil {
		log.Println("Error saving policy to database: ", err)
		return nil, err
	}

	return enforcer, nil
}
