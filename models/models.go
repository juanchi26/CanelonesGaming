package models

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
