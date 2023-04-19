package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/arjun/modules/go-dynamoDB/config"
	"github.com/arjun/modules/go-dynamoDB/internal/repository/adapter"
	"github.com/arjun/modules/go-dynamoDB/internal/repository/instance"
	"github.com/arjun/modules/go-dynamoDB/internal/routes"
	"github.com/arjun/modules/go-dynamoDB/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	fmt.Println("Main go")
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)
	logger.INFO("waiting for services", nil)
	errors := Migrate(connection)
	if len(errors) > 0 {
		for err := range errors {

			logger.PANIC("Error on migrate", err)
		}
	}
	logger.PANIC("", checkTables(connection))
	port := fmt.Sprintf("%v", configs.Port)
	router := routes.NewRouter().SetRouters(repository)
	logger.INFO("service is running", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error
	callMigrateAndAppendError(&errors, connection, &RulesProduct.Rules{})
	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface) {
	err := rule.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if response != nil {
		if len(response.TableNames) == 0 {
			logger.INFO("Table not found", nil)
		}
		for _, tableName := range response.TableNames {
			logger.INFO("table Names", tableName)
		}
	}
	return err
}
