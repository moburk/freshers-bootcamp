package main

import (
	"errors"
	"fmt"
	"sync"
)

type person struct{
	name string
	age int
}

type bankAccount struct {
	person
	mu sync.Mutex
	accountBalance int
}

func (account *bankAccount) deposit(money int, wg *sync.WaitGroup) {
	defer wg.Done()
	account.mu.Lock()
	account.accountBalance += money
	account.mu.Unlock()
}

//func (account *bankAccount) withdraw(money int) error{
//	account.mu.Lock()
//	defer account.mu.Unlock()
//	if account.accountBalance < money {
//		fmt.Println("err")
//		return nil
//		//return errors.New("could not withdraw money due to insufficient balance")
//	}
//	account.accountBalance += money
//	return nil
//}

func main() {
	var person1 = person{
		name : "Gopher",
		age : 22,
	}
	var account1 = bankAccount{
		person: person1,
		accountBalance: 500,
	}
	var wg sync.WaitGroup
	var amtList = []int{-500, -1000, -700, -300, -500}
	for _,amt:= range amtList {
		wg.Add(1)
		if amt >= 0 {
			go account1.deposit(amt, &wg)
		} else {
			go func() {
				account1.mu.Lock()
				defer wg.Done()
				defer account1.mu.Unlock()
				if account1.accountBalance < -1*amt {
					fmt.Println("Error withdrawing", amt, errors.New("could not withdraw money due to insufficient balance"))
				} else {
					account1.accountBalance += amt
				}
			}()
		}
	}
	wg.Wait()
	fmt.Println("Final account balance: ", account1.accountBalance)
}
