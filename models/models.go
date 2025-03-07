package models

import "time"

type SecretRDSjson struct {
	Username            string `json:"username"` //con backsticks
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"DbClusterIdentifier"`
}

type SignUp struct {
	UserEmail string `json:"userEmail"`
	UserUUID  string `json:"userUUID"`
}

type Category struct {
	CategID   int    `json:"categId"`
	CategName string `json:"categName"`
	CategPath string `json:"categPath"`
}

type Product struct {
	ProdId          int       `json:"prodID"`
	ProdTitle       string    `json:"prodTitle"`
	ProdDescription string    `json:"prodDescription"`
	ProdCreatedAt   time.Time `json:"prodCreatedAt"`
	ProdUpdated     time.Time `json:"prodUpdated"`
	ProdPrice       float64   `json:"prodPrice,omitempty"`
	ProdStock       int       `json:"prodStock"`
	ProdCategId     int       `json:"prodCategId"`
	ProdPath        string    `json:"prodPath"`
	ProdSearch      string    `json:"search,omitempty"`
	ProdCategPath   string    `json:"prodCategPath,omitempty"`
}

type ProductResp struct {
	TotalItems int       `json:"totalItems"`
	Data       []Product `json:"data"`
}
