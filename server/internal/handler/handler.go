package handler

import (
	"context"
	"log"
	"time"

	"github.com/Lalipopp4/grpc-uni/server/pkg/pb"
	"github.com/goombaio/namegenerator"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedStudentSearchServer
	pb.UnimplementedProfessorSearchServer
	pb.UnimplementedGroupSearchServer
	pb.UnimplementedDepartmentSearchServer
	pb.UnimplementedStudentsSearchServer
}

func InitServer(s grpc.ServiceRegistrar) {
	serv := &GRPCServer{}
	pb.RegisterDepartmentSearchServer(s, serv)
	pb.RegisterProfessorSearchServer(s, serv)
	pb.RegisterStudentSearchServer(s, serv)
	pb.RegisterStudentsSearchServer(s, serv)
	pb.RegisterGroupSearchServer(s, serv)
}

// *** Test data ***

var (
	studs  []*pb.Student
	profs  []*pb.Professor
	groups []*pb.Group
	deps   = make(map[string]*pb.Department)
)

func Maker() {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)
	for i := 1; i < 10; i++ {
		studs = append(studs, &pb.Student{
			Id:      int32(i),
			Name:    nameGenerator.Generate(),
			Surname: nameGenerator.Generate(),
		})
		profs = append(profs, &pb.Professor{
			Id:      int32(i),
			Name:    nameGenerator.Generate(),
			Surname: nameGenerator.Generate(),
		})

	}
	dep := [2]string{"bio", "math"}
	for i := 1; i < 2; i++ {
		groups = append(groups, &pb.Group{
			Number:   int32(i),
			Students: studs[i-1 : 5*i],
		})
		deps[dep[i-1]] = &pb.Department{
			Name:       dep[i-1],
			Groups:     groups[i-1 : i],
			Professors: profs[i-1 : 5*i],
		}
	}
}

// ***

func (s *GRPCServer) GetStudent(ctx context.Context, in *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetStudentResponse{
		Student: studs[in.Id],
	}, nil
}

func (s *GRPCServer) GetProfessor(ctx context.Context, in *pb.GetProfessorRequest) (*pb.GetProfessorResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetProfessorResponse{
		Professor: profs[in.Id],
	}, nil
}

func (s *GRPCServer) GetGroup(ctx context.Context, in *pb.GetGroupRequest) (*pb.GetGroupResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetGroupResponse{
		Students: studs[in.Number-1 : in.Number*5],
	}, nil
}

func (s *GRPCServer) GetDepartment(ctx context.Context, in *pb.GetDepartmentRequest) (*pb.GetDepartmentResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetDepartmentResponse{
		Groups:     deps[in.Name].Groups,
		Professors: deps[in.Name].Professors,
	}, nil
}

func (s *GRPCServer) GetStudents(ctx context.Context, in *pb.GetStudentAllRequest) (*pb.GetStudentAllResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetStudentAllResponse{
		Students: studs,
	}, nil
}

/*
func (s *GRPCServer) GetProfessor(ctx context.Context, in *pb.GetProfessorRequest) (*pb.GetProfessorResponse, error){
	return &pb.GetProfessorResponse{
		Professor: &pb.Professor{
			Id: 1,

		}
	}
}
*/
