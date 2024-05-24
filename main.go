package main

import (
	"fmt"
)

// improve projek:
// 1. return jumlah stock gas saat ini 
// 2. buat api post untuk coustomer yang menginputkan data gas nya 
// 3. simpan data nya ke database 

var maxGasStock int = 100

type Gas struct {
	ID int 
	Tersedia bool 
	Stock int
}

type CreateGasRequest struct {
	Tersedia bool `json:"tersedia"`
	Stock int	`json:"stock"`
}

func (g *Gas) UpdateAvailability() {
    if g.Stock <= maxGasStock {
        g.Tersedia = true
    } else {
        g.Tersedia = false
    }
}

func NewGas(req *CreateGasRequest) (*Gas, error) {
	if err := validatedInputGas(req); err != nil {
		return nil, err
	}
	// parts := strings.Split(strings.ToLower(req.Name), " ")
	// slug := strings.Join(parts, "-")
	gas := &Gas{
		Tersedia: req.Tersedia,
		Stock: req.Stock,
	}
	gas.UpdateAvailability()
	return gas, nil
}

func validatedInputGas(req *CreateGasRequest) error {
	if req.Stock >= maxGasStock {
		return fmt.Errorf("the stock of the gas is to short")
	}
	return nil
}

func main(){
	gas := &Gas{Stock: 110} // contoh nilai stock
	gas.UpdateAvailability() // memperbarui status ketersediaan
	fmt.Println("Gas tersedia:", gas.Tersedia) // output: Gas tersedia: false
}
