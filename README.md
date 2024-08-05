# Local installation

1. clone the code
2. create database and import sql file from database/data.sql file
3. adjust config.env according to your database setting
4. you can run the app locally from cmd folder then execute go run main.go . For the example, you'll run in localhost:5050 according to config.env but feel free to adjust your own
5. another option is running through docker but the database isn't configured locally yet. This only show a simple example on how to create the Dockerfile and integrated to CI/CD pipeline from github to create the artifact & deployed to GCP cloud run from github workflow which is located at .github/workflows/deploy.yml file

# Several Assumptions

1. One year is calculated as 50 weeks instead of 52 weeks which is common practice in financial industries.
2. GetOutstanding & MakePayment methods are implemented but IsDelinquent isn't because countObligation behaves to calculate how many weeks any customer hasn't made payment as well as checking whether delinquent or not.
3. Scheduler to check periodically any customer is delinquent or not is considered out of scope of this assignment so delinquent check is only done manually when customer make a payment through MakePayment method or check outstanding through GetOutStanding method.
4. Delinquent is based on each customer's loan to keep track customer's behavior.
5. TDD development is implemented inside test/unit folders. However due to the time constrain it covers most part with mostly positive case but it shows how to develop proper TDD behavior with hexagonal pattern.
6. Only simple validation is implemented due to the time constrain to fulfill the task assignment.
7. API doc is available in the root project and can be exported to postman named: billing.postman_collection.json.