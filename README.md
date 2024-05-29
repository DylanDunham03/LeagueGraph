# LeagueGraph

### League of Legends friend graph


### Getting Starteed

Install:

```
go get github.com/neo4j/neo4j-go-driver/v4/neo4j

go get github.com/joho/godotenv

go get -u google.golang.org/grpc

```



Step 1: Define API Endpoints
Begin by outlining the API endpoints that your services will expose. This will guide the development of both the backend and the frontend. For instance:

Authentication Service: Endpoints for user registration, login, and token refresh.
Player Graph Service: Endpoints to retrieve player matches, get player details, and fetch the graph of connections.
Step 2: Develop Backend Services
Authentication Service:

Implement user authentication using JWT (JSON Web Tokens).
Set up routes for login, registration, and token validation.
Integrate a database like PostgreSQL for user management.
Player Graph Service:

Implement the logic to interact with the Riot Games API to fetch match data.
Develop the functionality to parse and store this data in Neo4j.
Create algorithms to build and query the graph of player interactions.
gRPC Gateway:

Set up gRPC to handle requests and route them to appropriate services.
Define Protobufs for your services.
Implement gRPC methods in Go.
Step 3: Set Up the Database
For Authentication: Configure and initialize your PostgreSQL database.
For Player Graph: Set up Neo4j, define your graph schema, and ensure it's optimized for the queries you anticipate.
Step 4: Develop the Frontend
Create the UI components using React.
Implement state management with tools like Redux or Context API to handle data fetched from your backend.
Connect the frontend to your backend using REST API calls or gRPC-Web if you want to maintain consistency with gRPC across your stack.
Step 5: Testing
Write unit tests for both backend and frontend components.
Use integration tests to ensure that your services work well together.
Consider end-to-end tests for critical workflows like user registration, login, and graph querying.
Step 6: Continuous Integration/Continuous Deployment (CI/CD)
Enhance your GitHub Actions workflows to include:
Building and testing your application.
Docker image creation and pushing to a Docker registry.
Automated deployments to your chosen cloud provider.
Step 7: Documentation
Document your API using tools like Swagger or Postman.
Write user guides or READMEs for your repository to help future contributors understand and use your project.
Step 8: Review and Iterate
Review your application’s performance and usability.
Gather feedback from potential users and refine the application based on this feedback.
Iterate on new features and improvements.
Step 9: Deployment
Prepare for production by ensuring your application is secure and scalable.
Deploy your services to the cloud, and monitor their performance.