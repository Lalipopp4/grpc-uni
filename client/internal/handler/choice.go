package handler

import (
	"context"
	"fmt"
	"log"

	"github.com/Lalipopp4/grpc-uni/client/pkg/pb"
)

func resser[T any](res T, err error) {
	if err != nil {
		log.Printf("Error in request:  %v", err)
	}
	log.Printf("Result: %v", res)
}

func dialogue() {
	fmt.Print(`Options:
	1) Student by id.
	2) Professor by id.
	3) Department by name.
	4) Group by number.
	5) All students.
	6+) Exit.
Choose:`)
}

func (client *GRPCClient) Request() {
	var (
		ch, id int32
		name   string
		//loop   = true
	)
	for {
		dialogue()
		fmt.Scan(&ch)
		switch ch {
		case 1:
			fmt.Print("Enter id: ")
			fmt.Scan(&id)
			//res, err := client.GetStudent(context.Background(), &pb.GetStudentRequest{Id: id})
			resser(client.GetStudent(context.Background(), &pb.GetStudentRequest{Id: id}))
		case 2:
			fmt.Print("Enter id: ")
			fmt.Scan(&id)
			//res, err := client.GetProfessor(context.Background(), &pb.GetProfessorRequest{Id: id})
			resser(client.GetProfessor(context.Background(), &pb.GetProfessorRequest{Id: id}))
		case 3:
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			//res, err := client.GetDepartment(context.Background(), &pb.GetDepartmentRequest{Name: name})
			resser(client.GetDepartment(context.Background(), &pb.GetDepartmentRequest{Name: name}))
		case 4:
			fmt.Print("Enter number: ")
			fmt.Scan(&id)
			//res, err := client.GetGroup(context.Background(), &pb.GetGroupRequest{Number: id})
			resser(client.GetGroup(context.Background(), &pb.GetGroupRequest{Number: id}))
		case 5:
			//res, err := client.GetStudents(context.Background(), &pb.GetStudentAllRequest{})
			resser(client.GetStudents(context.Background(), &pb.GetStudentAllRequest{}))
		default:
			//return errors.New("connection closed")
			return
		}
	}
}
