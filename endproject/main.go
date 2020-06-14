package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/eiannone/keyboard"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	User    string
	Address string
	Brand   string
	ID      int
	GosNum  int
}

type Character struct {
	YearV     string
	Developer string
	UserW     string
	GosNum    int
}

type Malfunction struct {
	ID      int
	Malfunc string
	TimeC   string
	GosNum  int
}

var database *sql.DB

func main() {
	var choise_1 int
	pass := 12345
	var pass_check int

	db, err := sql.Open("mysql", "root:12345@/db15")
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	if err = db.Ping(); err != nil {
		log.Println(err)
	}

	database = db
	if err != nil {
		log.Println(err)
	}

	defer database.Close()
	fmt.Printf("Who are u?\n1-Admin\n2-User\n")
	_, err = fmt.Scan(&choise_1)
	fmt.Printf("Write Pass:\n")
	_, err = fmt.Scan(&pass_check)
	if pass_check == pass {
		switch choise_1 {
		case 1:
			{
				for {
					_, _, err := keyboard.GetSingleKey()
					if err != nil {
						log.Println(err)
					}

					cmd := exec.Command("cmd", "/c", "cls")
					cmd.Stdout = os.Stdout
					err = cmd.Run()
					if err != nil {
						log.Fatal(err)
					}
					if choiseVars() == 0 {
						break
					}
				}
			}

		case 2:
			{
				for {
					_, _, err := keyboard.GetSingleKey()
					if err != nil {
						log.Println(err)
					}
					cmd := exec.Command("cmd", "/c", "cls")
					cmd.Stdout = os.Stdout
					err = cmd.Run()
					if err != nil {
						log.Fatal(err)
					}
					if choiseVars_user() == 0 {
						break
					}
				}
			}
		}
	}
}

func choiseVars() int {
	var choise int

	fmt.Println("What do u want?")
	//	fmt.Println("1 - View all of database")
	fmt.Println("1 - Name of car's owner")
	fmt.Println("2 - Add a new User")
	fmt.Println("3 - Info about Car")
	fmt.Println("4 - Name and Address of car's owner")
	fmt.Println("5 - Find worker")
	fmt.Println("6 - Find malfunctions with Gosnumber")
	fmt.Println("7 - Find Users with Malfunction")
	fmt.Println("8 - Update GosNumber")
	fmt.Println("9 - Add new malfunction")
	fmt.Println("10 - View malfunctions with GosNum")
	fmt.Println("11 - get worker name with gosnumber")
	fmt.Println("12 - Delete the User")
	fmt.Println("0 - Exit")
	_, err := fmt.Fscan(os.Stdin, &choise)
	if err != nil {
		fmt.Println("Please write a correct number")
		return 1
	}

	switch choise {

	case 0:
		{
			return 0
		}

	case 1:
		{
			requestGosNum()
			return 1
		}
	case 2:
		{
			addNewUser()
			return 1
		}
	case 3:
		{
			infoAboutCar()
			return 1
		}
	case 4:
		{
			nameAddress()
			return 1
		}
	case 5:
		{
			searchWorker()
			return 1
		}
	case 6:
		{
			findMalf()
			return 1
		}
	case 7:
		{
			findMalfAuto()
			return 1
		}
	case 8:
		{
			updateGosNumber()
			return 1
		}
	case 9:
		{
			addMalf()
			return 1
		}
	case 10:
		{
			viewMalfFromGosNum()
			return 1
		}
	case 11:
		{
			getWorkerName()
			return 1
		}
	case 12:
		{
			deleteU()
			return 1
		}
	default:
		{
			fmt.Println("Please write a correct number")
			return 1
		}

	}
	return 1
}

