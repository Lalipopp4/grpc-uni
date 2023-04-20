package handler

import (
	"github.com/Lalipopp4/grpc-uni/client/pkg/pb"
	"google.golang.org/grpc"
)

/*
	type GRPCClient struct {
		Student    pb.StudentSearchClient
		Professor  pb.ProfessorSearchClient
		Group      pb.GroupSearchClient
		Department pb.DepartmentSearchClient
		Students   pb.StudentsSearchClient
	}

	func NewClient(conn *grpc.ClientConn) *GRPCClient {
		return &GRPCClient{
			Student:    pb.NewStudentSearchClient(conn),
			Professor:  pb.NewProfessorSearchClient(conn),
			Group:      pb.NewGroupSearchClient(conn),
			Department: pb.NewDepartmentSearchClient(conn),
			Students:   pb.NewStudentsSearchClient(conn),
		}
	}
*/
type GRPCClient struct {
	pb.StudentSearchClient
	pb.ProfessorSearchClient
	pb.GroupSearchClient
	pb.DepartmentSearchClient
	pb.StudentsSearchClient
}

func NewClient(conn *grpc.ClientConn) *GRPCClient {
	return &GRPCClient{
		pb.NewStudentSearchClient(conn),
		pb.NewProfessorSearchClient(conn),
		pb.NewGroupSearchClient(conn),
		pb.NewDepartmentSearchClient(conn),
		pb.NewStudentsSearchClient(conn),
	}
}
