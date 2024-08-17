// Create a new package for shared types
package shared

type Transaction struct {
    ID     string
    Amount int
}

type SmartContract struct {
    ID   string
    Code string
}
