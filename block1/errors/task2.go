package main

// не меняйте импорты, они нужны для проверки
//https://stepik.org/lesson/526881/step/7?unit=519600
import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

// account представляет счет
type account struct {
	balance   int
	overdraft int
}

func main() {
	var acc account
	var trans []int
	acc, trans, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Print("-> ")
	fmt.Println(acc, trans)
}

// parseInput считывает счет и список транзакций из os.Stdin.
func parseInput() (account, []int, error) {
	accSrc, transSrc := readInput()
	acc, err := parseAccount(accSrc)
	if err != nil {
		return acc, nil, err
	}
	trans, err := parseTransactions(transSrc)
	return acc, trans, err
}

// readInput возвращает строку, которая описывает счет
// и срез строк, который описывает список транзакций.
// эту функцию можно не менять
func readInput() (string, []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	accSrc := scanner.Text()
	var transSrc []string
	for scanner.Scan() {
		transSrc = append(transSrc, scanner.Text())
	}
	return accSrc, transSrc
}

// parseAccount парсит счет из строки
// в формате balance/overdraft.
func parseAccount(src string) (account, error) {
	parts := strings.Split(src, "/")
	balance, err := strconv.Atoi(parts[0])
	if err != nil {
		return account{0, 0}, err
	}
	overdraft, err := strconv.Atoi(parts[1])
	if err != nil {
		return account{0, 0}, err
	}
	if overdraft < 0 {
		err := fmt.Errorf("expect overdraft >= 0")
		return account{balance, overdraft}, err
	}
	if balance < -overdraft {
		err := fmt.Errorf("balance cannot exceed overdraft")
		return account{balance, overdraft}, err
	}
	return account{balance, overdraft}, nil
}

// parseTransactions парсит список транзакций из строки
// в формате [t1 t2 t3 ... tn].
func parseTransactions(src []string) ([]int, error) {
	trans := make([]int, len(src))
	for idx, s := range src {
		t, err := strconv.Atoi(s)
		if err != nil {
			return trans, err
		}
		trans[idx] = t
	}
	return trans, nil
}
