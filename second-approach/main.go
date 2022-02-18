package main


struct final struct {
	disk.DiskApi
	vm.VmApi
}

func (e *final) limit(ctx context.Context, req *pb.LimitRequest) (*pb.LimitResponse, error) {
	return &pb.LimitResponse{}, nil
}