func choiseVars_user() int {
	var choise int

	fmt.Println("What do u want?")
	fmt.Println("1 - Name of car's owner")
	fmt.Println("2 - Info about Car")
	fmt.Println("3 - Find worker")
	fmt.Println("4 - Find malfunctions with Gosnumber")
	fmt.Println("5 - Find Users with Malfunction")
	fmt.Println("6 - View malfunctions with GosNum")
	fmt.Println("7 - get worker name with gosnumber")
	fmt.Println("0 - Exit")
	_, err := fmt.Fscan(os.Stdin, &choise)
	if err != nil {
		fmt.Println("Please write a correct number")
		return 1
	}
	switch choise {
	case 1:
		{
			requestGosNum()
			return 1
		}
	case 2:
		{
			infoAboutCar()
			return 1
		}
	case 3:
		{
			searchWorker()
			return 1
		}
	case 4:
		{
			findMalf()
			return 1
		}
	case 5:
		{
			findMalfAuto()
			return 1
		}
	case 6:
		{
			viewMalfFromGosNum()
			return 1
		}
	case 7:
		{
			getWorkerName()
			return 1
		}
	default:
		{
			fmt.Println("Please write a correct number")
			return 1
		}

	}
}

func requestGosNum() {
	fmt.Println("Write the GosNumber of the car")
	var gosNum int
	fmt.Fscan(os.Stdin, &gosNum)
	checkResult := false
	rows, err := database.Query("select ID, User, Address, Brand, GosNum from User where GosNum =?", gosNum)
	if err != nil {
		fmt.Println("STUPID DB IS NOT WORKING")
		log.Println(err)
	}
	defer rows.Close()
	Users := []User{}
	for rows.Next() {
		checkResult = true
		u := User{}
		err := rows.Scan(&u.ID, &u.User, &u.Address, &u.Brand, &u.GosNum)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Users = append(Users, u)
		fmt.Printf("ID: %d User: %s Address: %s Brand: %s GosNum: %d \n", u.ID, u.User, u.Address, u.Brand, u.GosNum)

	}
	if checkResult == false {
		fmt.Println("Not found.")
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func addNewUser() {
	var newUser User

	fmt.Println("Write the Name of User, Address, Brand of his Car and Car's GosNum")

	_, err := fmt.Fscan(os.Stdin, &newUser.User)

	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fscan(os.Stdin, &newUser.Address)
	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fscan(os.Stdin, &newUser.Brand)

	if err != nil {
		log.Fatal(err)
	}

	_, err = fmt.Fscan(os.Stdin, &newUser.GosNum)

	if err != nil {
		log.Fatal(err)
	}

	result, err := database.Exec("INSERT  into User (User, Address, Brand, Gosnum) VALUES (?,?,?,?)", newUser.User, newUser.Address, newUser.Brand, newUser.GosNum)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total records in the table 'User' is %d", id)
}

func addMalf() {
	var Malf Malfunction
	var choise int
	fmt.Println("Write the Gosnumber of the Car, which malfunction u wanna add")

	_, err := fmt.Fscan(os.Stdin, &Malf.GosNum)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Write malfunctions:")

	_, err = fmt.Fscan(os.Stdin, &Malf.Malfunc)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Did u eliminate the malfunction?\n1-Y\n2-N\n")

	_, err = fmt.Fscan(os.Stdin, &choise)

	if err != nil {
		log.Fatal(err)
	}
	switch choise {
	case 1:
		{
			fmt.Println("Write time of eliminating the malfunction:")
			_, err = fmt.Fscan(os.Stdin, &Malf.TimeC)
			if err != nil {
				log.Fatal(err)
			}
			_, err := database.Exec("INSERT  into Malfunction (GosNum, Malfunc, TimeC) VALUES (?,?,?)", Malf.GosNum, Malf.Malfunc, Malf.TimeC)
			if err != nil {
				log.Fatal(err)
			}

		}
	case 2:
		{
			_, err := database.Exec("INSERT  into Malfunction (GosNum, Malfunc) VALUES (?,?)", Malf.GosNum, Malf.Malfunc)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func viewAll() {
	rows, err := database.Query("select	user.ID, user.User,	user.Address, user.Brand, user.GosNum, charact.YearV, charact.UserW, malfunction.MalfFunc,malfunction.TimeC from user inner join charact on user.GosNum = charact.GosNum inner join charact on charact.GosNum = malfunction.GosNum")
	if err != nil {
		fmt.Println("STUPID DB IS NOT WORKING")
		log.Println(err)
	}
	defer rows.Close()
	Users := []User{}
	Malfunctions := []Malfunction{}
	Characters := []Character{}
	for rows.Next() {
		u := User{}
		m := Malfunction{}
		c := Character{}
		err := rows.Scan(&u.ID, &u.User, &u.Address, &u.Brand, &u.GosNum, &c.YearV, &c.UserW, &m.Malfunc, &m.TimeC)
		if err != nil {
			fmt.Println(err)
			continue
		}
		Users = append(Users, u)
		Malfunctions = append(Malfunctions, m)
		Characters = append(Characters, c)
		fmt.Printf("ID: %d User: %s Address: %s Brand: %s GosNum: %d  YearV: %v Worker: %v Malfunc: %v TimeC: %v\n", u.ID, u.User, u.Address, u.Brand, u.GosNum, c.YearV, c.UserW, m.Malfunc, m.TimeC)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

type infoAbout struct {
	Brand     string
	Developer string
	YearV     string
}

func infoAboutCar() {
	var username string
	fmt.Println("Write Name of car's onwer:")
	fmt.Fscan(os.Stdin, &username)
	rows, err := database.Query("SELECT user.Brand, charact.Developer, charact.YearV from user inner join charact on user.GosNum = charact.GosNum where user.User = ?", username)
	if err != nil {
		fmt.Println("STUPID DB IS NOT WORKING")
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		i := infoAbout{}
		err = rows.Scan(&i.Brand, &i.Developer, &i.YearV)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("Brand: %s  Developer: %s YearV: %s\n", i.Brand, i.Developer, i.YearV)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}

func nameAddress() {
	checkResult := false
	fmt.Println("Write the GosNumber of the car")
	var gosNum int
	fmt.Fscan(os.Stdin, &gosNum)

	rows, err := database.Query("select  User, Address, GosNum from User where GosNum =?", gosNum)
	if err != nil {
		fmt.Println("STUPID DB IS NOT WORKING")
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		checkResult = true
		u := User{}
		err := rows.Scan(&u.User, &u.Address, &u.GosNum)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("ID: %d User: %s Address: %s Brand: %s GosNum: %d \n", u.ID, u.User, u.Address, u.Brand, u.GosNum)

	}
	if checkResult == false {
		fmt.Println("Not found.")
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func searchWorker() {
	checkResult := false
	fmt.Println("Write Worker's name:")
	var name string

	_, err := fmt.Fscan(os.Stdin, &name)
	if err != nil {
		log.Println(err)
	}
	rows, err := database.Query("select  Developer, GosNum from charact where UserW =?", name)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	fmt.Printf("Refurbished machine worker: %v\n", name)
	for rows.Next() {
		checkResult = true
		var dev string
		var gos int
		err := rows.Scan(&dev, &gos)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("Developer: %v   GosNumer: %v\n", dev, gos)

	}
	if checkResult == false {
		fmt.Println("Not found.")
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func findMalf() {
	checkResult := false
	fmt.Println("Write car's Gosnumer:")
	var gos int
	_, err := fmt.Fscan(os.Stdin, &gos)
	if err != nil {
		log.Println(err)
	}
	rows, err := database.Query("select  Malfunc from malfunction where GosNum =?", gos)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	fmt.Printf("Malfunctions of '%v':", gos)
	for rows.Next() {
		checkResult = true
		var Malf string
		err := rows.Scan(&Malf)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("  %v", Malf)

	}
	fmt.Printf(".\n")
	if checkResult == false {
		fmt.Println("Not found.")
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func findMalfAuto() {
	fmt.Println("Write Malfunction:")
	var Malf string
	check := false
	_, err := fmt.Fscan(os.Stdin, &Malf)
	if err != nil {
		log.Println(err)
	}
	gosArr := make([]int, 0)
	var gos int
	rows, err := database.Query("select  GosNum from malfunction where Malfunc =?", Malf)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		check = true
		err := rows.Scan(&gos)
		if err != nil {
			log.Println(err)
		}
		gosArr = append(gosArr, gos)
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	}
	if check != true {
		fmt.Println("Malfunction not found.")
		return
	}
	check = false
	var name string
	fmt.Printf("Users with Malfunction %v:\n", Malf)

	for _, value := range gosArr {
		rows, err := database.Query("select  User from user where GosNum =?", value)
		if err != nil {
			log.Println(err)
		}

		for rows.Next() {
			check = true
			err := rows.Scan(&name)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(name)
			if err := rows.Err(); err != nil {
				log.Fatal(err)
			}
		}
		defer rows.Close()
	}
}

func updateGosNumber() {
	fmt.Println("Write GosNumber, which u wanna update:")
	var gos, changedGos int
	_, err := fmt.Fscan(os.Stdin, &gos)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Write GosNumber, which u wanna see in DB")
	_, err = fmt.Fscan(os.Stdin, &changedGos)
	if err != nil {
		log.Println(err)
	}
	_, err = database.Exec("UPDATE user set GosNum = ? where GosNum = ?", changedGos, gos)
	if err != nil {
		log.Println(err)
	}
	_, err = database.Exec("UPDATE charact set GosNum = ? where GosNum = ?", changedGos, gos)
	if err != nil {
		log.Println(err)
	}
	_, err = database.Exec("UPDATE malfunction set GosNum = ? where GosNum = ?", changedGos, gos)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("The changes took effect.")
}

func viewMalfFromGosNum() {
	fmt.Println("Write GosNum which malfunctions u wanna find:")
	var gos int

	_, err := fmt.Fscan(os.Stdin, &gos)
	if err != nil {
		log.Println(err)
	}
	rows, err := database.Query("select Malfunc from malfunction where GosNum = ?", gos)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	fmt.Printf("Malfunctions of %v:", gos)
	for rows.Next() {
		var milfa string
		err = rows.Scan(&milfa)
		fmt.Printf(" %v", milfa)
	}
	fmt.Printf(". \n")

}

func getWorkerName() {
	var Time string
	var gos int
	var milfa, name string
	fmt.Println("Write Name of user")
	_, err := fmt.Fscan(os.Stdin, &name)
	if err != nil {
		log.Println(err)
	}
	rows, err := database.Query("select GosNum from user  where User = ?", name)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&gos)
	rows, err = database.Query("select UserW from charact where GosNum = ?", gos)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	rows.Next()
	err = rows.Scan(&name)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Write malfunction:")
	_, err = fmt.Fscan(os.Stdin, &milfa)
	if err != nil {
		log.Println(err)
	}
	rows, err = database.Query("select TimeC from malfunction where Malfunc = ? and GosNum = ?", milfa, gos)
	rows.Next()
	err = rows.Scan(&Time)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("GosNumber:%v  Malfunction:%v  Worker:%v\n", gos, milfa, name)
}

func deleteU() {
	var name string
	var gos int
	fmt.Println("Write Name of User which u wanna delete:")
	_, err := fmt.Fscan(os.Stdin, &name)
	if err != nil {
		log.Println(err)
	}
	rows, err := database.Query("select GosNum from user where User = ?", name)
	rows.Next()
	err = rows.Scan(&gos)
	result, err := database.Exec("DELETE from  user where GosNum = ?", gos)
	if err != nil {
		panic(err)
	}
	last, _ := result.RowsAffected()
	fmt.Printf("From Users has been deleted %v rows\n", last)
	result, err = database.Exec("DELETE from  charact where GosNum = ?", gos)
	if err != nil {
		panic(err)
	}
	last, _ = result.RowsAffected()
	fmt.Printf("From Character has been deleted %v rows\n", last)
	result, err = database.Exec("DELETE from  malfunction where GosNum = ?", gos)
	if err != nil {
		panic(err)
	}
	last, _ = result.RowsAffected()
	fmt.Printf("From Character has been deleted %v rows\n", last)

}
