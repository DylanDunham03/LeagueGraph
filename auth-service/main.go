package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	// "golang.org/x/oauth2/google"
	gOAuthGoogle "google.golang.org/api/oauth2/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	pb "github.com/DylanDunham03/LeagueGraph/auth-service/protos"
)

type server struct {
	pb.UnimplementedAuthServiceServer
	db *sql.DB
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	log.Println("Connecting to database with DSN:", os.Getenv("MYSQL_DSN"))

	// Initialize MySQL connection pool
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database.")

	// Set maximum number of concurrent open connections
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(25)

	// Create users table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INT AUTO_INCREMENT PRIMARY KEY,
        google_id VARCHAR(255) UNIQUE NOT NULL,
        email VARCHAR(255) NOT NULL
    )`)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	// Initialize gRPC server
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{db: db})
	reflection.Register(s)

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	log.Printf("Auth service listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (s *server) GoogleSignIn(ctx context.Context, req *pb.GoogleSignInRequest) (*pb.AuthResponse, error) {
	// Verify the Google ID token
	tokenInfo, err := verifyGoogleToken(req.IdToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid Google token: %v", err)
	}

	// Check if user exists in database, if not create a new user
	userID, err := s.getOrCreateUser(tokenInfo)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Database error: %v", err)
	}

	// Create a session token (JWT)
	sessionToken, err := createJWT(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create session token: %v", err)
	}

	return &pb.AuthResponse{
		Success:      true,
		UserId:       userID,
		SessionToken: sessionToken,
	}, nil
}

// func verifyGoogleToken(idToken string) (*oauth2.Tokeninfo, error) {
//     client := &http.Client{}
//     tokenInfoURL := "https://oauth2.googleapis.com/tokeninfo?id_token=" + idToken
//     resp, err := client.Get(tokenInfoURL)
//     if err != nil {
//         return nil, err
//     }
//     defer resp.Body.Close()

//     if resp.StatusCode != http.StatusOK {
//         return nil, fmt.Errorf("invalid token")
//     }

//     var tokenInfo oauth2.Tokeninfo
//     if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
//         return nil, err
//     }

//     return &tokenInfo, nil
// }

func verifyGoogleToken(idToken string) (*gOAuthGoogle.Tokeninfo, error) {
	// Initialize the OAuth2 service
	ctx := context.Background()
	oauth2Service, err := gOAuthGoogle.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create oauth2 service: %v", err)
	}

	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, fmt.Errorf("failed to get token info: %v", err)
	}

	return tokenInfo, nil
}

func (s *server) getOrCreateUser(tokenInfo *gOAuthGoogle.Tokeninfo) (string, error) {
	var userID string
	// Use UserId to get the user identifier from tokenInfo
	err := s.db.QueryRow("SELECT id FROM users WHERE google_id = ?", tokenInfo.UserId).Scan(&userID)
	if err == sql.ErrNoRows {
		// User doesn't exist, create a new one
		result, err := s.db.Exec("INSERT INTO users (google_id, email) VALUES (?, ?)", tokenInfo.UserId, tokenInfo.Email)
		if err != nil {
			return "", err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return "", err
		}
		userID = fmt.Sprintf("%d", id)
	} else if err != nil {
		return "", err
	}
	return userID, nil
}

func createJWT(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
