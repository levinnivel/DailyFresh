package repository

import (
	"log"

	dbHandler "DailyFresh-Backend/database"
	Model "DailyFresh-Backend/domain/model/user"

	_ "github.com/go-sql-driver/mysql"
)

// GetTickets...
func GetTickets(id string) []Model.Ticket {
	db := dbHandler.Connect()
	defer db.Close()

	query := "SELECT * FROM ticket"

	if id != "" {
		query += " WHERE id='" + id + "'"
	}

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
	}

	var ticket Model.Ticket
	var tickets []Model.Ticket
	for rows.Next() {
		if err := rows.Scan(&ticket.ID, &ticket.Category, &ticket.Inquiry, &ticket.Reply, &ticket.UserID); err != nil {
			log.Fatal(err.Error())
		} else {
			tickets = append(tickets, ticket)
		}
	}

	return tickets
}

// PostTicket...
func PostTicket(Ticket Model.Ticket) bool {
	db := dbHandler.Connect()
	defer db.Close()

	_, errQuery := db.Exec("INSERT INTO ticket (category, inquiry, user_id) values (?,?,?)",
		Ticket.Category,
		Ticket.Inquiry,
		Ticket.UserID,
	)

	if errQuery == nil {
		return true
	} else {
		return false
	}

}
