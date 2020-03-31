package main

import (
    "fmt"

    "github.com/2nolife/moduluschecking/models"
    "github.com/2nolife/moduluschecking/parsers"
    "github.com/2nolife/moduluschecking/resolvers"
)

func main() {
    // Read the modulus weight table and the sorting
    // code substitution table from the folder data
    parser := parsers.CreateFileParser()

    // The resolver handles the verification of the validity of
    // bank accounts according to the data obtained by the parser
    resolver := resolvers.NewResolver(parser)

    // This helper method handles special cases for
    // bank accounts from:
    // - National Westminster Bank plc (10 or 11 digits with possible presence of dashes, for account numbers)
    // - Co-Operative Bank plc (10 digits for account numbers)
    // - Santander (9 digits for account numbers)
    // - banks with 6 or 7 digits for account numbers
    bankAccount := models.CreateBankAccount("089999", "66374958")

    // Check if the created bank account is valid against the rules
    fmt.Print("Bank account 08-99-99 66374958. Valid: ")
    fmt.Println(resolver.IsValid(bankAccount))
}
