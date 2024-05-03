// package message

// import (
// 	"context"

// 	pb "github.com/ivanbatutin921/LiveChatGrpc/proto"
// )

// type server struct{}

// func (s *server) SendMessage(ctx context.Context, req *pb.LiveChatData) (*pb.LiveChatData, error) {
// 	return &pb.LiveChatData{Message: req.Message}, nil
// }

// func (s *server) ReceiveMessage(req *pb.LiveChatData, stream pb.LiveChatData) error {
// 	for {
// 		err := stream.Send(&pb.LiveChatData{Message: req.Message})
// 		if err != nil {
// 			return err
// 		}
// 	}
// }
