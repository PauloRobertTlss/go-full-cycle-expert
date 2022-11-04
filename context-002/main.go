package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	//simular a sessão de reserva | Hotel
	cxt := context.Background()
	//ctx, cancel := context.WithCancel(cxt)
	ctx, cancel := context.WithTimeout(cxt, time.Second*3)
	defer cancel()
	BookHotel(ctx)

}

func BookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
	}
}
