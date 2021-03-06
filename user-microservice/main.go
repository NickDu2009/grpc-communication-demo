package main





import (
	rolesPb "github.com/NickDu2009/grpc-communication-demo/roles-microservice/pb"
	pb "github.com/NickDu2009/grpc-communication-demo/user-microservice/pb"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	"github.com/pkg/errors"
	"net"
)

type Server struct {
	users []*pb.User
	rolesClient rolesPb.RolesClient
}

func getRolesClient() rolesPb.RolesClient {
	conn, err := grpc.Dial("localhost:6000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to start gRPC connection: %v", err)
	}

	return rolesPb.NewRolesClient(conn)
}

func (s *Server) GetUser(_ context.Context, req *pb.GetUserRequest) (*pb.UserReply, error) {
	if req.UserId < 0 || req.UserId > int32(len(s.users)) {
		return nil, errors.New("invalid user")
	}

	user := s.users[req.UserId]
	roleReq := &rolesPb.GetUserRoleRequest{
		UserId:req.UserId,
	}

	rolesReply, err := s.rolesClient.GetUserRole(context.Background(), roleReq)
	if err != nil {
		return  nil, err
	}

	roles := make([]*pb.Role, 0)
	for _, role := range rolesReply.Roles {
		roles = append(roles, &pb.Role{
			Id:role.Id,
			Name:role.Name,
		})
	}

	return &pb.UserReply{
		User:user,
		Roles:roles,
	}, nil
}

func main() {
	users := []*pb.User{
		{
			Id:1,
			Email:"bob@example.com",
			Name:"Bob",
		},
		{
			Id:2,
			Email:"amy@example.com",
			Name:"Amy",
		},
		{
			Id:3,
			Email:"geore@example.com",
			Name:"Geore",
		},
		{
			Id:4,
			Email:"jacob@example.com",
			Name:"Jacob",
		},
	}

	lis, err := net.Listen("tcp", "localhost:7000")
	if err != nil {
		log.Fatalf("failed to initializa TCP listen: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	roleServer := &Server{
		users:users,
		rolesClient:getRolesClient(),
	}

	pb.RegisterUsersServer(server, roleServer)

	server.Serve(lis)
}
